package router

import (
	"context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	bimbelRepo "UTS-LANJUTAN-ARYA/bimbel/Repository"
	bimbelUC "UTS-LANJUTAN-ARYA/bimbel/Usecase"
	bimbelHandler "UTS-LANJUTAN-ARYA/bimbel/Handler"
)

type Handlers struct {
	Ctx context.Context
	DB  *gorm.DB
	R   *gin.Engine
}

func (h *Handlers) Routes(){
	BimbelRepo := bimbelRepo.NewBimbelRepo(h.DB)
	BimbelUc := bimbelUC.NewBimbelUC(BimbelRepo)
	
	v1 := h.R.Group("api")
	bimbelHandler.BimbelRoute(BimbelUc, v1)
}