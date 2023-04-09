package db

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db/conv"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"time"
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
	tx := CaseDBHandler.Table(caseEntity.TableName()).Save(caseEntity)
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
	logrus.Infof("[MGetCase] queryparam: %v\n", xjson.ToString(queryParam))
	if len(queryParam.CaseIdList) > 0 {
		CaseDBHandler.Table(caseEntity.TableName()).Where("id IN ?", queryParam.CaseIdList).Find(&caseEntityList)
		return dbResHandler(ctx, caseEntityList, doList)
	}
	if queryParam.QueryTime != nil {
		begin := time.Unix(queryParam.QueryTime.BeginTime, 0)
		end := time.Unix(queryParam.QueryTime.EndTime, 0)
		CaseDBHandler.Table(caseEntity.TableName()).Where("create_time BETWEEN ? AND ?", begin.Format(time.DateTime), end.Format(time.DateTime))
		return dbResHandler(ctx, caseEntityList, doList)
	}
	if queryParam.CriminalId != nil {
		CaseDBHandler.Table(caseEntity.TableName()).Where("where criminal_id = ?", *queryParam.CriminalId)
		return dbResHandler(ctx, caseEntityList, doList)
	}
	return doList, nil
}

func (r *caseRepositoryImpl) GetCaseByFaceToken(ctx context.Context, faceToken string) (*model.Case, error) {
	caseEntity := &db_entity.CaseEntity{}
	tx := CaseDBHandler.Table(caseEntity.TableName()).Where("where face_token = ?", faceToken)
	if tx.Error != nil {
		logrus.Errorf("[GetCaseByFaceToken] err: %v\n", tx.Error)
		return nil, tx.Error
	}
	do, err := conv.ConvertCaseEntityToDO(ctx, caseEntity)
	if err != nil {
		logrus.Errorf("[GetCaseByFaceToken] convert err: %v\n", err)
		return nil, err
	}
	return do, nil
}

func dbResHandler(ctx context.Context, caseEntityList []*db_entity.CaseEntity, doList []*model.Case) ([]*model.Case, error) {
	for _, entity := range caseEntityList {
		do, err := conv.ConvertCaseEntityToDO(ctx, entity)
		if err != nil {
			logrus.Errorf("[MGetCase] convert from entity err: %v", err)
			return nil, err
		}
		doList = append(doList, do)
	}
	return doList, nil
}
