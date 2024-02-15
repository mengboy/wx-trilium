package wxmsg

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/mengboy/wx-trilium/conf"
	"log"
	"sort"
	"strings"
)

// CheckSignature 微信接入校验
func CheckSignature(c *gin.Context) {
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echostr := c.Query("echostr")

	ok := checkSignature(signature, timestamp, nonce, conf.GetWXToken())
	if !ok {
		log.Println("微信公众号签名校验失败!")
		return
	}
	_, _ = c.Writer.WriteString(echostr)
}

// checkSignature 微信公众号签名校验
func checkSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// Sha1 Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
