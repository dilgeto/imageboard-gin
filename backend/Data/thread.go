package data

type Thread struct {
	Id_thread   uint64
	Filepath    string
	Subject     string
	Username    string
	Timestampp  uint64
	Commenta    string
	Reply_count uint16
	Image_count uint16
	Is_archived bool
	Is_pinned   bool
	Code        string
}
