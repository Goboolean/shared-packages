proto-generate:
	protoc \
		-I api/proto \
		--go_out=pkg/kafka \
		--go_opt=paths=source_relative \
		stockaggs.proto

sqlc-generate:
	sqlc generate -f api/sql/sqlc.yml