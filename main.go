package main

import (
	"fmt"
	"os"

	"github.com/michalnovotny3/EconomistBlog-formatxml/formatxml"
)

func main() {
	Args := os.Args
	var path string
	if len(Args) < 2 {
		path = "formatxml/example.json"
	} else {
		path = Args[1]
	}
	Tmp_file, err := formatxml.OpenJSON(path)
	if err != nil {
		panic(err)
		return
	}
	obj, err := formatxml.ParseToXML(Tmp_file)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(obj))
}
