package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	adapter_entities "github.com/thapakazi/go-hex-arch/internal/adapter/entities"
	"github.com/thapakazi/go-hex-arch/internal/adapter/storage"
	"github.com/thapakazi/go-hex-arch/internal/core/entities"
)

type UserRepository struct {
	db *storage.DB
}

func NewUserRepository() *UserRepository {
	fmt.Println("NewUserRepository ++++++++++++++++++++++++++++++==", storage.Database)
	return &UserRepository{
		db: storage.Database,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entities.User) error {
	query := r.db.QueryBuilder.Insert("users").
		Columns("full_name", "username", "email", "password").
		Values(user.FullName, user.Username, user.Email, user.Password)
	fmt.Println("UserRepository CreateUser ------------", query)

	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, sql, args...)
	return err
}

func (r *UserRepository) GetUser(ctx context.Context, id int64) (user *entities.User, err error) {
	query := r.db.QueryBuilder.Select("*").From("users").Where(sq.Eq{"id": id})

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	user = &entities.User{}
	err = r.db.QueryRow(ctx, sql, args...).
		Scan(&user.ID, &user.FullName, &user.Username, &user.Email, &user.Password)

	fmt.Println("Lets see the query \n", r.db.Expr(sql, args...))
	return
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *entities.User) error {
	query := r.db.QueryBuilder.Update("users").
		Set("full_name", user.FullName).
		Set("username", user.Username).
		Set("email", user.Email).
		Set("password", user.Password).
		Where(sq.Eq{"id": user.ID})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	fmt.Println("Lets see the query \n", r.db.Expr(sql, args...))
	_, err = r.db.Exec(ctx, sql, args...)

	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, id int64) error {
	query := r.db.QueryBuilder.Delete("users").Where(sq.Eq{"id": id})
	sql, args, err := query.ToSql()
	if err != nil {
		return err
	}
	fmt.Println("Lets see the query \n", r.db.Expr(sql, args...))
	_, err = r.db.ExecContext(ctx, sql, args...)
	return err
}

func (r *UserRepository) GetAllUsers(ctx context.Context, params adapter_entities.QueryParams) ([]entities.User, error) {
	var users []entities.User

	query := r.db.QueryBuilder.Select("*").From("users")
	if params.FullName != "" {
		stmt := "%" + params.FullName + "%"
		query = query.Where(sq.Like{"full_name": stmt})
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return users, err
	}

	fmt.Println("Lets see the query \n", r.db.Expr(sql, args...), sql)
	rows, err := r.db.Query(ctx, sql, args...)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		user := entities.User{}
		err = rows.Scan(
			&user.ID, &user.FullName, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
		)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}

	return users, err
}
