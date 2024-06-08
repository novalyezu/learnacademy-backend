DATABASE_URL = postgresql://postgres:123qweasd@127.0.0.1:5433/db_learn_academy?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DATABASE_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DATABASE_URL)" -verbose down

.PHONY: migrateup migratedown
