createmigration:
	migrate create -ext=sql -dir=sql/migrate -seq init

migrate:
	migrate -path=sql/migrate -database "postgres://postgres:pass@localhost:5432/product-recommendation?sslmode=disable" -verbose up

migratedown:
	migrate -path=sql/migrate -database "postgres://postgres:pass@localhost:5432/product-recommendation?sslmode=disable" -verbose down

.PHONY: migrate migratedown createmigration
