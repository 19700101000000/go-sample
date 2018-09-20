package main

import (
	"fmt"
	"github.com/mattn/go-slim"
	"log"
	"os"
)

func main() {
	tmpls := []string{
		"template/good.slim",
		"template/good2.slim",
		"template/bad.slim",
		"template/bad2.slim",
		"template/bad3.slim",
		"template/bad4.slim",
	}
	for _, v := range tmpls {
		fmt.Printf("\n--- %s ---\n", v)
		tmpl, err := slim.ParseFile(v)
		if err != nil {
			logErr(v, err)
			continue
		}
		err = tmpl.Execute(os.Stdout, slim.Values{
			"nums": []int{1, 2, 3},
		})
		if err != nil {
			logErr(v, err)
		}
	}
}

func logErr(s string, e error) {
	log.Fatalf("%s:%v", s, e)
}
