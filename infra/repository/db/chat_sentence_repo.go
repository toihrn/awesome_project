package db

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db/conv"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/model"
)

type chatSentenceRepositoryImpl struct {
}

var ChatSentenceRepo *chatSentenceRepositoryImpl

func InitChatSentenceRepo() {
	ChatSentenceRepo = &chatSentenceRepositoryImpl{}
}

func (r *chatSentenceRepositoryImpl) GetSentence(ctx context.Context, chatId int64, sentenceId int64) (*model.ChatSentence, error) {
	chatSentenceEntity := &db_entity.ChatSentenceEntity{}
	tx := CaseDBHandler.Table(chatSentenceEntity.TableName()).Where("chat_id = ? AND sentence_id = ?", chatId, sentenceId).Find(&chatSentenceEntity)
	if tx.Error != nil {
		logrus.Errorf("[GetSentence] err: %v\n", tx.Error)
		return nil, tx.Error
	}
	if chatSentenceEntity == nil {
		logrus.Infof("[GetSentence] query nil by chatId %v, sentenceId: %v", chatId, sentenceId)
		return nil, nil
	}
	do, err := conv.ConvertChatSentenceEntityToDO(ctx, chatSentenceEntity)
	if err != nil {
		logrus.Errorf("[GetSentence] convert err: %v\n", err)
		return nil, errors.New("convert entity err")
	}
	return do, nil
}

func (r *chatSentenceRepositoryImpl) SaveSentence(ctx context.Context, chatSentence *model.ChatSentence) error {
	chatSentenceEntity, err := conv.ConvertChatSentenceDOToEntity(ctx, chatSentence)
	if err != nil {
		logrus.Errorf("[SaveSentence] convert entity err: %v\n", err)
		return err
	}
	tx := CaseDBHandler.Table(chatSentenceEntity.TableName()).Save(chatSentenceEntity)
	if tx.Error != nil {
		logrus.Errorf("[SaveSentence] save err: %v\n", err)
		return tx.Error
	}
	return nil
}
