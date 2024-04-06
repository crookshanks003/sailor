package data

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"

	"github.com/pritesh-mantri/sailor/config"
)

func TestUserModel_GetByID(t *testing.T) {
	testDB := connectDB(config.TestConfig())
	defer cleanupDB(testDB)

	var userID int64 = 1
	_, err := testDB.Exec(`INSERT INTO users (id, name, age, birthdate, gender, location, phone, password, created_at) VALUES (1, 'test_user1', 23, DATE '2000-10-06', 'male', '78.123,11.123', '0000099999', 'password', DATE '2022-10-06');`)
	if err != nil {
		panic("failed to insert into db")
	}

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *User
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy",
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  userID,
			},
			want: &User{
				ID:        userID,
				Name:      "test_user1",
				Age:       23,
				Location:  "78.123,11.123",
				Gender:    "male",
				Phone:     "0000099999",
				Birthdate: parseDate("2006-01-02", "2000-10-06"),
				CreatedAt: parseDate("2006-01-02", "2022-10-06"),
			},
			assertion: assert.NoError,
		},
		{
			name: "no record",
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  0,
			},
			want:      &User{},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserModel{
				db: tt.fields.db,
			}
			got, err := u.GetByID(tt.args.ctx, tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.want.Age, got.Age)
			assert.Equal(t, tt.want.Name, got.Name)
			assert.Equal(t, tt.want.Phone, got.Phone)
			assert.Equal(t, tt.want.Password, got.Password)
			assert.Equal(t, tt.want.Gender, got.Gender)
			assert.Equal(t, tt.want.ID, got.ID)
			assert.Equal(t, tt.want.Birthdate.UTC(), got.Birthdate.UTC())
			assert.Equal(t, tt.want.CreatedAt.UTC(), got.CreatedAt.UTC())
		})
	}
}

func TestUserModel_GetAuthInfo(t *testing.T) {
	testDB := connectDB(config.TestConfig())
	defer cleanupDB(testDB)

	var userID int64 = 1
	_, err := testDB.Exec(`INSERT INTO users (id, name, age, birthdate, gender, location, phone, password, created_at) VALUES (1, 'test_user1', 23, DATE '2000-10-06', 'male', '78.123,11.123', '0000099999', 'password', DATE '2022-10-06');`)
	if err != nil {
		panic("failed to insert into db")
	}

	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      *AuthInfo
		assertion assert.ErrorAssertionFunc
	}{
		{
			name: "happy",
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  userID,
			},
			want: &AuthInfo{
				ID:       userID,
				Phone:    "0000099999",
				Password: "password",
			},
			assertion: assert.NoError,
		},

		{
			name: "no record",
			fields: fields{
				db: testDB,
			},
			args: args{
				ctx: context.TODO(),
				id:  0,
			},
			want:      &AuthInfo{},
			assertion: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserModel{
				db: tt.fields.db,
			}
			got, err := u.GetAuthInfo(tt.args.ctx, tt.args.id)
			tt.assertion(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
