#!/bin/bash

set -e

# create Database for testing
echo 'start setup test db.'
apt-get -y install default-mysql-client
mysql -h "${DB_HOST}" -P "${DB_PORT}" -u "${DB_USER}" -p"${DB_PASSWORD}" -e "CREATE DATABASE IF NOT EXISTS "$TEST_DB_NAME";"
goose -env test up
echo "created & migrated test db"

# RDB migration
goose up
goose -env test up
echo "migrated."

# 起動
arelo -p '**/*.go' -p '**/*.toml' -- go run ./main.go