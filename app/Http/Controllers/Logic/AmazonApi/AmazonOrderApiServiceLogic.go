package AmazonApi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"gopkg.me/selling-partner-api-sdk/ordersV0"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	sp "gopkg.me/selling-partner-api-sdk/pkg/selling-partner"
	"gopkg.me/selling-partner-api-sdk/sellers"
)

type Account struct {
	RefreshToken string
	Area         string
	ProxyIp      string
	ProxyPort    string
	ClientID     string
	ClientSecret string
	AccessKeyID  string
	SecretKey    string
	Region       string
	RoleArn      string
	Endpoint     string
}

func (a Account) getApiInstace() {
	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     a.ClientID,
		ClientSecret: a.ClientSecret,
		RefreshToken: a.RefreshToken,
		AccessKeyID:  a.AccessKeyID,
		SecretKey:    a.SecretKey,
		Region:       a.Region,
		RoleArn:      a.RoleArn,
	})
	if err != nil {
		panic(err)
	}

	endpoint := "https://sellingpartnerapi-fe.amazon.com"

	seller, err := sellers.NewClientWithResponses(endpoint,
		sellers.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("X-Amzn-Requestid", uuid.New().String()) // tracking requests
			err = sellingPartner.SignRequest(req)
			if err != nil {
				return errors.Wrap(err, "sign error")
			}
			dump, err := httputil.DumpRequest(req, true)
			if err != nil {
				return errors.Wrap(err, "DumpRequest Error")
			}
			log.Printf("DumpRequest = %s", dump)
			return nil
		}),
		sellers.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
			dump, err := httputil.DumpResponse(rsp, true)
			if err != nil {
				return errors.Wrap(err, "DumpResponse Error")
			}
			log.Printf("DumpResponse = %s", dump)
			return nil
		}),
	)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	_, err = seller.GetMarketplaceParticipationsWithResponse(ctx)

	if err != nil {
		panic(err)
	}
}

func (a Account) GetOrderItems(orderId string) {
	sellingPartner, err := sp.NewSellingPartner(&sp.Config{
		ClientID:     a.ClientID,
		ClientSecret: a.ClientSecret,
		RefreshToken: a.RefreshToken,
		AccessKeyID:  a.AccessKeyID,
		SecretKey:    a.SecretKey,
		Region:       a.Region,
		RoleArn:      a.RoleArn,
	})
	c, _ := ordersV0.NewClient(a.Endpoint, ordersV0.WithRequestBefore(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("X-Amzn-Requestid", uuid.New().String()) // tracking requests
		err = sellingPartner.SignRequest(req)
		if err != nil {
			return errors.Wrap(err, "sign error")
		}
		dump, err := httputil.DumpRequest(req, true)
		if err != nil {
			return errors.Wrap(err, "DumpRequest Error")
		}
		log.Printf("DumpRequest = %s", dump)
		return nil
	}), ordersV0.WithResponseAfter(func(ctx context.Context, rsp *http.Response) error {
		dump, err := httputil.DumpResponse(rsp, true)
		if err != nil {
			return errors.Wrap(err, "DumpResponse Error")
		}
		log.Printf("DumpResponse = %s", dump)
		return nil
	}),
	)
	var NextToken string
	ctx := context.Background()
	result, _ := c.GetOrderItems(ctx, orderId, &ordersV0.GetOrderItemsParams{NextToken: &NextToken})
	fmt.Println(result)
}
