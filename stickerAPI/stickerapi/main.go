package main

import (
	"stickerapi/delivery/http"

	"stickerapi/repository"

	"clean-arch/usecase"

	"fmt"
	"log"
	"net/url"

	"clean-arch/domain"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-redis/redis"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/India")
	fmt.Println("a")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	
	stickerRepository := repository.NewStickerRepository(db)
	stickerUsecase := usecase.TrendingStickerUsecase(stickerRepository)
	

	e := echo.New()

	port := ":9090"
	
	log.Printf("Server listening on address %s\n", port)
	e.Logger.Fatal(e.Start(port))


	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

}

func migrateDatabase(db *gorm.DB) error {
	
	db.AutoMigrate(&domain.Trendingsticker{})
	return nil
}

