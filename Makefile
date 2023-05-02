build-kafka:
	docker compose -f ./build/kafka/docker-compose.yml up --build -d
clean-kafka:
	docker compose -f ./build/kafka/docker-compose.yml down

sqlc-generate:
	sqlc generate -f api/sql/sqlc.yml
