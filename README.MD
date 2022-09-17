# Go Todo App

This application is part of my developer portfolio, it is not intended to be used in production.
If you still want to do so, that you are free to do it on your own risk.

## Installation

At minimum you are required to have Docker and Docker Compose installed on your system.

To start the application just open a terminal and type in the following command:

```bash
docker-compose up -d 
```

## Usage

After starting up the project, point your REST client to the application URL

### Get all todos

```bash  
curl -X GET http://localhost:8080/api/v1/todo
```

### Create todo

```bash  
curl -X POST -H'Content-Type:application/json' -d'{"task": "My task goes here"}' http://localhost:8080/api/v1/todo
```

### Get a single todo item

```bash  
curl -X GET http://localhost:8080/api/v1/todo/d6a12b21-6c41-448d-87bc-f94ea8451994
```

### Update todo

```bash  
curl -X PUT -H'Content-Type:application/json' -d'{"task":"My updated task", "completed", true}' http://localhost:8080/api/v1/todo/d6a12b21-6c41-448d-87bc-f94ea8451994
```

### Delete todo

```bash  
curl -X DELETE http://localhost:8080/api/v1/todo/d6a12b21-6c41-448d-87bc-f94ea8451994
```

## Third-Party libraries

This project uses the following third-party libraries:

* GIN Framework [go to page](https://github.com/gin-gonic/gin)
* GoDotEnv [go to page](https://github.com/joho/godotenv)
* UUID [go to page](https://github.com/google/uuid)
