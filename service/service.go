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
	FindCourseByCategory(ctx context.Context, id string) model.Course
	FindAll(ctx context.Context) []web.CatResp
	Login(ctx context.Context, request web.CategoryRequest) web.CatResp
	Register(ctx context.Context, request web.CategoryRequest) web.CatResp
}
