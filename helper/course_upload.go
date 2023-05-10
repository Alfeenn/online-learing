package helper

import (
	"strconv"

	"github.com/Alfeenn/online-learning/model"
	"github.com/gin-gonic/gin"
)

func UploadFile(g *gin.Context) (model.Course, error) {
	course := model.Course{}
	g.Request.ParseMultipartForm(10 << 20)
	_, FileHeader, err := g.Request.FormFile("file")
	if err != nil {
		return course, err
	} else {
		err = g.ShouldBind(&course)
		if err != nil {
			return course, err
		} else {

			g.SaveUploadedFile(FileHeader, "./resources/"+FileHeader.Filename)
			price, _ := strconv.Atoi(g.Request.FormValue("price"))
			course = model.Course{
				Name:      g.Request.FormValue("name"),
				Thumbnail: FileHeader.Filename,
				File:      FileHeader,
				Price:     price,
				Category:  g.Request.FormValue("category"),
			}

		}
		return course, err
	}

}
