package config

import (
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"io"
	"log"
	"os"
	"strconv"
)

var jc JaegerConfig

type JaegerConfig struct {
	Enabled     bool
	ServiceName string
	Agent       interface{}
}

// NewTracer
func NewTracer(service string) (opentracing.Tracer, io.Closer) {
	ParseConfig()
	return newTracer(service, "")
}

// newTracer
func newTracer(service, collectorEndpoint string) (opentracing.Tracer, io.Closer) {
	// 参数详解 https://www.jaegertracing.io/docs/1.20/sampling/
	cfg := jaegerConfig.Configuration{
		ServiceName: jc.ServiceName,
		// 采样配置
		Sampler: &jaegerConfig.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegerConfig.ReporterConfig{
			LogSpans: true,
			// CollectorEndpoint:  CollectorEndpoint2, // 将span发往jaeger-collector的服务地址
			LocalAgentHostPort: jc.Agent.(map[string]string)["host"],
		},
	}
	// 不传递 logger 就不会打印日志
	tracer, closer, err := cfg.NewTracer(jaegerConfig.Logger(jaeger.StdLogger))
	// tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer
}

// ParseConfig 从 viper 中解析配置信息
func ParseConfig() JaegerConfig {
	enabled := viper.GetString("JAEGER_ENABLED")
	agent := viper.GetString("OPENTRACING_AGENT")
	name := viper.GetString("JAEGER_SERVICE_NAME")
	// 默认值
	if len(agent) == 0 {
		agent = "0.0.0.0:6831"
	}
	if len(enabled) == 0 {
		enabled = "false"
	}
	if enabled == "true" && name == "" {
		log.Println("JAEGER_SERVICE_NAME 不能为空")
		os.Exit(200)
	}
	jc.Enabled, _ = strconv.ParseBool(enabled)
	jc.Agent = map[string]string{"host": agent}
	jc.ServiceName = name
	return jc
}
