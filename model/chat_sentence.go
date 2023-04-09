package model

type ChatSentence struct {
	ChatId                         int64
	SentenceId                     int64
	CaseId                         int64
	CriminalId                     int64
	OriDescribeMsg                 string
	DescribeMsg                    string
	GenerateCriminalPictureUrlList []string
	ExtraInfo                      string
}

func (cs *ChatSentence) ConcatUrlListToStr() string {
	res := ""
	for _, s := range cs.GenerateCriminalPictureUrlList {
		res += s
	}
	return res
}
