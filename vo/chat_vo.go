package vo

type DescribeCriminalRequest struct {
	CaseId             int64  `json:"case_id"`
	ChatId             int64  `json:"chat_id,omitempty"`
	SentenceId         int64  `json:"sentence_id,omitempty"`
	PreviousSentenceId int64  `json:"previous_sentence_id"`
	Message            string `json:"message,omitempty"`
}

type DescribeCriminalResponse struct {
	CriminalPictureUrlList []string       `json:"criminal_picture_url_list"`
	ChatId                 int64          `json:"chat_id"`
	SentenceId             int64          `json:"sentence_id"`
	TagMap                 map[string]int `json:"tag_map"`
	BaseResponse           BaseResponse   `json:"base_response"`
}

type ConfirmPictureRequest struct {
	ChatId                          int64  `json:"chat_id"`
	SentenceId                      int64  `json:"sentence_id"`
	OriPictureUrl                   string `json:"ori_picture_url"`
	NeedVariablePicturesNum         int64  `json:"need_variable_picture_num"`
	NeedReasoningSimilarPicturesNum int64  `json:"need_reasoning_similar_picture_num"`
}

type ConfirmPictureResponse struct {
	ChatId                            int64        `json:"chat_id"`
	VariablePictureUrlList            []string     `json:"variable_picture_url_list"`
	ReasoningSimilarPictureBase64List []string     `json:"reasoning_similar_picture_url_list"`
	BaseResponse                      BaseResponse `json:"base_response"`
}

type SaveCriminalPictureRequest struct {
	CaseId         int64              `json:"case_id,omitempty"`
	ChatPictureMap map[int64][]string `json:"chat_picture_map,omitempty"`
}

type SaveCriminalPictureResponse struct {
	BaseResponse BaseResponse `json:"base_response"`
}

type MGetChatRequest struct {
	ChatIds []int
}
