package reply

import data "github.com/dilgeto/imageboard-gin/backend/Data"

type Reply = data.Reply

type IReplyRepository interface {
	saveReply(Reply) (*Reply, error)
	getReplyById(uint64) (*Reply, error)
	getAllReplies() ([]Reply, error)
	updateReply(Reply) error
	deleteReplyById(uint64) error
}

type Service struct {
	Repository IReplyRepository
}

func (serv *Service) saveReply(r Reply) (*Reply, error) {
	return serv.Repository.saveReply(r)
}

func (serv *Service) getReplyById(id uint64) (*Reply, error) {
	return serv.Repository.getReplyById(id)
}

func (serv *Service) getAllReplies() ([]Reply, error) {
	return serv.Repository.getAllReplies()
}

func (serv *Service) updateReply(r Reply) error {
	return serv.Repository.updateReply(r)
}

func (serv *Service) deleteReplyById(id uint64) error {
	return serv.Repository.deleteReplyById(id)
}
