# Golang clean-arch Web Template

## **Tech**
- Fiber
- Sqlc
- Viper

## **Running**
### [DOCS]

```shell
# openapi docs on server. N/IMPL
```

### [DEV]
**.env PROD=False**
```shell
# run locally (go v1.23.5 installed)
sudo docker-compose -f dev.yml up -d
make run
```
### [PROD] 
**.env PROD=True**
```shell
sudo docker-compose up -d
```


## **Usage**

### [MakeFile]

- **make run** - builds and runs server locally on appropriate port.
- **make fmt** - formats source code using gofmt. 
- **make sqlc** - generates sqlc code for newly added queries.


