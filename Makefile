RUN:
	go run ./cmd/api

CREATE_ANIMALS_TABLE_MIGRATION:
	migrate create -seq -ext=.sql -dir=./migrations create_animals_table

ADD_ANIMALS_CHECK_CONSTRAINTS_MIGRATION:
	migrate create -seq -ext=.sql -dir=./migrations add_animals_check_constraints

EXECUTE_UP_MIGRATIONS:
	migrate -path=./migrations -database=$$DATABASE_DSN up

EXECUTE_DOWN_MIGRATIONS:
	migrate -path=./migrations -database=$$DATABASE_DSN down

ADD_ANIMALS_INDEXES_MIGRATIONS:
	migrate create -seq -ext .sql -dir=./migrations add_animals_indexes