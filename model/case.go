package model

type CaseStatus int32

const (
	CaseInit       CaseStatus = 0
	CaseProcessing CaseStatus = 100
	CaseComplete   CaseStatus = 200
)

type Case struct {
	Id                 int64
	Name               string
	CriminalId         int64
	CriminalName       string
	CriminalAddress    *string
	CriminalPictureUrl *string
	CriminalPhone      *string
	CriminalGender     string
	Status             CaseStatus
	ExtraInfo          *string
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
