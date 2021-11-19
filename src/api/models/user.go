package models

type BasicUser struct {
	ID          string `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Phone       uint32 `json:"phone"`
	Photo       string `json:"photo"`
	Is_verified bool   `json:"is_verified"`
}

type User struct {
	ID          string `json:"id" binding:"required"`
	Email       string `json:"email" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	Role        string `json:"role" binding:"required"`
	Photo       string `json:"photo"`
	Name        string `json:"name" binding:"required"`
	Dni         string `json:"dni"`
	Birthday    string `json:"birthday"`
	Phone       uint32 `json:"phone"`
	Description string `json:"description"`
	// Events      []Event `json:event`
	// Recipes			[]Recipe	`json:"recipe"`

	// Nutricionist
	Is_verified bool        `json:"is_verified"`
	Education   []string    `json:"education"`
	Specialties []string    `json:"specialties"`
	Price       float64     `json:"price"`
	Likes       uint64      `json:"likes"`
	Clients     []BasicUser `json:"clients"`

	// Client
	Weight        uint        `json:"weight"`
	Height        uint        `json:"height"`
	Type_diet     string      `json:"type_diet"`
	Intolerances  []string    `json:"intolerances"`
	Nutricionists []BasicUser `json:"nutricionists"`
}