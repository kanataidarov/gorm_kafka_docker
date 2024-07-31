docker-compose down

export DOCKER_DEFAULT_PLATFORM=linux/amd64 \
DB_USER=postgres \
DB_PWD=changeme

docker-compose up --force-recreate