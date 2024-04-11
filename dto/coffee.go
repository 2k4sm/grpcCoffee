package dto

type MultiString []string

type Coffee struct {
	Id          uint        `json:"id"`
	Name        string      `json:"coffee_name"`
	Description string      `json:"description"`
	Origin      string      `json:"origin"`
	Contents    MultiString `json:"contents"`
	Cost        int64       `json:"cost"`
}

type CreateCoffee struct {
	Name        string      `json:"coffee_name"`
	Description string      `json:"description"`
	Origin      string      `json:"origin"`
	Contents    MultiString `json:"contents"`
	Cost        int64       `json:"cost"`
}
