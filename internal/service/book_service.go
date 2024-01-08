package service

import (
	"context"
	"fmt"
	"github.com/alichz2001/fidibo/internal/config"
	"github.com/alichz2001/fidibo/internal/model"
	"github.com/alichz2001/fidibo/internal/repository"
	"log/slog"
)

type BookService struct {
	bookRepo *repository.BookRepository
	logger   *slog.Logger
}

func InitBookService(appCfg *config.AppConfig, repo *repository.BookRepository) *BookService {
	return &BookService{bookRepo: repo, logger: appCfg.Logger}
}

func (srv *BookService) SearchByName(ctx context.Context, str string) (books []*model.Book, err error) {
	srv.logger.Info(fmt.Sprintf("try fetch books from cache, key: '%s'", str))
	books, err = srv.bookRepo.FetchFromCache(ctx, str)
	if err != nil {
		srv.logger.Info(fmt.Sprintf("try fetch books from db, key: '%s'", str))
		books, err = srv.bookRepo.SearchByName(ctx, str)
		if err != nil {
			srv.logger.Error(fmt.Sprintf("failed to fetch books form db, key: '%s'", str))
			return
		}
		srv.logger.Info(fmt.Sprintf("db hit, key: '%s'", str))

		srv.logger.Info(fmt.Sprintf("try put books to cache, key: '%s'", str))
		tmpErr := srv.bookRepo.PutToCache(ctx, str, books)
		if tmpErr != nil {
			srv.logger.Error(fmt.Sprintf("failed to put books to cache, key: '%s'", str))
		}
		return
	}
	srv.logger.Info(fmt.Sprintf("cache hit, key: '%s'", str))
	return
}
