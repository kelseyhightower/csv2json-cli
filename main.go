package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/kelseyhightower/csv2json"
)

var (
	infile string
)

func init() {
	flag.StringVar(&infile, "infile", "", "input file")
}

func main() {
	flag.Parse()
	if infile == "" {
		log.Fatal("Input file required")
	}
	f, err := os.Open(infile)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	jsonData, err := csv2json.Convert(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
