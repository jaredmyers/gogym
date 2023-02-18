package main

import (
	"flag"
	"log"

	"github.com/jaredmyers/gogym/api/testing/structure/api"
	"github.com/jaredmyers/gogym/api/testing/structure/storage"
)

func main() {

	listenAddr := flag.String("listenaddr", ":3000", "server address")
	flag.Parse()
	store := storage.NewMemoryStorage()

	server := api.NewServer(*listenAddr, store)
	log.Println("server running on port:", *listenAddr)
	log.Fatal(server.Start())

}
