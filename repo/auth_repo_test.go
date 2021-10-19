package repo_test

import (
	"testing"

	"authentication/model"
	"authentication/repo"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var u = model.User{
	Username: "erwindo",
	Password: "password",
	Name:     "Erwindo Sianipar",
}

func MockGormDB() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		logrus.Fatal(err)
	}

	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      sqlDB,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	return gormDB, mock
}

func TestCheckUsername(t *testing.T) {
	t.Run("test normal case repo check username", func(t *testing.T) {
		gormDB, mock := MockGormDB()

		rows := sqlmock.NewRows([]string{"count"}).AddRow(0)
		mock.ExpectQuery("SELECT count(*) FROM `users` WHERE username = ?").
			WithArgs(u.Username).
			WillReturnRows(rows)

		authRepo := repo.NewAuthRepo(gormDB)
		available := authRepo.CheckUsername(u.Username)

		t.Run("test username is available", func(t *testing.T) {
			assert.Equal(t, true, available)
		})
	})
}

func TestRegister(t *testing.T) {
	t.Run("test normal case repo register", func(t *testing.T) {
		gormDB, mock := MockGormDB()

		q := "INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`,`username`,`password`,`name`) VALUES (?,?,?,?,?,?)"
		mock.ExpectExec(q).
			WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), u.DeletedAt, u.Username, u.Password, u.Name).
			WillReturnResult(sqlmock.NewResult(1, 1))

		authRepo := repo.NewAuthRepo(gormDB)
		err := authRepo.Register(&u)

		t.Run("test store data with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}

func TestLogin(t *testing.T) {
	hashedPassword := "$2a$10$fk9IPSmo/VYhu5VJm.vPy.5.XVowBHU3otSDAzTBpMR3YpX2cqYwW"

	t.Run("test normal case repo login", func(t *testing.T) {
		gormDB, mock := MockGormDB()

		rows := sqlmock.NewRows([]string{"password"}).AddRow(hashedPassword)
		mock.ExpectQuery("SELECT * FROM `users` WHERE username = ? ORDER BY `users`.`id` LIMIT 1").
			WillReturnRows(rows)

		authRepo := repo.NewAuthRepo(gormDB)
		password, err := authRepo.Login(u.Username)
		assert.NoError(t, err)

		t.Run("test get stored password by username is hashed", func(t *testing.T) {
			assert.Equal(t, hashedPassword, password)
		})
	})
}

func TestCheckID(t *testing.T) {
	t.Run("test normal case repo check id", func(t *testing.T) {
		gormDB, mock := MockGormDB()

		rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
		mock.ExpectQuery("SELECT count(*) FROM `users` WHERE id = ?").WithArgs(0).WillReturnRows(rows)

		authRepo := repo.NewAuthRepo(gormDB)
		available := authRepo.CheckID(0)

		t.Run("test id is exist for case delete", func(t *testing.T) {
			assert.Equal(t, true, available)
		})
	})
}

func TestDelete(t *testing.T) {
	t.Run("test normal case repo delete", func(t *testing.T) {
		gormDB, mock := MockGormDB()

		mock.ExpectExec("DELETE FROM `users` WHERE `users`.`id` = ?").
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(1, 1))

		authRepo := repo.NewAuthRepo(gormDB)
		err := authRepo.Delete(1)

		t.Run("test data deleted with no error", func(t *testing.T) {
			assert.Equal(t, nil, err)
		})
	})
}
