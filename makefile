build:
	@echo "[+] BUILDING"
	@go build -C cmd -o ../bin/app

run: swag build
	@echo "[+] RUNNING"
	@./bin/app


fmt:
	gofmt -w .


sqlc:
	@echo "[+] GENERATING SQL"
	@cd internal/infrastructure/database/sqlc;	bash generate.sh

test:



swag:
	@echo "[+] GENERATING SWAGGER DOC"
	@swag fmt
	@swag init -g cmd/main.go


