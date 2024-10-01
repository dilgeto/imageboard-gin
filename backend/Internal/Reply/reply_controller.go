package reply

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type IReplyService interface {
	saveReply(Reply) (*Reply, error)
	getReplyById(uint64) (*Reply, error)
	getAllReplies() ([]Reply, error)
	updateReply(Reply) error
	deleteReplyById(uint64) error
}

type Controller struct {
	Service IReplyService
}

func (cntrl *Controller) postReply(c *gin.Context) {
	var reply Reply
	if err := c.BindJSON(&reply); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to reply: - %v", err)
		return
	}

	newReply, err := cntrl.Service.saveReply(reply)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed saving reply: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, newReply)
}

func (cntrl *Controller) getReplyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to int: - %v", err)
		return
	}

	reply, err := cntrl.Service.getReplyById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting reply by id: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, reply)
}

func (cntrl *Controller) getAllReplies(c *gin.Context) {
	replies, err := cntrl.Service.getAllReplies()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed getting all replies: - %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, replies)
}

func (cntrl *Controller) updateReply(c *gin.Context) {
	var reply Reply
	if err := c.BindJSON(&reply); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing JSON to reply: - %v", err)
		return
	}

	if err := cntrl.Service.updateReply(reply); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed updating reply: - %v", err)
		return
	}
	c.IndentedJSON(http.StatusOK, reply)
}

func (cntrl *Controller) deleteReplyById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed parsing id to unsigned int: - %v", err)
		return
	}

	err = cntrl.Service.deleteReplyById(uint64(id))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
		log.Printf("Failed deleting reply with id %d: - %v", id, err)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (cntrl *Controller) LinkPaths(rout *gin.Engine) {
	rout.POST("/replies", cntrl.postReply)
	rout.GET("/replies/:id", cntrl.getReplyById)
	rout.GET("/replies", cntrl.getAllReplies)
	rout.PUT("/replies", cntrl.updateReply)
	rout.DELETE("/replies/:id", cntrl.deleteReplyById)
}
