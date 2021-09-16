# pool_backend

```
## 初始化使用
* go mod init 
* go mod tidy  |  go install
* swag init  
* sh scripts/start.sh

## go.mod
* 使用viper/remote包会报错需添加:
* replace google.golang.org/grpc => google.golang.org/grpc v1.28.0

```