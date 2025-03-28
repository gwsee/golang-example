package gra

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang/freetype"
	"golang.org/x/image/font/gofont/goregular"
)

// 存储验证码的简单内存存储
type CaptchaStore struct {
	store map[string]string // captchaID -> code
	mu    sync.RWMutex
	ttl   time.Duration
}

func NewCaptchaStore(ttl time.Duration) *CaptchaStore {
	store := &CaptchaStore{
		store: make(map[string]string),
		ttl:   ttl,
	}
	// 启动清理过期验证码的协程
	go store.cleanupExpired()
	return store
}

func (s *CaptchaStore) Set(id, code string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.store[id] = code
}

func (s *CaptchaStore) Get(id string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	code, exists := s.store[id]
	return code, exists
}

func (s *CaptchaStore) Delete(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.store, id)
}

func (s *CaptchaStore) cleanupExpired() {
	ticker := time.NewTicker(s.ttl / 2)
	defer ticker.Stop()

	for range ticker.C {
		// 简单实现，实际应用中应该记录每个验证码的创建时间
		// 这里只是定期清空所有验证码
		s.mu.Lock()
		s.store = make(map[string]string)
		s.mu.Unlock()
	}
}

// 生成随机ID
func generateID() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	return base64.URLEncoding.EncodeToString(b)
}

// 生成指定长度的数字验证码
func generateCode(length int) string {
	var sb strings.Builder
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(10))
		sb.WriteString(n.String())
	}
	return sb.String()
}

// 生成验证码图片
func generateCaptchaImage(code string, width, height int) (image.Image, error) {
	// 创建图像
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// 填充背景色
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.RGBA{240, 240, 240, 255})
		}
	}

	// 加载字体
	font, err := freetype.ParseFont(goregular.TTF)
	if err != nil {
		return nil, err
	}

	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(font)
	c.SetFontSize(30)
	c.SetClip(img.Bounds())
	c.SetDst(img)
	c.SetSrc(image.NewUniform(color.RGBA{0, 0, 0, 255}))

	// 绘制文字
	//pt := freetype.Pt(10, height/2+10)
	//for _, char := range code {
	//	// 随机调整每个字符的位置
	//	offsetY, _ := rand.Int(rand.Reader, big.NewInt(10))
	//	y := pt.Y + fixed(offsetY.Int64()-5)
	//
	//	// 绘制字符
	//	_, err = c.DrawString(string(char), freetype.Pt(pt.X, y))
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	// 移动到下一个字符位置
	//	pt.X += c.PointToFixed(30)
	//}

	// 添加干扰线
	for i := 0; i < 4; i++ {
		x1, _ := rand.Int(rand.Reader, big.NewInt(int64(width)))
		y1, _ := rand.Int(rand.Reader, big.NewInt(int64(height)))
		x2, _ := rand.Int(rand.Reader, big.NewInt(int64(width)))
		y2, _ := rand.Int(rand.Reader, big.NewInt(int64(height)))

		drawLine(img, int(x1.Int64()), int(y1.Int64()), int(x2.Int64()), int(y2.Int64()), color.RGBA{100, 100, 100, 255})
	}

	// 添加噪点
	for i := 0; i < 100; i++ {
		x, _ := rand.Int(rand.Reader, big.NewInt(int64(width)))
		y, _ := rand.Int(rand.Reader, big.NewInt(int64(height)))
		img.Set(int(x.Int64()), int(y.Int64()), color.RGBA{0, 0, 0, 255})
	}

	return img, nil
}

// 辅助函数：绘制线段
func drawLine(img *image.RGBA, x1, y1, x2, y2 int, c color.Color) {
	dx := abs(x2 - x1)
	dy := abs(y2 - y1)
	sx, sy := 1, 1
	if x1 >= x2 {
		sx = -1
	}
	if y1 >= y2 {
		sy = -1
	}
	err := dx - dy

	for {
		img.Set(x1, y1, c)
		if x1 == x2 && y1 == y2 {
			break
		}
		e2 := 2 * err
		if e2 > -dy {
			err -= dy
			x1 += sx
		}
		if e2 < dx {
			err += dx
			y1 += sy
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//
//func fixed(x int64) freetype.Fixed {
//	return freetype.Fixed(x << 8)
//}

// 全局验证码存储
var captchaStore = NewCaptchaStore(5 * time.Minute)

func main() {
	// 设置CORS头
	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	// 生成验证码图片
	http.HandleFunc("/api/captcha", func(w http.ResponseWriter, r *http.Request) {
		// 生成验证码
		captchaID := generateID()
		captchaCode := generateCode(4) // 4位数字验证码

		// 存储验证码
		captchaStore.Set(captchaID, captchaCode)

		// 生成图片
		img, err := generateCaptchaImage(captchaCode, 120, 40)
		if err != nil {
			http.Error(w, "Failed to generate captcha", http.StatusInternalServerError)
			return
		}

		// 设置响应头
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Captcha-ID", captchaID)

		// 输出图片
		png.Encode(w, img)
	})

	// 验证验证码
	http.HandleFunc("/api/verify-captcha", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 解析请求
		var req struct {
			CaptchaID   string `json:"captchaId"`
			CaptchaCode string `json:"captchaCode"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// 验证
		w.Header().Set("Content-Type", "application/json")

		storedCode, exists := captchaStore.Get(req.CaptchaID)
		if !exists {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "验证码已过期",
			})
			return
		}

		// 验证后删除验证码，防止重复使用
		captchaStore.Delete(req.CaptchaID)

		if storedCode == req.CaptchaCode {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": true,
				"message": "验证成功",
			})
		} else {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "验证码错误",
			})
		}
	})

	// 发送短信/邮件的模拟接口
	http.HandleFunc("/api/send-message", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// 解析请求
		var req struct {
			CaptchaID   string `json:"captchaId"`
			CaptchaCode string `json:"captchaCode"`
			To          string `json:"to"`
			Content     string `json:"content"`
			Type        string `json:"type"` // "sms" or "email"
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// 验证验证码
		w.Header().Set("Content-Type", "application/json")

		storedCode, exists := captchaStore.Get(req.CaptchaID)
		if !exists {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "验证码已过期",
			})
			return
		}

		// 验证后删除验证码，防止重复使用
		captchaStore.Delete(req.CaptchaID)

		if storedCode != req.CaptchaCode {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"success": false,
				"message": "验证码错误",
			})
			return
		}

		// 模拟发送短信或邮件
		// 实际应用中，这里会调用短信或邮件服务的API
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": true,
			"message": fmt.Sprintf("成功发送%s到%s", req.Type, req.To),
		})
	})

	// 启动服务器
	fmt.Println("服务器启动在 http://localhost:8080")
	http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
}
