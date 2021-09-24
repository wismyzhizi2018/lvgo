package config

var nacos NacosConfig

type NacosConfig struct {
	Endpoint    string
	NamespaceId string
	AccessKey   string
	SecretKey   string
	Port        uint64
	DataId      string
	Group       string
}

// NewNacosConfig 从 viper 中解析配置信息
func NewNacosConfig(endpoint, namespaceId, accessKey, secretKey, dataId, group string, port uint64) NacosConfig {
	nacos.Endpoint = endpoint
	nacos.NamespaceId = namespaceId
	nacos.AccessKey = accessKey
	nacos.SecretKey = secretKey
	nacos.DataId = dataId
	nacos.Group = group
	nacos.Port = port
	return nacos
}

func GetNacosConfig() NacosConfig {
	return nacos
}
