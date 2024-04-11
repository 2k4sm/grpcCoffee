package dto

type CoffeeHouse struct {
	Id               uint      `json:"id"`
	Name             string    `json:"house_name"`
	UserCount        int64     `json:"user_count"`
	TopCoffee        string    `json:"top_coffee"`
	Revenue          int64     `json:"revenue"`
	AvailableCoffees []Coffee  `json:"coffees"`
	Payments         []Payment `json:"payments"`
}

type CreateCoffeeHouse struct {
	Name             string   `json:"house_name"`
	AvailableCoffees []string `json:"coffees"`
}
