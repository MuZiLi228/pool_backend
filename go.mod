module pool_backend

go 1.16

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/eoscanada/eos-go v0.9.0 => github.com/kimtony/eos-go v0.9.2

require (
	cloud.google.com/go v0.51.0 // indirect
	cloud.google.com/go/bigquery v1.3.0 // indirect
	cloud.google.com/go/pubsub v1.1.0 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.1116
	github.com/bwmarrin/snowflake v0.3.0
	github.com/ddliu/go-httpclient v0.6.9
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fsnotify/fsnotify v1.4.9
	github.com/getsentry/sentry-go v0.11.0
	github.com/gin-contrib/pprof v1.2.1
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.7.2
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.6.1
	github.com/go-redis/redis/v8 v8.11.1
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529 // indirect
	github.com/googollee/go-socket.io v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/hashicorp/consul/api v1.8.1
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/joho/godotenv v1.3.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/panjf2000/ants/v2 v2.4.6
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.5.1
	github.com/prometheus/common v0.10.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	github.com/sirupsen/logrus v1.6.0
	github.com/smartystreets/assertions v1.1.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/cast v1.3.0
	github.com/spf13/viper v1.7.1
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.7.0
	github.com/tidwall/gjson v1.8.0
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/zap v1.17.0
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/mod v0.4.2 // indirect
	golang.org/x/oauth2 v0.0.0-20200107190931-bf48bf16ab8d // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/time v0.0.0-20201208040808-7e3f01d25324 // indirect
	google.golang.org/genproto v0.0.0-20200108215221-bd8f9a0ef82f // indirect
	google.golang.org/grpc v1.36.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/gorm v1.21.10
	honnef.co/go/tools v0.1.3 // indirect
)
