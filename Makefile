run-local:
	. ./local_db_env_vars_init.sh && go run cmd/main.go
run:
	go run cmd/main.go
test:
	go test -v ./...