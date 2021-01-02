package service

import (
	"gover-server/common"
	"gover-server/model"
	"gover-server/util"
	"log"
)

func GetCommentsByDeptId(deptId int64) (*[]model.Comment, *util.Err) {
	var comments []model.Comment
	db := common.GetDB()
	err := db.Where("dept_id = ?", deptId).Find(&comments).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &comments, util.Success()
}

func FeedBack(commentId int64, reply string) *util.Err {

	comment, err := GetCommentById(commentId)
	if util.IsFailed(err) {
		log.Println(err)
		return err
	}
	comment.Reply = reply
	comment.Delete1 = "1"
	db := common.GetDB()
	db.Where("comment_id = ?", commentId).Delete(&model.Comment{})
	err1 := db.Save(&comment).Error
	if err1 != nil {
		log.Println(err1)
		return util.Fail(err1.Error())
	}
	return util.Success()
}

func GetCommentById(commentId int64) (*model.Comment, *util.Err) {
	var comment model.Comment
	db := common.GetDB()
	err := db.Where("comment_id = ?", commentId).First(&comment).Error
	if err != nil {
		log.Println(err)
		return nil, util.Fail(err.Error())
	}
	return &comment, util.Success()
}
