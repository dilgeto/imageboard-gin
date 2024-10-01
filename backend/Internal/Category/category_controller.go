package category

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ICategoryService interface {
	saveCategory(Category) (*Category, error)
	getCategoryById(uint64) (*Category, error)
	getAllCategories() ([]Category, error)
	updateCategory(Category) error
	deleteCategoryById(uint64) error
}

type Controller struct {
	Service ICategoryService
}

func (cntrl *Controller) postCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON  to category: - %v", err)
		return
	}

	newCategory, err := cntrl.Service.saveCategory(category)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving category: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, newCategory)
}

func (cntrl *Controller) getCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to int: - %v", err)
		return
	}

	category, err := cntrl.Service.getCategoryById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting category by id: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, category)
}

func (cntrl *Controller) getAllCategories(c *gin.Context) {
	categories, err := cntrl.Service.getAllCategories()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all categories: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}

func (cntrl *Controller) updateCategory(c *gin.Context) {
	var category Category
	if err := c.BindJSON(&category); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to category: - %v", err)
		return
	}

	if err := cntrl.Service.updateCategory(category); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating category: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, category)
}

func (cntrl *Controller) deleteCategoryById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to unsigned int: - %v", err)
		return
	}

	err = cntrl.Service.deleteCategoryById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting category with id %d: - %v", id, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/categories", cntrl.postCategory)
	rout.GET("/categories/:id", cntrl.getCategoryById)
	rout.GET("/categories", cntrl.getAllCategories)
	rout.PUT("/categories", cntrl.updateCategory)
	rout.DELETE("/categories/:id", cntrl.deleteCategoryById)
}
