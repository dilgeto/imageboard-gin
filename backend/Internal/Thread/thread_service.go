package thread

import (
	"time"

	data "github.com/dilgeto/imageboard-gin/backend/Data"
)

type Thread = data.Thread

type IThreadRepository interface {
	saveThread(Thread) (*Thread, error)
	getThreadById(uint64) (*Thread, error)
	getAllThreads() ([]Thread, error)
	updateThread(Thread) error
	deleteThreadById(uint64) error
}

type Service struct {
	Repository IThreadRepository
}

func (serv *Service) saveThread(t Thread) (*Thread, error) {
	t.Timestamp = uint64(time.Now().Unix())
	return serv.Repository.saveThread(t)
}

func (serv *Service) getThreadById(id uint64) (*Thread, error) {
	return serv.Repository.getThreadById(id)
}

func (serv *Service) getAllThreads() ([]Thread, error) {
	return serv.Repository.getAllThreads()
}

func (serv *Service) updateThread(t Thread) error {
	return serv.Repository.updateThread(t)
}

func (serv *Service) deleteThreadById(id uint64) error {
	return serv.Repository.deleteThreadById(id)
}
