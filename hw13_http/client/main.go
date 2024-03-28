package main

import (
	"flag"
	"log"

	client "github.com/TOsmanov/my-hw/hw13_http/client/pkg"
)

var (
	url    string
	path   string
	update string
)

func init() {
	if flag.Lookup("url") == nil {
		flag.StringVar(&url, "url",
			"http://localhost:10001", "Server URL")
	}
	if flag.Lookup("path") == nil {
		flag.StringVar(&path, "path",
			"", "Path")
	}
	if flag.Lookup("update") == nil {
		flag.StringVar(&update, "update",
			"", "Update message")
	}
}

func main() {
	flag.Parse()
	client := client.NewClient(url, path)

	if len(update) == 0 {
		msg, err := client.GetMessage()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(msg)
	} else {
		resp, err := client.PostMessage(update)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(resp)
	}
}
