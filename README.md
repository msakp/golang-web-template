# Golang clean-arch Web Template

## **Tech**
- Fiber
- Sqlc
- Viper

## **Running**
### [DOCS]
**if needed**, add your GOPATH to PATH 
```shell
[on linux machine]# export PATH=$GOPATH/bin:$PATH

```
```shell
# install go swag executable
go install github.com/swaggo/swag/cmd/swag@latest

# locally update openapi docs.
make swag

# look up :)
{SERVER_ADDRES}/swagger
```

### [DEV]
**.env PROD=False**
```shell
# run locally (go v1.23.5 installed)
sudo docker-compose -f dev.yml up -d

# includes re-generating openapi docs
make run
```
### [PROD] 
**.env PROD=True**
```shell
sudo docker-compose up -d
```


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

