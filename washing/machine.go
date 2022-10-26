package main

import (
	"fmt"
)

var allbrands [2]string
var prices []int
var sweets string = "milkkova"

func washingmachine() {

	var modelno string = "WA10BG4686BR"
	var price int = 30500
	var sweet = "milkkova"
	var features [10]string

	features[0] = "Ecobubble"
	features[1] = "Hygiene steam with in-built Heater"

	var color []string = make([]string, 2)
	color[0] = "blue"
	color[1] = "grey"

	var feature []string = make([]string, 2)
	feature[0] = "Ecobubble"
	feature[1] = "No Cost EMI starts from ₹ 4429.89/ month"

	var modelnos *string = new(string)
	*modelnos = "WA10BG4686BR"
	feature = append(feature, "1ss power")
	var sweets *string = new(string)
	*sweets = "milkkova"
	var sugar int
	if *sweets == "milkkova" {
		sugar = 1

	} else if *sweets == "palkova" {
		sugar = 2
	} else {
		fmt.Println("no sugar less sweet")

	}
	var prices int
	switch modelno {
	case "WA10BG4686BR":
		prices = 30500
		break
	case "CA10BG4686BR":
		prices = 45000
		break
	default:

	}
	//line 1        //line 2   //line 4 =>line 2
	for i := 0; i < 5; i++ {
		//conditon is to execute code in the block
		fmt.Println(features[i]) //line 3

	} //line 4
	for i, v := range color {
		//condition is to get the value in array
		fmt.Println(i, v)
	}

	fmt.Println(modelno, price, prices, features, color, feature, modelnos, sweet, sugar)

}

type feature struct {
	brand             string //have choice for the keyword can be allocate in stack or heap
	productdimensions string
	capacity          string //slice
	colour            string //slice
	specialfeature    []string
	washingprogrammes []string
	inbox             [4]string
	price             *int
	waterpurifier     *string
	noofclothes       int
}

func (p *feature) sellingcompanies() { //(string , int)means here we allocating  memory without variable
	p.brand = "samsung"
	p.colour = "blue"
	p.capacity = "8kg"
	p.productdimensions = "9*10 mm"
	p.specialfeature = []string{"superspeed", "have water purifier"}

	p.washingprogrammes = make([]string, 22)
	p.inbox[0] = "pipe for water"
	p.inbox[1] = "blots"
	p.inbox[2] = "cover"

	p.price = new(int)

	if p.brand == "samsung" {
		*p.price = 23000

	} else if p.brand != "LG" {
		*p.price = 22000
	} else {

	}

	switch p.capacity {
	case "8 kg":
		p.noofclothes = 10
	case "10 kg":
		p.noofclothes = 20
	case "7 kg":
		p.noofclothes = 15
	default:

	}
	for i, v := range p.inbox {
		fmt.Println(i, v)

	}

}
func main() {

	washingmachine()
	//var anjali features
	var p *feature = new(feature)

	p.sellingcompanies()


}
