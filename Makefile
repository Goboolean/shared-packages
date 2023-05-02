build-kafka:
	docker compose -f ./build/kafka/docker-compose.yml up --build -d
clean-kafka:
	docker compose -f ./build/kafka/docker-compose.yml down