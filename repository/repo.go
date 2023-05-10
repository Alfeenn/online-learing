package repository

import (
	"context"
	"database/sql"

	"github.com/Alfeenn/online-learning/model"
)

type Repository interface {
	CreateCourse(ctx context.Context, tx *sql.Tx, category model.Course) model.Course
	Update(ctx context.Context, tx *sql.Tx, category model.Course) model.Course
	Delete(ctx context.Context, tx *sql.Tx, id string)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
	FindCourseByCategory(ctx context.Context, tx *sql.Tx, category string) (model.Course, error)
	Login(ctx context.Context, tx *sql.Tx, category model.User) (model.User, error)
	Register(ctx context.Context, tx *sql.Tx, category model.User) model.User
}
