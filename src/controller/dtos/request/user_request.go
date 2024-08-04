package request

type UserRequest struct {
	Email    string `json:"email" binding:"required,email" example:"test@test.com"`
	Password string `json:"password" binding:"required,min=6,containsany=!@#$%*" example:"password#@#@!2121"`
	Name     string `json:"name" binding:"required,min=3,max=100"`
	Age      int8   `json:"age" binding:"required,min=1,max=140" example:"30"`
}
