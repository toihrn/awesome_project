package application

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/ai"
	"github.com/toxb11/awesome_project/vo"
	"sync"
)

func ConfirmPicture(ctx context.Context, req *vo.ConfirmPictureRequest) *vo.ConfirmPictureResponse {
	if err := checkConfirmPictureReq(ctx, req); err != nil {
		logrus.Errorf("[ConfirmPicture] req check err: %v\n", err)
		return genConfirmPictureErrResp(err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// variable
	res := true
	var vErr error
	variableUrlList := make([]string, 0)
	go func() {
		defer wg.Done()
		if req.NeedVariablePicturesNum > 0 {
			variableUrlList, vErr = ai.VariablePictures(ctx, req.OriPictureUrl, int(req.NeedVariablePicturesNum))
			if vErr != nil {
				res = false
			}
		}
	}()

	go func() {
		defer wg.Done()
		if req.NeedReasoningSimilarPicturesNum > 0 {
			logrus.Infof("[ConfirmPicture] not support similar reason yet")
		}
	}()
	wg.Wait()
	if !res {
		return genConfirmPictureErrResp(errors.New("VariablePictures err"))
	}

	resp := &vo.ConfirmPictureResponse{
		ChatId:                         req.ChatId,
		VariablePictureUrlList:         variableUrlList,
		ReasoningSimilarPictureUrlList: nil,
		BaseResponse: vo.BaseResponse{
			Status: vo.ResponseSuccess,
		},
	}
	return resp
}

func genConfirmPictureErrResp(err error) *vo.ConfirmPictureResponse {
	return &vo.ConfirmPictureResponse{
		BaseResponse: vo.BaseResponse{
			Status: vo.ResponseFailed,
			ErrMsg: err.Error(),
		},
	}
}

//func genConfirmPictureNormalResp() {
//
//}

func checkConfirmPictureReq(ctx context.Context, req *vo.ConfirmPictureRequest) error {
	if req == nil {
		return errors.New("req empty")
	}
	if req.OriPictureUrl == "" || req.SentenceId == 0 || req.ChatId == 0 {
		return errors.New("req invalid")
	}
	return nil
}
