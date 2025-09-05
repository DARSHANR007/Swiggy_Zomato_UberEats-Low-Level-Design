package main

import (
	"fmt"
)

type customer struct {
	name     string
	address  string
	phone    string
	email    string
	id       int
	password string
	orders   map[string][]orderHistory
}

type orderHistory struct {
	id           int
	customerId   int
	restaurantId int32
	dishes       map[int]*dish // Changed to *dish to match restaurant.go
	totalAmount  float64
}

var customerData = make(map[int]*customer)

func register(name, address, phone, email, password string, id int) bool {
	if _, exists := customerData[id]; exists {
		fmt.Println("Customer already exists with ID:", id)
		return false
	}

	c := customer{
		name:     name,
		address:  address,
		phone:    phone,
		email:    email,
		id:       id,
		password: password,
		orders:   make(map[string][]orderHistory), // Initialize the orders map
	}

	customerData[c.id] = &c
	return true
}

func searchRestaurant(restaurantname string) (*restaurant, bool) {
	if hotel, exists := restaurantData[restaurantname]; exists {
		return hotel, true
	}
	return nil, false
}

func getCustomerById(id int) *customer {
	return customerData[id]
}

func (r *restaurant) searchDish(restaurantname string, dishname string, nos int) (*dish, bool) {
	if _, exists := r.menu[dishname]; !exists {
		return nil, false
	}

	if r.menu[dishname].stock < nos {
		return nil, false
	}

	return r.menu[dishname], true
}

func (c *customer) getDish(r *restaurant, restaurantName string, dishNames []string, nos int) (bool, float64) {
	total := 0.0
	found := false
	orderedDishes := make(map[int]*dish) // Changed to *dish

	for _, dn := range dishNames {
		if dish, ok := r.searchDish(restaurantName, dn, nos); ok {
			total += dish.price * float64(nos)
			orderedDishes[dish.id] = dish
			found = true
		}
	}

	if found {
		orderID := len(c.orders[r.name]) + 1
		newOrder := orderHistory{
			id:           orderID,
			customerId:   c.id,
			restaurantId: r.id,
			dishes:       orderedDishes,
			totalAmount:  total,
		}

		if c.orders == nil {
			c.orders = make(map[string][]orderHistory)
		}
		c.orders[r.name] = append(c.orders[r.name], newOrder) // Use r.name instead of r.id
		return true, total
	}

	return false, 0
}
