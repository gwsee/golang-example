package gra

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"net/http"
	"sync"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 验证码会话存储
type CaptchaSession struct {
	ID        string    `json:"id"`
	Position  int       `json:"position"`
	CreatedAt time.Time `json:"createdAt"`
}

// 验证码请求响应
type CaptchaResponse struct {
	ID       string `json:"id"`
	ImageURL string `json:"imageUrl"`
	By       []byte `json:"by"`
}

// 验证请求
type VerifyRequest struct {
	ID       string `json:"id"`
	Position int    `json:"position"`
}

// 验证响应
type VerifyResponse struct {
	Success bool   `json:"success"`
	Token   string `json:"token,omitempty"`
	Message string `json:"message,omitempty"`
}

var (
	sessions     = make(map[string]*CaptchaSession)
	sessionMutex sync.Mutex
	jwtSecret    = []byte("your-secret-key") // 在生产环境中应该使用环境变量
)

// 生成随机位置
func generateRandomPosition(min, max int) (int, error) {
	if max <= min {
		return 0, fmt.Errorf("max must be greater than min")
	}

	diff := max - min
	n, err := rand.Int(rand.Reader, big.NewInt(int64(diff)))
	if err != nil {
		return 0, err
	}

	return min + int(n.Int64()), nil
}

// 生成随机ID
func generateRandomID2(length int) (string, error) {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)

	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result[i] = chars[n.Int64()]
	}

	return string(result), nil
}

// 生成JWT令牌
func generateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"verified": true,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func Main() {
	r := gin.Default()
	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 获取验证码
	r.GET("/api/captcha", func(c *gin.Context) {
		// 生成随机位置 (20-280范围内)
		position, err := generateRandomPosition(20, 280)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate position"})
			return
		}

		// 生成会话ID
		id, err := generateRandomID2(32)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session ID"})
			return
		}

		// 存储会话
		sessionMutex.Lock()
		sessions[id] = &CaptchaSession{
			ID:        id,
			Position:  position,
			CreatedAt: time.Now(),
		}
		sessionMutex.Unlock()

		// 返回验证码信息
		// 注意：在实际应用中，你应该生成真实的图片
		c.JSON(http.StatusOK, CaptchaResponse{
			ID:       id,
			ImageURL: fmt.Sprintf("/api/captcha/image/%s", id),
		})
	})
	// 获取验证码
	r.GET("/api/captcha/img", func(c *gin.Context) {
		s := GetRandStr(4)
		d := ImgText(200, 100, s)
		// 返回验证码信息
		// 注意：在实际应用中，你应该生成真实的图片
		c.JSON(http.StatusOK, CaptchaResponse{
			ID: s,
			By: d,
		})
	})
	// 验证码图片端点 (简化示例)
	r.GET("/api/captcha/image/:id", func(c *gin.Context) {
		id := c.Param("id")

		sessionMutex.Lock()
		session, exists := sessions[id]
		sessionMutex.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "Captcha not found"})
			return
		}

		// 在实际应用中，这里应该生成带有滑块的图片
		// 这里简化处理，只返回位置信息
		c.JSON(http.StatusOK, gin.H{
			"id":       session.ID,
			"position": session.Position,
		})
	})

	// 验证滑块位置
	r.POST("/api/captcha/verify", func(c *gin.Context) {
		var req VerifyRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		sessionMutex.Lock()
		session, exists := sessions[req.ID]
		if exists {
			// 验证完成后删除会话
			defer func() {
				delete(sessions, req.ID)
			}()
		}
		sessionMutex.Unlock()

		if !exists {
			c.JSON(http.StatusNotFound, VerifyResponse{
				Success: false,
				Message: "Captcha not found or expired",
			})
			return
		}

		// 检查会话是否过期 (5分钟)
		if time.Since(session.CreatedAt) > 5*time.Minute {
			c.JSON(http.StatusOK, VerifyResponse{
				Success: false,
				Message: "Captcha expired",
			})
			return
		}

		// 验证位置 (允许5像素的误差)
		diff := math.Abs(float64(session.Position - req.Position))
		if diff <= 5 {
			// 验证成功，生成令牌
			token, err := generateToken()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
				return
			}

			c.JSON(http.StatusOK, VerifyResponse{
				Success: true,
				Token:   token,
			})
		} else {
			c.JSON(http.StatusOK, VerifyResponse{
				Success: false,
				Message: "Position mismatch",
			})
		}
	})

	// 启动服务器
	log.Println("Server running on :8080")
	r.Run(":8081")
}
