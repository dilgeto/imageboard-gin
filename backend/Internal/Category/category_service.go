package category

import data "github.com/dilgeto/imageboard-gin/backend/Data"

type Category = data.Category

type ICategoryRepository interface {
	saveCategory(Category) (*Category, error)
	getCategoryById(uint64) (*Category, error)
	getAllCategories() ([]Category, error)
	updateCategory(Category) error
	deleteCategoryById(uint64) error
	getCategoryByName(string) (*Category, error)
}

type Service struct {
	Repository ICategoryRepository
}

func (serv *Service) saveCategory(c Category) (*Category, error) {
	existe, err := serv.Repository.getCategoryByName(c.Name)
	if existe != nil {
		return nil, err
	}
	return serv.Repository.saveCategory(c)
}

func (serv *Service) getCategoryById(id uint64) (*Category, error) {
	return serv.Repository.getCategoryById(id)
}

func (serv *Service) getAllCategories() ([]Category, error) {
	return serv.Repository.getAllCategories()
}

func (serv *Service) updateCategory(c Category) error {
	return serv.Repository.updateCategory(c)
}

func (serv *Service) deleteCategoryById(id uint64) error {
	return serv.Repository.deleteCategoryById(id)
}
