# Penggunaan
- ganti file `example.env` menjadi `.env` sesuikan isinya
- `migrate create -ext sql -dir migrations create_user_tabel` untuk membuat file migrate sql
- `migrate -path ./migrations -database "postgres://user:password@localhost:5432/pos_kasir?sslmode=disable" up` untuk melakukan migrasi
