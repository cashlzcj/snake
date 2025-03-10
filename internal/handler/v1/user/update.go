package user

import (
	"context"

	"github.com/1024casts/snake/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"github.com/1024casts/snake/pkg/errno"
	"github.com/1024casts/snake/pkg/log"
)

// Update 更新用户信息
// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags 用户
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserBaseModel true "The user info"
// @Success 200 {object} app.Response "{"code":0,"message":"OK","data":null}"
// @Router /users/{id} [put]
func Update(c *gin.Context) {
	// Get the user id from the url parameter.
	userID := cast.ToUint64(c.Param("id"))

	// Binding the user data.
	var req UpdateRequest
	if err := c.Bind(&req); err != nil {
		log.Warnf("bind request param err: %+v", err)
		response.Error(c, errno.ErrBind)
		return
	}
	log.Infof("user update req: %#v", req)

	userMap := make(map[string]interface{})
	userMap["avatar"] = req.Avatar
	userMap["sex"] = req.Sex
	err := service.UserSvc.UpdateUser(context.TODO(), userID, userMap)
	if err != nil {
		log.Warnf("[user] update user err, %v", err)
		response.Error(c, errno.InternalServerError)
		return
	}

	response.Success(c, userID)
}
