package handlers

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/render"

	"github.com/TOsmanov/my-hw/hw15_go_sql/internal/lib/api/response"
	"github.com/TOsmanov/my-hw/hw15_go_sql/internal/storage"
)

type Response struct {
	response.Response
	Message any `json:"data,omitempty"`
}

func UsersHandler(log *slog.Logger, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.UsersHandler"

		log := log.With(
			slog.String("op", op),
		)
		var resp response.Response
		switch r.Method {
		case http.MethodGet:
			GetUsers(log, storage, w, r, &resp)
		case http.MethodPut:
			InsertUser(log, storage, w, r, &resp)
		case http.MethodPost:
			UpdateUser(log, storage, w, r, &resp)
		case http.MethodDelete:
			DeleteUser(log, storage, w, r, &resp)
		}
	}
}

func UserStatHandler(log *slog.Logger, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.UsersHandler"

		log := log.With(
			slog.String("op", op),
		)
		var resp response.Response
		switch r.Method {
		case http.MethodGet:
			GetUserStat(log, storage, w, r, &resp)
		}
	}
}

func ProductsHandler(log *slog.Logger, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.ProductsHandler"

		log := log.With(
			slog.String("op", op),
		)
		var resp response.Response
		switch r.Method {
		case http.MethodGet:
			GetProducts(log, storage, w, r, &resp)
		case http.MethodPut:
			InsertProduct(log, storage, w, r, &resp)
		case http.MethodPost:
			UpdateProduct(log, storage, w, r, &resp)
		case http.MethodDelete:
			DeleteProduct(log, storage, w, r, &resp)
		}
	}
}

func OrdersHandler(log *slog.Logger, storage *storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.users.OrdersHandler"

		log := log.With(
			slog.String("op", op),
		)
		var resp response.Response
		switch r.Method {
		case http.MethodGet:
			GetOrders(log, storage, w, r, &resp)
		case http.MethodPut:
			InsertOrder(log, storage, w, r, &resp)
		case http.MethodDelete:
			DeleteOrder(log, storage, w, r, &resp)
		}
	}
}

func responseOK(w http.ResponseWriter, r *http.Request, msg string) {
	render.JSON(w, r, Response{
		Response: response.OK(),
		Message:  msg,
	})
}
