package main

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu"
)

func main() {
	oldFile := "D:/360Downloads/storage_label.pdf"
	newFile := "D:/360Downloads/storage_label_new.pdf"
	waterFile := "D4 Midin China"
	water(oldFile, newFile, waterFile)
}
func water(oldFile string, newFile string, waterFile string) {
	onTop := true
	wm, _ := pdfcpu.ParseTextWatermarkDetails(waterFile, "ac:2 abs, d:2, op:1, pos:c", onTop, 1)
	api.AddWatermarksFile(oldFile, newFile, nil, wm, nil)
}
