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
	orders   map[int][]orderHistory
}

type orderHistory struct {
	id           int
	customerId   int
	restaurantId int
	dishes       map[int]dish
	totalAmount  float64
}

var customerData = make(map[int]customer)

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
	}

	customerData[c.id] = c
	return true
}

func searchRestaurant(restaurantname string) (*restaurant, bool) {
	if hotel, exists := restaurantData[restaurantname]; exists {

		return hotel, true
	}

	return nil, false
}

func getCustomerById(id int) *customer {
	c, ok := customerData[id]
	if !ok {
		return nil
	}
	return &c
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

func (r *restaurant) getDish(restaurantname string, dishname string, nos int) (bool, int) {

	if dish, check := r.searchDish(restaurantname, dishname, nos); check {
		return true, int(dish.price)
	}

	return false, 0

}
