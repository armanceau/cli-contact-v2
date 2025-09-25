package main

import (
	"github.com/armanceau/cli-contact-v2/internal/app"
	"github.com/armanceau/cli-contact-v2/internal/storage"
)

func main() {
	var store storage.Storer = storage.NewJsonStore("./")
	app.Run(store)
}
