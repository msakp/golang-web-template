# Golang clean-arch Web Template


### [+] See [CHANGELOG](CHANGELOG.md)
### [+] See [CONTRIBUTING](CONTRIBUTING.md)



# Running

### Production Mode
```bash
# inside .env
PROD=True
...
```
```bash
sudo docker-compose up --build
```
### Dev Mode
```bash
# inside .env
PROD=False
...
```

### (Dev Mode) with MakeFile
```bash
make up # for creating dependency containers (postgres, etc)
make down # for removing  dependency containers (postgres, etc)
make server # formats source code, regenerates all code generation, compiles and runs server
make test # runs e2e tests (empty tests)
make sqlc # for sqlc code generation
make swag # for openapi spec generation 
```

### (Dev Mode) with TaskFile
```bash
task -l # list all tasks with their descpitions
# but all same here
task up
task down
task server
task test
task sqlc
task swag
```
