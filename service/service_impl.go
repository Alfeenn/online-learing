package service

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/online-learning/helper"
	"github.com/Alfeenn/online-learning/model"
	"github.com/Alfeenn/online-learning/model/web"
	"github.com/Alfeenn/online-learning/repository"
)

type ServiceImpl struct {
	Rep repository.Repository
	DB  *sql.DB
}

func NewService(c repository.Repository, DB *sql.DB) Service {
	return &ServiceImpl{
		Rep: c,
		DB:  DB,
	}
}

func (s *ServiceImpl) Create(ctx context.Context, req web.CategoryRequest) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	request := model.User{
		Id:       req.Id,
		Username: req.Username,
		Password: req.Password,
		Name:     req.Name,
		Age:      req.Age,
		Phone:    req.Phone,
		Role:     req.Role,
	}

	User := s.Rep.Create(ctx, tx, request)

	return helper.ConvertModel(User)

}

func (s *ServiceImpl) Update(ctx context.Context, req web.UpdateRequest) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	id := req.Id
	findId, err := s.Rep.Find(ctx, tx, id)
	helper.PanicIfErr(err)
	updateArticle := s.Rep.Update(ctx, tx, findId)
	return helper.ConvertModel(updateArticle)
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.Find(ctx, tx, id)
	helper.PanicIfErr(err)
	s.Rep.Delete(ctx, tx, req.Id)

}

func (s *ServiceImpl) Find(ctx context.Context, id string) web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.Find(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	return helper.ConvertModel(model)

}

func (s *ServiceImpl) FindAll(ctx context.Context) []web.CatResp {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	slicemodel := s.Rep.FindAll(ctx, tx)

	var webResp []web.CatResp

	for _, v := range slicemodel {
		webResp = append(webResp, helper.ConvertModel(v))
	}
	return webResp
}

func (s *ServiceImpl) Login(ctx context.Context, request web.CategoryRequest) web.CatResp {
	tx, err := s.DB.Begin()
	if err != nil {
		panic(err)
	}
	defer helper.CommitorRollback(tx)
	category := model.User{
		Username: request.Username,
		Password: request.Password,
	}
	category, err = s.Rep.Login(ctx, tx, category)
	if err != nil {

		panic(err.Error())
	}

	return helper.ConvertModel(category)

}
