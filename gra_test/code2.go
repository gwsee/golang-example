package gra

import (
	"bytes"
	"crypto/rand"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/big"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang/freetype"
)

var rdb *redis.Client

// 图形验证码配置
const (
	captchaWidth   = 150
	captchaHeight  = 50
	captchaLength  = 4
	captchaCharset = "ABCDEFGHJKLMNPQRSTUVWXYZ23456789" // 去除了易混淆字符
	captchaExpire  = 5 * time.Minute
)

// 生成图形验证码
func generateCaptcha(c *gin.Context) {
	// 生成随机验证码
	code := make([]byte, captchaLength)
	for i := range code {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(captchaCharset))))
		code[i] = captchaCharset[num.Int64()]
	}
	captchaCode := string(code)

	// 创建图片
	img := image.NewRGBA(image.Rect(0, 0, captchaWidth, captchaHeight))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{245, 245, 245, 255}}, image.Point{}, draw.Src)

	// 添加干扰元素
	addNoise(img)
	addCurve(img)

	// 添加文字
	fontSize := 24.0
	ctx := freetype.NewContext()
	ctx.SetDPI(72)
	ctx.SetFontSize(fontSize)
	ctx.SetClip(img.Bounds())
	ctx.SetDst(img)
	ctx.SetSrc(image.Black)

	// 计算文字位置
	totalWidth := float64(captchaLength) * fontSize * 0.8
	startX := (float64(captchaWidth) - totalWidth) / 2

	// 绘制每个字符
	for i, char := range captchaCode {
		//angle := math.Pi * (randFloat(-0.2, 0.2))
		x := startX + float64(i)*fontSize*0.8
		y := 30 + randFloat(-5, 5)

		// 字符扭曲
		ctx.SetFontSize(fontSize + randFloat(-2, 2))
		//ctx.Rotate(angle)
		ctx.DrawString(string(char), freetype.Pt(int(x), int(y)))
		//ctx.Rotate(-angle)
	}

	// 存储到Redis
	captchaID := generateRandomID(16)
	rdb.Set(c, "captcha:"+captchaID, captchaCode, captchaExpire)

	// 返回PNG图片
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

// 验证图形验证码
func verifyCaptcha(c *gin.Context) {
	type Request struct {
		Phone       string `form:"phone" binding:"required"`
		CaptchaID   string `form:"captcha_id" binding:"required"`
		CaptchaCode string `form:"captcha_code" binding:"required"`
	}

	var req Request
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
		return
	}

	// 验证图形验证码
	storedCode, err := rdb.Get(c, "captcha:"+req.CaptchaID).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid captcha"})
		return
	}

	if storedCode != req.CaptchaCode {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid captcha code"})
		return
	}

	// 验证成功后删除验证码
	rdb.Del(c, "captcha:"+req.CaptchaID)

	// 执行后续短信发送逻辑
	// ...（参考之前的短信发送代码）
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// 工具函数
func randFloat(min, max float64) float64 {
	r, _ := rand.Int(rand.Reader, big.NewInt(1<<52))
	return min + (max-min)*float64(r.Int64())/float64(1<<52)
}

func generateRandomID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		b[i] = charset[num.Int64()]
	}
	return string(b)
}

// 添加干扰线
func addNoise(img *image.RGBA) {
	for i := 0; i < 10; i++ {
		y := randInt(0, captchaHeight)
		img.Set(randInt(0, captchaWidth), y, color.RGBA{uint8(randInt(0, 255)), uint8(randInt(0, 255)), uint8(randInt(0, 255)), 255})
	}
}

func randInt(i int, height int) int {
	return 0
}

// 添加曲线干扰
func addCurve(img *image.RGBA) {
	// 实现曲线绘制逻辑...
}
