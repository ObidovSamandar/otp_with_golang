package v1

import (
	"database/sql"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/minio/minio-go"

	"github.com/gin-gonic/gin"
	"github.com/obidovsamandar/go-task-auth/api/helpers"
	"github.com/obidovsamandar/go-task-auth/api/models"
	"github.com/obidovsamandar/go-task-auth/config"
	"github.com/obidovsamandar/go-task-auth/pkg/jwt"
	"github.com/obidovsamandar/go-task-auth/pkg/utils"
)

func GenerateOTPCode(c *gin.Context) {

	cfg := config.Load()
	var (
		payload models.SignUpModel
	)

	err := c.ShouldBindJSON(&payload)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to bind the JSON into struct model",
		})
		return
	}

	if len(payload.PhoneNumber) < 12 {
		c.JSON(400, gin.H{
			"error":   false,
			"message": "Phone number is incorrect!",
		})
		return
	}

	code, err := utils.GenerateRandomStringByPool(cfg.OTPDigit, cfg.OTPPool)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to generate 6 digit code with given pool",
		})
		return
	}

	fmt.Println("CODE", code)

	m := map[string]interface{}{
		"code":  code,
		"phone": payload.PhoneNumber,
	}

	passcodeToken, err := jwt.GenerateJWT(m, time.Duration(270)*time.Second, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to generate passcodetoken!",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"token": passcodeToken,
	})
}

func SignUp(c *gin.Context) {
	var (
		payload  models.SignUpModel
		userInfo models.User
	)
	cfg := config.Load()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to bind JSON",
		})
		return
	}

	row := helpers.DBClient.QueryRow("SELECT * FROM users_db WHERE phone_number=$1", payload.PhoneNumber)
	err := row.Scan(&userInfo)

	if err != sql.ErrNoRows {
		c.JSON(400, gin.H{
			"errors":  true,
			"message": "User already exists!",
		})
		return
	}

	claims, err := jwt.ExtractClaims(payload.PassCodeToken, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(400, gin.H{
			"errors":  true,
			"message": "Token is experied!",
		})
		return
	}

	user_phone := claims["phone"].(string)
	sent_code := claims["code"].(string)

	if user_phone != payload.PhoneNumber {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Phone number is incorrect!",
		})
		return
	}

	if sent_code != payload.OTPCode {
		c.JSON(400, gin.H{
			"errors":  true,
			"message": "Code is incorrect!",
		})
		return
	}
	id := uuid.New()

	_, err = helpers.DBClient.Exec("INSERT INTO users_db (id, first_name,phone_number) VALUES ($1,$2,$3);", id, payload.FirstName, payload.PhoneNumber)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to save user data!",
		})
		return
	}

	m := map[string]interface{}{
		"user_id":      id,
		"phone_number": payload.PhoneNumber,
	}

	accessToken, err := jwt.GenerateJWT(m, time.Duration(30*84600)*time.Second, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to generate access token",
		})
		return
	}

	c.JSON(200, gin.H{
		"erros": false,
		"token": accessToken,
	})
}

func SignIn(c *gin.Context) {
	var (
		payload models.SignInModel
		user    models.User
	)
	cfg := config.Load()

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to bind JSON",
		})
		return
	}

	row := helpers.DBClient.QueryRow("SELECT id, phone_number, user_img, first_name FROM users_db WHERE phone_number=$1", payload.PhoneNumber)
	err := row.Scan(&user.ID, &user.PhoneNumber, &user.UserImg, &user.FirstName)

	if err != nil {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "User not found!",
		})
		return
	}

	claims, err := jwt.ExtractClaims(payload.PassCodeToken, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(400, gin.H{
			"errors":  true,
			"message": err.Error(),
		})
		return
	}

	user_phone := claims["phone"].(string)
	sent_code := claims["code"].(string)

	if user_phone != payload.PhoneNumber {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Phone number is incorrect!",
		})
		return
	}

	if sent_code != payload.OTPCode {
		c.JSON(400, gin.H{
			"errors":  true,
			"message": "Code is incorrect!",
		})
		return
	}

	m := map[string]interface{}{
		"user_id": user.ID,
	}
	accessToken, err := jwt.GenerateJWT(m, time.Duration(30*84600)*time.Second, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to generate passcodetoken!",
		})
		return
	}

	c.JSON(200, gin.H{
		"error": false,
		"token": accessToken,
	})

}

