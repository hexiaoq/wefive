package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
	"wefive/model"
	"wefive/response"
	"wefive/service"
	"wefive/util"
)

func GetHotChats(ctx *gin.Context) {
	chats, err := service.GetHotChats()
	if util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"chats": *chats,
	}, "获取热门评论成功！")
}

func SendChat(ctx *gin.Context) {
	chatId, err := strconv.Atoi(ctx.Param("chatId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	chat, err1 := service.GetChatDto(int64(chatId))
	if util.IsFailed(err1) {
		response.Fail(ctx, nil, err1.Message)
		return
	}
	response.Success(ctx, gin.H{
		"chat": *chat,
	}, "获取评论成功！")
}

func SendSubChat(ctx *gin.Context) {
	chatId, err := strconv.Atoi(ctx.Param("chatId"))
	if err != nil {
		response.Fail(ctx, nil, err.Error())
		return
	}
	subChats, err1 := service.GetSubChatDtos(int64(chatId))
	if util.IsFailed(err1) {
		response.Fail(ctx, nil, err1.Message)
		return
	}
	response.Success(ctx, gin.H{
		"subChats": *subChats,
	}, "获取子评论成功！")
}

func CreateChat(ctx *gin.Context) {
	var mMap = make(map[string]string)
	var chat model.Chat
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	userId, err1 := strconv.Atoi(mMap["userId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	chat.UserId = int64(userId)
	chat.Content = mMap["content"]
	chat.Picture = mMap["picture"]
	chat.Title = mMap["title"]
	if err := service.CreateChat(&chat); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "发表评论成功！")
}

func CreateSubChat(ctx *gin.Context) {
	var mMap = make(map[string]string)
	var subChat model.SubChat
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	userId, err1 := strconv.Atoi(mMap["userId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	chatId, err1 := strconv.Atoi(mMap["chatId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	subChat.UserId = int64(userId)
	subChat.Content = mMap["content"]
	subChat.Picture = mMap["picture"]
	subChat.ChatId = int64(chatId)
	if err := service.CreateSubChat(&subChat); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "回复评论成功！")
}

func DeleteSubChat(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	subId, err1 := strconv.Atoi(mMap["subId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.DeleteSubChat(int64(subId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "删除回复成功！")
}

func LikeChat(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	chatId, err1 := strconv.Atoi(mMap["chatId"])
	if err1 != nil {
		response.Fail(ctx, nil, err1.Error())
		return
	}
	if err := service.LikeChat(int64(chatId)); util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, nil, "点赞成功！")
}

func SendChatByTitle(ctx *gin.Context) {
	var mMap = make(map[string]string)
	json.NewDecoder(ctx.Request.Body).Decode(&mMap)
	title := mMap["title"]
	chats, err := service.GetChatsByTitle(title)
	if util.IsFailed(err) {
		response.Fail(ctx, nil, err.Message)
		return
	}
	response.Success(ctx, gin.H{
		"chats": *chats,
	}, "按标题获取动态成功！")
}
