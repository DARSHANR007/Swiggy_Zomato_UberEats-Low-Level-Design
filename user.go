package main

import "fmt"

type user struct {
	userName    string
	userId      int
	userAddress string
	userEmail   string
	user_phone  int
}

var userData map[int]*user

func Create_NewCustomer(userName string, userId int, userAddress string, userEmail string, user_phone int) *user {

	newuser := &user{
		userName:    userName,
		userId:      userId,
		userAddress: userAddress,
		userEmail:   userEmail,
		user_phone:  user_phone,
	}

	userData[userId] = newuser

	return newuser
}

func (this *user) GetUserData(id int) *user {
	return userData[id]
}

func (u *user) order_dish(restaurant_name string, dish_names []string) {

	res_value, exists := restaurant_data[restaurant_name]

	if exists == false {
		fmt.Println("Restaurant not found")
		return

	}

	for _, dish_name := range dish_names {
		_, exists := res_value.restaurantMenu[dish_name]
		if exists == false {
			fmt.Println("Dish not found")
			return
		}
	}
	generate_bill(res_value, dish_names, u)
}
