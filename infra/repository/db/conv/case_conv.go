package conv

import (
	"context"
	"github.com/toxb11/awesome_project/infra/repository/db/db_entity"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"math/rand"
)

func ConvertCaseDOToEntity(ctx context.Context, caseDO *model.Case) (*db_entity.CaseEntity, error) {
	d := &db_entity.CaseEntity{
		ID:             caseDO.Id,
		Name:           caseDO.Name,
		CriminalId:     rand.Int63n(10),
		CriminalGender: caseDO.CriminalGender,
		CriminalName:   caseDO.CriminalName,
		Status:         int32(caseDO.Status),
		CriminalPictureUrl: func() string {
			if caseDO.CriminalPictureUrlStr == nil {
				return ""
			}
			return *caseDO.CriminalPictureUrlStr
		}(),
		ExtraInfo: xjson.ToString(caseDO.ExtraInfo),
	}
	if caseDO.CriminalAddress != nil {
		d.CriminalAddress = *caseDO.CriminalAddress
	}
	if caseDO.CriminalPhone != nil {
		d.CriminalPhone = *caseDO.CriminalPhone
	}
	return d, nil
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
		CriminalPictureUrlStr: &caseEntity.CriminalPictureUrl,
		CriminalPhone:         &caseEntity.CriminalPhone,
		CriminalGender:        caseEntity.CriminalGender,
		Status:                model.CaseStatus(caseEntity.Status),
		ExtraInfo:             &caseEntity.ExtraInfo,
	}, nil
}
