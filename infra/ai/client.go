package ai

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/utils/http_utils"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"github.com/toxb11/awesome_project/model"
)

var OpenAIClient *openai.Client

const (
	openaiAPIKey = "T0wcbTSzr07H2ElfhT9kyJ0T3BlbkFJy6CMb3wktQQDm0YUhsk-"
	//gpt4Model    = "gpt-4-32k"
	gpt35Model = "gpt-3.5-turbo"
)

func InitOpenAIClient() {
	OpenAIClient = openai.NewClient(skvus(openaiAPIKey))
}

func skvus(key string) string {
	prefix := key[0:3]
	suffix := key[len(key)-3:]
	res := suffix + key[3:len(key)-3] + prefix
	logrus.Infof("[skvus] res: %v\n", res)
	return res
}

func GeneratePromptBySum(ctx context.Context, prevMsg string, curMsg string) (string, error) {
	if prevMsg == "" {
		return curMsg, nil
	}
	req := openai.ChatCompletionRequest{
		Model: gpt35Model,
		Messages: []openai.ChatCompletionMessage{{
			Role:    "user",
			Content: fmt.Sprintf("下面两句话都是关于人脸的描述，请将这两句合并。\n第一句：%v\n第二句：%v\n", curMsg, prevMsg),
		}},
	}
	response, err := OpenAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		logrus.Errorf("[GeneratePromptBySum] err: %v, use force contact msg\n", err)
		return fmt.Sprintf("%v,%v", curMsg, prevMsg), err
	}
	logrus.Infof("[GeneratePromptBySum] req: %v,resposne: %v\n", xjson.ToString(req), xjson.ToString(response))
	completionMessage := response.Choices[0].Message
	return completionMessage.Content, nil
}

func GetPromptDescribeTags(ctx context.Context, prompt string) (*model.FaceTag, error) {
	req := openai.ChatCompletionRequest{
		Model: gpt35Model,
		Messages: []openai.ChatCompletionMessage{{
			Role:    "user",
			Content: fmt.Sprintf("请分析下面句话是否有形容到一个人的年纪，性别，身体体型，脸型，头发样式，头发颜色，是否有胡须，眉毛样式，眉毛方向，眼睛形状，眉眼间距，是否双眼皮、是否有眼袋，鼻子形状，鼻眼间距，法令纹，皱纹，嘴巴大小,嘴唇厚薄,神态描述。请使用只有一层的英文json结构表示哪些属性被形容到（1代表被形容到），哪些属性没有被形容到（2代表没有被形容到）。我只需要json结构，不需要文本说明。\n%v", prompt),
		}},
	}
	response, err := OpenAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		logrus.Errorf("[GetPromptDescribeTags] err: %v\n", err)
		return nil, err
	}
	logrus.Infof("[GetPromptDescribeTags] req: %v, response: %v\n", xjson.ToString(req), xjson.ToString(response))
	tagStr := response.Choices[0].Message.Content
	faceTag := &model.FaceTag{}
	err = sonic.UnmarshalString(tagStr, &faceTag)
	if err != nil {
		logrus.Errorf("[GetPromptDescribeTags] unmarshal err: %v\n", err)
		return nil, err
	}
	return faceTag, nil
}

func GeneratePictures(ctx context.Context, prompt string, generateNum int64) ([]string, error) {
	prompt = fmt.Sprintf("按照下面关键词描述生成半身证件照。%v", prompt)
	req := openai.ImageRequest{
		Prompt:         prompt,
		N:              int(generateNum),
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatURL,
	}
	response, err := OpenAIClient.CreateImage(ctx, req)
	if err != nil {
		logrus.Errorf("[GeneratePictures] err: %v\n", err)
		return nil, err
	}
	logrus.Infof("[GeneratePictures] req: %v , response picture num: %v\n", xjson.ToString(req), len(response.Data))
	urls := make([]string, 0, generateNum)
	for _, data := range response.Data {
		urls = append(urls, data.URL)
	}
	return urls, nil
}

func VariablePictures(ctx context.Context, oriPictureUrl string, variableNum int) ([]string, error) {
	tempFile, err := http_utils.GetPngImageFileByUrl(oriPictureUrl)
	if tempFile != nil {
		defer tempFile.Close()
	}
	if err != nil {
		return nil, err
	}
	req := openai.ImageVariRequest{
		Image: tempFile,
		N:     variableNum,
		Size:  openai.CreateImageSize256x256,
	}
	response, err := OpenAIClient.CreateVariImage(ctx, req)
	if err != nil {
		logrus.Errorf("[VariablePictures] err; %v\n", err)
		return nil, err
	}
	resUrlList := make([]string, 0, variableNum)
	for _, datum := range response.Data {
		resUrlList = append(resUrlList, datum.URL)
	}
	return resUrlList, nil
}
