package database

import (
	"database/sql"
	"log"
	"strings"

	dto "meli.test/dtos"
)

const QRY_CREATE_TABLE_DNA = `
	CREATE TABLE IF NOT EXISTS dnas (
  		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  		sequence TEXT NOT NULL UNIQUE,
  		is_mutant INTEGER NOT NULL
  	);`

const QRY_INSERT_INTO_DNA = `
	INSERT INTO dnas
	VALUES(NULL, ?, ?);`

const QRY_SELECT_DNA_BY_SEQUENCE = `
	SELECT
		id,
		sequence,
		is_mutant
	FROM dnas
	WHERE sequence=?`

const QRY_SELECT_DNA_STATS = `
	SELECT
		( SELECT COUNT(id) FROM dnas WHERE is_mutant = 1 ) as mutant_count,
		( SELECT COUNT(id) FROM dnas WHERE is_mutant = 0 ) as human_count;`

func AddDna(dna *dto.Dna) {
	db := GetConnection()
	res, err := db.Exec(
		QRY_INSERT_INTO_DNA,
		strings.Join(dna.Sequence, ","),
		dna.IsMutant,
	)

	defer db.Close()

	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			log.Println("the dna couldn't be added to the database because already exists")
		} else {
			log.Printf("the dna couldn't be added to the database - details: %s", err.Error())
		}

		return
	}

	var id int64

	if id, err = res.LastInsertId(); err != nil {
		log.Printf("the dna couldn't be added to the database - details: %s", err.Error())
		return
	}

	dna.Id = int(id)

	log.Println("the dna was added successfully to the database ( id=", id, ")")
}

func FindDnaBySequence(sequence []string) *dto.Dna {
	db := GetConnection()
	row := db.QueryRow(
		QRY_SELECT_DNA_BY_SEQUENCE,
		strings.Join(sequence, ","),
	)

	defer db.Close()

	p_id := 0
	p_sequence := ""
	p_is_mutant := 0

	if err := row.Scan(&p_id, &p_sequence, &p_is_mutant); err == sql.ErrNoRows {
		log.Println("the given dna sequence doesn't exists and could be analized")
		return nil
	}

	return &dto.Dna{
		Id:       p_id,
		Sequence: strings.Split(p_sequence, ","),
		IsMutant: p_is_mutant == 1,
	}
}

func GetDnaStats() *dto.DnaStats {
	db := GetConnection()
	row := db.QueryRow(QRY_SELECT_DNA_STATS)

	defer db.Close()

	p_mutant_count := 0
	p_human_count := 0

	_ = row.Scan(&p_mutant_count, &p_human_count)

	ratio := calculateRatio(p_mutant_count, p_human_count)

	return &dto.DnaStats{
		MutantsQty: p_mutant_count,
		HumansQty:  p_human_count,
		Ratio:      ratio,
	}
}

func calculateRatio(mutantsQty int, humansQty int) float64 {
	if humansQty != 0 {
		return float64(int((float64(mutantsQty)/float64(humansQty))*100)) / 100
	}

	return float64(mutantsQty)
}
