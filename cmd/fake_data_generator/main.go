package main

import (
	"context"
	"flag"
	appconfig "github.com/alichz2001/fidibo/internal/config"
	"github.com/alichz2001/fidibo/internal/model"
	"github.com/alichz2001/fidibo/internal/repository"
	"github.com/google/uuid"
	"log"
	"log/slog"
	"math/rand"
	"os"
)

var (
	count         uint64
	perChunkCount uint64
)

func init() {
	flag.Uint64Var(&count, "count", 1000, "count of fake books")
	flag.Uint64Var(&perChunkCount, "per_chunk_count", 10, "count of every chunk fake books")
}

func main() {
	flag.Parse()

	_ = os.Setenv("PORT", "1234")

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appCfg := appconfig.LoadConfig(appconfig.LoadConfigOptions{PanicByErr: false, RunMysql: true, RunRedis: false})
	appCfg.Logger = logger

	repo := repository.InitRepo(appCfg)

	bookRepo := repository.NewBookRepository(appCfg, repo)

	var tmpBooks []model.Book

	log.Printf("start generating fake data")
	for i := uint64(0); i < count; {
		tmpBooks = make([]model.Book, 0)

		for j := uint64(0); j < perChunkCount; j++ {
			tmpStr := uuid.New().String()
			tmpBook := model.Book{
				Title: tmpStr + "-title",
				Cover: tmpStr + "-cover",
				Type:  uint8(rand.Intn(3) + 1),
			}
			tmpBooks = append(tmpBooks, tmpBook)
		}

		err := bookRepo.BulkInsert(context.Background(), tmpBooks)
		if err != nil {
			panic("error inserting")
		}
		i += perChunkCount
	}

	log.Printf("%d fake book inserted to db", count)
	os.Exit(0)

}
