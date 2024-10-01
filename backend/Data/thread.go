package data

type Thread struct {
	Id         uint64
	Subject    string
	Username   string
	Timestamp  uint64
	Comment    string
	ReplyCount uint16
	ImageCount uint16
	IsArchived bool
	IsPinned   bool
	BoardCode  string
}
