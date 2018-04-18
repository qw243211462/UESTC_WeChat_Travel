package configs

import (
	"flag"
	"runtime"
	"log"
	"github.com/larspensjo/config"
	"fmt"
)

var(
	//value 配置文件的路径
	configFile = flag.String("configfile","/Users/laiye/Documents/wangmaotong/laiye/UESTC_WeChat_Travel/src/configs/conf.ini","General configuration file")
	TOPIC = make(map[string]string)
)



func LoadConfig(section string) (map[string]string){
	runtime.GOMAXPROCS(runtime.NumCPU())
	flag.Parse()

	cfg,err := config.ReadDefault(*configFile)   //读取配置文件，并返回其Config

	if err != nil {
		log.Fatalf("Fail to find %v,%v",*configFile,err)
	}

	if cfg.HasSection(fmt.Sprintf("%v",section)) { //判断配置文件中是否有section（一级标签）
		options, err := cfg.SectionOptions(fmt.Sprintf("%v",section)) //获取一级标签的所有子标签options（只有标签没有值）
		if err == nil {
			for _, v := range options {
				optionValue, err := cfg.String(fmt.Sprintf("%v",section), v) //根据一级标签section和option获取对应的值
				if err == nil {
					TOPIC[v] = optionValue
				}
			}
		}
	}
	return TOPIC
}