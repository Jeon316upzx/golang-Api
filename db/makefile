mysqldb:
	docker run --name mysqldb -p 5432:3306 -e MYSQL_ROOT_PASSWORD=bankapi18  -d mysql:latest

createdb:
	docker exec -it mysqldb mysql -u root -pbankapi18 -e "create database bankdb"

dropdb:
	docker exec -it mysqldb mysql -u root -pbankapi18 -e "drop database bankdb"

usedb:
	docker exec -it mysqldb mysql -u root -pbankapi18 -e "use bankdb"

stopdb:
	docker stop mysqldb

removedb:
	docker rm mysqldb -f

startdb:
	docker start mysqldb

migrateup:
	migrate -path migrations -database "mysql://root:bankapi18@tcp(localhost:5432)/bankdb" -verbose up

migratedown:
	migrate -path migrations -database "mysql://root:bankapi18@tcp(localhost:5432)/bankdb" -verbose down

.PHONY: mysqldb createdb dropdb migrateup migratedown
