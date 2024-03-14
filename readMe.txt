1. Create database in PosgreSql
postgre_db.go 
URL: connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable" 
if DB exists: DROP table delivery, items, order_table, payment
createDB.sql
2. go run "L0-main\runner\main.go"
3. go run src/publisher.go