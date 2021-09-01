package Consul

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"

	"github.com/satori/go.uuid"
)

func TestConsulServiceRegistry(t *testing.T) {
	host := "192.168.0.29"
	port := 8500

	registerClient := NewRegistryClient(host, port)
	serviceId := fmt.Sprintf("%s", uuid.NewV4())
	err := registerClient.Register("192.168.0.49", 8099, "global.ServerConfig.Name", []string{"php", "go"}, serviceId)
	if err != nil {
		t.Error(err)
	}
	r := gin.Default()
	// 健康检测接口，其实只要是 200 就认为成功了
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	if errs := r.Run(":8099"); errs != nil {
		registerClient.DeRegister(serviceId)
	}
}
