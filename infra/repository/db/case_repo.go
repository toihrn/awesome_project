package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db/conv"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/model"
)

type caseRepositoryImpl struct {
}

var CaseRepo *caseRepositoryImpl

func InitCaseRepo() {
	CaseRepo = &caseRepositoryImpl{}
}

func (r *caseRepositoryImpl) SaveCase(ctx context.Context, caseDO *model.Case) error {
	caseEntity, err := conv.ConvertCaseDOToEntity(ctx, caseDO)
	if err != nil {
		logrus.Errorf("[SaveCase] err: %v\n", err)
		return err
	}
	tx := CaseDBHandler.Table(caseEntity.TableName()).Create(caseEntity)
	if tx.Error != nil {
		logrus.Errorf("[SaveCase] save err; %v\n", tx.Error)
		return tx.Error
	}
	return nil
}

func (r *caseRepositoryImpl) MGetCase(ctx context.Context, queryParam *model.CaseQueryParam) ([]*model.Case, error) {
	caseEntity := &db_entity.CaseEntity{}
	caseEntityList := make([]*db_entity.CaseEntity, 0)
	doList := make([]*model.Case, 0)
	if len(queryParam.CaseIdList) > 0 {
		CaseDBHandler.Table(caseEntity.TableName()).Where("id IN ?", queryParam.CaseIdList).Find(&caseEntityList)
		for _, entity := range caseEntityList {
			do, err := conv.ConvertCaseEntityToDO(ctx, entity)
			if err != nil {
				logrus.Errorf("[MGetCase] convert from entity err: %v", err)
				return nil, err
			}
			doList = append(doList, do)
		}
	}
	return doList, nil
}
