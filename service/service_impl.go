package service

import (
	"context"
	"database/sql"
	"log"

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

func (s *ServiceImpl) CreateCourse(ctx context.Context, req model.Course) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	request := model.Course{
		Id:        req.Id,
		Name:      req.Name,
		Price:     req.Price,
		Category:  req.Category,
		Thumbnail: req.Thumbnail,
	}
	Course := s.Rep.CreateCourse(ctx, tx, request)

	return Course

}

func (s *ServiceImpl) Update(ctx context.Context, req model.Course) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	id := req.Id
	findId, err := s.Rep.FindCourseById(ctx, tx, id)
	helper.PanicIfErr(err)
	updateCourse := s.Rep.Update(ctx, tx, findId)
	return updateCourse
}

func (s *ServiceImpl) DeleteUser(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	s.Rep.Delete(ctx, tx, id)
}

func (s *ServiceImpl) Delete(ctx context.Context, id string) {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	req, err := s.Rep.FindCourseById(ctx, tx, id)
	s.Rep.DeleteCourse(ctx, tx, req.Id)
	helper.PanicIfErr(err)
	s.Rep.Delete(ctx, tx, req.Id)

}

func (s *ServiceImpl) FindCourseById(ctx context.Context, id string) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindCourseById(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	return model

}

func (s *ServiceImpl) FindCourseByCategory(ctx context.Context, id string) model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	model, err := s.Rep.FindCourseByCategory(ctx, tx, id)
	if err != nil {
		panic(err)
	}
	return model

}

func (s *ServiceImpl) FindAll(ctx context.Context) []model.Course {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)

	slicemodel := s.Rep.FindAll(ctx, tx)

	var sliceCourse []model.Course

	for _, v := range slicemodel {
		sliceCourse = append(sliceCourse, v)
	}
	return sliceCourse
}

func (s *ServiceImpl) Login(ctx context.Context, request web.RequestLogin) web.CatResp {
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

func (s *ServiceImpl) Register(ctx context.Context, request web.CategoryRequest) web.CatResp {
	tx, err := s.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer helper.CommitorRollback(tx)
	category := model.User{
		Id:       request.Id,
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
		Age:      request.Age,
		Phone:    request.Phone,
		Role:     request.Role,
	}
	if category.Role == "" {
		category.Role = "user"
	}
	category = s.Rep.Register(ctx, tx, category)
	return helper.ConvertModel(category)
}

func (s *ServiceImpl) GetCourse(ctx context.Context, req model.Class, id string) model.Class {
	tx, err := s.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitorRollback(tx)
	Class := s.Rep.GetCourse(ctx, tx, req, id)
	Class.CourseId = id
	return Class

}
