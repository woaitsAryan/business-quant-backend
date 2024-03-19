#!/bin/bash

# Generate random values
JWT_SECRET=$(openssl rand -hex 32)
DB_PASSWORD=$(openssl rand -hex 32)
REDIS_PASSWORD=$(openssl rand -hex 32)

echo "Please enter the project name:"
read PROJECT_NAME

mkdir logs
touch logs/debug.log
touch logs/error.log
touch logs/fatal.log
touch logs/info.log
touch logs/panic.log
touch logs/warn.log

# Copy .env.template to .env and replace the placeholders
cp .env.template .env
sed -i "s|your_jwt_secret|$JWT_SECRET|g" .env
sed -i "s|postgres_password|$DB_PASSWORD|g" .env
sed -i "s|redis_password|$REDIS_PASSWORD|g" .env

# Run the Docker commands
docker run --name $PROJECT_NAME-postgres -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=$DB_PASSWORD -p 5432:5432 -v ./init.sql:/docker-entrypoint-initdb.d/init.sql -d postgres
docker run --name $PROJECT_NAME-redis -p 6379:6379 -d redis redis-server --requirepass $REDIS_PASSWORD