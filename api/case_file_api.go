package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/application"
	"github.com/toxb11/awesome_project/vo"
	"net/http"
)

func SetupCaseFileRouter(engine *gin.Engine) *gin.Engine {
	engine.POST("/case/m_get", func(c *gin.Context) {
		req := &vo.MGetCaseRequest{}
		err := c.Bind(&req)
		if err != nil {
			logrus.Errorf("[MGetCase]: Bind request err: %v\n", err)
			return
		}
		response := application.MGetCase(c, req)
		c.JSON(http.StatusOK, response)
	})

	engine.POST("/case/create", func(c *gin.Context) {
		req := &vo.CreateCaseRequest{}
		err := c.Bind(&req)
		if err != nil {
			logrus.Errorf("[CreateCase]: bind req err: %v\n", err)
			return
		}
		response := application.CreateCase(c, req)
		c.JSON(http.StatusOK, response)
	})

	engine.POST("/chat/describe", func(c *gin.Context) {
		req := &vo.DescribeCriminalRequest{}
		err := c.Bind(&req)
		if err != nil {
			logrus.Errorf("[DescribeCriminal] bind req err: %v\n", err)
			return
		}
		response := application.DescribeCriminal(c, req)
		c.PureJSON(http.StatusOK, response)
	})

	engine.POST("/chat/confirm_picture", func(c *gin.Context) {
		req := &vo.ConfirmPictureRequest{}
		err := c.Bind(&req)
		if err != nil {
			logrus.Errorf("[ConfirmPicture] bind req err: %v\n", err)
			return
		}
		response := application.ConfirmPicture(c, req)
		c.PureJSON(http.StatusOK, response)
	})

	engine.POST("/chat/save_picture", func(c *gin.Context) {
		req := &vo.SaveCriminalPictureRequest{}
		err := c.Bind(&req)
		if err != nil {
			logrus.Errorf("[SavePicture] bind req err: %v\n", err)
			return
		}
		response := application.SavePicture(c, req)
		c.PureJSON(http.StatusOK, response)
	})

	//engine.POST("/chat/m_get", func(c *gin.Context) {
	//	req := &vo.MGetChatRequest{}
	//	err := c.Bind()
	//})

	return engine

}
