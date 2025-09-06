package main

import (
	"fmt"
)

type Restaurant struct {
	name string
	id   int32
	menu map[string]*dish
}

type dish struct {
	id    int
	name  string
	price float64
	stock int
}

var restaurantData = make(map[string]*Restaurant)

type Customer struct {
	name     string
	address  string
	phone    string
	email    string
	id       int
	password string
	orders   map[string][]OrderHistory
}

type OrderHistory struct {
	id           int
	CustomerId   int
	RestaurantId int32
	dishes       map[int]*dish // Changed to *dish to match Restaurant.go
	totalAmount  float64
}

var CustomerData = make(map[int]*Customer)

func register(name, address, phone, email, password string, id int) bool {
	if _, exists := CustomerData[id]; exists {
		fmt.Println("Customer already exists with ID:", id)
		return false
	}

	c := Customer{
		name:     name,
		address:  address,
		phone:    phone,
		email:    email,
		id:       id,
		password: password,
		orders:   make(map[string][]OrderHistory), // Initialize the orders map
	}

	CustomerData[c.id] = &c
	return true
}

func searchRestaurant(Restaurantname string) (*Restaurant, bool) {
	if hotel, exists := restaurantData[Restaurantname]; exists {
		return hotel, true
	}
	return nil, false
}

func getCustomerById(id int) *Customer {
	return CustomerData[id]
}

func (r *Restaurant) searchDish(Restaurantname string, dishname string, nos int) (*dish, bool) {
	if _, exists := r.menu[dishname]; !exists {
		return nil, false
	}

	if r.menu[dishname].stock < nos {
		return nil, false
	}

	return r.menu[dishname], true
}

func (c *Customer) getDish(r *Restaurant, RestaurantName string, dishNames []string, nos int) (bool, float64) {
	total := 0.0
	found := false
	orderedDishes := make(map[int]*dish) // Changed to *dish

	for _, dn := range dishNames {
		if dish, ok := r.searchDish(RestaurantName, dn, nos); ok {
			total += dish.price * float64(nos)
			orderedDishes[dish.id] = dish
			found = true
		}
	}

	if found {
		orderID := len(c.orders[r.name]) + 1
		newOrder := OrderHistory{
			id:           orderID,
			CustomerId:   c.id,
			RestaurantId: r.id,
			dishes:       orderedDishes,
			totalAmount:  total,
		}

		if c.orders == nil {
			c.orders = make(map[string][]OrderHistory)
		}
		c.orders[r.name] = append(c.orders[r.name], newOrder) // Use r.name instead of r.id
		return true, total
	}

	return false, 0
}
