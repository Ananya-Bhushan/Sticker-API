package repository

import (
	"stickerapi/domain"

	"time"

	"gorm.io/gorm"

	"context"

)

type stickerRepository struct {
	db *gorm.DB
}

func NewStickerRepository(db *gorm.DB) domain.StickerRepository {
	return &stickerRepository{
	db:db,
	}
}

func (tu *stickerRepository) GetTrendingStickers(ctx context.Context,currentTime time.Time) ([]domain.StickerResponse, error) {
	
	t := make([]domain.StickerResponse, 0)
	err := tu.db.Model(&domain.Trendingsticker{}).
		Select("trendingstickers.id, trendingstickers.Name, trendingstickers.Priority").
		Where("trendingstickers.StartTime <= ? AND trendingstickers.EndTime >=?", currentTime, currentTime).
		Order("trendingstickers.Priority desc").
		Scan(&t).Error
	return t,err
	
}

