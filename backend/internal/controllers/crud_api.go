package controllers

import (
	"sso/internal/errors"
	"sso/internal/models"
	"sso/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CRUDController interface {
	Users(c *gin.Context)
	UserByID(c *gin.Context)
	UserScopes(c *gin.Context)
	UserSessions(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	DeleteUserSession(c *gin.Context)

	AddClient(c *gin.Context)
	Clients(c *gin.Context)
	ClientByID(c *gin.Context)
	UpdateClient(c *gin.Context)
	DeleteClient(c *gin.Context)

	AddClientScope(c *gin.Context)
	ClientScopes(c *gin.Context)
	UpdateClientScope(c *gin.Context)
	DeleteClientScope(c *gin.Context)
}

type crud struct {
	userService   service.UserService
	clientService service.ClientService
}

func NewCRUD(us service.UserService, cs service.ClientService) CRUDController {
	return crud{us, cs}
}

// @Summary Get users
// @Description paginated user query
// @Tags CRUD
// @Param limit query int false "Page size"
// @Param page query int false "Page number"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users [get]
func (cr crud) Users(c *gin.Context) {
	limit := c.GetInt("limit")
	page := c.GetInt("page")

	users := cr.userService.GetUsers(true, limit, page)
	c.JSON(200, models.UsersToResponses(users))
}

// @Summary Get user
// @Description single user query
// @Tags CRUD
// @Param id path string true "User id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id} [get]
func (cr crud) UserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := cr.userService.GetUserByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.AppErr(404, err.Error()))
		default:
			c.Error(errors.AppErr(500, err.Error()))
		}
		return
	}
	c.JSON(200, user.ToResponse())
}

// @Summary Get user scopes
// @Description Query user scopes for a client
// @Tags CRUD
// @Param id path string true "User id"
// @Param client_id path string true "Client id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id}/scopes/{client_id} [get]
func (cr crud) UserScopes(c *gin.Context) {
	userID := c.Param("id")
	clientID := c.Param("client_id")
	scopes := cr.userService.GetUserScopes(userID, clientID)
	c.JSON(200, models.ScopesToResponses(scopes))
}

// @Summary Get user sessions
// @Description Query sessions for user
// @Tags CRUD
// @Param id path string true "User id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id}/sessions [get]
func (cr crud) UserSessions(c *gin.Context) {
	id := c.Param("id")
	sessions := cr.userService.GetUserSessions(id)
	c.JSON(200, sessions)
}

// @Summary Update user
// @Description Updates a user
// @Tags CRUD
// @Param id path string true "User id"
// @Param form body models.UpdateUserForm true "Form"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id} [put]
func (cr crud) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var form models.UpdateUserForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	newUser, err := cr.userService.UpdateUser(id, form)
	if err != nil {
		c.Error(errors.AppErr(500, err.Error()))
	}

	c.JSON(204, newUser.ToResponse())
}

// @Summary Delete user
// @Description Deletes a user
// @Tags CRUD
// @Param id path string true "User id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/users/{id} [delete]
func (cr crud) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := cr.userService.DeleteUser(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.AppErr(404, err.Error()))
		default:
			c.Error(errors.AppErr(500, err.Error()))
		}
		return
	}
	c.Status(201)
}

// @Summary Delete session
// @Description Deletes a session
// @Tags CRUD
// @Param id path string true "Session id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/sessions/{id} [delete]
func (cr crud) DeleteUserSession(c *gin.Context) {
	id := c.Param("id")
	err := cr.userService.DeleteSession(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.AppErr(404, err.Error()))
		default:
			c.Error(errors.AppErr(500, err.Error()))
		}
		return
	}
	c.Status(201)
}

// @Summary Add client
// @Description adds a client to db and returns it
// @Tags CRUD
// @accept json
// @Param form body models.AddClientForm true "Add client form"
// @Success 201
// @Failure 400
// @Failure 401
// @Failure 403
// @Failure 500
// @Router /api/v1/clients [post]
func (cr crud) AddClient(c *gin.Context) {
	var form models.AddClientForm
	if err := c.ShouldBindJSON(&form); err != nil {
		c.Error(errors.AppErr(400, err.Error()))
		return
	}

	client, err := cr.clientService.AddClient(&form)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.AppErr(404, err.Error()))
		default:
			c.Error(errors.AppErr(500, err.Error()))
		}
		return
	}

	c.JSON(201, client)
}

// @Summary Get clients
// @Description paginated client query
// @Tags CRUD
// @Param limit query int false "Page size"
// @Param page query int false "Page number"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/clients [get]
func (cr crud) Clients(c *gin.Context) {
	limit := c.GetInt("limit")
	page := c.GetInt("page")

	clients := cr.clientService.GetClients(true, limit, page)
	c.JSON(200, models.ClientsToResponses(clients))
}

// @Summary Get client
// @Description single client query
// @Tags CRUD
// @Param id path string true "User id"
// @Success 200
// @Failure 400
// @Failure 500
// @Router /api/v1/clients/{id} [get]
func (cr crud) ClientByID(c *gin.Context) {
	id := c.Param("id")
	client, err := cr.clientService.GetClientByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.Error(errors.AppErr(404, err.Error()))
		default:
			c.Error(errors.AppErr(500, err.Error()))
		}
		return
	}
	c.JSON(200, client.ToResponse())
}

// UpdateClient implements CRUDController.
func (crud) UpdateClient(c *gin.Context) {
	panic("unimplemented")
}

// DeleteClient implements CRUDController.
func (crud) DeleteClient(c *gin.Context) {
	panic("unimplemented")
}

// ClientScopes implements CRUDController.
func (cr crud) ClientScopes(c *gin.Context) {
	panic("unimplemented")
}

// AddClientScope implements CRUDController.
func (cr crud) AddClientScope(c *gin.Context) {
	panic("unimplemented")
}

// UpdateClientScope implements CRUDController.
func (crud) UpdateClientScope(c *gin.Context) {
	panic("unimplemented")
}

// DeleteClientScope implements CRUDController.
func (crud) DeleteClientScope(c *gin.Context) {
	panic("unimplemented")
}