func UpdateUser(c *gin.Context) {
	var (
		userInfo models.User
		payload  models.UpdateUser
	)

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to bind JSON to struct model",
		})
		return
	}

	cfg := config.Load()
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.JSON(401, gin.H{
			"error":   true,
			"message": "The request is not authorized!",
		})
		return
	}

	claims, err := jwt.ExtractClaims(authToken, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(401, gin.H{
			"error":   true,
			"message": "Token is expired!",
		})
		return
	}

	id := claims["user_id"].(string)

	row := helpers.DBClient.QueryRow("SELECT id, user_img, phone_number, first_name FROM users_db WHERE id=$1", id)

	err = row.Scan(&userInfo.ID, &userInfo.UserImg, &userInfo.PhoneNumber, &userInfo.FirstName)

	if err != nil {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "User not found!",
		})
		return
	}

	_, err = helpers.DBClient.Exec("UPDATE users_db SET first_name=$1, phone_number=$2, user_img=$3 WHERE id=$4", payload.FirstName, payload.PhoneNumber, payload.UserImg, id)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to update user!",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":   false,
		"message": "User crediantials updated!",
	})
}

func ImageUpload(c *gin.Context) {

	cfg := config.Load()

	minioClient, err := minio.New(cfg.MINIO_ENDPOINT, cfg.MINIO_ACCESSKEY_ID, cfg.MINIO_SECRET_KEY, cfg.MINIO_SSL_MODE)

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while connecting with minio!",
		})
		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while getting file!",
		})
		return
	}

	contentType := " image/png"

	object, err := file.Open()

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Something went wrong while getting object!",
		})
		return
	}

	defer object.Close()

	c.Request.ParseForm()

	fname := uuid.NewString()

	file.Filename = fname

	_, err = minioClient.PutObject(cfg.MINIO_BUCKET_NAME, file.Filename, object, file.Size, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Failed to upload file !",
		})
		return
	}

	c.JSON(200, gin.H{
		"error":    false,
		"filename": file.Filename,
	})
}

func GetUser(c *gin.Context) {

	var (
		userInfo models.User
	)

	type userPayload struct {
		ID          string `json:"id"`
		Username    string `json:"first_name"`
		PhoneNumber string `json:"phone_number"`
		UserImg     string `json:"user_img"`
	}

	cfg := config.Load()
	authToken := c.GetHeader("Authorization")

	if authToken == "" {
		c.JSON(401, gin.H{
			"error":   true,
			"message": "The request is not authorized!",
		})
		return
	}

	claims, err := jwt.ExtractClaims(authToken, cfg.JWT_SECRET_KEY)

	if err != nil {
		c.JSON(401, gin.H{
			"error":   true,
			"message": "Token is expired!",
		})
		return
	}

	id := claims["user_id"].(string)

	row := helpers.DBClient.QueryRow("SELECT id, user_img, phone_number, first_name FROM users_db WHERE id=$1", id)

	err = row.Scan(&userInfo.ID, &userInfo.UserImg, &userInfo.PhoneNumber, &userInfo.FirstName)

	if err != nil {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "User not found!",
		})
		return
	}

	var userImg string

	if userInfo.UserImg != nil {
		cfg := config.Load()

		minioClient, _ := minio.New(cfg.MINIO_ENDPOINT, cfg.MINIO_ACCESSKEY_ID, cfg.MINIO_SECRET_KEY, cfg.MINIO_SSL_MODE)

		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "attachment; filename=\"user.png\"")
		presignedURL, _ := minioClient.PresignedGetObject(cfg.MINIO_BUCKET_NAME, *userInfo.UserImg, time.Second*24*60*60, reqParams)

		userImg = presignedURL.String()

	}
	type UserData struct {
		PhoneNumber string
		ID          string
		FirstName   string
		UserImg     string
	}

	c.JSON(200, gin.H{
		"error": false,
		"data": UserData{
			PhoneNumber: userInfo.PhoneNumber,
			ID:          userInfo.ID,
			FirstName:   userInfo.FirstName,
			UserImg:     userImg,
		},
	})
}
