package http

import (
	"clean-arch/domain"
	"stickerapi/domain"
	"stickerapi/usecase"

	// "fmt"
	"net/http"
	// "strconv"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"

	"stickerapi/delivery/http/middleware"
)

type ResponseError struct {
	Message string `json:"message"`
}

type StickerHandler struct {
	stickerUsecase domain.StickerUsecase
}

func InitHTTPServer(port string, stickerUsecase domain.StickerUsecase) {
	e := echo.New()
	midd_ware := middleware.InitMiddleware()
	e.Use(midd_ware.CORS)
	stickerHandler := NewStickerHandler(stickerUsecase)
	e.GET("/v1/trendingStickers", stickerHandler.GetTrendingStickersHandler)
	e.Logger.Fatal(e.Start(port))
}


func NewStickerHandler(e *echo.Echo, us usecase.NewTrendingStickerUsecase) {

	handler := &StickerHandler{
		stickerUsecase: us,
	}
	
	e.GET("/v1/trendingStickers", handler.GetTrendingStickers)
}

	art,err := h.stickerUsecase.GetTrendingStickers(currentTime time.Time)
	
	if err != nil {
		return c.JSON(getStatusCode(err), ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusOK, art)


func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	logrus.Error(err)
	switch err {
	case domain.ErrNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

