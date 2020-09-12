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
	MySQL  MySQLConfig  `ini:"mysql"`
	Redis  RedisConfig  `ini:"redis"`
	Log    LogConfig    `ini:"log"`
}

type ServerConfig struct {
	Port string `ini:"port"`
	Mode string `ini:"mode"`
}

type MySQLConfig struct {
	IP       string `ini:"ip"`
	Port     string `ini:"port"`
	User     string `ini:"user"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type RedisConfig struct {
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
	cfg, err = ini.Load(filepath.Join(dir, "web/conf", "app.conf"))
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
