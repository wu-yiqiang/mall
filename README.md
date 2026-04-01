# mall

goctl api go -api user.api -dir . -style=goZero

model
    带缓存
goctl model mysql datasource -url="root:root@root@tcp(127.0.0.1:3306)/go-zero" -table="user" -dir="./model" -c
    不带缓存
goctl model mysql datasource -url="root:root@root@tcp(127.0.0.1:3306)/go-zero" -table="user" -dir="./model"