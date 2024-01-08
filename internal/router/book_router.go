package router

import (
	"context"
	"encoding/json"
	"github.com/alichz2001/fidibo/internal/server"
	"github.com/alichz2001/fidibo/internal/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"time"
)

type BookRouter struct {
	bookService *service.BookService
}

func AddBookRoutes(app *server.WebApp, bookService *service.BookService) {
	br := &BookRouter{bookService: bookService}
	bookServiceGroup := app.Group("/v1/books")
	bookServiceGroup.Get("/search", br.Search)
}

func (r *BookRouter) Search(ctx *fiber.Ctx) error {
	q := ctx.Query("q")
	if q == "" {
		return ctx.SendStatus(http.StatusUnprocessableEntity)
	}

	ctx2, cancelFunc := context.WithTimeout(ctx.Context(), time.Second*10)
	defer cancelFunc()

	books, err := r.bookService.SearchByName(ctx2, q)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	resp, err := json.Marshal(books)
	if err != nil {
		return ctx.SendStatus(http.StatusInternalServerError)
	}

	return ctx.Send(resp)
}
