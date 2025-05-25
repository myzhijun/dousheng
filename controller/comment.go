package controller

import (
	"encoding/json"
	"github.com/RaymondCode/simple-demo/model"
	"github.com/RaymondCode/simple-demo/model/request"
	"github.com/RaymondCode/simple-demo/model/response"
	"github.com/RaymondCode/simple-demo/service"
	"github.com/RaymondCode/simple-demo/utils/verify"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CommentAction add comment or delete comment
func CommentAction(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")

	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var commentRequest request.CommentRequest
	if err := c.ShouldBind(&commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error " + err.Error()})
		return
	}

	// verify
	if err := verify.Comment(commentRequest); err != nil {
		c.JSON(http.StatusBadRequest, response.Response{StatusCode: 1, StatusMsg: "非法数据 "})
		return
	}

	cs := service.CommentService{}
	var comment *model.Comment
	if commentRequest.ActionType == "1" {
		comment, err = cs.CommentAction(userInfoVar.ID, &commentRequest)
	} else if commentRequest.ActionType == "2" {
		err = cs.DeleteCommentAction(&commentRequest)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentAction: " + err.Error()})
		return
	}
	c.JSON(http.StatusOK, response.CommentActionResponse{
		Response: response.Response{StatusCode: 0},
		Comment:  comment,
	})

}

// CommentList get comments list of a video
func CommentList(c *gin.Context) {

	// authentication
	UserStr, _ := c.Get("UserStr")
	var userInfoVar model.User
	if err := json.Unmarshal([]byte(UserStr.(string)), &userInfoVar); err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "error: session unmarshal error"})
		return
	}

	// bind request var
	var commentListRequest request.CommentListRequest
	if err := c.ShouldBind(&commentListRequest); err != nil {
		c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "bind error "})
		return
	}

	cs := service.CommentService{}
	commentList, err := cs.CommentList(userInfoVar.ID, &commentListRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{StatusCode: 1, StatusMsg: "error in commentList: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.CommentListResponse{
		Response:    response.Response{StatusCode: 0},
		CommentList: commentList,
	})
}
