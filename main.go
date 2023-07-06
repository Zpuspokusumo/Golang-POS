package main

import (
	"encoding/json"
	"fmt"

	Objects "example.com/postest/objects"
	Products "example.com/postest/products"
)

func Show_products() {
	x, _ := json.MarshalIndent(Products.PData, "", "    ")
	/*if err != nil {
		err.Error()
	}*/
	fmt.Print(string(x))
}

func main() {
	fmt.Println("Welcome to the POS System")
	/*Broid := 0
	for {
		if Broid == 1 {
			break
		}
		Toid := time.Now().UnixNano()
		Trox := Objects.Transaction{}
		Trox.Trxid = Toid
		Trox.Date = time.Now().Unix()
	}*/
	//Show_products()
	Trx := Objects.NewTrax()
	Trx.Add_items("Tango", 5, 6000)
	Trx.Add_items("Mainan Mobil", 1, 200000)
	Trx.Add_items("Pasta Gigi", 2, 15000)
	Trx.Check_order()

}
