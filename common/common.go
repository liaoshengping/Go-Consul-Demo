package common

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
	"strconv"
)

func GetConsulConfig(host string ,port int64,prefix string)(config.Config,error)  {
	consulSource := consul.NewSource(
		consul.WithAddress(host+":"+strconv.FormatInt(port,10)),
		consul.WithPrefix(prefix),
		consul.StripPrefix(false),
	)

	config,err := config.NewConfig()

	if err != nil {
		return config,err
	}

	err = config.Load(consulSource);

	return config, err


}
