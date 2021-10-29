package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/micro/go-micro/errors"
)

var errJsonUnmarshal = errors.InternalServerError("", "序列化失败,请重试")
var errInvalidPostsId = func(postsId int64) error {
	return errors.New("", fmt.Sprintf("invalid posts id %d", postsId), 20001)
}

type PostsServicer interface {
	NewPostsId(ctx context.Context) (int64, error)
	Save(ctx context.Context, params *SaveParams) (int64, error)
	Get(ctx context.Context, postsId int64) (*PostsData, error)
	IncrView(ctx context.Context, postsId int64) (int64, error)
}

type PostsService struct {
	rdb *redis.Client
}

func NewPostsService() PostsServicer {
	s := new(PostsService)
	s.rdb = getRdb()
	return s
}

// IncrView 获取并递增文章的访问数量
func (s *PostsService) IncrView(ctx context.Context, postsId int64) (int64, error) {
	if postsId <= 0 {
		return 0, errInvalidPostsId(postsId)
	}
	key := fmt.Sprintf("posts:%d:view", postsId)
	view, err := s.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return view, nil
}

func (s *PostsService) Get(ctx context.Context, postsId int64) (*PostsData, error) {
	if postsId <= 0 {
		return nil, errInvalidPostsId(postsId)
	}

	var postData PostsData
	key := fmt.Sprintf("posts:%d:data", postsId)
	err := s.rdb.HGetAll(ctx, key).Scan(&postData)
	//result, err := s.rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("get posts key not found:%s", key)
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	//postData := new(PostsData)
	//err = json.Unmarshal([]byte(result), &postData)
	//if err != nil {
	//	return nil, err
	//}

	return &postData, nil
}

// Save 保存
func (s *PostsService) Save(ctx context.Context, params *SaveParams) (int64, error) {
	postsData := new(PostsData)
	postsData.Id = params.Id
	postsData.Title = params.Title
	postsData.Content = params.Content
	postsData.Author = params.Author
	postsData.Time = time.Now().Unix()

	if params.Id <= 0 {
		postsId, err := s.NewPostsId(ctx)
		if err != nil {
			return 0, err
		}
		postsData.Id = postsId
	}

	//postBytes, err := json.Marshal(postsData)
	//if err != nil {
	//	return 0, err
	//}

	key := fmt.Sprintf("posts:%d:data", postsData.Id)
	values := postsData.ToMap()
	err := s.rdb.HMSet(ctx, key, values).Err()
	//err = s.rdb.Set(ctx, key, string(postBytes), 0).Err()
	if err != nil {
		return 0, err
	}

	return postsData.Id, nil
}

// NewPostsId 从redis获取自增数字作为id,类似mysql中的自增
func (s *PostsService) NewPostsId(ctx context.Context) (int64, error) {
	key := fmt.Sprintf("posts:id")
	result, err := s.rdb.Incr(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	return result, nil
}
