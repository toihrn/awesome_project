package application

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db"
	"github.com/toxb11/awesome_project/infra/utils/id_gen"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"github.com/toxb11/awesome_project/vo"
	"math/rand"
)

func CreateCase(ctx context.Context, req *vo.CreateCaseRequest) *vo.CreateCaseResponse {
	if req == nil || req.Name == "" {
		logrus.Errorf("[CreateCase] invalid req: %v\n", xjson.ToString(req))
		return genCreateCaseResp(ctx, errors.New("invalid request"))
	}
	do, err := convertCreateCaseReqToDO(ctx, req)
	if err != nil {
		logrus.Errorf("[CreateCase] convert req err: %v\n", err)
		return genCreateCaseResp(ctx, errors.New("convert req err"))
	}
	err = db.CaseRepo.SaveCase(ctx, do)
	if err != nil {
		logrus.Errorf("[CreateCase] save err: %v\n", err)
		return genCreateCaseResp(ctx, errors.New("save failed"))
	}
	return genCreateCaseResp(ctx, nil)
}

func convertCreateCaseReqToDO(ctx context.Context, req *vo.CreateCaseRequest) (*model.Case, error) {
	return &model.Case{
		Id:              id_gen.DoGen(),
		Name:            req.Name,
		CriminalName:    req.CriminalName,
		CriminalAddress: req.CriminalAddress,
		CriminalPhone:   req.CriminalPhone,
		CriminalGender:  req.CriminalGender,
		CriminalId:      int64(rand.Uint64() / 100),
		Status:          0,
		ExtraInfo:       req.Description,
	}, nil
}

func genCreateCaseResp(ctx context.Context, err error) *vo.CreateCaseResponse {
	if err != nil {
		resp := &vo.CreateCaseResponse{BaseResponse: vo.BaseResponse{
			Status: "failed",
			ErrMsg: err.Error(),
		}}
		return resp
	}
	return &vo.CreateCaseResponse{BaseResponse: vo.BaseResponse{
		Status: "success",
		ErrMsg: "",
	}}
}
