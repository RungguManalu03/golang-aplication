package main

import (
	"fmt"
	"goaplication/auth"
	"goaplication/campaign"
	"goaplication/handler"
	"goaplication/helper"
	"goaplication/user"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=123 dbname=goaplication port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	} 

	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
// 2a6db32a-9294-474d-9f41-1e907e69e985
// 7b5f26c8-3743-400d-b7fd-d8e6573843cf

	campaigns, err := campaignRepository.FindAll()
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println(len(campaigns))
	for _, campaign := range campaigns {
		fmt.Println(campaign.Name)
	}

	uuidStr := "2a6db32a-9294-474d-9f41-1e907e69e985"
	uuidValue, err := uuid.Parse(uuidStr)
	if err != nil {
		// Handle error
	}

	campaignsById, err := campaignRepository.FindByUserID(uuidValue)
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println("=========")
	fmt.Println(len(campaignsById))
	for _, campaignId := range campaignsById {
		fmt.Println(campaignId.Description)
		if len(campaignId.CampaignImages) > 0 {
			fmt.Println(campaignId.CampaignImages[0].FileName)
		}
	}

	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService) ,userHandler.UploadAvatar)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
		return func (c *gin.Context) {
			authHeader := c.GetHeader("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			tokenString := ""
			arrayToken := strings.Split(authHeader, " ")
			if len(arrayToken) == 2 {
				tokenString = arrayToken[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			claim, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			userValue := claim["user_id"]
			userIDStr, ok := userValue.(string)
			if !ok {
				// Penanganan kesalahan jika asersi tipe gagal
			}

			userID, err := uuid.Parse(userIDStr)
			if err != nil {
				// Penanganan kesalahan jika parsing UUID gagal
			}

			user, err := userService.GetUserByID(userID)
			if err != nil {
				response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
				c.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			}

			c.Set("currentUser", user)
		}
}