package mock

import (
	"fmt"

	"github.com/gookit/color"
)

type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	color.Error.Printf("Retriever {Contents=%s}", r.Contents)
	return fmt.Sprintf("Retriever {Contents=%s}", r.Contents)
}

func (r *Retriever) Get(s string) string {
	return r.Contents
}

func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
