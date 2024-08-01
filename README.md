# Golang REST API with GORM + Kafka + Docker

A microservice in Go that will accept messages via HTTP API, save them in PostgreSQL using GORM, and then send them to Kafka for further processing. \
Processed messages are read by another Kafka consumer, and the message is marked as processed in the database. \
The service also provides an API to retrieve statistics on processed messages. \
The solution is packaged in a Docker deployment: containers will be created for PostgreSQL, Apache Kafka, and the application itself.

## Local Setup

### Prerequisites
- Golang version 1.22+. [Installation instructions](https://go.dev/doc/install)
- Docker. [Installation instructions](https://docs.docker.com/get-docker/)
- docker-compose. [Installation instructions](https://docs.docker.com/compose/install/standalone/)

### Launch
- For `Windows`, in CMD, run `./start.bat` from the project root.
- For `Linux / MacOS`, in the terminal, run `./start.sh` from the project root.

After that API will be available at http://localhost:44049. \
Further it will referred as `<Base URL>`.

### Sending with HTTP API
Send HTTP POST request `<Base URL>/applications` with payload: 
```json
{
    "name": "Kanat Aidarov",
    "email": "kanataidarov@yahoo.com",
    "position": "Junior Go Developer"
}
```

Example of sending with `curl`:
```shell
curl --location 'https://localhost:44049/applications' \                                                                      git:main*
--request POST \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "Канат Айдаров",
    "email": "kanataidarov@yahoo.com",
    "position": "Junior Go Developer"
}'

{"assignment":{"position":"Junior Go Developer","version":3},"message":"Application processed successfully"}
```

### Retrieving statistics 
Send HTTP GET request `<Base URL>/application`. \

xample of sending with `curl`:
```shell
url --location 'https://localhost:44049/applications' \                                                                      git:main*
--request GET --header 'Content-Type: application/json'

[{"name":"Канат Айдаров","email":"kanataidarov@yahoo.com","position":"Junior Go Developer","processed":true},{"name":"Иван Федеров","email":"ivanf@yahoo.com","position":"Junior Go Developer","processed":true}]
```
