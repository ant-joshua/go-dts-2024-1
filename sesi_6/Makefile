include .env

ARGS = $(filter-out $@,$(MAKECMDGOALS))

db_status:
	goose -dir=./migrations ${DB_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" status
migrate:
	goose -dir=./migrations ${DB_DRIVER} "host=${DB_HOST} port=${DB_PORT} user=${DB_USER} password=${DB_PASS} dbname=${DB_NAME}" $(ARGS)

# This target is a workaround to allow passing arguments to Makefile targets
%:
	@: