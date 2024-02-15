package conf

import (
	"strings"

	"github.com/spf13/viper"
)

// GetTriliumHost ...
func GetTriliumHost() string {
	return viper.GetString("trilium.host")
}

// GetTriliumEapiToken ...
func GetTriliumEapiToken() string {
	return viper.GetString("trilium.eapi_token")
}

// GetTriliumParentNoteID ...
func GetTriliumParentNoteID() string {
	return viper.GetString("trilium.parent_note_id")
}

// GetWXToken 获取wx token
func GetWXToken() string {
	return viper.GetString("wx.wx_token")
}

// IsSelfMsg 是否个人消息
func IsSelfMsg(fromOpenID string) bool {
	ids := strings.Split(viper.GetString("wx.self_open_ids"), ",")
	for _, id := range ids {
		if id == fromOpenID {
			return true
		}
	}
	return false
}

// GetNotePrefix 获取笔记消息前缀
func GetNotePrefix() string {
	return viper.GetString("wx.note_prefix")
}
