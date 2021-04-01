package main

import (
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"log"
	"strings"
)

type Specification struct {
	ManualOverride1 string `envconfig:"manual_override_1"`
	DefaultVar      string `default:"foobar"`
	RequiredVar     string `required:"true"`
	IgnoredVar      string `ignored:"true"`
	AutoSplitVar    string `split_words:"true"`
	RequiredAndAutoSplitVar    string `required:"true" split_words:"true"`
}

func main() {
	strings.Index("aaaa", "aaa")
	var s Specification
	err := envconfig.Process("myapp", &s)
	if err != nil {
		log.Fatal(err.Error())
	}
	_, err = fmt.Printf("---> %v\n", s.ManualOverride1)
	if err != nil {
		log.Fatal(err.Error())
	}
}