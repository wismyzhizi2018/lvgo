package main

import (
	"fmt"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	endpoint := "addr-hz-internal.edas.aliyun.com"
	namespaceId := "9be813e2-b734-485e-afb1-24173a33f366"
	accessKey := "LTAI4G3ER6msFBtVYLF2wEvR"
	secretKey := "298PGQrCMA8FDKpGSoSscmiNnmg7Q1"

	clientConfig := constant.ClientConfig{
		Endpoint:       endpoint + ":8080",
		NamespaceId:    namespaceId,
		AccessKey:      accessKey,
		SecretKey:      secretKey,
		TimeoutMs:      5 * 1000,
		ListenInterval: 30 * 1000,
	}

	// Initialize client.
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig": clientConfig,
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	dataId := "app.yaml"
	group := "test"

	// Get plain content from ACM.
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	fmt.Println("Get config：" + content)
}
