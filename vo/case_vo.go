package vo

type CompleteChatRequest struct {
	UserId  int64
	Message string
}

type CompleteChatResponse struct {
	Answer string
}

type MGetCaseRequest struct {
	CaseIds    []int64    `json:"case_ids,omitempty"`
	TimeRange  *TimeRange `json:"time_range,omitempty"`
	CriminalId *int64     `json:"criminal_id,omitempty"`
}

type MGetCaseResponse struct {
	CaseList     []*CaseVo    `json:"case_list,omitempty"`
	BaseResponse BaseResponse `json:"base_response"`
}

type CaseVo struct {
	CaseId             int64   `json:"case_id,omitempty"`
	CaseName           string  `json:"case_name,omitempty"`
	CaseStatus         int32   `json:"case_status,omitempty"`
	CriminalId         int64   `json:"criminal_id,omitempty"`
	CriminalName       string  `json:"criminal_name,omitempty"`
	CriminalPictureUrl *string `json:"criminal_picture_url,omitempty"`
	ExtraDescription   *string `json:"extra_description,omitempty"`
}

type TimeRange struct {
	BeginTime int64 `json:"begin_time"`
	EndTime   int64 `json:"end_time"`
}

type CreateCaseRequest struct {
	Name            string  `json:"case_name"`
	Description     *string `json:"extra_description,omitempty"`
	CriminalName    string  `json:"criminal_name"`
	CriminalAddress *string `json:"criminal_address"`
	CriminalPhone   *string `json:"criminal_phone"`
	CriminalGender  string  `json:"criminal_gender"`
}

type ResponseStatus string

const (
	ResponseSuccess ResponseStatus = "success"
	ResponseFailed  ResponseStatus = "failed"
)

type BaseResponse struct {
	Status ResponseStatus `json:"status,omitempty"`
	ErrMsg string         `json:"err_msg,omitempty"`
}

type CreateCaseResponse struct {
	BaseResponse BaseResponse `json:"base_response"`
}
