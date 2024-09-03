package request

type UserRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int8   `json:"age" validate:"required"`
}
