package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wduartebr/goexpert/apis/internal/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jhon Do", "j@mail.com", "123456")

	userDB := NewUser(db)

	err = userDB.CreateUser(user)

	assert.Nil(t, err)

	var userFound entity.User

	err = db.First(&userFound, "id = ?", user.ID).Error

	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}

func TestFindByEmailUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.User{})
	user, _ := entity.NewUser("Jhon Do", "j@mail.com", "123456")

	userDB := NewUser(db)

	err = userDB.CreateUser(user)
	assert.Nil(t, err)

	userFound, err := userDB.FindByEmail(user.Email)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.Equal(t, user.ID, userFound.ID)
	assert.NotNil(t, userFound.Password)

}
