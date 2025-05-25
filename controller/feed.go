package controller

import (
	"github.com/RaymondCode/simple-demo/global"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/respToDTO"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

// Feed token is optional here
func Feed(c *gin.Context) {

	// bind request var
	var feedRequest request.FeedRequest
	if err := c.ShouldBind(&feedRequest); err != nil {
		global.App.DY_LOG.Error("feed bind error!", zap.Error(err))
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "feed bind error"})
		return
	}

	fs := service.FeedService{}
	var videos *[]model.Video
	if feedRequest.Token != "" {
		videos, err = fs.FeedWithToken(&feedRequest)
	} else {
		videos, err = fs.FeedWithoutToken()
	}
	if err != nil {
		global.App.DY_LOG.Error("feed service error!", zap.Error(err))
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	feedListReturn := respToDTO.GetVideoListDTO(*videos)
	if feedListReturn == nil {
		global.App.DY_LOG.Info("get feedList null")
	}
	c.JSON(http.StatusOK, response.FeedResponse{
		Response:  response.Response{StatusCode: 0},
		VideoList: feedListReturn,
		NextTime:  time.Now().Unix(),
	})
	return
}
