package board

import data "github.com/dilgeto/imageboard-gin/backend/Data"

type Board = data.Board

type IBoardRepository interface {
	saveBoard(Board) (*Board, error)
	getBoardById(uint64) (*Board, error)
	getAllBoards() ([]Board, error)
	updateBoard(Board) error
	deleteBoardById(uint64) error
}

type Service struct {
	Repository IBoardRepository
}

func (serv *Service) saveBoard(b Board) (*Board, error) {
	return serv.Repository.saveBoard(b)
}

func (serv *Service) getBoardById(id uint64) (*Board, error) {
	return serv.Repository.getBoardById(id)
}

func (serv *Service) getAllBoards() ([]Board, error) {
	return serv.Repository.getAllBoards()
}

func (serv *Service) updateBoard(b Board) error {
	return serv.Repository.updateBoard(b)
}

func (serv *Service) deleteBoardById(id uint64) error {
	return serv.Repository.deleteBoardById(id)
}
