package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-gogo/protobuf/pb"

	"github.com/gogo/protobuf/proto"
)

const fname string = "./data/address_book.txt"

func main() {
	p := pb.Person{
		Id:    1234,
		Name:  "John Doe",
		Email: "jdoe@example.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	book := &pb.AddressBook{}
	book.People = append(book.People, &p)

	// Write the new address book back to disk.
	out, err := proto.Marshal(book)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}

	book_read := &pb.AddressBook{}
	if err := proto.Unmarshal(in, book_read); err != nil {
		log.Fatalln("Failed to parse address book_read:", err)
	}

	// print
	for _, person := range book_read.People {
		fmt.Printf("id : %d\n", person.Id)
		fmt.Printf("name : %s\n", person.Name)
		fmt.Printf("emial : %s\n", person.Email)
	}
}
