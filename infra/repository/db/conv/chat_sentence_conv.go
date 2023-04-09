package conv

import (
	"context"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/model"
)

func ConvertChatSentenceEntityToDO(ctx context.Context, entity *db_entity.ChatSentenceEntity) (*model.ChatSentence, error) {
	return &model.ChatSentence{
		ChatId:                         entity.ChatId,
		SentenceId:                     entity.SentenceId,
		CaseId:                         entity.CaseId,
		CriminalId:                     entity.CriminalId,
		OriDescribeMsg:                 entity.OriDescribeMsg,
		DescribeMsg:                    entity.DescribeMsg,
		GenerateCriminalPictureUrlList: entity.GetGenerateCriminalPictureUrlList(),
		ExtraInfo:                      entity.ExtraInfo,
	}, nil
}

func ConvertChatSentenceDOToEntity(ctx context.Context, chatSentenceDO *model.ChatSentence) (*db_entity.ChatSentenceEntity, error) {
	entity := &db_entity.ChatSentenceEntity{
		ChatId:                         chatSentenceDO.ChatId,
		SentenceId:                     chatSentenceDO.SentenceId,
		CaseId:                         chatSentenceDO.CaseId,
		CriminalId:                     chatSentenceDO.CriminalId,
		OriDescribeMsg:                 chatSentenceDO.OriDescribeMsg,
		DescribeMsg:                    chatSentenceDO.DescribeMsg,
		GenerateCriminalPictureUrlList: chatSentenceDO.ConcatUrlListToStr(),
	}
	return entity, nil
}
