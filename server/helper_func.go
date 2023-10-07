package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/GopherNet-golang/dto"
)

// Wraps error nicely
func (svc *HttpService) ErrorWithJson(c *gin.Context, code int, err error) {

	c.AbortWithStatusJSON(code, &dto.SuccessResponse{
		Success: false,
		Error:   err,
	})

}
