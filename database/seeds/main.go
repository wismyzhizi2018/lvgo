package main

import (
	"fmt"
	"github.com/gookit/color"
	"order/database/seeds/mock"
	"order/database/seeds/real"
	"time"
)

type Retriever interface {
	Get(string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	Retriever
	Poster
}

func postDownload(r RetrieverPoster) {
	r.Get("https://www.imooc.com")
	r.Post("https://www.imooc.com", map[string]string{"name": "yeweipng", "mobile": "15999640682"})
}
func download(r Retriever) string {
	return r.Get("https://www.imooc.com")
}
func session(r RetrieverPoster) string {
	r.Post("https://www.imooc.com", map[string]string{"name": "yeweipng", "mobile": "15999640682", "contents": "application/x-www-form-urlencoded"})
	return r.Get("https://www.imooc.com")
}
func main() {
	var r Retriever
	retriever := mock.Retriever{Contents: "this is mock"}
	r = &retriever
	printMsg(r)

	//os.Open(name string) (*File, error)
	//File也实现了Read方法，所以也就是实现了io.Reader接口
	//func (f *File) Read(b []byte) (n int, err error)
	if mr, ok := r.(*mock.Retriever); ok {
		fmt.Println(mr.Contents)
	} else {
		fmt.Println("not mock Retriever")
	}
	r = &real.Retriever{UserAgent: "PostmanRuntime/7.28.3", ContentType: "application/x-www-form-urlencoded", TimeOut: time.Second * 5}
	printMsg(r)

	fmt.Println(session(&retriever))
}

func printMsg(r Retriever) {
	switch r.(type) {
	case *mock.Retriever:
		color.Debug.Printf("%T %v\n", r, r)
	case *real.Retriever:
		color.Danger.Printf("%T %v\n", r, r)
	default:
		color.Question.Printf("%T %v\n", r, r)
	}

}
