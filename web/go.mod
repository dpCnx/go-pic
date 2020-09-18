module go-pic

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/google/uuid v1.1.2 // indirect
	github.com/juju/ratelimit v1.0.1
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.7
	go.etcd.io/etcd v3.3.25+incompatible
	go.uber.org/ratelimit v0.1.0
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.32.0
	gopkg.in/ini.v1 v1.61.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace github.com/coreos/bbolt v1.3.5 => go.etcd.io/bbolt v1.3.5

replace google.golang.org/grpc v1.32.0 => google.golang.org/grpc v1.26.0
