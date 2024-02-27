package usecase

import (
	"stickerapi/domain"
	"time"
	"context"
	"github.com/go-redis/redis/v8"
	"encoding/json"
)

type NewTrendingStickerUsecase interface {
	GetTrendingStickers(currentTime time.Time) ([]domain.Trendingsticker, error)

}

type newTrendingStickerUsecase struct {
	StickerRepository domain.StickerRepository
	redisClient       *redis.Client
	
}

func TrendingStickerUsecase(stickerRepository domain.StickerRepository) domain.StickerUsecase {
	return &newTrendingStickerUsecase{
		stickerRepository:stickerRepository,
		redisClient:      redisClient
		
	}
}

func (tu *newTrendingStickerUsecase) GetTrendingStickers(ctx context.Context) ([]domain.StickerResponse, error) {
	cacheKey := "trendingStickers"
	val, err := tu.redisClient.Get(ctx, cacheKey).Result()
	if err == redis.Nil {
		// The stickers are not in Redis, get them from the DB
		currentTime := time.Now().UTC()
		stickers, err := tu.StickerRepository.GetTrendingStickers(ctx, currentTime)
		if err != nil {
			return nil, err
		}
		stickersJson, _ := json.Marshal(stickers)
		err = tu.redisClient.Set(ctx, cacheKey, stickersJson, time.Hour).Err()
		if err != nil {
			
			return nil, err
		}	
		return stickers, nil
	} else if err != nil {
		
		return nil, err
	}
	var stickers []domain.StickerResponse
	err = json.Unmarshal([]byte(val), &stickers)
	if err != nil {
		
		return nil, err
	}
	return stickers, nil

 }
