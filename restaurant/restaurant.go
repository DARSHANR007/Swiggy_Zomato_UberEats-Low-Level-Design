package restaurant

type Restaurant struct {
	Name     string
	ID       int32
	Address  []string
	Phone    string
	Email    string
	Password string
	Menu     map[string]*Dish
}

type Dish struct {
	Name         string
	ID           int
	Price        float64
	Description  string
	RestaurantID int
	Available    bool
	Category     string
	Stock        int
	Cuisine      string
	Rating       float64
	Reviews      []string
	Tags         []string
}

var RestaurantData = make(map[string]*Restaurant)

func AddRestaurant(address []string, name, phone, email, password string, id int32) (bool, string, *Restaurant) {
	if _, exists := RestaurantData[name]; exists {
		return false, "Restaurant already exists with ID: " + string(id), nil
	}

	r := &Restaurant{
		Name:     name,
		Address:  address,
		Phone:    phone,
		Email:    email,
		ID:       id,
		Password: password,
		Menu:     make(map[string]*Dish),
	}

	RestaurantData[r.Name] = r
	return true, "Restaurant created successfully", r
}

func (r *Restaurant) AddDishToMenu(restaurantId int, dishName string, dishId int, price float64, description, category, cuisine string, stock int) (bool, *Dish) {
	if _, exists := r.Menu[dishName]; exists {
		return false, nil
	}

	d := &Dish{
		RestaurantID: restaurantId,
		Name:         dishName,
		ID:           dishId,
		Price:        price,
		Description:  description,
		Category:     category,
		Cuisine:      cuisine,
		Stock:        stock,
		Available:    true,
		Rating:       0.0,
		Reviews:      []string{},
		Tags:         []string{},
	}

	r.Menu[dishName] = d
	return true, d
}
