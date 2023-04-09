package application

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/ai"
	"github.com/toxb11/awesome_project/infra/repository/db"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
	"github.com/toxb11/awesome_project/vo"
	"sync"
)

func DescribeCriminal(ctx context.Context, req *vo.DescribeCriminalRequest) *vo.DescribeCriminalResponse {
	if err := checkDescribeCriminal(ctx, req); err != nil {
		return genDescribeCriminalErrResp(err)
	}
	curSentence, err := convertReqToCurSentence(ctx, req)
	prevSentence, err := db.ChatSentenceRepo.GetSentence(ctx, req.ChatId, req.PreviousSentenceId)
	if err != nil {
		return genDescribeCriminalErrResp(err)
	}
	prompt := ""
	if prevSentence != nil {
		prompt, err = ai.GeneratePromptBySum(ctx, prevSentence.DescribeMsg, curSentence.OriDescribeMsg)
		if err != nil {
			prompt = curSentence.DescribeMsg
		}
	}
	curSentence.DescribeMsg = prompt
	res := true
	var faceTag *model.FaceTag
	pictureUrlList := make([]string, 0)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	// 取脸部tag弱依赖
	go func() {
		defer wg.Done()
		faceTag, _ = ai.GetPromptDescribeTags(ctx, prompt)
	}()

	// 生成图片
	go func() {
		defer wg.Done()
		pictureUrlList, err = ai.GeneratePictures(ctx, prompt, 4)
		if err != nil {
			res = false
			return
		}
	}()
	wg.Wait()
	if !res {
		return genDescribeCriminalErrResp(errors.New("generate picture err"))
	}

	curSentence.GenerateCriminalPictureUrlList = pictureUrlList

	// 当前chatId-sentenceId对话记录插数据库
	err = db.ChatSentenceRepo.SaveSentence(ctx, curSentence)
	if err != nil {
		return genDescribeCriminalErrResp(err)
	}

	resultTag := faceTag.ToMap()

	return &vo.DescribeCriminalResponse{
		CriminalPictureUrlList: pictureUrlList,
		ChatId:                 req.ChatId,
		SentenceId:             req.SentenceId,
		TagMap:                 resultTag,
		BaseResponse:           vo.BaseResponse{Status: vo.ResponseSuccess},
	}
}

func convertReqToCurSentence(ctx context.Context, req *vo.DescribeCriminalRequest) (*model.ChatSentence, error) {
	curSentence := &model.ChatSentence{
		ChatId:         req.ChatId,
		CaseId:         req.CaseId,
		SentenceId:     req.SentenceId,
		OriDescribeMsg: req.Message,
	}
	return curSentence, nil
}

func genDescribeCriminalErrResp(err error) *vo.DescribeCriminalResponse {
	return &vo.DescribeCriminalResponse{
		BaseResponse: vo.BaseResponse{
			Status: vo.ResponseFailed,
			ErrMsg: err.Error(),
		},
	}
}

func genDescribeCriminalResp(pictureUrlList []string, chatId int64, sentenceId int64) *vo.DescribeCriminalResponse {
	return &vo.DescribeCriminalResponse{
		CriminalPictureUrlList: pictureUrlList,
		ChatId:                 chatId,
		SentenceId:             sentenceId,
		BaseResponse: vo.BaseResponse{
			Status: vo.ResponseSuccess,
		},
	}
}

func checkDescribeCriminal(ctx context.Context, req *vo.DescribeCriminalRequest) error {
	if req == nil {
		logrus.Errorf("[checkDescribeCriminal] req empty")
		return errors.New("empty req")
	}
	if req.ChatId < 0 || req.SentenceId < 0 {
		logrus.Errorf("[checkDescribeCriminal] req invalid, chat or sentence nil, req: %v\n", xjson.ToString(req))
		return errors.New("invalid req")
	}
	if req.Message == "" {
		logrus.Warnf("[checkDescribeCriminal] message in req empty, req: %v\n", xjson.ToString(req))
		return errors.New("message empty")
	}
	return nil
}
