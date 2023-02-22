package handler

import (
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary 登录
// @Tags user
// @Description 邮箱密码登录
// @Accept application/json
// @Produce application/json
// @Param login_request body handler.loginRequest true "login_request"
// @Success 200 {object} handler.Response "{"message":"登录成功"}"
// @Failure 400 {object} handler.Response "{"message":"登录失败"}"
// @Router /api/v1/login
func Login(r *gin.Context) {
	req := new(loginRequest)
	r.ShouldBindJSON(req)
	if req.Email == "" || req.Password == "" {
		SendBadRequest(r, nil, req)
		return
	}

}

func UpdateUserInfo(c *gin.Context) {

}
