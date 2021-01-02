package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gover-server/response"
	"gover-server/service"
	"gover-server/util"
	"log"
	"strconv"
)

func SendDeptComment(ctx *gin.Context) {
	deptId, err := strconv.Atoi(ctx.Param("deptId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}

	comments, err1 := service.GetCommentsByDeptId(int64(deptId))
	if util.IsFailed(err1) {
		response.Fail(ctx, nil, err1.Message)
		return
	}

	response.Success(ctx, gin.H{
		"comments": *comments,
	}, "获取部门评论成功！")
}

func FeedBack(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	reply := mMap["reply"]
	commentId, err1 := strconv.Atoi(mMap["commentId"])
	if err1 != nil {
		log.Println(err1)
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.FeedBack(int64(commentId), reply); util.IsFailed(err) {
		log.Println(err)
		response.Fail(ctx, nil, err.Message)
		return
	}

	response.Success(ctx, nil, "反馈成功！")
}
