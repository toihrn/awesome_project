package vo

type CompleteChatRequest struct {
	UserId  int64
	Message string
}

type CompleteChatResponse struct {
	Answer string
}
