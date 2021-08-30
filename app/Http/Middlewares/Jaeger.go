package Middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"order/config"
)

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parentSpan opentracing.Span
		tracer, closer := config.NewTracer("")
		defer closer.Close()
		// 直接从 c.Request.Header 中提取 span,如果没有就新建一个
		spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {
			parentSpan = tracer.StartSpan(c.Request.URL.Path)
			defer parentSpan.Finish()
		} else {
			parentSpan = opentracing.StartSpan(
				c.Request.URL.Path,
				opentracing.ChildOf(spCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
				ext.SpanKindRPCServer,
			)
			defer parentSpan.Finish()
		}
		// 然后存到 g.ctx 中 供后续使用
		c.Set("tracer", tracer)
		c.Set("ctx", opentracing.ContextWithSpan(context.Background(), parentSpan))
		c.Next()
	}
}
