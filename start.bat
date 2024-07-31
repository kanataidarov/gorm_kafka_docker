docker-compose down

for /f "tokens=3" %%i in ('docker images ^| findstr "gorm_kafka_docker_app"') do docker image rm -f %%i

set DOCKER_DEFAULT_PLATFORM=linux/amd64
set DB_USER=postgres
set DB_PWD=changeme
REM след. раскомментить если запускаете саму прикладу отдельно или из IDE
REM set KFK_BROKERS=localhost:9092

docker-compose up --force-recreate