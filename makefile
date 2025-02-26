up:
	@echo "[+] CREATING CONTAINERS"
	@sudo docker-compose -f dev.yml up -d

down:
	@echo "[+] REMOVING CONTAINERS"
	@sudo docker-compose -f dev.yml down

build: fmt sqlc swag
	@echo "[+] BUILDING"
	@go build -C cmd -o ../bin/server

server: build
	@echo "[+] RUNNING"
	@./bin/server


fmt:
	gofmt -w .


sqlc:
	@echo "[+] GENERATING SQL"
	@cd pkg/sqlc;	bash generate.sh


test:
	@echo "[+] TESTING"
	@go test ./tests/e2e -v -count=1 -run Main


swag:
	@echo "[+] GENERATING SWAGGER DOC"
	@pkg/swag/swag fmt
	@pkg/swag/swag init -g cmd/main.go


