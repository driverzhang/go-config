package app

import (
	"fmt"
	"testing"

	"gitHub/go-config/config" // 引入你自己的GOPAHT下的对应config包路径
)

// 使用者根据自己需要修改这个结构体
// 实际开发中 可将此文件在根目录专门建一个文件夹包里放该文件内容
var C struct {
	Debug    bool `yaml:"debug"`
	LogDebug bool `yaml:"log_debug"`
	Mysql    struct {
		Bamboo string `yaml:"bamboo_website"`
	} `yaml:"mysql"`
	GrpcService struct {
		User string `yaml:"user"`
	} `yaml:"grpc_service"`
	HttpAddr string `yaml:"http_addr"`
	GrpcAddr string `yaml:"grpc_addr"`
}

func TestConfig(t *testing.T) {
	config.Init(&C)
	fmt.Println(config.C.Debug, config.C.Mysql.Bamboo, config.C.Mysql)
	t.Log(config.C.Debug)
}
