protobuf-stock-generate:
	protoc \
		-I api/proto \
		--go_out=pkg/broker \
		--go_opt=paths=source_relative \
		stockaggs.proto

protobuf-event-generate:
	protoc \
		-I api/proto \
		--go_out=pkg/kafka \
		--go_opt=paths=source_relative \
		event.proto

build-kafka:
	docker compose -f ./build/kafka/docker-compose.yml up --build -d
clean-kafka:
	docker compose -f ./build/kafka/docker-compose.yml down

sqlc-generate:
	sqlc generate -f api/sql/sqlc.yml

test-pkg:
	docker compose -f ./

test-app:
	if docker-compose -f ./build/test/docker-compose.yml up --build ; then \
		docker-compose -f ./build/test/docker-compose.yml down; \
		cd ../; \
		rm -r test; \
		exit 1; \
	else \
		docker-compose -f ./build/test/docker-compose.yml down; \
		cp -r ./* ../; \
		cd ../; \
		rm -r test; \
		exit 0; \
	fi