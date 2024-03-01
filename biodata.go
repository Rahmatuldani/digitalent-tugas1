package main

import (
	"assignment1/friend"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// variabel untuk menyimpan data collection friend
	friendCollection := friend.FriendsCollection{}

	// menambah data friend dalam collection
	friendCollection.AddFriend(friend.NewFriend(1, "Nama 1", "alamat 1", "pekerjaan 1", "alasan 1"))
	friendCollection.AddFriend(friend.NewFriend(2, "Nama 2", "alamat 2", "pekerjaan 2", "alasan 2"))
	friendCollection.AddFriend(friend.NewFriend(3, "Nama 3", "alamat 3", "pekerjaan 3", "alasan 3"))

	// ambil data inputan argumen cli
	args := os.Args

	// check jika argumen kosong
	if len(args) == 1 {
		fmt.Printf("Please input id")
		return
	}

	// konversi string menjadi integer
	id, err := strconv.ParseInt(args[1], 10, 8)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	// mencari friend dengan id
	friends, err := friendCollection.GetFriend(int8(id))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(friends)
}