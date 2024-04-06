package data

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/pritesh-mantri/sailor/internal/query"
)

type UserModel struct {
	db *sqlx.DB
}

type gender string

const (
	Male   gender = "male"
	Female gender = "female"
)

type User struct {
	ID        int64     `db:"id"`
	Name      string    `db:"name"`
	Age       int64     `db:"age"`
	Phone     string    `db:"phone"`
	Birthdate time.Time `db:"birthdate"`
	Password  string
	Location  string    `db:"location"`
	Gender    gender    `db:"gender"`
	CreatedAt time.Time `db:"created_at"`
}

var userSelectableColumns = []string{"id", "name", "age", "phone", "birthdate", "location", "gender", "created_at"}

// this selects id, phone and password for authentication
type AuthInfo struct {
	ID       int64  `db:"id"`
	Phone    string `db:"phone"`
	Password string `db:"password"`
}

func (u UserModel) Create(ctx context.Context, user User) error {
	q, args, err := query.InsertInto(userTable).Rows(user).ToSQL()
	if err != nil {
		return err
	}

	_, err = u.db.ExecContext(ctx, q, args...)

	return err
}

// returns error if the result is empty
func (u UserModel) GetByID(ctx context.Context, id int64) (*User, error) {
	q, args, err := query.From(userTable).Select(userSelectableColumns...).Where(query.M{"id": id}).ToSQL()
	if err != nil {
		return nil, err
	}
	user := &User{}
	err = u.db.GetContext(ctx, user, q, args...)

	return user, err
}

// returns user info required for authentication
func (u UserModel) GetAuthInfo(ctx context.Context, id int64) (*AuthInfo, error) {
	q, args, err := query.From(userTable).Select("id", "phone", "password").Where(query.M{"id": id}).ToSQL()
	if err != nil {
		return nil, err
	}

	authInfo := &AuthInfo{}
	err = u.db.GetContext(ctx, authInfo, q, args...)

	return authInfo, err
}

func (u UserModel) Delete(ctx context.Context, id int64) error {
	return nil
}

func (u UserModel) Update(ctx context.Context) *User {
	return &User{}
}
