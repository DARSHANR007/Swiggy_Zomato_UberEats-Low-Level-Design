package main

type restaurant struct {
	name     string
	id       int32
	address  []string
	phone    string
	email    string
	password string
	menu     map[string]*dish
}

type dish struct {
	name         string
	id           int
	price        float64
	description  string
	restaurantId int
	available    bool
	category     string
	stock        int
	cuisine      string
	rating       float64
	reviews      []string
	tags         []string
}

var restaurantData = make(map[string]*restaurant)

func addRestaurant(address []string, name, phone, email, password string, id int32) (bool, string, *restaurant) {
	if _, exists := restaurantData[name]; exists {
		return false, "	restaurant already exists with ID: " + string(id), nil
	}

	r := &restaurant{
		name:     name,
		address:  address,
		phone:    phone,
		email:    email,
		id:       id,
		password: password,
		menu:     make(map[string]*dish),
	}

	restaurantData[r.name] = r
	return true, " restaurant created successfully", r

}

func (r *restaurant) addDishToMenu(restaurantId int, dishName string, dishId int, price float64, description, category, cuisine string, stock int) (bool, *dish) {

	if _, exists := r.menu[dishName]; exists {
		return false, nil

	}

	d := &dish{
		restaurantId: restaurantId,
		name:         dishName,
		id:           dishId,
		price:        price,
		description:  description,
		category:     category,
		cuisine:      cuisine,
		stock:        stock,
		available:    true,
		rating:       0.0,
		reviews:      []string{},
		tags:         []string{},
	}

	r.menu[dishName] = d

	return true, d

}

func (r *restaurant) recieveOrder(customerData customer, dishname string, nos int) (int, string) {

}
