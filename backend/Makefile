.PHONY: all

generate-abi:
	abigen --abi abi.json --pkg wager --out pkg/wager/wager.go

docker-db:
	docker run --name pg --rm -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=admin -e POSTGRES_DB=bets -p 5432:5432 postgres:13

db-connect:
	docker exec -it backend-db-1 /usr/bin/psql -h backend-db-1 -p 5432 -U admin -d bets

migrate-up:
	docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database 'postgresql://admin:admin@localhost/bets?sslmode=disable' up