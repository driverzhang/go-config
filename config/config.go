package config

import (
	"io/ioutil"
	"os"

	"log"

	"gopkg.in/yaml.v2"
)

// 传入一个结构体指针
func Init(config interface{}) {
	configFileName := "config.yml"          // 用于开发者自己本地调试
	configFileNameBack := "config.test.yml" // 用于测试服务器中
	var findConfig, findConfigBack bool
	// 向上层查找配置文件
	// 在项目的任何地方运行(test时)都能加载到配置文件
	// 有备配置文件, 当主配置文件没找到时就会使用备配置文件
	// 优先使用最近的主备配置文件
	for i := 0; i < 10; i++ {

		// Stat返回一个描述name指定的文件对象的FileInfo。 FileInfo用来描述一个文件对象
		_, err := os.Stat(configFileName)
		if err != nil {
			if os.IsNotExist(err) { // 返回一个布尔值说明该错误是否表示一个文件或目录不存在。
				configFileName = "../" + configFileName // 如果错误原因是该文件不存在，则往上层文件查找
			} else {
				panic(err)
			}
		} else {
			findConfig = true
			break
		}

		_, err = os.Stat(configFileNameBack)
		if err != nil {
			if os.IsNotExist(err) {
				configFileNameBack = "../" + configFileNameBack
			} else {
				panic(err)
			}
		} else {
			findConfigBack = true
			break
		}
	}

	var fileName string
	// 如果两个文件都存在，则首先满足 config.yml 文件用于本地调式
	if findConfig {
		fileName = configFileName
	} else if findConfigBack {
		fileName = configFileNameBack
	} else {
		log.Panicf("can't find 'config.yml' or 'config.test.yml', Please write the config file in the project root directory")
		return
	}

	log.Printf("found config file: %s", fileName)

	// 具体读取文件内的内容
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("can't read '%s' file", fileName)
		return
	}

	// 将二进制包内容 转化成 interface{},这里是传进来的 结构体指针
	err = yaml.Unmarshal(bs, config)
	if err != nil {
		log.Panicf("yaml.Unmarshal err:%v; row:%s", err, bs)
		return
	}
}
