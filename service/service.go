package service

import (
	"context"

	"github.com/Alfeenn/online-learning/model/web"
)

type Service interface {
	Create(ctx context.Context, req web.CategoryRequest) web.CatResp
	Update(ctx context.Context, req web.UpdateRequest) web.CatResp
	Delete(ctx context.Context, id string)
	Find(ctx context.Context, id string) web.CatResp
	FindAll(ctx context.Context) []web.CatResp
	Login(ctx context.Context, request web.CategoryRequest) web.CatResp
}
