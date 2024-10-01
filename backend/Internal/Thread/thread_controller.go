package thread

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IThreadService interface {
	saveThread(Thread) (*Thread, error)
	getThreadById(uint64) (*Thread, error)
	getAllThreads() (*Thread, error)
	updateThread(Thread) error
	deleteThreadById(uint64) error
}

type Controller struct {
	Service IThreadService
}

func (cntrl *Controller) postThread(c *gin.Context) {
	var thread Thread
	if err := c.BindJSON(&thread); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON  to thread: - %v", err)
		return
	}

	newThread, err := cntrl.Service.saveThread(thread)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving thread: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, newThread)
}

func (cntrl *Controller) getThreadById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to int: - %v", err)
		return
	}

	thread, err := cntrl.Service.getThreadById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting thread by id: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, thread)
}

func (cntrl *Controller) getAllThreads(c *gin.Context) {
	boards, err := cntrl.Service.getAllThreads()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all threads: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, boards)
}

func (cntrl *Controller) updateThread(c *gin.Context) {
	var thread Thread
	if err := c.BindJSON(&thread); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to thread: - %v", err)
		return
	}

	if err := cntrl.Service.updateThread(thread); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating thread: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, thread)
}

func (cntrl *Controller) deleteThreadById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to unsigned int: - %v", err)
		return
	}

	err = cntrl.Service.deleteThreadById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting thread with id %d: - %v", id, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/threads", cntrl.postThread)
	rout.GET("/threads/:id", cntrl.getThreadById)
	rout.GET("/threads", cntrl.getAllThreads)
	rout.PUT("/threads", cntrl.updateThread)
	rout.DELETE("/threads/:id", cntrl.deleteThreadById)
}
