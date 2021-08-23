package main

import (
	"context"
	"github.com/smallnest/rpcx/share"

	"github.com/smallnest/rpcx/client"
	"log"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C       int
	Message string
}

func main() {
	var addr = "tcp@127.0.0.1:8972"
	d, _ := client.NewPeer2PeerDiscovery(addr, "")
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()
	args := &Args{
		A: 10,
		B: 20,
	}
	reply := &Reply{}
	//传给服务器元数据
	ctx := context.WithValue(context.Background(), share.ReqMetaDataKey, map[string]string{"ip": "127.0.0.1"})
	//读取客户端的数据
	ctx = context.WithValue(ctx, share.ResMetaDataKey, make(map[string]string))
	err := xclient.Call(ctx, "Mul", args, reply)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	log.Printf("Message%s: %d * %d = %d", reply.Message, args.A, args.B, reply.C)
}
