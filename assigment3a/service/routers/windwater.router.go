package routers

import (
	"assigment3a/connection"
	handler "assigment3a/service/module/windwater/handler"
	repo "assigment3a/service/module/windwater/repository"
	"assigment3a/service/module/windwater/usecase"

	"github.com/gin-gonic/gin"
)

type windWaterRoutes struct {
	Handler handler.WindWaterHandler
	Router  *gin.RouterGroup
}

func InitWindWaterRoute(
	router *gin.RouterGroup, gormDB *connection.GormDB,
) *windWaterRoutes {
	handler := handler.InitWindWaterHandler(
		usecase.InitWindWaterUsecase(
			repo.InitWindWaterRepository(
				gormDB,
			),
			gormDB,
		),
	)
	return &windWaterRoutes{
		Handler: *handler,
		Router:  router,
	}
}

func (r *windWaterRoutes) Routes() {
	router := r.Router.Group("/windwater")
	{
		router.POST("/update", r.Handler.Update)
	}
}
