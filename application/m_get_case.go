package application

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"github.com/toxb11/awesome_project/vo"
)

func MGetCase(ctx context.Context, req *vo.MGetCaseRequest) *vo.MGetCaseResponse {
	if err := checkMGetCaseReq(ctx, req); err != nil {
		return genMGetCaseResponse(nil, err)

	}
	param, err := convertReqToQueryParam(ctx, req)
	if err != nil {
		logrus.Errorf("[MGetCase] convert req err : %v\n", err)
		return genMGetCaseResponse(nil, err)
	}
	caseList, err := db.CaseRepo.MGetCase(ctx, param)
	if err != nil {
		logrus.Errorf("[MGetCase] db handler query err: %v\n", err)
		return genMGetCaseResponse(nil, err)
	}
	return genMGetCaseResponse(caseList, nil)
}

func convertReqToQueryParam(ctx context.Context, request *vo.MGetCaseRequest) (*model.CaseQueryParam, error) {
	ids := request.CaseIds
	return &model.CaseQueryParam{
		CaseIdList: ids,
	}, nil
}

func genMGetCaseResponse(caseList []*model.Case, err error) *vo.MGetCaseResponse {
	if err != nil {
		return &vo.MGetCaseResponse{
			BaseResponse: vo.BaseResponse{
				Status: "failed",
				ErrMsg: err.Error(),
			},
		}
	}
	caseVos := make([]*vo.CaseVo, 0, len(caseList))
	for _, do := range caseList {
		caseVo := &vo.CaseVo{
			CaseId:             do.Id,
			CaseName:           do.Name,
			CaseStatus:         int32(do.Status),
			CriminalId:         do.CriminalId,
			CriminalName:       do.CriminalName,
			CriminalPictureUrl: do.CriminalPictureUrlStr,
			ExtraDescription:   do.ExtraInfo,
		}
		caseVos = append(caseVos, caseVo)
	}
	return &vo.MGetCaseResponse{
		CaseList: caseVos,
		BaseResponse: vo.BaseResponse{
			Status: "success",
		},
	}
}

func checkMGetCaseReq(ctx context.Context, req *vo.MGetCaseRequest) error {
	if req == nil {
		logrus.Errorf("[checkMGetCaseReq] request empty\n")
		return errors.New("empty request")
	}
	logrus.Infof("[checkMGetCaseReq] req: %v", xjson.ToString(req))
	if len(req.CaseIds) == 0 && req.CriminalId == nil && req.TimeRange == nil {
		logrus.Errorf("[checkMGetCaseReq] request invalid\n")
		return errors.New("request invalid")
	}
	return nil
}
