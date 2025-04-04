package main

import "fmt"

type restaurant struct {
	restaurantName    string
	restaurantId      int
	restaurantAddress string
	restaurantEmail   string
	restaurantPhone   int
	restaurantRating  int
	restaurantMenu    map[string]dishes
}

type dishes struct {
	restaurant_id   int
	restaurantName  string
	dishName        string
	dishId          int
	dishPrice       int
	dishDescription string
	dishRating      int
	dishType        string
}

var restaurant_data = make(map[string]*restaurant)
var dishesData = make(map[int]*dishes)

func register_Restaurant(
	restaurantName string,
	restaurantId int,
	restaurantAddress string,
	restaurantEmail string,
	restaurantPhone int,
	restaurantRating int,
) *restaurant {
	new_restaurant := &restaurant{
		restaurantName:    restaurantName,
		restaurantId:      restaurantId,
		restaurantAddress: restaurantAddress,
		restaurantEmail:   restaurantEmail,
		restaurantPhone:   restaurantPhone,
		restaurantRating:  restaurantRating,
		restaurantMenu:    make(map[string]dishes),
	}
	restaurant_data[restaurantName] = new_restaurant
	return new_restaurant
}

func addDish(
	restaurant_id int,
	restaurantName string,
	dishName string,
	dishId int,
	dishPrice int,
	dishDescription string,
	dishRating int,
	dishType string,
) *dishes {
	value, exists := restaurant_data[restaurantName]
	if !exists {
		fmt.Println("Restaurant not found")
		return nil
	}

	new_dish := &dishes{
		restaurant_id:   restaurant_id,
		restaurantName:  restaurantName,
		dishName:        dishName,
		dishId:          dishId,
		dishPrice:       dishPrice,
		dishDescription: dishDescription,
		dishRating:      dishRating,
		dishType:        dishType,
	}

	dishesData[dishId] = new_dish

	value.restaurantMenu[dishName] = *new_dish

	return new_dish
}
