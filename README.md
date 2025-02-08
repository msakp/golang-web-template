# Golang clean-arch Web Template

## **Tech**
**Web Framework**: Fiber  
**Postgersql Driver**: Sqlc  
**Config Engine:** Viper  
**Docs**: swaggo/swag  

## **Running**
### [DOCS] READ
```shell
# launch backend and lookup:)
{SERVER_ADDRESS}/api/v1/docs
```
### [DOCS] GENERATE
**if needed**, add your GOPATH to PATH
```shell
[on linux machine]$ export PATH=$GOPATH/bin:$PATH

# install go swag executable
go install github.com/swaggo/swag/cmd/swag@latest

# locally update openapi docs.
make swag

```

### [DEV]
**.env PROD=False**
```shell
# run locally (go v1.23.5 installed)
sudo docker-compose -f dev.yml up -d

# regenerates openapi docs also
make run
```
### [PROD] 
**.env PROD=True**
```shell
sudo docker-compose up --build -d
```

i
## **Usage**

### [MakeFile]

- **make run** - builds, re-generates openapi docs and runs server locally on appropriate port.
- **make fmt** - formats source code using gofmt. 
- **make sqlc** - generates sqlc code for newly added queries.
- **make swag** - generates openapi docs.


## TODO
1) Validation Service
2) Custom Errors
3) Custom Logger
4) Container Registry

