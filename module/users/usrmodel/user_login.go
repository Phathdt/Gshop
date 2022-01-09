package usrmodel

type UserLogin struct {
	Username string `json:"username" form:"username" validate:"required,min=6,max=32"`
	Password string `json:"password" form:"password" validate:"required,min=6,max=32"`
}
