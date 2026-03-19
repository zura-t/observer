new_migration:
	migrate create -ext sql -dir db/migrations -seq $(name)

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5423/observer?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5423/observer?sslmode=disable" -verbose up 1
	
migratedown:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5423/observer?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5423/observer?sslmode=disable" -verbose down 1

server:
	go run cmd/app/main.go

.PHONY: new_migration migrateup migrateup1 migratedown migratedown1 server