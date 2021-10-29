package handler

import (
	"context"

	"github.com/taadis/cms-api/internal/posts/service"
	postspb "github.com/taadis/cms-api/proto/posts"
)

type PostsHandler struct {
	postsService service.PostsServicer
}

func NewPostsHandler() postspb.PostsHandler {
	h := new(PostsHandler)
	h.postsService = service.NewPostsService()
	return h
}

func (h *PostsHandler) Save(ctx context.Context, req *postspb.SaveRequest, resp *postspb.SaveResponse) error {
	params := &service.SaveParams{}
	params.Id = req.Id
	params.Title = req.Title
	params.Content = req.Content
	params.Author = req.Author
	postsId, err := h.postsService.Save(ctx, params)
	if err != nil {
		return err
	}

	resp.PostsId = postsId
	return nil
}
