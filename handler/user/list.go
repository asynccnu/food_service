package user

import (
	. "github.com/asynccnu/food_service/handler"
	"github.com/asynccnu/food_service/pkg/errno"
	"github.com/asynccnu/food_service/service"

	"github.com/gin-gonic/gin"
)

// List list the users in the database.
func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendBadRequest(c, errno.ErrBind, nil, err.Error())
		return
	}

	infos, count, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		SendError(c, err, nil, err.Error())
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		UserList:   infos,
	})
}
