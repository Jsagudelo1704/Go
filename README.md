# Mutant Challenge

El contenido de este repositorio incluye el codigo fuente como solucion al desafio Mutante para Mercadolibre. Este proyecto se trata de una API REST
hecha en GO para validar una cadena de ADN y retornar si esta es de un Humano o Mutante.

## Comenzando 🚀

Para la realización de este ejercicion se tuvo en cuenta las definiciones del enunciado recibido, se plantearon supuestos en los escenarios no definidos y se 
realizaron planteamientos para las implementaciones no realizadas.

### Definiciones

* _Se recibe un objeto JSON con un array de string_
* _Este Array se descompone en una matriz simetrica (N*N)_
* _La cadena solo puede contener las letras (A,C,G,T)_
* _la secuencia a buscar es de 4 caracteres iguales_
* _Con mas de una secuencia encontrada se considera Mutante_
* _Metodo POST para ingresar la cadena a validar (/mutant)_
* _Metodo GET para consultar las estadísticas (/stats)_
* _Dondigos de retorno definidos para las respuestas de los servicios_


### Supuestos 

* _No se encontraron supuestos en el enunciado._


### Planteamientos 

* _Para manejar las fluctuaciones agresivas de trafico y los picos de concurrencia, se plantea la implementación de un modelo de escalabilidad Horizantal,_ 
  _ya sea replicando los contenedores o servidores según el esquema de infraestructura presente._
* _Debido a que este escalamiento es costoso, se plantea el uso de un servicio en la nube que nos permita escalasr de forma inmediata durante el pico y de_
  _esta forma solo pagar por el consumo generado._
* _Acompañando el planteamiento previo, para el almacenamiento en la BD se plantea el uso de colas para almacenar las peticiones, esto debido a que en el enunciado_
  _no se evidencia que esta creación sea parte de la operación crítica del ejercicio ya que es un servicio de consulta. Al tener una cola de peticiones a lka BD,_
  _estas se pueden tratar mediante reglas: 1. Definir 5 reintentos, uno cada 5 minutos. 2. Cuando la petición ya cumple los reintentos definidos se guarda en un repositorio._
  _3. Crear un proceso periodico para realizar un cargue masivo de las peticiones fallidas. 4. Si la petición vuelve a fallar se descarta completamente._
  

Teniendo en cuenta estos planteamientos y las definiciones, a continuacion se relacionas las definiciones técnias implementadas.
* _Para manejar las fluctuaciones agresivas de trafico y los picos de concurrencia, se plantea la implementación de un modelo de escalabilidad Horizantal,_ 
  _ya sea replicando los contenedores o servidores según el esquema de infraestructura presente._
* _Debido a que este escalamiento es costoso, se plantea el uso de un servicio en la nube que nos permita escalasr de forma inmediata durante el pico y de_
  _esta forma solo pagar por el consumo generado._
* _Acompañando el planteamiento previo, para el almacenamiento en la BD se plantea el uso de colas para almacenar las peticiones, esto debido a que en el enunciado_
  _no se evidencia que esta creación sea parte de la operación crítica del ejercicio ya que es un servicio de consulta. Al tener una cola de peticiones a lka BD,_
  _estas se pueden tratar mediante reglas: 1. Definir 5 reintentos, uno cada 5 minutos. 2. Cuando la petición ya cumple los reintentos definidos se guarda en un repositorio._
  _3. Crear un proceso periodico para realizar un cargue masivo de las peticiones fallidas. 4. Si la petición vuelve a fallar se descarta completamente._



_Para este ejercicio se parte de las siguientes definiciones:

_Para generar una copia de este proyecto se puede descargar el archivo .zip de los componentes y pegarlos en el WorkSpace de GO._

Mira **Despliegue** para conocer como desplegar el proyecto.





## Ejecutando las pruebas ⚙️

_Explica como ejecutar las pruebas automatizadas para este sistema_



## Despliegue 📦

_Agrega notas adicionales sobre como hacer deploy_

## Construido con 🛠️

_Menciona las herramientas que utilizaste para crear tu proyecto_

* [Dropwizard](http://www.dropwizard.io/1.0.2/docs/) - El framework web usado
* [Maven](https://maven.apache.org/) - Manejador de dependencias
* [ROME](https://rometools.github.io/rome/) - Usado para generar RSS


## Versionado 📌

Usamos [SemVer](http://semver.org/) para el versionado. Para todas las versiones disponibles, mira los [tags en este repositorio](https://github.com/tu/proyecto/tags).

## Autores ✒️

_Menciona a todos aquellos que ayudaron a levantar el proyecto desde sus inicios_

* **Andrés Villanueva** - *Trabajo Inicial* - [villanuevand](https://github.com/villanuevand)
* **Fulanito Detal** - *Documentación* - [fulanitodetal](#fulanito-de-tal)

También puedes mirar la lista de todos los [contribuyentes](https://github.com/your/project/contributors) quíenes han participado en este proyecto. 


## Expresiones de Gratitud 🎁

* Comenta a otros sobre este proyecto 📢
* Invita una cerveza 🍺 o un café ☕ a alguien del equipo. 
* Da las gracias públicamente 🤓.
* etc.



