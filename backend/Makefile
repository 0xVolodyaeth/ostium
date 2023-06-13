.PHONY: all

generate-abi:
	abigen --abi abi.json --pkg wager --out pkg/wager/wager.go

docker-db:
	docker run --name pg --rm -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=bets -p 5432:5432 postgres:13

db-connect:
	docker exec -it pg /usr/bin/psql -h localhost -p 5432 -U admin -d bets

migrate-up:
	migrate -database "postgresql://admin:admin@localhost/bets?sslmode=disable" -path migrations up