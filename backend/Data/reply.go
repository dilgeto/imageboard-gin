package data

type Reply struct {
	Id        uint32
	Number    uint32
	File      string
	Username  string
	Timestamp uint64
	Comment   string
	Id_thread uint64
}
