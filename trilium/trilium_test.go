package trilium

import (
	"context"
	"testing"
)

func TestClient_CreateNote(t *testing.T) {
	type args struct {
		ctx      context.Context
		content  string
		title    string
		noteType NoteType
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "testCreateNote", args: args{
			ctx:      context.Background(),
			content:  "测试",
			title:    "测试",
			noteType: NoteTypeText,
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{}
			if err := c.CreateNote(tt.args.ctx, tt.args.content, tt.args.title, tt.args.noteType); (err != nil) != tt.wantErr {
				t.Errorf("CreateNote() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
