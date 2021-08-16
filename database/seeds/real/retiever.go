package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type Retriever struct {
	UserAgent   string
	ContentType string
	TimeOut     time.Duration
}

func (r *Retriever) Get(url string) string {
	httpreq, _ := http.Get(url)
	rep, _ := httputil.DumpResponse(httpreq, true)
	defer httpreq.Body.Close()
	return string(rep)
}

//func (r *Retriever) Post(url string, form map[string]string) string {
//	httpreq, _ := http.Post(url, r.ContentType, strings.NewReader("name=cjb"))
//	rep, _ := httputil.DumpResponse(httpreq, true)
//	defer httpreq.Body.Close()
//	return string(rep)
//}
