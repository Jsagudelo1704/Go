# Mutant Challenge

El contenido de este repositorio incluye el codigo fuente como solucion al desafio Mutante para Mercadolibre. Este proyecto se trata de una API REST
hecha en GO para validar una cadena de ADN y retornar si esta es de un Humano o Mutante.

## Comenzando 馃殌

Para la realizaci贸n de este ejercicion se tuvo en cuenta las definiciones del enunciado recibido, se plantearon supuestos en los escenarios no definidos y se 
realizaron planteamientos para las implementaciones no realizadas.

### Definiciones

* _Se recibe un objeto JSON con un array de string_
* _Este Array se descompone en una matriz simetrica (N*N)_
* _La cadena solo puede contener las letras (A,C,G,T)_
* _la secuencia a buscar es de 4 caracteres iguales_
* _Con mas de una secuencia encontrada se considera Mutante_
* _Metodo POST para ingresar la cadena a validar (/mutant)_
* _Metodo GET para consultar las estad铆sticas (/stats)_
* _Dondigos de retorno definidos para las respuestas de los servicios_


### Supuestos 

* _No se encontraron supuestos en el enunciado._


### Planteamientos 

* _Para manejar las fluctuaciones agresivas de trafico y los picos de concurrencia, se plantea la implementaci贸n de un modelo de escalabilidad Horizantal,_ 
  _ya sea replicando los contenedores o servidores seg煤n el esquema de infraestructura presente._
* _Debido a que este escalamiento es costoso, se plantea el uso de un servicio en la nube que nos permita escalasr de forma inmediata durante el pico y de_
  _esta forma solo pagar por el consumo generado._
* _Acompa帽ando al planteamiento previo, para el almacenamiento en la BD se plantea el uso de colas para mantener las peticiones sin saturar el servidor de BD, esto debido a que en el enunciado_
  _no se evidencia que esta creaci贸n sea parte de la operaci贸n cr铆tica del ejercicio ya que es un servicio de consulta. Al tener una cola de peticiones a lka BD,_
  _estas se pueden tratar mediante reglas: 1. Definir 5 reintentos, uno cada 5 minutos. 2. Cuando la petici贸n ya cumple los reintentos definidos se guarda en un repositorio._
  _3. Crear un proceso periodico para realizar un cargue masivo de las peticiones fallidas. 4. Si la petici贸n vuelve a fallar se descarta completamente._
  

Teniendo en cuenta estos planteamientos y las definiciones, a continuacion se relacionan las definiciones t茅cnias **implementadas**.
* _Se procuran retornos temprano para mejorar el tiempo de respuesta de la aplicacion_
* _Se establecen 4 tipos de busqueda de secuencia: Horizontal a la derecha, Vertical hacia abajo, Diagonal derecha y abajo, Diagonal derecha y arriba_
* _Como las cadenas son palindromes solo se busca en un sentido para no generar conteos erroneos._
* _Una vez recibida la peticion, lo primero es validar la cadena recibida seg煤n las definiciones estipuladas._
* _PAra analizar un dna, primero se consulta en la BD si este ya existe y se retorna la respuesta almacenada._
* _Solo se envian peticiones de creacion a la BD cuando este no existe._
* _La peticion de creaci贸n a la base de datos se realiza de manera as铆crona._
* _Existen otras consideraciones t茅cnicas que se documentan en el c贸gido fuente ya que es mas entendible hacerlo ah铆._



## Ejecutando las pruebas 鈿欙笍

Se crearon pruebas unitarias usando las librerias estandar de la herramienta. Estas se implementaron sobre los 3 paquetes que contienen las funciones principales
de la aplicaci贸n.
* Mutant -- [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/mutant_coverage.html)
* Stats  -- [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/stats_coverage.html)
* Validate --  [Coverage](https://github.com/Jsagudelo1704/Go/blob/main/docs/validate_coverage.html)



## Despliegue 馃摝

El despliegue de la aplicaci贸n se realiz贸 en **Heroku**. La URL del servicion es:
```
[URL - API](https://mutantcha.herokuapp.com/)
```

## Construido con 馃洜锔?

Estas son las herrameientas y software usado para la construccion

* Lenguaje de programacion GO
* Visual Studio Code - Editor
* Mongo DB - Motor de BD NoSQL
* Postman - conexiones REST
* SourceTree - conexion con github
* Github - Reposotorio de versionamiento
* Heroku - Plataforma cloud de PAAS


## Versionado 馃搶

Se us贸 Github como repositorio, este se conecta desde windows a traves de la herramienta SourceTree. Se pueden ver los tags creados para este proyecto.

## Autores 鉁掞笍

* **Juan Sebastian Agudelo** - *Trabajo Completo* 
