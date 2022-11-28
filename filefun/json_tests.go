// Golang
// file fun
// for reference
// json marshal/unmarshal testing

package filefun

import (
	"encoding/json"
	"fmt"
	//"io"
	"log"
	//"os"
)

type Worker struct {
	FirstName string      `json:"first_name"`
	LastName  string      `json:"last_name"`
	Desc      interface{} `json:"desc"`
}

type Homeland struct {
	Country string `json:"country"`
	City    string `json:"city"`
	Power   []int  `json:"power"`
}

type Post struct {
	ID      string   `json:"id"`
	Message string   `json:"message"`
	Forward []string `json:"forward"`
}

func test2() {

	p1 := &Worker{}
	bs1 := `{"first_name": "david", "last_name": "Heart", "desc": {"country": "UK", "city": "London", "power":[1,3,5,8]}}`

	if err := json.Unmarshal([]byte(bs1), &p1); err != nil {
		log.Panic(err)
	}
	fmt.Println(p1)
	var msg json.RawMessage
	p2 := Worker{Desc: &msg}
	if err := json.Unmarshal([]byte(bs1), &p2); err != nil {
		log.Panic(err)
	}
	fmt.Println(p2)
	fmt.Println(p1)

	switch p2.FirstName {
	case "david":
		deets := &Homeland{}
		if err := json.Unmarshal(msg, &deets); err != nil {
			log.Panic(err)
		}
		fmt.Println(deets)

		descCountry, descCity, descPower := deets.Country, deets.City, deets.Power
		fmt.Println(descCountry, descCity, descPower)

		p2.Desc = deets

		fmt.Println(p2, p2.Desc)
	default:
		log.Printf("unknown first name: %q", p1.FirstName)

	}
}

func test1() {

	p := &Post{"234", "Hello!", []string{"Ema", "David"}}

	jp, err := json.Marshal(p)
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%T\n", jp)
	fmt.Println(jp)
	fmt.Println(string(jp))

	post := `{"id":"212", "message":"Hello", "forward":["Lily", "Fred"]}`
	p2 := &Post{}

	err = json.Unmarshal([]byte(post), &p2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(p2)
	fmt.Printf("%T\n", p2)

	p2.ID, p2.Message = "1234", "nahhDog"
	p3, _ := json.Marshal(p2)

	fmt.Printf("%T\n", p3)
	fmt.Println(string(p3))
}

