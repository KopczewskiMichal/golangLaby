package main

import (
	"fmt"
	"flag"
)

func main() {
	zupa := flag.String("napis", "initVal", "podaj napis")
	flag.Parse()
	fmt.Println(zupa)
	fmt.Println(*zupa)
}