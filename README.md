# mall

goctl api go -api user.api -dir . -style=goZero

model
    带缓存
goctl model mysql datasource -url="root:root@root@tcp(127.0.0.1:3306)/go-zero" -table="user" -dir="./model" -c
    不带缓存
goctl model mysql datasource -url="root:root@root@tcp(127.0.0.1:3306)/go-zero" -table="user" -dir="./model"


RPC
goctl rpc protoc user.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.


启动etcd或者consule
./etcd

rpc测试工具启动命令
./grpcui -plaintext localhost:8080


swagger
   goctl api swagger --api ./system.api --dir ./docs/swagger
