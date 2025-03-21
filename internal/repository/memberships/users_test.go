package memberships

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/joshuatheokurniawansiregar/music_catalog_2/internal/models/memberships"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRepository_CreateUser(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}))

	assert.NoError(t, err)

	type args struct {
		model memberships.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mockFn  func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "succcess",
			args: args{
				model: memberships.User{
					Email:     "joshuatheo19@gmail.com",
					Username:  "joshua",
					Password:  "password",
					CreatedBy: "joshuatheo19@gmail.com",
					UpdatedBy: "joshuatheo19@gmail.com",
				},
			},
			wantErr: false,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "user" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
				mock.ExpectCommit()
			},
		},

		{
			name: "failed",
			args: args{
				model: memberships.User{
					Email:     "joshuatheo19@gmail.com",
					Username:  "joshua",
					Password:  "password",
					CreatedBy: "joshuatheo19@gmail.com",
					UpdatedBy: "joshuatheo19@gmail.com",
				},
			},
			wantErr: true,
			mockFn: func(args args) {
				mock.ExpectBegin()
				mock.ExpectQuery(`INSERT INTO "users" (.+) VALUES (.+)`).WithArgs(
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					sqlmock.AnyArg(),
					args.model.Email,
					args.model.Username,
					args.model.Password,
					args.model.CreatedBy,
					args.model.UpdatedBy,
				).WillReturnError(assert.AnError)
				mock.ExpectRollback()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			r := &Repository{
				db: gormDB,
			}
			if err := r.CreateUser(tt.args.model); (err != nil) != tt.wantErr {
				t.Errorf("Repository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestRepository_GetUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()
	
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db,} ))
	assert.NoError(t, err)
	type args struct {
		email    string
		username string
		id       uint
	}
	tests := []struct {
		name    string
		args    args
		want    *memberships.User
		wantErr bool
		mockFn func(args args)
	}{
		// TODO: Add test cases.
		{
			name: "success",
			args: args{
				email: "test@gmail.com",
				username: "testusername",
				id: 0,
			},
			want: &memberships.User{
				Model: &gorm.Model{
					ID: 1,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
				Email: "test@gmail.com",
				Username: "testusername",
				Password: "testpassword",
				CreatedBy: "test@gmail.com",
				UpdatedBy: "test@gmail.com",
			},
			wantErr: false,
			mockFn: func(args args){
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).
				WithArgs(
					args.email,
					args.username,
					args.id,
					1,
				).
				WillReturnRows(sqlmock.NewRows([]string{
					"id","created_at", "updated_at", "email", "username", "password", "created_by", "updated_by",
				}).AddRow(
					1, time.Now(), time.Now(), "test@gmail.com", "testusername", "testpassword", "test@gmail.com", "test@gmail.com",
				))

			},
		},

		{
			name: "failed",
			args: args{
				email: "test@gmail.com",
				username: "testusername",
				id: 0,
			},
			want: nil,
			wantErr: true,
			mockFn: func(args args){
				mock.ExpectQuery(`SELECT \* FROM "users" .+`).WithArgs(
					args.email,
					args.username,
					args.id,
					1,
				).WillReturnError(assert.AnError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockFn(tt.args)
			var r *Repository = &Repository{
				db: gormDb,
			}
			got, err := r.GetUser(tt.args.email, tt.args.username, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetUser() = %v, want %v", got, tt.want)
			}

			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}
