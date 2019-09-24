DATABASE_URL = "postgres://testuser:testpass@localhost:5432/testuser?sslmode=disable"

.PHONY: test-it
test-it:
	docker run --name test-postgres -p 5432:5432 -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=testpass --rm -d postgres
	until pg_isready -h localhost -p 5432 -U testuser; do sleep 0.2; done ;
	migrate -database $(DATABASE_URL) -path db-migrations up || true # Ignore error: EOF error
	go test test/*.go -v -count=1
	# docker rm -f test-postgres

.PHONY: up
up:
	docker run --name test-postgres -p 5432:5432 -e POSTGRES_USER=testuser -e POSTGRES_PASSWORD=testpass --rm -d postgres
	until pg_isready -h localhost -p 5432 -U testuser; do sleep 0.2; done ;
	migrate -database $(DATABASE_URL) -path db-migrations up || true # Ignore error: EOF error
	DATABASE_URL=$(DATABASE_URL) PORT=8080 go run cmd/highscore-server/*.go


.PHONY: down
down:
	docker rm -f test-postgres || true

build: ## build application binaries
	go build -o highscore-server -race *.go 