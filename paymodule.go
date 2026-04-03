// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type Payment struct {
// 	Description string `json:"description"`
// 	USD         int    `json:"usd"`
// 	FullName    string `json:"fullname"`
// 	Address     string `json:"address"`
// 	Time        time.Time
// }

// type HttpResponse struct {
// 	Money          int
// 	PaymentHistory []Payment
// }

// func (p Payment) Println() {
// 	fmt.Println("Description:", p.Description)
// 	fmt.Println("USD:", p.USD)
// 	fmt.Println("FullName:", p.FullName)
// 	fmt.Println("Address:", p.Address)
// }

// var money = 1000
// var paymentsHistory = make([]Payment, 0)
// var mtx = sync.Mutex{}

// func payHandler(w http.ResponseWriter, r *http.Request) {
// 	var payment Payment
// 	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
// 		fmt.Println("err :", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	payment.Time = time.Now()

// 	payment.Println()

// 	mtx.Lock()
// 	if money-payment.USD >= 0 {
// 		money -= payment.USD
// 	}

// 	paymentsHistory = append(paymentsHistory, payment)

// 	HttpResponse := HttpResponse{
// 		Money:          money,
// 		PaymentHistory: paymentsHistory,
// 	}

// 	b, err := json.Marshal(HttpResponse)
// 	if err != nil {
// 		fmt.Println("err:", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}

// 	if _, err := w.Write(b); err != nil {
// 		fmt.Println("err:", err)
// 		return
// 	}

// 	mtx.Unlock()
// }

// func main() {
// 	http.HandleFunc("/pay", payHandler)

// 	if err := http.ListenAndServe(":9091", nil); err != nil {
// 		fmt.Println("Ошибка во время запуска HTTP сервера", err)
// 	}
// }
