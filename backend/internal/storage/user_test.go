package storage_test

import (
	"database/sql"
	"testing"

	"sso/internal/storage"
	"sso/internal/storage/models"
	"sso/internal/utils"
	"sso/pkg/config"

	"github.com/davecgh/go-spew/spew"
)

func TestUserCRUD(t *testing.T) {
	config := config.MustLoad()

	config.DB.Name = "test"
	config.DB.Host = "localhost"
	db := storage.MustLoad(config)
	repo := storage.NewUserRepo(db)

	init := models.User{
		Username:     utils.RandomString(8),
		Email:        sql.NullString{String: utils.RandomString(8) + "@test.com", Valid: true},
		PasswordHash: sql.NullString{String: "hash", Valid: true},
	}
	if err := repo.CreateUser(&init); err != nil {
		t.Error(err)
	} else {
		t.Log("CREATED USER:", spew.Sdump(init))
	}

	user, err := repo.GetUserByID(init.ID.String())
	if err != nil {
		t.Error(err)
	} else {
		t.Log("FOUND USER:", spew.Sdump(user))
	}

	user.EmailVerified = true
	if err := repo.UpdateUser(user); err != nil {
		t.Error(err)
	} else {
		t.Log("UPDATED USER:", spew.Sdump(user))
	}

	if err := repo.SoftDeleteUser(user.ID.String()); err != nil {
		t.Error(err)
	} else {
		t.Log("SOFT DELETED USER:", spew.Sdump(user))
	}

	if err := repo.RestoreUser(user.ID.String()); err != nil {
		t.Error(err)
	} else {
		t.Log("RESTORED USER:", spew.Sdump(user))
	}

	if err := repo.DeleteUser(user.ID.String()); err != nil {
		t.Error(err)
	} else {
		t.Log("HARD DELETED USER:", spew.Sdump(user))
	}
}
