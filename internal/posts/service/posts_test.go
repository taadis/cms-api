package service

import (
	"context"
	"testing"
)

// 测试获取并递增文章的访问数量
func TestPostsService_IncrView(t *testing.T) {
	postsId := int64(3)
	ctx := context.Background()
	postsService := NewPostsService()
	view, err := postsService.IncrView(ctx, postsId)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("posts incr view success: %d", view)
}

// 测试获取文章数据
func TestPostsService_Get(t *testing.T) {
	postsId := int64(3)
	ctx := context.Background()
	postsService := NewPostsService()
	postsData, err := postsService.Get(ctx, postsId)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("got postsData:%+v", postsData)
}

// 测试创建新的文章
func TestPostsService_Save(t *testing.T) {
	params := new(SaveParams)
	params.Title = "测试文章标题"
	params.Content = "测试文章内容"
	params.Author = "测试人员"
	ctx := context.Background()
	postsService := NewPostsService()
	postsId, err := postsService.Save(ctx, params)
	if err != nil {
		t.Fatal(err)
	}
	if postsId <= 0 {
		t.Fatal("fail: posts create")
	}
	t.Logf("posts create success postsId:%d", postsId)
}

// 测试获取新文章的Id
func TestPostsService_NewPostsId(t *testing.T) {
	ctx := context.Background()
	postsService := NewPostsService()
	postId, err := postsService.NewPostsId(ctx)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("postId:%d", postId)
}
