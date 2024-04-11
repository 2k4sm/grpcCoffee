package dto

type User struct {
	Id            uint          `json:"id"`
	Name          string        `json:"user_name"`
	Email         string        `json:"email"`
	Password      string        `json:"password"`
	LastOrder     string        `json:"last_order"`
	Favourite     string        `json:"favourite"`
	Revenue       string        `json:"-"`
	Orders        []Payment     `json:"orders"`
	VisitedHouses []CoffeeHouse `json:"visited_houses"`
}

type CreateUser struct {
	Name     string `json:"user_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
