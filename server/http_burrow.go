package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hamdiBouhani/GopherNet-golang/dto"
	"github.com/hamdiBouhani/GopherNet-golang/storage/model"
)

// @Summary Show Burrow
// @Accept  json
// @Produce json
// @Param uuid path int64 true "UUID"
// @Success 200 {object} model.Burrow
// @Router /burrow/{id} [get]
func (svc *HttpService) Show(c *gin.Context) {
	// uid := c.Param("id")
	// if _, err := uuid.Parse(uid); err != nil {
	// 	svc.ErrorWithJson(c, 400, err)
	// 	return
	// }
	var resp model.Burrow

	c.JSON(200, resp)
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
