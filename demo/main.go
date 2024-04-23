package main

import (
	"fmt"
	"github.com/razer96/docxt"
)

type TestStruct struct {
	FileName string
	Items    []TestItemStruct
}

type TestItemStruct struct {
	Column1  string
	Column2  string
	SubItems []TestItemStruct2
}

type TestItemStruct2 struct {
	Column1 string
	Column4 string
}

func main() {
	template, err := docxt.OpenTemplate("./example.docx")
	if err != nil {
		fmt.Println(err)
		return
	}
	t := make(map[string]any)

	t["Items"] = []map[string]any{
		{"Column1": "1", "Column2": "2",
			"SubItems": []map[string]any{
				{"Column1": "3", "Column4": "4"},
				{"Column1": "5", "Column4": "6"},
			},
		},
	}
	t["FileName"] = "example.docx"

	if err := template.RenderTemplate(t); err != nil {
		fmt.Println(err)
		return
	}
	if err := template.Save("result.docx"); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Success")
}
