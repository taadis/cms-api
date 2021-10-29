package handler

import (
	"context"
	"testing"

	postspb "github.com/taadis/cms-api/proto/posts"
)

// 测试保存新文章
func TestPostsHandler_Save(t *testing.T) {
	req := &postspb.SaveRequest{}
	req.Title = "测试文章标题"
	req.Content = "测试文章内容"
	req.Author = "测试文章作者"
	resp := &postspb.SaveResponse{}
	ctx := context.Background()
	postsHandler := NewPostsHandler()
	err := postsHandler.Save(ctx, req, resp)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("save got postsId %d", resp.PostsId)
}
