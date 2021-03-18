
# basic crud with golang
do the basic crud with this app. this app will provide api to store, update, delete infomation

## api list

| METHOD        | api            | Output         | body
| --------------| ---------------| ---------------| ------------------------------------------
| Get           | /users         | all user       |                                           
| Get           | /users/{id}    | single user    |                                           
| Post          | /users         | store user     | {"ID":1,"FirstName":"foo","LastName":"bar"}
| Put           | /users/{id}    | update user    | {"ID":1,"FirstName":"foo","LastName":"bar"}
| Delete        | /users/{id}    | delete user    |                                            

## Usage

Compile with `go build github.com/TafhimFaisal/basic_crud_with_golang `
or run with `go run github.com/TafhimFaisal/basic_crud_with_golang `
this app will run at `http://localhost:3000`

## blog
check out my blog https://learninggolangwithtafhim.blogspot.com/