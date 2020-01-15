package user

import (
	. "github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/model"
	"github.com/asynccnu/food_service/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Get gets an user by the user identifier.
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the user by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendError(c, errno.ErrUserNotFound, nil, err.Error())
		return
	}
	SendResponse(c, nil, user)
}
