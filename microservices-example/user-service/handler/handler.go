package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	order "my.go/microservicesgrpc/service/order/generated"
)

func Init(orderClient order.OrderServiceClient) {
	http.Handle("/makeorder", makeOrder(orderClient))
}

// create a static order
func makeOrder(orderClient order.OrderServiceClient) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fail := r.URL.Query().Get("fail")
		username := "Guts"
		quantity := 2
		if fail == "yes" {
			quantity = 0
		}
		itemName := "mech keyboard"
		additions := []string{
			"color:red",
			"layout:tkl",
		}
		ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
		defer cancel()
		res, err := orderClient.CreateOrder(ctx, &order.OrderRequest{
			Username: username,
			Quantiity: int64(quantity),
			Item: &order.Item{
				Name: itemName,
				Additions: additions,
			},
		})
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err.Error())
			return
		}		
		fmt.Fprint(w, res)
	})	
}