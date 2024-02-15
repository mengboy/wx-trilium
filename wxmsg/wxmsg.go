package wxmsg

import (
	"encoding/xml"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mengboy/wx-trilium/conf"
	"github.com/mengboy/wx-trilium/trilium"
)

const (
	MsgOptionGetOpenID = "openid"
)

// ReqTextMsg 微信文本消息请求
type ReqTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	MsgId        int64
}

// Receive 微信消息接收
func Receive(c *gin.Context) {
	var textMsg ReqTextMsg
	err := c.ShouldBindXML(&textMsg)
	if err != nil {
		log.Printf("XML数据包解析失败: %v\n", err)
		return
	}
	if strings.ToLower(textMsg.Content) == MsgOptionGetOpenID {
		Response(c, textMsg.ToUserName, textMsg.FromUserName, textMsg.FromUserName)
		return
	}
	log.Printf("receive msg %v \n", textMsg)
	if !conf.IsSelfMsg(textMsg.FromUserName) {
		return
	}
	if strings.HasPrefix(textMsg.Content, conf.GetNotePrefix()) {
		cli := &trilium.Client{}
		title, content := parseNoteMsg(textMsg.Content)
		err = cli.CreateNote(c, content, title, trilium.NoteTypeText)
		if err != nil {
			Response(c, textMsg.ToUserName, textMsg.FromUserName, "add note failed")
			return
		}
		Response(c, textMsg.ToUserName, textMsg.FromUserName, "add note succ")
	}

	// 对接收的消息进行被动回复
	Response(c, textMsg.ToUserName, textMsg.FromUserName, "成功接收到消息")
}

func parseNoteMsg(content string) (string, string) {
	contentLines := strings.Split(content, "\n")
	if len(contentLines) > 2 {
		return contentLines[1], strings.Join(contentLines[2:], "\n")
	}
	return time.Now().Format(time.DateTime), strings.Join(contentLines[1:], "\n")
}

// ResTextMsg 微信文本消息回复
type ResTextMsg struct {
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
	XMLName      xml.Name `xml:"xml"`
}

// Response 微信消息回复
func Response(c *gin.Context, fromUser, toUser string, content string) {
	repTextMsg := ResTextMsg{
		ToUserName:   toUser,
		FromUserName: fromUser,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      content,
	}

	msg, err := xml.Marshal(&repTextMsg)
	if err != nil {
		log.Printf("XML编码出错: %v\n", err)
		return
	}
	_, _ = c.Writer.Write(msg)
}
