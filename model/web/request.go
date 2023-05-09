package web

type CategoryRequest struct {
	Id       string `json:"id"`
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Name     string `form:"name" json:"Name" binding:"required"`
	Age      int64  `form:"age" json:"Age" binding:"required"`
	Phone    int64  `form:"phone" json:"Phone" binding:"required"`
	Role     string `form:"role" json:"role"`
}
