package kit

import (
	"fmt"
	"testing"
)

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for n, ia := range a {
		ib := b[n]
		if ia != ib {
			return false
		}
	}
	return true
}

func TestNewTrafficLight_isGreen(t *testing.T) {
	light, err := NewTrafficLight(&TrafficLightOption{1, 3})
	if err != nil {
		t.Fatal(err)
	}
	green := []int{0, 3, 6, 9}
	g1 := make([]int, 0)
	for i := 0; i < 10; i++ {
		if light.IsGreen() {
			g1 = append(g1, i)
		}
	}
	if !equal(green, g1) {
		t.Fatal(green, g1)
	}
	// 2
	light, err = NewTrafficLight(&TrafficLightOption{2, 3})
	if err != nil {
		t.Fatal(err)
	}
	green = []int{0, 1, 3, 4, 6, 7, 9}
	g1 = make([]int, 0)
	for i := 0; i < 10; i++ {
		if light.IsGreen() {
			g1 = append(g1, i)
		}
	}
	if !equal(green, g1) {
		t.Fatal(green, g1)
	}
	// 3
	light, err = NewTrafficLight(&TrafficLightOption{8, 10})
	if err != nil {
		t.Fatal(err)
	}
	green = []int{0, 1, 2, 3, 4, 5, 6, 7}
	g1 = make([]int, 0)
	for i := 0; i < 10; i++ {
		if light.IsGreen() {
			g1 = append(g1, i)
		}
	}
	if !equal(green, g1) {
		t.Fatal(green, g1)
	}
	// 4
	fmt.Println("4......")
	light, err = NewTrafficLight(&TrafficLightOption{1, 2})
	if err != nil {
		t.Fatal(err)
	}
	green = []int{0, 2, 4, 6, 8}
	g1 = make([]int, 0)
	for i := 0; i < 10; i++ {
		if light.IsGreen() {
			g1 = append(g1, i)
		}
	}
	if !equal(green, g1) {
		t.Fatal(green, g1)
	}
	// 5
	fmt.Println("5......")
	light, err = NewTrafficLight(&TrafficLightOption{1, 1})
	if err != nil {
		t.Fatal(err)
	}
	green = []int{0, 1, 2, 3, 4}
	g1 = make([]int, 0)
	for i := 0; i < 5; i++ {
		if light.IsGreen() {
			g1 = append(g1, i)
		}
	}
	if !equal(green, g1) {
		t.Fatal(green, g1)
	}
}
