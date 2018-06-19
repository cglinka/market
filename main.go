package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

// product contains the information about each product
type product struct {
	name  string
	price int64
	code  string
}

// Order is a map of product codes to quantities
type order struct {
	// items maps the product code to the quantity of that item.
	orderList []string
	priceList []int64
	items     map[string]int
	total     int64
}

var prods = make(map[string]product)

func init() {
	prods["CH1"] = product{
		name:  "Chai",
		price: 311,
		code:  "CH1",
	}
	prods["AP1"] = product{
		name:  "Apples",
		price: 600,
		code:  "AP1",
	}
	prods["CF1"] = product{
		name:  "Coffee",
		price: 1123,
		code:  "CF1",
	}
	prods["MK1"] = product{
		name:  "Milk",
		price: 475,
		code:  "MK1",
	}
	prods["OM1"] = product{
		name:  "Oatmeal",
		price: 369,
		code:  "OM1",
	}
}
func main() {
	// Build order with command line args
	l := os.Args[1:]
	o := buildOrder(l)

	// apply discounts
	o = bogo(o)
	o = aapl(o)
	o = chmk(o)
	o = apom(o)

	// Print out basket and order total
	fmt.Printf("Basket: %v\n", strings.Join(o.orderList, ", "))
	fmt.Printf("Total: %v\n", makeMoneyString(o.total))
}

func buildOrder(ol []string) *order {
	o := &order{
		items: map[string]int{},
	}
	o.orderList = ol

	for _, code := range o.orderList {
		// add item to order
		if _, ok := o.items[code]; ok {
			o.items[code]++
		} else {
			o.items[code] = 1
		}
		o.priceList = append(o.priceList, prods[code].price)
		o.total += prods[code].price
	}
	return o
}

// makeMoneyString converts the 'cents' into dollars and adds a dollar sign.
func makeMoneyString(price int64) string {
	p := float64(price) / float64(100)
	return fmt.Sprintf("$%.2f", p)
}

// 1. BOGO -- Buy-One-Get-One-Free Special on Coffee. (Unlimited)
func bogo(o *order) *order {
	// discount{
	// 	discountCode:  "BOGO",
	// 	discount:      -11.23,
	// 	qualification: map[string]int{"CF1": 2},
	// }

	if o.items["CF1"] >= 2 {
		numDiscounts := o.items["CF1"] / 2
		o.total -= (int64(numDiscounts) * int64(1123))
	}
	return o
}

// 2. APPL -- If you buy 3 or more bags of Apples, the price drops to $4.50.
func aapl(o *order) *order {
	// discount{
	// 	discountCode:  "AAPL",
	// 	discount:      -1.5,
	// 	qualification: map[string]int{"AP1": 3},
	// }
	if o.items["AP1"] >= 3 {
		for _, item := range o.orderList {
			if item == "AP1" {
				o.total -= int64(150)
			}
		}
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
		for _, item := range o.orderList {
			if item == "MK1" {
				o.total -= int64(475)
				break
			}
		}
	}
	return o
}

// 4. APOM -- Purchase a bag of Oatmeal and get 50% off a bag of Apples
// Interpretation: 1 bag of oatmeal = 1 apple discount, 2 oatmeal = 2 apple discount
func apom(o *order) *order {
	// discount{
	// 	discountCode:  "APOM",
	// 	discount:      -3,
	// 	qualification: map[string]int{"OM1": 1, "AP1": 1},
	// }
	numOM := float64(o.items["OM1"])
	numAP := float64(o.items["AP1"])
	if numOM >= 1 && numAP >= 1 {
		if numAP == numOM {
			o.total -= int64(numOM) * 300
		} else {
			numDiscounts := math.Min(numOM, numAP)
			o.total -= int64(numDiscounts) * 300
		}
	}
	return o
}
