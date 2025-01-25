1. **Create the migrations.** <br>
   migrate create -ext sql -dir cmd/migrate/migrations -seq create_users_table
2. **Run the migrations.** <br>
    migrate -path cmd/migrate/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose up
3. **Rollback the migrations.** <br>
    migrate -path cmd/migrate/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose down
4. **Create the schema**. <br>
    migrate -path cmd/migrate/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose create -ext sql -dir cmd/migrate/migrations -seq create_users_table
5. **Drop the schema.** <br>
    migrate -path cmd/migrate/migrations -database "postgres://postgres:password@localhost:5432/postgres?sslmode=disable" -verbose drop
