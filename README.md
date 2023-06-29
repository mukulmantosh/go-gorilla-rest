
# Golang Gorilla/Mux CRUD REST

![golang_icon](./misc/golang-1.svg)
In this application we are going to create CRUD endpoints using Gorilla/Mux and GORM.


## Installing Dependencies
```
go get -u github.com/gorilla/mux
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```


When you're creating the `user` package make sure to run the below command

```
god mod init user
god mod tidy
```

Run the below command in the root

```
go mod edit -replace user=./user
```






