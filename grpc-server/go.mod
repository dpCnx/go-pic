module grpc-server

go 1.14

require (
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/golang/protobuf v1.4.3
	github.com/google/uuid v1.1.2 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.2
	github.com/natefinch/lumberjack v2.0.0+incompatible
	go.uber.org/zap v1.16.0
	google.golang.org/grpc v1.34.0
	gopkg.in/ini.v1 v1.62.0
	gorm.io/driver/mysql v1.0.3
	gorm.io/gorm v1.20.8
)

replace google.golang.org/grpc v1.34.0 => google.golang.org/grpc v1.26.0
