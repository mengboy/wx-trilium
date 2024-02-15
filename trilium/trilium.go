package trilium

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mengboy/wx-trilium/conf"
)

const (
	// CreateNoteAPI 创建笔记api
	CreateNoteAPI = "/etapi/create-note"
)

type NoteType string

const (
	NoteTypeText        NoteType = "text"
	NoteTypeCode        NoteType = "code"
	NoteTypeFile        NoteType = "file"
	NoteTypeImage       NoteType = "image"
	NoteTypeSearch      NoteType = "search"
	NoteTypeBook        NoteType = "book"
	NoteTypeRelationMap NoteType = "relationMap"
	NoteTypeRender      NoteType = "render"
)

type CreateNoteReq struct {
	ParentNoteId string   `json:"parentNoteId"`
	Title        string   `json:"title"`
	Type         NoteType `json:"type"`
	Content      string   `json:"content"`
}

type Client struct {
}

func (c *Client) CreateNote(ctx *gin.Context, content string, title string, noteType NoteType) error {
	httpCli := http.DefaultClient
	reqParams := &CreateNoteReq{
		ParentNoteId: conf.GetTriliumParentNoteID(),
		Title:        title,
		Type:         noteType,
		Content:      content,
	}
	reqJsonByte, _ := json.Marshal(reqParams)
	reqURL := conf.GetTriliumHost() + CreateNoteAPI
	req, err := http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(reqJsonByte))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", conf.GetTriliumEapiToken())
	req.Header.Add("Content-Type", "application/json")
	rsp, err := httpCli.Do(req)
	if err != nil {
		return err
	}
	if rsp.Body != nil {
		rsp.Body.Close()
	}
	return nil
}
