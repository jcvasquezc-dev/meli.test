# Exámen Mercadolibre

---

## Tabla de Contenido

- [Sobre este Proyecto](#about)
- [Preparando el Proyecto](#getting_started)
- [Cómo usar la Aplicación](#usage)
- [Consideraciones](#considerations)

---

## Sobre este Proyecto <a name = "about"></a>

El desarrollo de este proyecto tiene como finalidad cumplir con ciertos requerimientos demandados por el equipo técnico de Mercadolibre. De todos los lenguajes de mi stack, este es completamente nuevo para mi, y lo planteé como forma de reto.

### El Enunciado

Magneto quiere reclutar la mayor cantidad de mutantes para poder luchar contra los X-Men.

Te ha contratado a ti para que desarrolles un proyecto que detecte si un humano es mutante basándose en su  ecuencia de ADN.

Para eso te ha pedido crear un programa con un método o función con la siguiente firma (En alguno de los siguiente lenguajes: Java / Golang / C-C++ / Javascript (node) / Python / Ruby):

```
boolean isMutant(String[] dna); // Ejemplo Java
```

En donde recibirás como parámetro un array de Strings que representan cada fila de una tabla de (NxN) con la secuencia del ADN. Las letras de los Strings solo pueden ser: (A,T,C,G), las cuales representa cada base nitrogenada del ADN.

<table>
    <thead>
        <tr><th>No-Mutante</th><th>Mutante</th></tr>
    </thead>
    <tbody>
        <tr>
            <td>
                <table>
                <tr><td>A</td><td>T</td><td>G</td><td>C</td><td>G</td><td>A</td></tr>
                <tr><td>C</td><td>A</td><td>G</td><td>T</td><td>G</td><td>C</td></tr>
                <tr><td>T</td><td>T</td><td>A</td><td>T</td><td>T</td><td>T</td></tr>
                <tr><td>A</td><td>G</td><td>A</td><td>C</td><td>G</td><td>G</td></tr>
                <tr><td>G</td><td>C</td><td>G</td><td>T</td><td>C</td><td>A</td></tr>
                <tr><td>T</td><td>C</td><td>A</td><td>C</td><td>T</td><td>G</td></tr>
                </table>
            </td>
            <td>
                <table>
                <tr><td style="color: green">A</td><td>T</td><td>G</td><td>C</td><td style="color: blue">G</td><td>A</td></tr>
                <tr><td>C</td><td style="color: green">A</td><td>G</td><td>T</td><td style="color: blue">G</td><td>C</td></tr>
                <tr><td>T</td><td>T</td><td style="color: green">A</td><td>T</td><td style="color: blue">G</td><td>T</td></tr>
                <tr><td>A</td><td>G</td><td>A</td><td style="color: green">A</td><td style="color: blue">G</td><td>G</td></tr>
                <tr><td style="color: red">C</td><td style="color: red">C</td><td style="color: red">C</td><td style="color: red">C</td><td>T</td><td>A</td></tr>
                <tr><td>T</td><td>C</td><td>A</td><td>C</td><td>T</td><td>G</td></tr>
                </table>
            </td>
        </tr>
    </tbody>
</table>

Sabrás si un humano es mutante, si encuentras más de una secuencia de cuatro letras iguales, de forma  blicua, horizontal o vertical.

#### Ejemplo (Caso mutante):

```
String[] dna = {"ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"};
```

En este caso el llamado a la función isMutant(dna) devuelve “true”.

Desarrolla el algoritmo de la manera más eficiente posible.

#### Desafíos:

***Nivel 1:***

Programa (en cualquier lenguaje de programación) que cumpla con el método pedido por Magneto.

***Nivel 2:***

Crear una API REST, hostear esa API en un cloud computing libre (Google App Engine, Amazon AWS, etc),  crear el servicio “/mutant/” en donde se pueda detectar si un humano es mutante enviando la secuencia de ADN mediante un HTTP POST con un Json el cual tenga el siguiente formato:

```
POST → /mutant/
{
“dna”:["ATGCGA","CAGTGC","TTATGT","AGAAGG","CCCCTA","TCACTG"]
}
```

En caso de verificar un mutante, debería devolver un HTTP 200-OK, en caso contrario un 403-Forbidden

***Nivel 3:***

Anexar una base de datos, la cual guarde los ADN’s verificados con la API. Solo 1 registro por ADN. Exponer un servicio extra “/stats” que devuelva un Json con las estadísticas de las
verificaciones de ADN:

```
{“count_mutant_dna”:40, “count_human_dna”:100: “ratio”:0.4}
```

Tener en cuenta que la API puede recibir fluctuaciones agresivas de tráfico (Entre 100 y 1
millón de peticiones por segundo).

Test-Automáticos, Code coverage > 80%

***Entregar:***

- Código Fuente (Para Nivel 2 y 3: En repositorio github).
- Instrucciones de cómo ejecutar el programa o la API. (Para Nivel 2 y 3: En README de
github).
- URL de la API (Nivel 2 y 3).

---

## Preparando el Proyecto <a name = "getting_started"></a>

Las instrucciones a continuación les ofrecerán una copia del proyecto ya configurado y preparado para ejecutarse en un entorno local en un sistema operativo linux/ubuntu.

### Prerequisites

El desarrollo fue realizado bajo el lenguaje GO. No se utilizaron frameworks en particular. Para saber cómo instalar y configurar GO en tu entorno local revisa la guía oficial en https://go.dev/doc/install.

Versión de GO:

```
go version go1.18.1 linux/amd64
```

### Configuración

Necesitaremos instalar **gcc** en nuestro entorno (linux/ubuntu), y para ello vamos a ejecutar los siguientes comandos:

Actualizamos la lista de paquetes del sistema operativo

```
$ sudo apt update
```

Instalamos **build-essential**

```
$ sudo apt-get install manpages-dev
```

Comprobamos que tenemos instalado **gcc** ejecutando

```
$ gcc --version
```

Hasta aquí, todo ok.

Ahora, el proyecto usa una base de datos local sqlite3. Esta no requiere de otro tipo de instalación para que funcione, a menos que deseen instalar algún visor de su parte e ir verificando los registros y la estructura de la base de datos. Si queremos caambiar el nombre del fichero que crea sqlite3, este se encuentra en el fichero **database/database.go**

```
const dbFile = "xmen.db"
```

El último parámetro que tal vez nos insterese modificar, puede ser el puerto donde levanta nuestra api. Por defecto levanta en el puerto 8080. En caso de que querramos modificar este parámetro, podemos hacerlo en nuestro fichero **main.go**

```
srv := server.Create(":8080")
```

Creo que está de más comentar que el puerto debe estar habilitado para recibir peticiones en donde le sea configurado. Recordemos que usar iptables para redireccionar las peticiones de un puerto oficial (como el 80 o 443), a nuestro puerto personlizado, también es una práctica aceptable.

---

## Cómo usar la Aplicación <a name = "usage"></a>

Para ejecutar la aplicación nos bastaría con ejecutar en nuestra terminal

```
$ go run main.go
```

Esto lo que hará será levantar un servidor que escuchará las peticiones que le enviemos. Para detener la ejecución, bastaría con presionar **Ctrl+C**.

Para ejecutar las pruebas nos bastaría con ejecutar en nuestra terminal

```
$ go test -v
```

Adicionalmente, contamos con una colección postman (**api.meli.test.postman_collection.json**) que contienen las peticiones que podríamos ejecutar de acuerdo al desarrollo de nuestro proyecto, el cual está ubicado en la carpeta **postman/** en el directorio raíz del proeycto. **IMPORTANTE** tener en cuenta que en la colección postman están configuradas las variables para que la url del api apunte al cloud o a nuestro entorno local de desarrollo.

---

## Consideraciones <a name = "usage"></a>

Para desarrollar la lógica del proyecto se tomaron en cuenta varios aspectos que no se detallaban de manera explícita en el planteamiento inicial:

- Los valores de una secuencia bajo la misma ruta de análisis (horizontal/vertical/oblicua), no únicos. Es decir, que si tenemos para el ejemplo horizontal [ A A A A A ], se contará como una coincidencia; esto -> [ **A A A A** A ] - [ A **A A A A** ] no aplicaría. El siguiente caso sí daría 2 coincidencias -> [ A A A A A A A A ].
- Se considerá como válida una secuencia de ADN cuya traspolación resulte en una matriz de 2 dimensiones cuya columnas tengan el mismo tamaño.
- Es posible evaluar una matriz que no necesariamente sea cuadrada NxN. Lo único es que para que esto se cumpla, alguno de los dos lados de la matriz debe resultar en un tamaño mínimo al requerido por el eneunciado, es decir, que para que se pueda cumplir la condición de evaluación del ADN al menos uno de los dos lados debe tener mínimo cuatro (4) elementos.
- En los casos previamente descritos, donde la muestra no cumpla con los criterios de evaluación, se devolverá un http 400 al usar el método /mutant/ de la api.
- Al realizar un llamado a //mutant/ se verifica si el ADN enviado es nuevo para analizarlo. En caso de que sea uno repetido, se devolverá un error http 400.
