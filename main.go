package main

import (
	"log"
	"os"

	"github.com/hesidoryn/jt/client"
)

func main() {
	client, err := client.NewClient("config.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	res, err := client.Exists("k")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(res)

	err = client.Set("k", "value")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resGet, err := client.Get("k")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(resGet)

	resDel, err := client.Del("k")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	log.Println(resDel)

	err = client.Rpush("list", "first")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	m := map[string]string{
		"field1": "value1",
		"field2": "value2",
	}
	err = client.DSet("dict", m)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	resDGet, err := client.DGet("dict", "field1", "field2")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println(resDGet)
}
