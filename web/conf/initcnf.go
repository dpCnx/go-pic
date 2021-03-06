package conf

import (
	"gopkg.in/ini.v1"
	"os"
	"path/filepath"
)

var (
	cfg *ini.File
	C   Config
)

type Config struct {
	Server ServerConfig `ini:"server"`
	Grpc   GrpcConfig   `ini:"grpc"`
	Etcd   EtcdConfig   `ini:"etcd"`
	Log    LogConfig    `ini:"log"`
}

type GrpcConfig struct {
	Ip   string `ini:"ip"`
	Port string `ini:"port"`
}

type ServerConfig struct {
	Port string `ini:"port"`
	Mode string `ini:"mode"`
}

type EtcdConfig struct {
	IP   string `ini:"ip"`
	Port string `ini:"port"`
}

type LogConfig struct {
	Filename   string `ini:"filename"`
	MaxSize    int    `ini:"max_size"`
	MaxBackups int    `ini:"max_backups"`
	MaxAge     int    `ini:"max_age"`
}

func init() {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg, err = ini.Load(filepath.Join(dir, "/conf", "app.conf"))
	if err != nil {
		panic(err)
	}

	if err = cfg.MapTo(&C); err != nil {
		panic(err)
	}

}

func GetString(sec, key string) string {
	return cfg.Section(sec).Key(key).String()
}

func GetBool(sec, key string) (bool, error) {
	return cfg.Section(sec).Key(key).Bool()
}

func GetInt(sec, key string) (int, error) {
	return cfg.Section(sec).Key(key).Int()
}
