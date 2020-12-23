package service

import (
	"log"
	"wefive/common"
	"wefive/model"
	"wefive/util"
)

// 获取热门评论
func GetHotChats() (*[]model.Chat, *util.Err) {
	// 未实现，获取所有评论
	db := common.GetDB()
	var chats []model.Chat
	err := db.Find(&chats).Error
	if err != nil {
		log.Println(err.Error())
		return nil, util.Fail(err.Error())
	}
	return &chats, util.Success()
}

func GetChatDto(chatId int64) (*model.ChatDto, *util.Err) {
	chat, err := GetChatByChatId(chatId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}
	chatDto, err := toChatDto(chat)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}
	return chatDto, util.Success()
}

func GetChatByChatId(chatId int64) (*model.Chat, *util.Err) {
	db := common.GetDB()
	var chat model.Chat
	err := db.Where("chat_id = ?", chatId).First(&chat).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &chat, util.Success()
}

func GetSubChatDtos(chatId int64) (*[]model.SubChatDto, *util.Err) {
	var subChats []model.SubChat
	var subChatDtos []model.SubChatDto
	db := common.GetDB()
	err := db.Where("chat_id = ?", chatId).Find(&subChats).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	for _, subChat := range subChats {
		subChatDto, err := toSubChatDto(&subChat)
		if util.IsFailed(err) {
			log.Println(err)
			return nil, err
		}
		subChatDtos = append(subChatDtos, *subChatDto)
	}
	return &subChatDtos, util.Success()
}

func CreateChat(chat *model.Chat) *util.Err {
	db := common.GetDB()
	err := db.Create(chat).Error
	if err != nil {
		log.Println(err)
		return util.Fail(err.Error())
	}
	return util.Success()
}

func toChatDto(chat *model.Chat) (*model.ChatDto, *util.Err) {
	userId := chat.UserId
	user, err := GetUserByUserId(userId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}
	var chatDto model.ChatDto
	chatDto.UserId = userId
	chatDto.Content = chat.Content
	chatDto.Picture = chat.Picture
	chatDto.ChatId = chat.ChatId
	chatDto.Discussions = chat.Discussions
	chatDto.Likes = chat.Likes
	chatDto.Name = user.Name
	chatDto.Avatar = user.Avatar
	return &chatDto, util.Success()
}

func toSubChatDto(subChat *model.SubChat) (*model.SubChatDto, *util.Err) {
	userId := subChat.UserId
	user, err := GetUserByUserId(userId)
	if util.IsFailed(err) {
		log.Println(err)
		return nil, err
	}
	var subChatDto model.SubChatDto
	subChatDto.UserId = userId
	subChatDto.Content = subChat.Content
	subChatDto.Picture = subChat.Picture
	subChatDto.ChatId = subChat.ChatId
	subChatDto.SubId = subChat.SubId
	subChatDto.Name = user.Name
	subChatDto.Avatar = user.Avatar
	return &subChatDto, util.Success()
}
