// Defines the user service
package services

import (
	"github.com/jordanparker6/go-api/src/core"
	"github.com/jordanparker6/go-api/src/models"
)

type UserServiceI interface {
	Create()
	Delete()
}

type userService struct {
	db core.DatabaseI
	UserServiceI
}

// UserSerivce.Create: Creates a User.
func (us *userService) Create(FirstName string, LastName string) {
	user := models.User{Id: core.NewId(), FirstName: FirstName, LastName: LastName}
	us.db.Orm().Create(user)
}

// UserService.Delete: Deletes a User.
func (us *userService) Delete(Id string) {
	us.db.Orm().Delete(models.User{Id: Id})
}

// UserService.GetOne: Gets a User by User ID.
func (us *userService) GetOne(Id string) (models.User, error) {
	var user models.User
	if err := us.db.Orm().Where("id = ?", Id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// NewUserService: The UserService constructor.
func NewUserService(db core.DatabaseI) *userService {
	service := new(userService)
	service.db = db
	return service
}
