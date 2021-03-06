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
				priceList: []int64{311, 600, 600, 600, 475},
				items:     map[string]int{"CH1": 1, "AP1": 3, "MK1": 1},
				total:     2586,
			},
		},
		{
			list: []string{"CH1", "AP1", "CF1", "MK1"},
			expectedOrder: &order{
				orderList: []string{"CH1", "AP1", "CF1", "MK1"},
				priceList: []int64{311, 600, 1123, 475},
				items:     map[string]int{"CH1": 1, "AP1": 1, "CF1": 1, "MK1": 1},
				total:     2509,
			},
		},
		{
			list: []string{"MK1", "AP1"},
			expectedOrder: &order{
				orderList: []string{"MK1", "AP1"},
				priceList: []int64{475, 600},
				items:     map[string]int{"AP1": 1, "MK1": 1},
				total:     1075,
			},
		},
		{
			list: []string{"CF1", "CF1"},
			expectedOrder: &order{
				orderList: []string{"CF1", "CF1"},
				priceList: []int64{1123, 1123},
				items:     map[string]int{"CF1": 2},
				total:     2246,
			},
		},
		{
			list: []string{"AP1", "AP1", "CH1", "AP1"},
			expectedOrder: &order{
				orderList: []string{"AP1", "AP1", "CH1", "AP1"},
				priceList: []int64{600, 600, 311, 600},
				items:     map[string]int{"CH1": 1, "AP1": 3},
				total:     2111,
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

func TestGivenCases(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal int64
	}{
		{
			name: "CH1 AP1 AP1 AP1 MK1",
			o: &order{
				orderList: []string{"CH1", "AP1", "AP1", "AP1", "MK1"},
				priceList: []int64{311, 600, 600, 600, 475},
				items:     map[string]int{"CH1": 1, "AP1": 3, "MK1": 1},
				total:     2586,
			},
			expectedTotal: 1661,
		},
		{
			name: "CH1, AP1, CF1, MK1",
			o: &order{
				orderList: []string{"CH1", "AP1", "CF1", "MK1"},
				priceList: []int64{311, 600, 1123, 475},
				items:     map[string]int{"CH1": 1, "AP1": 1, "CF1": 1, "MK1": 1},
				total:     2509,
			},
			expectedTotal: 2034,
		},
		{
			name: "MK1 AP1",
			o: &order{
				orderList: []string{"MK1", "AP1"},
				priceList: []int64{475, 600},
				items:     map[string]int{"AP1": 1, "MK1": 1},
				total:     1075,
			},
			expectedTotal: 1075,
		},
		{
			name: "CF1 CF1",
			o: &order{
				orderList: []string{"CF1", "CF1"},
				priceList: []int64{1123, 1123},
				items:     map[string]int{"CF1": 2},
				total:     2246,
			},
			expectedTotal: 1123,
		},
		{
			name: "AP1 AP1 CH1 AP1",
			o: &order{
				orderList: []string{"AP1", "AP1", "CH1", "AP1"},
				priceList: []int64{600, 600, 311, 600},
				items:     map[string]int{"CH1": 1, "AP1": 3},
				total:     2111,
			},
			expectedTotal: 1661,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := bogo(tst.o)
			ordr = aapl(ordr)
			ordr = chmk(ordr)
			ordr = apom(ordr)
			if !reflect.DeepEqual(ordr.total, tst.expectedTotal) {
				t.Errorf("Order total of %v did not match expected order total of %v", ordr.total, tst.expectedTotal)
			}
		})
	}
}

func TestBOGODiscount(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal int64
	}{
		{
			name: "CF1 CF1",
			o: &order{
				orderList: []string{"CF1", "CF1"},
				priceList: []int64{1123, 1123},
				items:     map[string]int{"CF1": 2},
				total:     2246,
			},
			expectedTotal: 1123,
		},
		{
			name: "CF1 CF1 CF1",
			o: &order{
				orderList: []string{"CF1", "CF1", "CF1"},
				priceList: []int64{1123, 1123, 1123},
				items:     map[string]int{"CF1": 3},
				total:     3369,
			},
			expectedTotal: 2246,
		},
		{
			name: "CF1 CF1 CF1 CF1",
			o: &order{
				orderList: []string{"CF1", "CF1", "CF1", "CF1"},
				priceList: []int64{1123, 1123, 1123, 1123},
				items:     map[string]int{"CF1": 4},
				total:     4492,
			},
			expectedTotal: 2246,
		},
		{
			name: "CF1",
			o: &order{
				orderList: []string{"CF1"},
				priceList: []int64{1123},
				items:     map[string]int{"CF1": 1},
				total:     1123,
			},
			expectedTotal: 1123,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := bogo(tst.o)
			if ordr.total != tst.expectedTotal {
				t.Errorf("Expected tota: %v, recieved total: %v", tst.expectedTotal, ordr.total)
			}
		})
	}
}

