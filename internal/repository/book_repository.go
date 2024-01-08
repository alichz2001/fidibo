package repository

import (
	"context"
	"encoding/json"
	"github.com/alichz2001/fidibo/internal/config"
	"github.com/alichz2001/fidibo/internal/model"
	"time"
)

type BookRepository struct {
	repo           *AppRepo
	searchCacheTTL time.Duration
}

type BookRepositoryInterface interface {
	FetchFromCache(ctx context.Context, key string) (books []*model.Book, err error)
	SearchByName(ctx context.Context, str string) (books []*model.Book, err error)
	PutToCache(ctx context.Context, key string, books []*model.Book) error
	BulkInsert(ctx context.Context, books []model.Book) error
}

func NewBookRepository(appCfg *config.AppConfig, repo *AppRepo) *BookRepository {
	return &BookRepository{repo: repo, searchCacheTTL: appCfg.CacheTTL}
}

func (b *BookRepository) FetchFromCache(ctx context.Context, key string) (books []*model.Book, err error) {
	resp, err := b.repo.cache.Get(ctx, key).Result()
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(resp), &books)
	return
}

func (b *BookRepository) PutToCache(ctx context.Context, key string, books []*model.Book) (err error) {
	v, _ := json.Marshal(books)
	_, err = b.repo.cache.Set(ctx, key, v, b.searchCacheTTL).Result()
	return
}

func (b *BookRepository) SearchByName(ctx context.Context, str string) (books []*model.Book, err error) {
	b.repo.db.WithContext(ctx).Model(&model.Book{}).Where("title LIKE ?", "%"+str+"%").Find(&books)
	return
}

func (b *BookRepository) BulkInsert(ctx context.Context, books []model.Book) error {
	result := b.repo.db.WithContext(ctx).Create(books)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
