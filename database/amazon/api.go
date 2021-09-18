package main

import "order/app/Http/Controllers/Logic/AmazonApi"

func main() {
	api := AmazonApi.Account{
		RefreshToken: "Atzr|IwEBIP_Uys3-igMeC2v09MH_3bNwQCcb0u-u6KoEBjNmg6NjSebt9WgXGEMX3YN9Uk70PhCa97xrAV66jjTvFH4IYhFQuKE3SCySyAli1u-h5z75fgYRZ2e7oq4-nJV8zxWPLQmrl8lfCI4ClGDWAY3fX5tLCzHzgYVvIgMqMHVlBjUgVqERUlyG0r6KbqaJDH3O2wz3jLyM2MKbxC6kRggkMNlklk8QGONsk2qhwOz4rHexZW6vskriuZFdtVfUdWmijYlbkw5BKPdjYAUuxBsGgcua9kwjfeuz6h8TLDuro5pIvtI9bQt-whm7JHPfmkTSAZM",
		Area:         "EU",
		ClientID:     "amzn1.application-oa2-client.0dcfe86970404e008baa6be6f27f32bb",
		ClientSecret: "fc05ce7768383ab8b403af1818256fa2dbc7fbc82521b6e385501e18529c4db5",
		AccessKeyID:  "AKIASMC5IIUOMROK2RF3",
		SecretKey:    "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",
		Region:       "eu-west-1",
		RoleArn:      "arn:aws:iam::163404334364:role/nantang_role",
		Endpoint:     "https://sellingpartnerapi-eu.amazon.com",
		ProxyIp:      "",
		ProxyPort:    "",
	}
	api.GetOrderItems("408-4843951-2863535")
}
