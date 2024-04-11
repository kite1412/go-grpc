package handler

import (
	"fmt"
	"net/http"

	"my.go/microservicesgrpc/service/order/repository"
)

func HandlerInit() {
	http.Handle("/get", getOrders())
}

func getOrders() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		repo := repository.MockRepositoryInstance()
		fmt.Fprint(w, repo.GetAll())
	})
}