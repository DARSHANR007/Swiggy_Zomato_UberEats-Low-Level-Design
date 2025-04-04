package main

import "time"

type bill struct {
	restaurant_name string
	total_price     int
	dish_names      []string
	customer_name   string
	time            time.Time
	total           int
}

func generate_bill(res_value *restaurant, dish_names []string, customer_data *user) *bill {

	var final_total int

	for _, x := range dish_names {
		final_total += res_value.restaurantMenu[x].dishPrice
	}

	new_bill := &bill{
		restaurant_name: res_value.restaurantName,
		total_price:     final_total,
		dish_names:      dish_names,
		customer_name:   customer_data.userName,
		time:            time.Now(),
	}

	return new_bill

}
