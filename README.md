## Mysql form

A signup/registration form that accepts user details through a frontend or API call and stores the details in a mysql database
## Install
### Any version of [golang](https://go.dev/dl/)
### `gorilla/mux` package
```sh
go get -u github.com/gorilla/mux
```
### `gorm` package
```sh
 go get -u "github.com/jinzhu/gorm"
```
### `mysql` package
```sh
 go get -u "github.com/jinzhu/gorm/dialects/mysql"
```

#### include your dialect and argument in  `gorm.Open()`e.g.
```
func connect(){
gorm.Open("mysql","user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
}
```