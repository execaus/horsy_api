package handler

import (
	"errors"
	"fmt"
	"horsy_api/config"
	"horsy_api/internal/service"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	saga service.SagaRunner
}

func NewHandler(saga service.SagaRunner) Handler {
	return Handler{
		saga: saga,
	}
}

func (h *Handler) GetRouter(serverConfig *config.ServerConfig) *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{serverConfig.Origin},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := router.Group("/api")

	auth := api.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	herd := api.Group("/herd", h.authMiddleware)
	{
		herd.POST("", h.createHerd)
		herd.GET("", h.getHerds)

		herdID := herd.Group("/:id")
		{
			herdID.GET("", h.getHerdByID)
			herdID.GET("/horses", h.getHerdHorses)
		}
	}

	horseGender := api.Group("/horse-gender")
	{
		horseGender.GET("", h.getHorseGenderList)
	}

	horseColor := api.Group("/horse-color")
	{
		horseColor.GET("", h.getColors)
	}

	horseBirthplace := api.Group("/horse-birthplace")
	{
		horseBirthplace.GET("", h.getBirthplaces)
	}

	horseGeneticMarker := api.Group("/horse-genetic-marker")
	{
		horseGeneticMarker.GET("", h.getGeneticMarkers)
	}

	horseBreed := api.Group("/horse-breed")
	{
		horseBreed.GET("", h.getBreeds)
	}

	horse := api.Group("/horse")
	{
		horse.GET("/:id", h.getHorse)
		horse.POST("", h.createHorse)
	}

	return router
}

func (h *Handler) getParameterUUID(c *gin.Context, key string) (uuid.UUID, error) {
	value := c.Param(key)
	if value == "" {
		return uuid.Nil, errors.New(fmt.Sprintf("invalid path parameter %s", key))
	}

	parsed, err := uuid.Parse(value)
	if err != nil {
		zap.L().Error(err.Error())
		return uuid.Nil, err
	}

	return parsed, nil
}
