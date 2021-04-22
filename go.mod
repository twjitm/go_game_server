module go_game_server

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/coreos/etcd v3.3.13+incompatible
	github.com/fogleman/gg v1.3.0
	github.com/gin-contrib/sse v0.0.0-20190301062529-5545eab6dad3 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/gorilla/websocket v1.4.2 // indirect
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/mailru/easyjson v0.7.6
	github.com/mattn/go-isatty v0.0.7 // indirect
	github.com/onsi/ginkgo v1.14.0 // indirect
	github.com/spf13/viper v1.7.0
	github.com/stretchr/testify v1.5.1 // indirect
	github.com/ugorji/go v1.1.4 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.24.0
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace (
	github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0
	google.golang.org/grpc => google.golang.org/grpc v1.27.0
)
