package main

import (
	"math"
	"os"
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
	// TODO: print order + total?
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
		// TODO: update item/price lists
	}
	return o
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
		holderPriceList := []int64{}
		for i, item := range o.orderList {
			if item == "AP1" {
				holderItemList = append(holderItemList, item, "AAPL")
				holderPriceList = append(holderPriceList, int64(-150), o.priceList[i])
				o.total -= int64(150)
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
		holderPriceList := []int64{}
		for i, item := range o.orderList {
			// TODO: limit times it can be applied
			if item == "MK1" {
				holderItemList = append(holderItemList, item, "CHMK")
				holderPriceList = append(holderPriceList, int64(-475), o.priceList[i])
				o.total -= int64(475)
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
