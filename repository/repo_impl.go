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

func (r *RepoImpl) Create(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "INSERT INTO user(id,email,password,role,created_at,updated_at) VALUES(?,?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Username, category.Password,
		category.Role, category.Phone, category.Password)
	helper.PanicIfErr(err)
	return category
}

func (r *RepoImpl) Update(ctx context.Context, tx *sql.Tx, category model.User) model.User {
	SQL := "UPDATE article SET name=? WHERE id=?"

	_, err := tx.ExecContext(ctx, SQL, category.Username, category.Id)
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
		err := rows.Scan(&article.Id, &article.Username, &article.Name, &article.Age,
			&article.Password, &article.Role, &article.Phone, &article.Role)
		helper.PanicIfErr(err)
		sliceArticle = append(sliceArticle, article)
	}
	return sliceArticle
}

func (r *RepoImpl) Find(ctx context.Context, tx *sql.Tx, id string) (model.User, error) {
	SQL := "SELECT *FROM user WHERE id =?"

	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)
	defer rows.Close()
	article := model.User{}
	if rows.Next() {
		rows.Scan(&article.Id, &article.Username,
			&article.Password, &article.Role, &article.Phone, &article.Role)

		return article, nil
	} else {
		return article, err
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
	SQL := "INSERT INTO user(id,email,password,role,created_at,updated_at) VALUES(?,?,?,?,?,?)"
	category.Id = uuid.NewString()
	_, err := tx.ExecContext(ctx, SQL,
		category.Id, category.Username, category.Password,
		category.Role, category.Age, category.Phone)
	helper.PanicIfErr(err)
	return category
}
