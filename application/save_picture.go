package application

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/repository/db"
	"github.com/toxb11/awesome_project/model"
	"github.com/toxb11/awesome_project/vo"
)

func SavePicture(ctx context.Context, req *vo.SaveCriminalPictureRequest) *vo.SaveCriminalPictureResponse {
	if err := checkSavePictureReq(ctx, req); err != nil {
		logrus.Errorf("[SavePicture] check err: %v\n", err)
		return genSavePictureErrResp(err)
	}
	caseId := req.CaseId
	caseFileList, err := db.CaseRepo.MGetCase(ctx, &model.CaseQueryParam{
		CaseIdList: []int64{caseId},
	})
	if err != nil {
		logrus.Errorf("[SavePicture] err: %v\n", err)
		return genSavePictureErrResp(err)
	}

	cf := caseFileList[0]
	existUrlList := cf.CriminalPictureUrlList()
	chatPictureMap := req.ChatPictureMap
	for _, urlList := range chatPictureMap {
		existUrlList = append(existUrlList, urlList...)
	}
	cf.SetCriminalPictureByStrSlice(existUrlList)
	err = db.CaseRepo.SaveCase(ctx, cf)
	if err != nil {
		logrus.Errorf("[SavePicture] err: %v\n", err)
		return genSavePictureErrResp(err)
	}
	resp := &vo.SaveCriminalPictureResponse{BaseResponse: vo.BaseResponse{
		Status: vo.ResponseSuccess,
	}}
	return resp
}

func checkSavePictureReq(ctx context.Context, req *vo.SaveCriminalPictureRequest) error {
	if req == nil {
		return errors.New("empty req")
	}
	if req.CaseId < 0 || len(req.ChatPictureMap) == 0 {
		return errors.New("invalid req")
	}
	return nil
}

func genSavePictureErrResp(err error) *vo.SaveCriminalPictureResponse {
	return &vo.SaveCriminalPictureResponse{
		BaseResponse: vo.BaseResponse{
			Status: vo.ResponseFailed,
			ErrMsg: err.Error(),
		},
	}
}
