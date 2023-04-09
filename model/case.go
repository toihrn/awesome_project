package model

import "strings"

type CaseStatus int32

const (
	CaseInit       CaseStatus = 0
	CaseProcessing CaseStatus = 100
	CaseComplete   CaseStatus = 200
)

type Case struct {
	Id                    int64
	Name                  string
	CriminalId            int64
	CriminalName          string
	CriminalAddress       *string
	CriminalPictureUrlStr *string
	CriminalPictureBase64 *string
	CriminalPhone         *string
	CriminalGender        string
	Status                CaseStatus
	ExtraInfo             *string
}

func (c *Case) CriminalPictureUrlList() []string {
	if c.CriminalPictureUrlStr != nil {
		urlList := strings.Split(*c.CriminalPictureUrlStr, ",")
		return urlList
	}
	return nil
}

func (c *Case) SetCriminalPictureByStrSlice(urlList []string) {
	str := ""
	for i, s := range urlList {
		str += s
		if i < len(urlList)-1 {
			str += ","
		}
	}
	c.CriminalPictureUrlStr = &str
}

type CaseQueryParam struct {
	CaseIdList []int64
	QueryTime  *TimeUnion
	CriminalId *int64
}

type TimeUnion struct {
	BeginTime int64
	EndTime   int64
}
