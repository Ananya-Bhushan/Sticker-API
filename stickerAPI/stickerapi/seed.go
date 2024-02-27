package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Trendingsticker struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Priority  int       `json:"priority"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

var sampleData= [] struct {
	TrendingstickerID       int
    TrendingstickerName     string
	TrendingstickerPriority int
	TrendingstickerStartTime time.Time 
	TrendingstickerEndTime   time.Time 
	
}
{
	{	
		TrendingstickerID:   1,
		TrendingstickerName:     "Animated Sticker",
		TrendingstickerPriority: 57,
		TrendingstickerStartTime: "9:00:00",
		TrendingstickerEndTime: "17:00:00",
    },
    {
		TrendingstickerID:   2,
		TrendingstickerName:     "Friends Sticker",
		TrendingstickerPriority: 58,
		TrendingstickerStartTime: "8:30:00",
		TrendingstickerEndTime: "18:00:00",
    }

}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3308)/stickers?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Seed the database with sample data
	for _, data := range sampleData {
		sticker := Trendingsticker{
			Name:     data.TrendingstickerName,
			Priority: data.TrendingstickerPriority,
			ID:	      data.TrendingstickerID,
			StartTime:data.TrendingstickerStartTime,
			EndTime:  data.TrendingstickerEndTime
		}

		

		if err := db.Create(&sticker).Error; err != nil {
			panic("failed to seed database")
		}
	}

	fmt.Println("Database seeding completed successfully!")
}
