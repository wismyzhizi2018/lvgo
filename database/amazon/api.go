package main

import "order/app/Http/Controllers/Logic/AmazonApi"

func main() {
	api := AmazonApi.Account{
		RefreshToken: "<refreshtoken>",
		Area:         "EU",
		ClientID:     "<clientid>",
		ClientSecret: "<clientSecret>",
		AccessKeyID:  "<AccessKeyID>",
		SecretKey:    "<SecretKey>",
		Region:       "eu-west-1",
		RoleArn:      "<RoleArn>",
		Endpoint:     "<Endpoint>",
		ProxyIp:      "",
		ProxyPort:    "",
	}
	api.GetOrderItems("408-4843951-2863535")
}
