sqlc:
	sqlc generate

start:
	docker start my-postgres

stop:
	docker stop my-postgres

build:
	docker run --name my-postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

connect:
	docker exec -it my-postgres psql -U root

test:
	go test -v --cover ./...
 
migration:
	psql -d postgresql://root@localhost/postgres -a -f ./db/schema/shema.sql

	
