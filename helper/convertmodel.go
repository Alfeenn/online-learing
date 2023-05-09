package helper

import (
	"github.com/Alfeenn/online-learning/model"
	"github.com/Alfeenn/online-learning/model/web"
)

func ConvertModel(req model.User) web.CatResp {
	return web.CatResp{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Age:      req.Age,
		Phone:    req.Phone,
		Role:     req.Role,
	}
}
