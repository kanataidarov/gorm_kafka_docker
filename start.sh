docker-compose down

export DOCKER_DEFAULT_PLATFORM=linux/amd64 \
DB_USER=postgres \
DB_PWD=changeme \
BROKERS="broker:9092"

docker-compose up --force-recreate