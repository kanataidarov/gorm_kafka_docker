docker-compose down

docker image rm $(docker images | tr -s ' ' : | grep gorm_kafka_docker_app | cut -f3 -d:)

export DOCKER_DEFAULT_PLATFORM=linux/amd64 \
DB_USER=postgres \
DB_PWD=changeme \
KFK_BROKERS="broker:9092"

docker-compose up --force-recreate