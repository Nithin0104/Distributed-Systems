package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	Title string
	Body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection error: ", err)
	}

	a := Item{"First", "First File content"}
	b := Item{"Second", "Second File content"}
	c := Item{"Third", "Third File content"}

	client.Call("API.CreateItem", a, &reply)
	client.Call("API.CreateItem", b, &reply)
	client.Call("API.CreateItem", c, &reply)
	client.Call("API.GetListOfFile", "", &db)

	fmt.Println("List of Files: ", db)

	client.Call("API.DeleteItem", c, &reply)
	client.Call("API.GetListOfFile", "", &db)
	fmt.Println("List of Files after Deleting ", db)

	client.Call("API.GetByName", "First", &reply)
	fmt.Println("first file body: ", reply.Body)

}
