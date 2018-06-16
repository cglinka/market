package main

import (
	"fmt"
	"os"
)

// product contains the information about each product
type product struct {
	name  string
	price float32
	code  string
}

// Order is a map of product codes to quantities
type order struct {
	// items maps the product code to the quantity of that item.
	orderList []string
	priceList []float32
	items     map[string]int
	total     float32
}

// discount contains the info about available discounts
type discount struct {
	discountCode string
	discount     float32
	// qualification is a map of the required products and their quantities
	qualification map[string]int
}

// func applyDiscount(o order, d []discount) (order, error) {
// 	// iterate over available discounts
// 	// update total
// }

func main() {

	var prods = make(map[string]product)

	prods["CH1"] = product{
		name:  "Chai",
		price: 3.11,
		code:  "CH1",
	}
	prods["AP1"] = product{
		name:  "Apples",
		price: 6.00,
		code:  "AP1",
	}
	prods["CF1"] = product{
		name:  "Coffee",
		price: 11.23,
		code:  "CF1",
	}
	prods["MK1"] = product{
		name:  "Milk",
		price: 4.75,
		code:  "MK1",
	}
	prods["OM1"] = product{
		name:  "Oatmeal",
		price: 3.69,
		code:  "OM1",
	}

	discounts := make(map[string]discount)
	// 1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
	// 2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
	// 3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
	// 4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples
	discounts["BOGO"] = discount{
		discountCode:  "BOGO",
		discount:      -11.23,
		qualification: map[string]int{"CF1": 2},
	}
	discounts["AAPL"] = discount{
		discountCode:  "AAPL",
		discount:      -1.5,
		qualification: map[string]int{"AP1": 3},
	}
	discounts["CHMK"] = discount{
		discountCode:  "CHMK",
		discount:      -4.75,
		qualification: map[string]int{"MK1": 1, "CH1": 1},
	}
	discounts["APOM"] = discount{
		discountCode:  "APOM",
		discount:      -3,
		qualification: map[string]int{"OM1": 1, "AP1": 1},
	}

	// CH1, AP1, AP1, AP1, MK1
	o := &order{
		items: map[string]int{},
	}
	o.orderList = os.Args[1:]
	// o.orderList = []string{"CH1", "AP1", "AP1", "AP1", "MK1"}

	fmt.Printf("%+v\n", o)
	for _, code := range o.orderList {
		// fmt.Printf("%+v\n", o)
		// add item to order
		if _, ok := o.items[code]; ok {
			o.items[code] += 1
		} else {
			o.items[code] = 1
		}
		o.priceList = append(o.priceList, prods[code].price)
		o.total += prods[code].price
	}
	fmt.Printf("subtotal: %d\n", o.total)
	fmt.Printf("%+v\n", o)

	o = aapl(o)
	o = chmk(o)
	fmt.Printf("%+v\n", o)
}

// 2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
func aapl(o *order) *order {
	// d := discount{
	// 	discountCode:  "AAPL",
	// 	discount:      -1.5,
	// 	qualification: map[string]int{"AP1": 3},
	// }
	if o.items["AP1"] >= 3 {
		holderItemList := []string{}
		holderPriceList := []float32{}
		for i, item := range o.orderList {
			if item == "AP1" {
				holderItemList = append(holderItemList, item, "AAPL")
				holderPriceList = append(holderPriceList, float32(-1.5), o.priceList[i])
				o.total -= 1.5
			} else {
				holderItemList = append(holderItemList, item)
				holderPriceList = append(holderPriceList, o.priceList[i])
			}
		}
		o.orderList = holderItemList
		o.priceList = holderPriceList
	}
	return o
}

// 3. CHMK -- Purchase a box of Chai and get milk free. (Limit 1)
func chmk(o *order) *order {
	// discount{
	// 	discountCode:  "CHMK",
	// 	discount:      -4.75,
	// 	qualification: map[string]int{"MK1": 1, "CH1": 1},
	// }

	if o.items["MK1"] >= 1 && o.items["CH1"] >= 1 {
		holderItemList := []string{}
		holderPriceList := []float32{}
		for i, item := range o.orderList {
			if item == "MK1" {
				holderItemList = append(holderItemList, item, "CHMK")
				holderPriceList = append(holderPriceList, float32(-4.75), o.priceList[i])
				o.total -= 1.5
			} else {
				holderItemList = append(holderItemList, item)
				holderPriceList = append(holderPriceList, o.priceList[i])
			}
		}
		o.orderList = holderItemList
		o.priceList = holderPriceList
	}
	return o
}
