package common

import "github.com/micro/go-micro/v2/config"

type MysqlConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Pwd      string `json:"pwd"`
	Port     string `json:"port"`
	Database string `json:"database"`
}

//获取mysql 配置

func GetMysqlFromConsul(config config.Config, path ...string) *MysqlConfig {

	mysqlConfig := &MysqlConfig{}

	config.Get(path...).Scan(mysqlConfig);

	return mysqlConfig;

}
