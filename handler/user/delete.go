package user

import (
	"strconv"

	. "github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/model"
	"github.com/asynccnu/food_service/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user identifier.
func Delete(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userID)); err != nil {
		SendError(c, errno.ErrDatabase, nil, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}
