package board

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IBoardService interface {
	saveBoard(Board) (*Board, error)
	getBoardById(uint64) (*Board, error)
	getAllBoards() (*Board, error)
	updateBoard(Board) error
	deleteBoardById(uint64) error
}

type Controller struct {
	Service IBoardService
}

func (cntrl *Controller) postBoard(c *gin.Context) {
	var board Board
	if err := c.BindJSON(&board); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON  to board: - %v", err)
		return
	}

	newBoard, err := cntrl.Service.saveBoard(board)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving board: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, newBoard)
}

func (cntrl *Controller) getBoardById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to int: - %v", err)
		return
	}

	board, err := cntrl.Service.getBoardById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting board by id: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, board)
}

func (cntrl *Controller) getAllBoards(c *gin.Context) {
	boards, err := cntrl.Service.getAllBoards()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all borads: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, boards)
}

func (cntrl *Controller) updateBoard(c *gin.Context) {
	var board Board
	if err := c.BindJSON(&board); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to board: - %v", err)
		return
	}

	if err := cntrl.Service.updateBoard(board); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating board: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, board)
}

func (cntrl *Controller) deleteBoardById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to unsigned int: - %v", err)
		return
	}

	err = cntrl.Service.deleteBoardById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting board with id %d: - %v", id, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/boards", cntrl.postBoard)
	rout.GET("/boards/:id", cntrl.getBoardById)
	rout.GET("/boards", cntrl.getAllBoards)
	rout.PUT("/boards", cntrl.updateBoard)
	rout.DELETE("/boards/:id", cntrl.deleteBoardById)
}
