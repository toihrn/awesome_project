package bd_face

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/sirupsen/logrus"
	"github.com/toxb11/awesome_project/infra/utils/http_utils"
	"github.com/toxb11/awesome_project/infra/utils/xjson"
	"image/png"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	bdkey                  = "0C9GBihTCG63tl4hZ79vDBdm"
	bdsk                   = "SrQD4QEiS3ltPLeOKiIs9mcuXZylP1YP"
	bduid                  = "108974fedc0a4ee38517a22490081be1"
	bdrgFaceUrlPrefix      = "https://aip.baidubce.com/rest/2.0/face/v3/faceset/user/add?access_token="
	bdSimilarFaceUrlPrefix = "https://aip.baidubce.com/rest/2.0/face/v3/search?access_token="
)

var bdApiClient *http.Client

func InitBdClient() {
	bdApiClient = &http.Client{}
}

func RegisterFace(ctx context.Context, imageUrl string) (faceToken string, err error) {
	file, err := http_utils.GetImageFileByUrl(imageUrl)
	if err != nil {
		return "", err
	}
	imgDecode, err := png.Decode(file)
	if err != nil {
		logrus.Errorf("[RegisterFace] decode err: %v\n", err)
		return "", err
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, imgDecode)
	if err != nil {
		logrus.Errorf("[RegisterFace] encode err: %v\n", err)
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

	httpUrl := fmt.Sprintf("%v%v", bdrgFaceUrlPrefix, GetAccessToken())
	payload := strings.NewReader(fmt.Sprintf(`{"group_id":"1","image":"%v","image_type":"BASE64","user_id":"%v"}`, base64Str, bduid))
	req, err := http.NewRequest("POST", httpUrl, payload)
	if err != nil {
		logrus.Errorf("[RegisterFace] err: %v\n", err)
		return "", err
	}
	res, err := bdApiClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("[RegisterFace] read resp body err: %v\n", err)
		return "", err
	}
	s := string(body)
	resp := &BdFaceResp{}
	err = sonic.UnmarshalString(s, &resp)
	if err != nil {
		logrus.Errorf("[RegisterFace] unmarshal err:%v\n", err)
		return "", err
	}
	if resp == nil || resp.ErrorMsg != "SUCCESS" || resp.Result.FaceToken == "" {
		logrus.Errorf("[RegisterFace] err, resp: %v\n", xjson.ToString(resp))
		return "", errors.New("api rg face failed")
	}
	return resp.Result.FaceToken, nil
}

func SimilarFace(ctx context.Context, imageUrl string) (faceToken string, err error) {
	file, err := http_utils.GetImageFileByUrl(imageUrl)
	if err != nil {
		return "", err
	}
	imgDecode, err := png.Decode(file)
	if err != nil {
		logrus.Errorf("[SimilarFace] decode err: %v\n", err)
		return "", err
	}
	var buf bytes.Buffer
	err = png.Encode(&buf, imgDecode)
	if err != nil {
		logrus.Errorf("[SimilarFace] encode err: %v\n", err)
		return "", err
	}
	base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
	httpUrl := fmt.Sprintf("%v%v", bdSimilarFaceUrlPrefix, GetAccessToken())
	payload := strings.NewReader(fmt.Sprintf(`{"group_id":"1","image":"%v","image_type":"BASE64","user_id":"%v"}`, base64Str, bduid))
	req, err := http.NewRequest("POST", httpUrl, payload)
	if err != nil {
		logrus.Errorf("[SimilarFace] err: %v\n", err)
		return "", err
	}
	res, err := bdApiClient.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logrus.Errorf("[SimilarFace] read resp body err: %v\n", err)
		return "", err
	}
	s := string(body)
	resp := &BdFaceResp{}
	err = sonic.UnmarshalString(s, &resp)
	if err != nil {
		logrus.Errorf("[SimilarFace] unmarshal err:%v\n", err)
		return "", err
	}
	if resp == nil || resp.ErrorMsg != "SUCCESS" || len(resp.Result.UserList) == 0 {
		return "", nil
	}
	if resp.Result.UserList[0].Score > 10.0 {
		logrus.Infof("[SimilarFace] get response greater than 10.0, score: %v\n", resp.Result.UserList[0].Score)
		return resp.Result.FaceToken, nil
	}
	return "", nil
}

/*
func main() {

    url := "https://aip.baidubce.com/rest/2.0/face/v3/search?access_token=" + GetAccessToken()
    payload := strings.NewReader(`{"group_id_list":"1,2,3,4,5,6","image":"123456789","image_type":"BASE64"}`)
    client := &http.Client {}
    req, err := http.NewRequest("POST", url, payload)

    if err != nil {
        fmt.Println(err)
        return
    }
    req.Header.Add("Content-Type", "application/json")

    res, err := client.Do(req)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer res.Body.Close()

    body, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(body))
}

*/

type BdFaceResp struct {
	ErrorMsg string `json:"error_msg,omitempty"`
	Result   Result `json:"result"`
}

type Result struct {
	FaceToken string    `json:"face_token,omitempty"`
	UserList  []*BdUser `json:"user_list,omitempty"`
}

type BdUser struct {
	Score float64 `json:"score,omitempty"`
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", bdkey, bdsk)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
