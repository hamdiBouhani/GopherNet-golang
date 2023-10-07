package server

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/GopherNet-golang/dto"
)

// @Summary Rent Burrow
// @Accept  json
// @Produce json
// @Param uuid path int64 true "UUID"
// @Success 200 {object} model.Burrow
// @Router /rent-burrow/{id} [put]
func (svc *HttpService) RentBurrow(c *gin.Context) {

	id := c.Param("id")
	burrowId, err := strconv.ParseInt(id, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", burrowId, burrowId)
		svc.ErrorWithJson(c, 400, err)
		return
	}

	err = svc.BurrowServiceInstance.RentBurrow(burrowId)
	if err != nil {
		svc.ErrorWithJson(c, 500, err)
		return
	}

	c.JSON(200, dto.SuccessResponse{Success: true, Error: nil})
}

// @Summary Show Burrow Status
// @Accept  json
// @Produce json
// @Param id path int64 true "ID"
// @Success 200 {object} dto.IndexResponse{results=[]model.Burrow}
// @Router /burrows [get]
func (svc *HttpService) BurrowStatus(c *gin.Context) {

	burrows, err := svc.BurrowServiceInstance.BurrowStatus()
	if err != nil {
		svc.ErrorWithJson(c, 500, err)
		return
	}

	resp := &dto.IndexResponse{Results: burrows, Count: len(burrows)}
	c.JSON(200, resp)
}
