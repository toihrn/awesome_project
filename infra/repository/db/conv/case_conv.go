package conv

import (
	"context"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"math/rand"
)

func ConvertCaseDOToEntity(ctx context.Context, caseDO *model.Case) (*db_entity.CaseEntity, error) {
	return &db_entity.CaseEntity{
		ID:              caseDO.Id,
		Name:            caseDO.Name,
		CriminalId:      rand.Int63n(10),
		CriminalAddress: *caseDO.CriminalAddress,
		CriminalPhone:   *caseDO.CriminalPhone,
		CriminalGender:  caseDO.CriminalGender,
		CriminalName:    caseDO.CriminalName,
		Status:          int32(caseDO.Status),
		CriminalPictureUrl: func() string {
			if caseDO.CriminalPictureUrl == nil {
				return ""
			}
			return *caseDO.CriminalPictureUrl
		}(),
		ExtraInfo: xjson.ToString(caseDO.ExtraInfo),
	}, nil
}

func ConvertCaseEntityToDO(ctx context.Context, caseEntity *db_entity.CaseEntity) (*model.Case, error) {
	return &model.Case{
		Id:           caseEntity.ID,
		Name:         caseEntity.CriminalAddress,
		CriminalId:   caseEntity.CriminalId,
		CriminalName: caseEntity.CriminalName,
		CriminalAddress: func() *string {
			if caseEntity.CriminalAddress == "" {
				return nil
			}
			return &caseEntity.CriminalAddress
		}(),
		CriminalPictureUrl: &caseEntity.CriminalPictureUrl,
		CriminalPhone:      &caseEntity.CriminalPhone,
		CriminalGender:     caseEntity.CriminalGender,
		Status:             model.CaseStatus(caseEntity.Status),
		ExtraInfo:          &caseEntity.ExtraInfo,
	}, nil
}
