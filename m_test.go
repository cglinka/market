package main

import (
	"reflect"
	"testing"
)

func TestBuildOrder(t *testing.T) {
	tests := []struct {
		list          []string
		expectedOrder *order
	}{
		{
			list: []string{"CH1", "AP1", "AP1", "AP1", "MK1"},
			expectedOrder: &order{
				orderList: []string{"CH1", "AP1", "AP1", "AP1", "MK1"},
				priceList: []float32{3.11, 6, 6, 6, 4.75},
				items:     map[string]int{"CH1": 1, "AP1": 3, "MK1": 1},
				total:     25.86,
			},
		},
		{
			list: []string{"CH1", "AP1", "CF1", "MK1"},
			expectedOrder: &order{
				orderList: []string{"CH1", "AP1", "CF1", "MK1"},
				priceList: []float32{3.11, 6, 11.23, 4.75},
				items:     map[string]int{"CH1": 1, "AP1": 1, "CF1": 1, "MK1": 1},
				total:     25.09,
			},
		},
		{
			list: []string{"MK1", "AP1"},
			expectedOrder: &order{
				orderList: []string{"MK1", "AP1"},
				priceList: []float32{4.75, 6},
				items:     map[string]int{"AP1": 1, "MK1": 1},
				total:     10.75,
			},
		},
		{
			list: []string{"CF1", "CF1"},
			expectedOrder: &order{
				orderList: []string{"CF1", "CF1"},
				priceList: []float32{11.23, 11.23},
				items:     map[string]int{"CF1": 2},
				total:     22.46,
			},
		},
		{
			list: []string{"AP1", "AP1", "CH1", "AP1"},
			expectedOrder: &order{
				orderList: []string{"AP1", "AP1", "CH1", "AP1"},
				priceList: []float32{6, 6, 3.11, 6},
				items:     map[string]int{"CH1": 1, "AP1": 3},
				total:     21.11,
			},
		},
	}

	for _, tst := range tests {
		o := buildOrder(tst.list)
		if !reflect.DeepEqual(o, tst.expectedOrder) {
			t.Errorf("Order did not match expected order: %+v, %+v", o, tst.expectedOrder)
		}
	}
}

func TestAllDiscounts(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal float32
	}{
		{
			name: "CH1 AP1 AP1 AP1 MK1",
			o: &order{
				orderList: []string{"CH1", "AP1", "AP1", "AP1", "MK1"},
				priceList: []float32{3.11, 6, 6, 6, 4.75},
				items:     map[string]int{"CH1": 1, "AP1": 3, "MK1": 1},
				total:     25.86,
			},
			expectedTotal: 16.61,
		},
		{
			name: "CH1, AP1, CF1, MK1",
			o: &order{
				orderList: []string{"CH1", "AP1", "CF1", "MK1"},
				priceList: []float32{3.11, 6, 11.23, 4.75},
				items:     map[string]int{"CH1": 1, "AP1": 1, "CF1": 1, "MK1": 1},
				total:     25.09,
			},
			expectedTotal: 20.34,
		},
		{
			name: "MK1 AP1",
			o: &order{
				orderList: []string{"MK1", "AP1"},
				priceList: []float32{4.75, 6},
				items:     map[string]int{"AP1": 1, "MK1": 1},
				total:     10.75,
			},
			expectedTotal: 10.75,
		},
		{
			name: "CF1 CF1",
			o: &order{
				orderList: []string{"CF1", "CF1"},
				priceList: []float32{11.23, 11.23},
				items:     map[string]int{"CF1": 2},
				total:     22.46,
			},
			expectedTotal: 11.23,
		},
		{
			name: "AP1 AP1 CH1 AP1",
			o: &order{
				orderList: []string{"AP1", "AP1", "CH1", "AP1"},
				priceList: []float32{6, 6, 3.11, 6},
				items:     map[string]int{"CH1": 1, "AP1": 3},
				total:     21.11,
			},
			expectedTotal: 16.61,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := bogo(tst.o)
			ordr = aapl(ordr)
			ordr = chmk(ordr)
			if !reflect.DeepEqual(ordr.total, tst.expectedTotal) {
				t.Errorf("Order total of %v did not match expected order total of %v", ordr.total, tst.expectedTotal)
			}
		})
	}
}
