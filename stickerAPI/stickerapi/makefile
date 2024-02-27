MYSQL_USER ?= user
MYSQL_PASSWORD ?= password
MYSQL_ADDRESS ?= 127.0.0.1:9090
MYSQL_DATABASE ?= dummy

export PATH   := $(PWD)/bin:$(PATH)

export SHELL  := bash

export OSTYPE := $(shell uname -s)

seed-db: $(MIGRATE) 
	@ echo "Seeding the database..."
	@ docker-compose -f docker-compose.yml up -d mysql
	@ sleep 5  # Give MySQL some time to start
	@ docker exec -i $(shell docker-compose ps -q mysql) mysql -u $(MYSQL_USER) -p$(MYSQL_PASSWORD) $(MYSQL_DATABASE) < ./sticker.sql
	@ echo "Database seeding completed."

up: 	
	@ docker-compose up             

down: 	
	@ docker-compose down   

migrate up:	
	@ migrate -path db/migration -database "mysql://$$DB_USER:$$DB_PASS@tcp($$DB_HOST:$$DB_PORT)/$$DB_NAME?charset=utf8mb4&parseTime=True&loc=Local" -verbose up

migrate down:	
	@ migrate -path db/migration -database "mysql://$$DB_USER:$$DB_PASS@tcp($$DB_HOST:$$DB_PORT)/$$DB_NAME?charset=utf8mb4&parseTime=True&loc=Local" -verbose down         
