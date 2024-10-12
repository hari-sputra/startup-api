# Memuat file .env jika ada
ifeq (,$(wildcard .env))
    $(error ".env file not found!")
endif

include .env
export $(shell sed 's/=.*//' .env)

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down

migrate-force:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations force 1