package db_entity

import (
	"strings"
	"time"
)

type ChatSentenceEntity struct {
	ID                             int64     `gorm:"column:id;primary_key"`
	ChatId                         int64     `gorm:"column:chat_id"`
	SentenceId                     int64     `gorm:"column:sentence_id"`
	CaseId                         int64     `gorm:"column:case_id"`
	CriminalId                     int64     `gorm:"column:criminal_id"`
	OriDescribeMsg                 string    `gorm:"column:ori_describe_msg"`
	DescribeMsg                    string    `gorm:"column:describe_msg"`
	GenerateCriminalPictureUrlList string    `gorm:"column:criminal_picture_url_list"` // 当前对话生成的urllist，逗号分割
	CreateTime                     time.Time `gorm:"column:create_time;autoCreateTime"`
	Update                         time.Time `gorm:"column:update_time;autoUpdateTime"`
	ExtraInfo                      string    `gorm:"column:extra_info"`
}

func (c *ChatSentenceEntity) TableName() string {
	return "t_chat_sentence"
}

func (c *ChatSentenceEntity) GetGenerateCriminalPictureUrlList() []string {
	urlListStr := c.GenerateCriminalPictureUrlList
	if urlListStr == "" {
		return nil
	}
	split := strings.Split(urlListStr, ",")
	return split
}
