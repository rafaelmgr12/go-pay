package web

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rafaelmgr12/go-pay/internal/infra/repository"
	"github.com/rafaelmgr12/go-pay/internal/usecase/transactions"
)

func HandleRequests(port string, conn *sql.DB) {

	var standardAdress string = "/api/v1"
	repo := repository.NewChatRepositoryMySQL(conn)
	usecase := transactions.NewTransactionsUseCase(repo)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the CRUD API, go to /docs for more info"))
	})

	r.Get(standardAdress+"/balance/{_id}", func(w http.ResponseWriter, r *http.Request) {
		user_id := chi.URLParam(r, "_id")
		result, err := usecase.GetAmountById(user_id, r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error getting balance"))
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"balance": %f}`, result)))
	})

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
