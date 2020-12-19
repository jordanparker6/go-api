// Dependency Injection
// The Container struct builds singletons
// during the initiation of the container and provides
// methods that build services reliant upon dependencies
package services

import "github.com/jordanparker6/go-api/src/core"

type Container struct {
	DB core.DatabaseI
}

// NewContainer
// constructs the container and init singletons
func NewContainer() *Container {
	c := new(Container)
	c.DB = core.NewDatabase(
		core.Config.Database.Dialect,
		core.Config.Database.Uri,
	)
	return c
}

// CreateUserService
// builds a UserService and wires dependencies
func (c *Container) CreateUserService() *userService {
	return NewUserService(c.DB)
}
