package service

import (
	"context"

	"github.com/Alfeenn/online-learning/model"
	"github.com/Alfeenn/online-learning/model/web"
)

type Service interface {
	CreateCourse(ctx context.Context, req model.Course) model.Course
	Update(ctx context.Context, req model.Course) model.Course
	Delete(ctx context.Context, id string)
	FindCourseById(ctx context.Context, id string) model.Course
	FindCourseByCategory(ctx context.Context, id string) model.Course
	FindAll(ctx context.Context) []model.Course
	Login(ctx context.Context, request web.RequestLogin) web.CatResp
	Register(ctx context.Context, request web.CategoryRequest) web.CatResp
	GetCourse(ctx context.Context, req model.Class, id string) model.Class
	DeleteUser(ctx context.Context, id string)
}
