include .env

ARGS = $(filter-out $@,$(MAKECMDGOALS))

db_status:
	goose -dir=./migrations ${DB_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" status
migrate:
	goose -dir=./migrations ${DB_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" $(ARGS)
create-migration:
	goose -dir=./migrations ${DB_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" create $(ARGS) sql
build-docker:
	docker build -t hacktiv_8_api .
run-docker:
	docker run --network postgres-16_postgres  -p 5000:5000 --env-file .env hacktiv_8_api

# This target is a workaround to allow passing arguments to Makefile targets
%:
	@: