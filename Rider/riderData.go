package rider

import (
	"food_delivery/customer"
	"food_delivery/restaurant"
)

type Rider struct {
	Name        string
	ID          int
	Address     string
	Available   bool
	Phone       string
	Rating      float64
	RideHistory map[string][]Order
}

type Order struct {
	CustomerInfo   *customer.Customer
	RestaurantInfo *restaurant.Restaurant
	OrderAddress   string
	Price          float64
}
