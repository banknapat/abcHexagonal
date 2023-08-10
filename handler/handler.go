package handler

import (
	"abc/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type handlerAdaper struct {
	s service.IServicePort
}

func NewHandlerAdaper(s service.IServicePort) *handlerAdaper {
	return &handlerAdaper{s: s}
}

// Get all data
func (h handlerAdaper) GetHand(c *gin.Context) {
	data, err := h.s.GetSer()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}

// Get data by id
func (h handlerAdaper) GetHandById(c *gin.Context) {
	id := c.Param("id")
	data, err := h.s.GetSerById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"data": data})
}
