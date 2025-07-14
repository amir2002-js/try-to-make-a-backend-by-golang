package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"myProject/auth"
	"myProject/models"
	"myProject/store"
	"net/http"
)

type StoreStruct struct {
	DB *pgxpool.Pool
}

func (r *StoreStruct) GetProducts(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hello World"})
}

func (r *StoreStruct) CreateProduct(c *gin.Context)  {}
func (r *StoreStruct) UpdateProduct(c *gin.Context)  {}
func (r *StoreStruct) DelProductById(c *gin.Context) {}
func (r *StoreStruct) GetProductById(c *gin.Context) {}

func (r *StoreStruct) CreateUser(c *gin.Context) {
	clientInfoUser := &struct {
		Username        string `json:"username" required:"true"`
		Password        string `json:"password" required:"true"`
		ConfirmPassword string `json:"confirm_password" required:"true"`
		Email           string `json:"email" required:"true"`
	}{}

	err := c.BindJSON(clientInfoUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "wrong info -- 35 handler.go"})
		return
	}

	if clientInfoUser.ConfirmPassword != clientInfoUser.Password {
		c.JSON(http.StatusBadRequest, gin.H{"message": "password and confirm password not match  -- 40 handler.go"})
		return
	}

	exist, err := store.CheckUserExist(clientInfoUser.Username, clientInfoUser.Email, r.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error  -- 46 handler.go", "err": err})
		return
	}

	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"message": "userName or email already exist  -- 51 handler.go"})
		return
	}

	hashedPass, err := auth.HashPassword(clientInfoUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error  -- 57 handler.go", "err": err})
		return
	}

	newUserRole := auth.Role(clientInfoUser.Username, clientInfoUser.Email)
	newUser := &models.User{
		Username: clientInfoUser.Username,
		Email:    clientInfoUser.Email,
		Password: hashedPass,
		Role:     newUserRole,
	}

	err = store.CreateUser(newUser, r.DB)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error  -- 71 handler.go", "err": err})
		return
	}

	tkn, err := auth.GenerateJWTTkn(newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error  -- 77 handler.go", "err": err})
		return
	}
	c.SetCookie("jwt-tkn", tkn, 3600*24*30, "/", "", false, true)
	c.JSON(200, gin.H{"message": "success create user"})
}
