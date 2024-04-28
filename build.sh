#!/bin/bash

DB_NAME="storiesservice"
DB_USER="postgres"

# Команда для проверки существования базы данных
if psql -U "$DB_USER" -lqt | cut -d \| -f 1 | grep -qw "$DB_NAME"; then
    echo "База данных $DB_NAME существует. Все ОК"
else
    echo "База данных $DB_NAME не найдена."
	createdb $DB_NAME
	psql -U $DB_USER -d $DB_NAME -f create_tables.sql
fi

export CGO_ENABLED=1
arch=$(uname -m)
if [ "$arch" = "arm64" ]; then
  echo "Архитектура процессора - ARM64."
  go build -trimpath -tags release -buildmode exe -ldflags "-s -w -linkmode external" -o server.out cmd/server/server.go cmd/server/globals.go \
	cmd/server/handlers.go cmd/server/i18n.go cmd/server/templates.go
else
  echo "Архитектура процессора не является ARM64, текущая архитектура: $arch."
  go build -trimpath -tags release -buildmode exe -ldflags "-s -w -linkmode external -extldflags -static" -o server.out cmd/server/server.go cmd/server/globals.go \
	cmd/server/handlers.go cmd/server/i18n.go cmd/server/templates.go
fi

