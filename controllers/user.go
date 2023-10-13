package controllers

import (
	"fmt"
	"go-chat/httputil"
	"go-chat/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

// Users
// @Summary 用户列表
// @Tags 用户模块
// @Success 200 {string} json{"code", "users"}
// @Router /users [get]
func Users(c *gin.Context) {
	users := models.Users()

	fmt.Println(users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

// Users
// @Summary 新增用户
// @Tags 用户模块
// @Param User body models.CreateUserAttributes true "用户"
// @Success 200 {string} json{"code", "users"}
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var createUser models.CreateUserAttributes
	if err := c.ShouldBindJSON(&createUser); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := createUser.Validation(); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	user := models.User{
		Name:     createUser.Name,
		Password: createUser.Password,
	}

	models.CreateUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User created.",
		"user":    user,
	})
}

// Users
// @Summary 删除用户
// @Tags 用户模块
// @Param id path int true "User ID" Format(int64)
// @Success 200 {string} json{"code", "users"}
// @Router /user/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	user.ID = uint(aid)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}

	models.DeleteUser(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted.",
		"user":    user,
	})
}

// Users
// @Summary 编辑用户
// @Tags 用户模块
// @Param id path int true "User ID" Format(int64)
// @Param User body models.CreateUserAttributes true "用户"
// @Success 200 {string} json{"code", "users"}
// @Router /user/{id} [patch]
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	user.ID = uint(aid)
	var createUser models.CreateUserAttributes
	if err := c.ShouldBindJSON(&createUser); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	if err := createUser.Validation(); err != nil {
		httputil.NewError(c, http.StatusBadRequest, err)
		return
	}
	user.Name = createUser.Name
	user.Password = createUser.Password
	models.UpdateUser(user)
	fmt.Println(user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated.",
		"user":    user,
	})
}
