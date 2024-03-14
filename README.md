# Nats_Streaming_Server
1. Create database in PosgreSql
postgre_db.go 
URL: connStr := "postgres://postgres:admin@localhost:5432/postgres?sslmode=disable" 
if DB exists: DROP table delivery, items, order_table, payment
createDB.sql
2. go run "L0-main\runner\main.go"
3. go run src/publisher.go

Result:

![image](https://github.com/ernsterfickfacker/Nats_Streaming_Server/assets/93219479/85250367-2992-4559-82e1-a184e66950ed)




