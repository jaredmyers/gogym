package main

import (
	"flag"
	"fmt"
	"log"
)

// optional database seeding //
func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Println("NewAccount Error")
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Println("CreateAccount Error")
		log.Fatal(err)
	}

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "Larry", "Finkle", "pw123")
	seedAccount(s, "Monkey", "Man", "pwYay")
}

// Entry point //
func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		log.Println("seeding the db")
		seedAccounts(store)
	}

	fmt.Printf("%+v\n", store)

	server := NewAPIServer(":3000", store)
	server.Run()

}
