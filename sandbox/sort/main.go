package main

import "fmt"
import "sort"

type food struct {
	Name  string
	Price int
}

func main() {
	foods := make([]food, 3)
	foods[0] = food{Name: "みかん", Price: 150}
	foods[1] = food{Name: "バナナ", Price: 100}
	foods[2] = food{Name: "りんご", Price: 120}

	sort.Slice(foods, func(i, j int) bool {
		return foods[i].Price < foods[j].Price
	})

	for _, food := range foods {
		fmt.Printf("%+v\n", food)
	}
}
