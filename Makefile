PARSER_NAME=parser
SERVER_NAME=server
MIGRATE_NAME=migrate

build:
	go build -o ${PARSER_NAME} cmd/${PARSER_NAME}/main.go
	go build -o ${SERVER_NAME} cmd/${SERVER_NAME}/main.go
	go build -o ${MIGRATE_NAME} cmd/${MIGRATE_NAME}/main.go

run_parser:
	go run cmd/${PARSER_NAME}/main.go

run_migrate:
	go run cmd/${MIGRATE_NAME}/main.go

run_server:
	go run cmd/${SERVER_NAME}/main.go

clean:
	go clean
	rm ${MIGRATE_NAME}
	rm ${SERVER_NAME}
	rm ${PARSER_NAME}