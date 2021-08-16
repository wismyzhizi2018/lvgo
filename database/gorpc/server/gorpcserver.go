package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"github.com/smallnest/rpcx/share"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C       int
	Message string
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reqMeta := ctx.Value(share.ReqMetaDataKey).(map[string]string)
	fmt.Println(reqMeta)
	reply.C = args.A * args.B
	reply.Message = "返回成功"
	return nil
}
func main() {
	s := server.NewServer()
	s.RegisterName("Arith", new(Arith), "")
	s.Serve("tcp", ":8972")
}