func TestAAPLDiscount(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal int64
	}{
		{
			name: "AP1",
			o: &order{
				orderList: []string{"AP1"},
				priceList: []int64{600},
				items:     map[string]int{"AP1": 1},
				total:     600,
			},
			expectedTotal: 600,
		},
		{
			name: "AP1 AP1 AP1",
			o: &order{
				orderList: []string{"AP1", "AP1", "AP1"},
				priceList: []int64{600, 600, 600},
				items:     map[string]int{"AP1": 3},
				total:     1800,
			},
			expectedTotal: 1350,
		},
		{
			name: "AP1 AP1 AP1",
			o: &order{
				orderList: []string{"AP1", "AP1", "AP1"},
				priceList: []int64{600, 600, 600},
				items:     map[string]int{"AP1": 3},
				total:     1800,
			},
			expectedTotal: 1350,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := aapl(tst.o)
			if ordr.total != tst.expectedTotal {
				t.Errorf("Expected tota: %v, recieved total: %v", tst.expectedTotal, ordr.total)
			}
		})
	}
}

func TestCHMKDiscount(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal int64
	}{
		{
			name: "CH1 MK1",
			o: &order{
				orderList: []string{"CH1", "MK1"},
				priceList: []int64{311, 475},
				items:     map[string]int{"CH1": 1, "MK1": 1},
				total:     786,
			},
			expectedTotal: 311,
		},
		{
			name: "CH1 MK1 CH1 MK1",
			o: &order{
				orderList: []string{"CH1", "MK1", "CH1", "MK1"},
				priceList: []int64{311, 475, 311, 475},
				items:     map[string]int{"CH1": 2, "MK1": 2},
				total:     1572,
			},
			expectedTotal: 1097,
		},
		{
			name: "CH1 MK1 CH1",
			o: &order{
				orderList: []string{"CH1", "MK1", "CH1"},
				priceList: []int64{311, 475, 311},
				items:     map[string]int{"CH1": 2, "MK1": 1},
				total:     1097,
			},
			expectedTotal: 622,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := chmk(tst.o)
			if ordr.total != tst.expectedTotal {
				t.Errorf("Expected tota: %v, recieved total: %v", tst.expectedTotal, ordr.total)
			}
		})
	}
}

func TestAPOMDiscount(t *testing.T) {
	tests := []struct {
		name          string
		o             *order
		expectedTotal int64
	}{
		{
			name: "OM1 AP1",
			o: &order{
				orderList: []string{"OM1", "AP1"},
				priceList: []int64{369, 600},
				items:     map[string]int{"OM1": 1, "AP1": 1},
				total:     369 + 600,
			},
			expectedTotal: 669,
		},
		{
			name: "OM1 OM1 AP1",
			o: &order{
				orderList: []string{"OM1", "OM1", "AP1"},
				priceList: []int64{369, 369, 600},
				items:     map[string]int{"OM1": 2, "AP1": 1},
				total:     1338,
			},
			expectedTotal: 1038,
		},
		{
			name: "OM1 OM1 AP1 AP1",
			o: &order{
				orderList: []string{"OM1", "OM1", "AP1", "AP1"},
				priceList: []int64{369, 369, 600, 600},
				items:     map[string]int{"OM1": 2, "AP1": 2},
				total:     1938,
			},
			expectedTotal: 1338,
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			ordr := apom(tst.o)
			if ordr.total != tst.expectedTotal {
				t.Errorf("Expected tota: %v, recieved total: %v", tst.expectedTotal, ordr.total)
			}
		})
	}
}
