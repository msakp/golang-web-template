build:
	@echo "[+] BUILDING"
	@go build -C cmd -o ../bin/app

run: build
	@echo "[+] RUNNING"
	@./bin/app


fmt:
	gofmt -w.


sqlc:
	@echo "[+] GENERATING SQL"
	@cd internal/infrastructure/database/sqlc;	bash generate.sh

test:




