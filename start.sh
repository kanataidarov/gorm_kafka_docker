docker-compose down

docker image rm $(docker images | tr -s ' ' : | grep gorm_kafka_docker_app | cut -f3 -d:)

export DOCKER_DEFAULT_PLATFORM=linux/amd64 \
DB_USER=postgres \
DB_PWD=changeme \
# KFK_BROKERS="localhost:9092" # раскомментить если запускаете саму прикладу отдельно или из IDE

docker-compose up --force-recreate