package helper

import (
	"github.com/Alfeenn/online-learning/model"
	"github.com/Alfeenn/online-learning/model/web"
)

func ConvertModel(req model.User) web.CatResp {
	return web.CatResp{
		Id:        req.Id,
		Email:     req.Email,
		Password:  req.Password,
		Role:      req.Role,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
}
