# Load environment variables from .env file
include .env

MIGRATION_DIR = ./domain/migration
DB_DSN = mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp\(${DB_HOST}:${DB_PORT}\)/${DB_DATABASE}

.PHONY: help

help:
	@echo "Available targets:"
	@echo " migrate-create	: Create new migration"
	@echo " migrate-up	: Apply database migrations"
	@echo " migrate-down	: Drop all migration's table"
	@echo " migrate-drop	: Drop entire schema's tables"
	@echo " migrate-version: Stdout current migration version"
	@echo " up		: Setup and run the app"
	@echo " cert		: Generate SSL Certificate"

migrate-create:
	migrate create -ext sql -dir ${MIGRATION_DIR} -seq create_$(name)_table

migrate-up:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} up

migrate-down:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} down

migrate-drop:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} drop

migrate-version:
	migrate -database ${DB_DSN} -path ${MIGRATION_DIR} version

up:
	go run main.go

cert:
	openssl genrsa -out server.key 2048 && openssl ecparam -genkey -name secp384r1 -out server.key && openssl req -new -x509 -sha256 -key server.key -out server.crt -days 3650
