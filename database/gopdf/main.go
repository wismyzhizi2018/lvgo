//package main
//
//import (
//	"github.com/pdfcpu/pdfcpu/pkg/api"
//	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
//)

package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
)

func main() {
	err := ChromedpPrintPdf("https://www.google.com", "/path/to/file.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func ChromedpPrintPdf(url string, to string) error {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var buf []byte
	err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitReady("body"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			buf, _, err = page.PrintToPDF().
				Do(ctx)
			return err
		}),
	})
	if err != nil {
		return fmt.Errorf("chromedp Run failed,err:%+v", err)
	}

	if err := ioutil.WriteFile(to, buf, 0644); err != nil {
		return fmt.Errorf("write to file failed,err:%+v", err)
	}

	return nil
}

//func main() {
//	oldFile := "D:/360Downloads/storage_label.pdf"
//	newFile := "D:/360Downloads/storage_label_new.pdf"
//	waterFile := "D4 Midin China"
//	water(oldFile, newFile, waterFile)
//}
//func water(oldFile string, newFile string, waterFile string) {
//	onTop := true
//	wm, _ := pdfcpu.ParseTextWatermarkDetails(waterFile, "ac:2 abs, d:2, op:1, pos:c", onTop, 1)
//	api.AddWatermarksFile(oldFile, newFile, nil, wm, nil)
//}
//
