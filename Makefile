# Konfigurasi umum
DB_URL=postgresql://postgres:richard10@localhost:5432/user_service_gol?sslmode=disable
MIGRATE_CMD=migrate -path db/migrations -database "$(DB_URL)"

# Menjalankan migrasi ke versi terbaru
migrate-up:
	$(MIGRATE_CMD) -verbose up

# Rollback satu langkah
migrate-down:
	$(MIGRATE_CMD) -verbose down 1

# Membuat file migrasi baru
migrate-create:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir internal/db/migrations -seq $$name

# Menampilkan status migrasi
migrate-status:
	$(MIGRATE_CMD) version

# Menjalankan aplikasi
run:
	go run ./cmd/server/main.go
