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
	FindAll(ctx context.Context, tx *sql.Tx) []model.Course
	FindCourseById(ctx context.Context, tx *sql.Tx, id string) (model.Course, error)
	FindCourseByCategory(ctx context.Context, tx *sql.Tx, category string) (model.Course, error)
	Login(ctx context.Context, tx *sql.Tx, category model.User) (model.User, error)
	Register(ctx context.Context, tx *sql.Tx, category model.User) model.User
	GetCourse(ctx context.Context, tx *sql.Tx, category model.Class, id string) model.Class
	DeleteCourse(ctx context.Context, tx *sql.Tx, id string)
}
