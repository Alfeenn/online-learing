package web

type CategoryRequest struct {
	Id       string `json:"id"`
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role"`
}
