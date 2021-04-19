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
* _Acompañando al planteamiento previo, para el almacenamiento en la BD se plantea el uso de colas para mantener las peticiones sin saturar el servidor de BD, esto debido a que en el enunciado_
  _no se evidencia que esta creación sea parte de la operación crítica del ejercicio ya que es un servicio de consulta. Al tener una cola de peticiones a lka BD,_
  _estas se pueden tratar mediante reglas: 1. Definir 5 reintentos, uno cada 5 minutos. 2. Cuando la petición ya cumple los reintentos definidos se guarda en un repositorio._
  _3. Crear un proceso periodico para realizar un cargue masivo de las peticiones fallidas. 4. Si la petición vuelve a fallar se descarta completamente._
  

Teniendo en cuenta estos planteamientos y las definiciones, a continuacion se relacionan las definiciones técnias **implementadas**.
* _Se procuran retornos temprano para mejorar el tiempo de respuesta de la aplicacion_
* _Se establecen 4 tipos de busqueda de secuencia: Horizontal a la derecha, Vertical hacia abajo, Diagonal derecha y abajo, Diagonal derecha y arriba_
* _Como las cadenas son palindromes solo se busca en un sentido para no generar conteos erroneos._
* _Una vez recibida la peticion, lo primero es validar la cadena recibida según las definiciones estipuladas._
* _PAra analizar un dna, primero se consulta en la BD si este ya existe y se retorna la respuesta almacenada._
* _Solo se envian peticiones de creacion a la BD cuando este no existe._
* _La peticion de creación a la base de datos se realiza de manera asícrona._
* _Existen otras consideraciones técnicas que se documentan en el cógido fuente ya que es mas entendible hacerlo ahí._



## Ejecutando las pruebas ⚙️

Se crearon pruebas unitarias usando las librerias estandar de la herramienta. Estas se implementaron sobre los 3 paquetes que contienen las funciones principales
de la aplicación.
* Mutant -- [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/mutant_coverage.html)
* Stats  -- [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/stats_coverage.html)
* Validate --  [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/validate_coverage.html)



## Despliegue 📦

El despliegue de la aplicación se realizó en **Heroku**. La URL del servicion es:
```
[URL - API](https://mutantcha.herokuapp.com/)
```

## Construido con 🛠️

Estas son las herrameientas y software usado para la construccion

* Lenguaje de programacion GO
* Visual Studio Code - Editor
* Mongo DB - Motor de BD NoSQL
* Postman - conexiones REST
* SourceTree - conexion con github
* Github - Reposotorio de versionamiento
* Heroku - Plataforma cloud de PAAS


## Versionado 📌

Se usó Github como repositorio, este se conecta desde windows a traves de la herramienta SourceTree. Se pueden ver los tags creados para este proyecto.

## Autores ✒️

* **Juan Sebastian Agudelo** - *Trabajo Completo* 
