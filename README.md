# search-api
Evaluacion de competencias tecnicas - API test

La API tiene el proposito de centralizar multiples busquedas de diferentes fuentes.

![Heroku](https://heroku-badge.herokuapp.com/?app=search-api-tribal.herokuapp.com/v1/search?query=Newton)

### Configuracion

Si desea trabajar con el proyecto debera instalar Go(Golang) en su equipo segun el sistema operativo que  utilice y Git.

Ya instalado las dos herramientas se puede dirigir a una carpeta y clonar el repositorio:

    git clone https://github.com/gonzalezlrjesus/search-api.git

luego de clonado:

    cd search-api && go run main.go

con los comandos de arriba ya deberia estar ejecutandose la API.

### Configuracion con Docker

Si desea puede utilizar un entorno utilizando docker, para ello debe instalar Docker primero y Git.

Ya instalados puede clonar el repositorio:
    
    git clone https://github.com/gonzalezlrjesus/search-api.git
    cd search-api

Con el siguiente comando se creara una imagen del repositorio:

    docker build --tag search-api .

luego se crea un container para finalmente comenzar a hacer pruebas localmente.
    
    docker run -it -p 8000:8000 --name api-test search-api

### Alojamiento de la API

la API se encuentra alojada en una app de HEROKU: https://search-api-tribal.herokuapp.com/

#### Version : V1

### Endpoint:

    /v1/search?query={criterio}

#### Header

    Content-Type = application/json
    
#### Parametro

    query = indica el criterio de la busqueda a realizar.

#### Ejemplo de Request: Curl

    curl -i https://search-api-tribal.herokuapp.com/v1/search?query=Newton --header 'Content-Type: application/json'

#### Ejemplo de Request: Navegador

https://search-api-tribal.herokuapp.com/v1/search?query=Newton

#### Ejemplo de Response

    {
    "query": [
        {
            "name": "All In with Cam Newton",
            "kind": "show",
            "date": "2016-06-03",
            "originapi": "TVMAZE"
        },
        {
            "name": "Defending Jacob",
            "kind": "ebook",
            "date": "2012-01-31T08:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Dig That, Zeebo Newton",
            "kind": "feature-movie",
            "date": "2018-06-28T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Don't Stop Believin'",
            "kind": "song",
            "date": "2010-06-01T12:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Juice Newton: Every Road Leads Back to You",
            "kind": "feature-movie",
            "date": "2001-09-10T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Les Norton",
            "kind": "show",
            "date": "2019-08-04",
            "originapi": "TVMAZE"
        },
        {
            "name": "Master Keaton",
            "kind": "show",
            "date": "1998-10-06",
            "originapi": "TVMAZE"
        },
        {
            "name": "Midsummer In Newtown",
            "kind": "feature-movie",
            "date": "2017-01-27T08:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Nesten voksen",
            "kind": "show",
            "date": "2018-01-06",
            "originapi": "TVMAZE"
        },
        {
            "name": "Nestor Burma",
            "kind": "show",
            "date": "1991-09-29",
            "originapi": "TVMAZE"
        },
        {
            "name": "Newborn Russia",
            "kind": "show",
            "date": "2014-03-26",
            "originapi": "TVMAZE"
        },
        {
            "name": "Newton",
            "kind": "ebook",
            "date": "2008-04-15T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Newton's Apple",
            "kind": "show",
            "date": "1983-10-15",
            "originapi": "TVMAZE"
        },
        {
            "name": "Newton's Law",
            "kind": "show",
            "date": "2017-02-09",
            "originapi": "TVMAZE"
        },
        {
            "name": "Newton,Dave R.",
            "kind": "person",
            "date": "2000-03-20",
            "originapi": "CRCIND"
        },
        {
            "name": "Newton,Mario B.",
            "kind": "person",
            "date": "2004-11-17",
            "originapi": "CRCIND"
        },
        {
            "name": "Newtown",
            "kind": "feature-movie",
            "date": "2017-01-07T08:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Olivia Newton-John: Hopelessly Devoted To You",
            "kind": "show",
            "date": "2018-05-13",
            "originapi": "TVMAZE"
        },
        {
            "name": "Peston",
            "kind": "show",
            "date": "2016-05-08",
            "originapi": "TVMAZE"
        },
        {
            "name": "Philosophiae Naturalis Principia Mathematica",
            "kind": "ebook",
            "date": "1543-01-11T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Streamline",
            "kind": "song",
            "date": "2006-07-10T12:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Summer of '64",
            "kind": "song",
            "date": "2010-04-16T12:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "The Newton Boys",
            "kind": "feature-movie",
            "date": "2002-05-21T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "The minute mathematician: or, the free-thinker no just-thinker. Set forth in a second letter to the author of The analyst; containing a defence of Sir Isaac Newton and the British mathematicians, ... By Philalethes Cantabrigiensis",
            "kind": "ebook",
            "date": "1735-01-01T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "We're All Alone",
            "kind": "song",
            "date": "1997-01-01T12:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "Yes Please",
            "kind": "ebook",
            "date": "2014-10-28T07:00:00Z",
            "originapi": "APPLE"
        },
        {
            "name": "You're the One That I Want",
            "kind": "song",
            "date": "1978-04-14T12:00:00Z",
            "originapi": "APPLE"
        }
    ],
    "status": 200
    }

##### Estructura del proyecto
    .
    ├── bin
    │   └── main
    ├── config
    │   └── config.go
    ├── controllers
    │   └── search-controller.go
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── LICENSE
    ├── main.go
    ├── Makefile
    ├── models
    │   ├── APIApple.go
    │   ├── APICrcind.go
    │   ├── APITvMaze.go
    │   └── response.go
    ├── README.md
    ├── routes
    │   └── routes.go
    ├── services
    │   ├── crcind.go
    │   ├── itunes-apple.go
    │   └── tvmaze.go
    └── utils
        └── utils.go

    7 directorios, 19 archivos