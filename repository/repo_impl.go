package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Alfeenn/online-learning/helper"
	"github.com/Alfeenn/online-learning/model"
	"github.com/google/uuid"
)

type RepoImpl struct{}

func NewRepository() Repository {
	return &RepoImpl{}
}

func (r *RepoImpl) CreateCourse(ctx context.Context, tx *sql.Tx, category model.Course) model.Course {
	SQL := "INSERT INTO courses(id,name,price,category,thumbnail) VALUES(?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Name, category.Price,
		category.Category, category.Thumbnail)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) Update(ctx context.Context, tx *sql.Tx, category model.Course) model.Course {
	SQL := "UPDATE article SET name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfErr(err)

	return category

}

func (r *RepoImpl) Delete(ctx context.Context, tx *sql.Tx, id string) {
	SQL := "DELETE FROM user WHERE id=?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfErr(err)
}

func (r *RepoImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	sql := "SELECT *FROM user"
	rows, err := tx.QueryContext(ctx, sql)
	helper.PanicIfErr(err)
	defer rows.Close()

	var sliceArticle []model.User

	for rows.Next() {
		article := model.User{}
		err := rows.Scan(&article.Id, &article.Username, &article.Password, &article.Name,
			&article.Age, &article.Phone, &article.Role)
		helper.PanicIfErr(err)
		sliceArticle = append(sliceArticle, article)
	}
	return sliceArticle
}

func (r *RepoImpl) FindCourseByCategory(ctx context.Context, tx *sql.Tx, category string) (model.Course, error) {
	SQL := "SELECT *FROM courses WHERE category =?"

	rows, err := tx.QueryContext(ctx, SQL, category)
	helper.PanicIfErr(err)
	defer rows.Close()
	model := model.Course{}
	if rows.Next() {
		rows.Scan(&model.Id, &model.Name,
			&model.Category, &model.Thumbnail, &model.Price, &model.File)

		return model, nil
	} else {
		return model, errors.New("no data")
	}

}

func (m *RepoImpl) Login(ctx context.Context, tx *sql.Tx, category model.User) (model.User, error) {
	SQL := `SELECT email,password FROM user WHERE email=?`
	rows, err := tx.QueryContext(ctx, SQL, category.Username)
	helper.PanicIfErr(err)
	defer rows.Close()
	user := model.User{}
	if rows.Next() {
		err := rows.Scan(&user.Username, &user.Password)
		if err != nil {
			panic(err)
		}
		return user, nil
	} else {

		return user, errors.New("no data")
	}
}

func (r *RepoImpl) Register(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "INSERT INTO users(id,username,password,name,age,phone,role) VALUES(?,?,?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Username, category.Password,
		category.Name, category.Age, category.Phone, category.Role)
	helper.PanicIfErr(err)
	return category
}
