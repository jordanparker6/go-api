package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/jordanparker6/go-api/src/core"
)

// * Schema Definitions
type CreateUserSchema struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// * User Controllers

// @Summary GetUser
// @Description Gets a user by user ID.
// @ID get-user-by-id
// @Tags User
// @Router /user/{id} [get]
// @Security BasicAuth
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} models.User
// @Failure 400 {object} core.HTTPError
// @Failure 404 {object} core.HTTPError
// @Failure 500 {object} core.HTTPError
func (c *controller) GetUser(ctx *gin.Context) {
	UserService := c.depends.CreateUserService()
	data, err := UserService.GetOne(ctx.Param("id"))
	if err != nil {
		core.NewHTTPError(ctx, http.StatusNotFound, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"result": data})
}

// @Summary CreateUser
// @Description Creates a new user record.
// @ID create-user
// @Tags User
// @Router /user/create [post]
// @Security BasicAuth
// @Param id body CreateUserSchema true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} core.HTTPSuccess
// @Failure 400 {object} core.HTTPError
// @Failure 404 {object} core.HTTPError
// @Failure 500 {object} core.HTTPError
func (c *controller) CreateUser(ctx *gin.Context) {
	var body CreateUserSchema
	if err := ctx.ShouldBindJSON(&body); err != nil {
		core.NewHTTPError(ctx, http.StatusBadRequest, err)
		return
	}
	UserService := c.depends.CreateUserService()
	UserService.Create(body.FirstName, body.LastName)
	// should fetch user from db and return
	ctx.JSON(http.StatusOK, gin.H{"message": "user created"})
}

// @Summary DeleteUser
// @Description Delete a user record by user ID.
// @ID delete-user
// @Tags User
// @Router /user/{id} [delete]
// @Security BasicAuth
// @Param id path int true "User ID"
// @Accept  json
// @Produce  json
// @Success 200 {object} core.HTTPSuccess
// @Failure 400 {object} core.HTTPError
// @Failure 404 {object} core.HTTPError
// @Failure 500 {object} core.HTTPError
func (c *controller) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	UserService := c.depends.CreateUserService()
	UserService.Delete(id)
	ctx.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("deleted user: %s", id)})
}
