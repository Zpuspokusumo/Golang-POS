package Objects

import (
	"fmt"
	"time"

	Products "example.com/postest/products"
)

type Product struct {
	Name  string
	Qty   uint32
	Price uint64
}

type Transaction struct {
	Trxid int64
	Date  int64
	Cart  map[string]Product
}

func NewTrax() Transaction {
	newtrx := &Transaction{}
	newtrx.Trxid = time.Now().UnixNano()
	newtrx.Date = time.Now().Unix()
	newtrx.Cart = make(map[string]Product)
	return *newtrx
}

func NewProduct(name string, qty uint32, price uint64) Product {
	newprod := Product{name, qty, price}
	return newprod
}

func (trx *Transaction) Add_items(name string, qty uint32, price uint64) {
	newtrx := trx.Cart[name]
	newtrx.Name = name
	newtrx.Qty = qty
	newtrx.Price = price
	trx.Cart[name] = newtrx
	fmt.Printf("added item %s, %d, %d\n", name, qty, price)
}

func (trx *Transaction) Update_item_name(name string, name2 string) {
	prod, ok := trx.Cart[name]
	if !ok { // check if product is in cart
		fmt.Println("Item is not in cart")
		return
	}
	prod2, ok := trx.Cart[name2]
	if ok { //item2 is in cart, add qty to prod2, delete prod
		prod2.Qty += prod.Qty
		delete(trx.Cart, name)
	} else { //item2 is not in cart, move name item qty to name2 item
		prod2 = prod
		delete(trx.Cart, name)
	}
	fmt.Printf("Updated product %s name to %s\n", name, name2)
	//
}

func (trx *Transaction) Update_item_qty(name string, qty uint32) {
	prod, ok := trx.Cart[name]
	if !ok { // check if product is in cart
		fmt.Println("Item is not in cart")
		return
	}
	prod.Qty = qty
	trx.Cart[name] = prod
	fmt.Printf("Updated product %s qty to %d\n", name, qty)
}

func (trx *Transaction) Update_item_price(name string, price uint64) {
	prod, ok := trx.Cart[name]
	if !ok { // check if product is in cart
		fmt.Println("Item is not in cart")
		return
	}
	prod.Price = price
	trx.Cart[name] = prod
	fmt.Printf("Updated product %s price to %d\n", name, price)
}

func (trx *Transaction) Delete_item(name string) {
	delete(trx.Cart, name)
	fmt.Printf("Deleted product %s\n", name)
}

func (trx *Transaction) Reset_Trx() {
	trx.Cart = make(map[string]Product)
	fmt.Printf("Transaction has been reset\n")
}

func (trx *Transaction) Check_order() {
	//maybe load invalid products in a slice
	//and later reference it as in a table of invalid data
	var invaliddata []Product
	fmt.Println("| No |  Nama Item  | Jumlah Item | Harga/Item | Total Harga |")
	fmt.Println("|----|-------------|-------------|------------|-------------|")
	for item, prod := range trx.Cart {
		i := 1
		_, ok := Products.PData[item]
		if !ok { // doesnt exist in data
			invaliddata[len(invaliddata)] = trx.Cart[item]
			continue
		}
		fmt.Printf("|%3d ", i)           //no, padding to the left
		fmt.Printf("| %-12s", prod.Name) //item name, padding to the right
		fmt.Printf("| %-11d ", prod.Qty)
		fmt.Printf("| %-10d ", prod.Price)
		fmt.Printf("| %-11d |\n", (prod.Qty * uint32(prod.Price)))

		i++
	}
	//print invalid data here
	//fmt.Printf("Invalid product found %s, please change\n", item)
}

func (trx *Transaction) Total_price() {
	var total int64
	var disc int64
	total = 0
	for item, iteminfo := range trx.Cart {
		total += (Products.PData[item] * int64(iteminfo.Qty))
	}
	if total >= 100000 {
		disc = 5
	} else if total >= 200000 {
		disc = 10
	} else if total >= 500000 {
		disc = 20
	}
	fmt.Printf("Subtotal is %d with a discount of %d%% (%d)", total, disc, ((total / 100) * disc))
	total = total - ((total / 100) * disc)
	fmt.Printf("Total: %d", total)
}
