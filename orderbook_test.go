package main

import (
	"fmt"
	"reflect"
	"testing"
)

func assert(t *testing.T, a, b any) {
	if !reflect.DeepEqual(a, b) {
		t.Errorf("%+v != %+v", a, b)
	}
}

func TestLimit(t *testing.T) {
	l := NewLimit(10_000)
	buyOrderB := NewOrder(true, 5)
	buyOrderA := NewOrder(true, 8)
	buyOrderC := NewOrder(true, 1)

	l.AddOrder(buyOrderB)
	l.AddOrder(buyOrderA)
	l.AddOrder(buyOrderC)

	l.DeleteOrder(buyOrderA)

	fmt.Println(l)
}

func TestPlaceLimitOrder(t *testing.T) {
	ob := NewOrderbook()

	sellOrderA := NewOrder(false, 10)
	sellOrderB := NewOrder(false, 5)

	ob.PlaceLimitOrder(10_000, sellOrderA)
	ob.PlaceLimitOrder(9_000, sellOrderB)

	assert(t, len(ob.asks), 2)
}
