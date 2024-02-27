package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrNotFound = errors.New("your requested Item is not found")
)

type Trendingsticker struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Priority  int       `json:"priority"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type StickerResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type StickerUsecase interface {
	GetTrendingStickers(ctx context.Context) ([]Trendingsticker, error)

}

type StickerRepository interface {
	GetTrendingStickers(ctx context.Context, currentTime time.Time) ([]Trendingsticker, error)
	
}
