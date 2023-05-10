package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Alfeenn/online-learning/helper"
	"github.com/Alfeenn/online-learning/middleware"
	"github.com/Alfeenn/online-learning/model"
	"github.com/Alfeenn/online-learning/model/web"
	"github.com/Alfeenn/online-learning/service"
	"github.com/gin-gonic/gin"
)

type ControllerImpl struct {
	ServiceModel service.Service
}

func NewController(c service.Service) Controller {
	return &ControllerImpl{
		ServiceModel: c,
	}
}

func (c *ControllerImpl) Create(g *gin.Context) {
	req, err := helper.UploadFile(g)
	if err != nil {

		g.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
	} else {

		resp := c.ServiceModel.CreateCourse(g.Request.Context(), req)
		response := web.WebResponse{
			Code:   http.StatusCreated,
			Status: "CREATED",
			Data:   resp,
		}
		g.JSON(http.StatusOK, response)
	}
}

func (c *ControllerImpl) Update(g *gin.Context) {
	req := model.Course{}
	err := g.ShouldBind(&req)
	req.Id = g.Param("id")
	//check if bind json error
	if err != nil {

		g.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
	} else {
		result := c.ServiceModel.Update(g.Request.Context(), req)
		g.JSON(http.StatusOK, result)
	}

}

func (c *ControllerImpl) Delete(g *gin.Context) {
	id := g.Param("id")
	c.ServiceModel.Delete(g.Request.Context(), id)
	g.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "Successfully delete data"})
}

func (c *ControllerImpl) Find(g *gin.Context) {
	id := g.Params.ByName("id")
	if id == "" {
		g.AbortWithStatusJSON(http.StatusNotFound,
			gin.H{
				"code": http.StatusNotFound,
				"msg":  "Id not found"})
	}

	result := c.ServiceModel.FindCourseByCategory(g.Request.Context(), id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) FindAll(g *gin.Context) {

	result := c.ServiceModel.FindAll(g.Request.Context())
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) UserSignIn(g *gin.Context) {
	key := strconv.AppendBool([]byte(model.Key), true)
	requestservice := web.CategoryRequest{}
	//check form input
	err := g.ShouldBind(&requestservice)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"code": http.StatusInternalServerError,
				"msg":  err.Error(),
			})
	} else {
		//proceed to login
		user := c.ServiceModel.Login(g.Request.Context(), requestservice)
		match := helper.CheckHashPassword(user.Password, requestservice.Password)
		var data map[string]interface{}
		//check if password match
		if !match {
			g.AbortWithStatusJSON(http.StatusInternalServerError,
				gin.H{
					"code": http.StatusInternalServerError,
					"msg":  "Password not match"})
		} else {
			//set token
			tokenstring := helper.GenerateToken(g, key, user)
			data = map[string]interface{}{
				"Authorization": tokenstring,
			}
			g.JSON(http.StatusOK, web.WebResponse{
				Code:   200,
				Status: "OK",
				Data:   data,
			})
		}
	}
}

func (c *ControllerImpl) GetAccessList(g *gin.Context) {
	enforcer := middleware.UserPolicy()

	adapter := enforcer.GetAllObjects()

	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   adapter,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) Register(g *gin.Context) {
	enforcer := middleware.UserPolicy()
	req := web.CategoryRequest{}
	age, _ := strconv.Atoi(g.Request.FormValue("age"))
	req.Age = int64(age)
	phone, _ := strconv.Atoi(g.Request.FormValue("phone"))
	req.Phone = int64(phone)

	err := g.ShouldBind(&req)
	log.Print(req)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"code": http.StatusBadRequest,
				"msg":  err.Error(),
			})
	} else {

		req.Password, _ = helper.HashPassword(req.Password)
		resp := c.ServiceModel.Register(g.Request.Context(), req)
		enforcer.AddGroupingPolicy(fmt.Sprint(resp.Username), resp.Role)
		response := web.WebResponse{
			Code:   http.StatusCreated,
			Status: "CREATED",
			Data:   resp,
		}
		g.JSON(http.StatusOK, response)
	}

}
