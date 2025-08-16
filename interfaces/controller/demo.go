package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	adto "youras/application/dto"
	"youras/application/service"
	"youras/interfaces/dto"
)

type DemoController struct {
	svc *service.DemoService
}

func (d *DemoController) Query(c *gin.Context) {
	id := c.Query("id")
	intId, _ := strconv.Atoi(id)
	res, err := d.svc.Get(intId)
	if err != nil {
		log.Printf("Get error: %v", err)
	}
	c.JSON(200, res)
}

func NewDemoController(demoSvc *service.DemoService) *DemoController {
	return &DemoController{svc: demoSvc}
}

func (d *DemoController) Update(c *gin.Context) {
	var req dto.UpdateDemoReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := d.svc.UpdateName(adto.UpdateDemoCommand{Id: req.Id, Name: req.Name})
	if err != nil {
		c.JSON(200, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{})
}
