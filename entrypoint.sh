#!/bin/bash

set -e

# testRDB作成
if [ "$GO_ENV" = "development" ]; then
  echo 'start setup test db.'
  apt-get -y install default-mysql-client
  mysql -h "${DB_HOST}" -P "${DB_PORT}" -u "${DB_USER}" -p"${DB_PASSWORD}" -e "CREATE DATABASE IF NOT EXISTS "$TEST_DB_NAME";"
  goose -env test up
  echo "created & migrated test db"
fi

# RDB migration
goose up
echo "migrated."

# 起動
if [ "$GO_ENV" = "production" ]; then
  /app/build
elif [ "$GO_ENV" = "staging" ]; then
  /app/build
elif [ "$GO_ENV" = "development" ]; then
  goose -env test up
  arelo -p '**/*.go' -p '**/*.toml' -- go run ./main.go
fi
