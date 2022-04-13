package model

import (
	"encoding/json"
	"log"
	"time"
)

type Model struct {
	OrderUID string
	Json     struct {
		OrderUID    string `json:"order_uid"`
		TrackNumber string `json:"track_number"`
		Entry       string `json:"entry"`
		Delivery    struct {
			Name    string `json:"name"`
			Phone   int64  `json:"phone"`
			Zip     int    `json:"zip"`
			City    string `json:"city"`
			Address string `json:"address"`
			Region  string `json:"region"`
			Email   string `json:"email"`
		} `json:"delivery"`
		Payment struct {
			Transaction  string `json:"transaction"`
			RequestID    string `json:"request_id"`
			Currency     string `json:"currency"`
			Provider     string `json:"provider"`
			Amount       int    `json:"amount"`
			PaymentDt    int    `json:"payment_dt"`
			Bank         string `json:"bank"`
			DeliveryCost int    `json:"delivery_cost"`
			GoodsTotal   int    `json:"goods_total"`
			CustomFee    int    `json:"custom_fee"`
		} `json:"payment"`
		Items []struct {
			ChrtID      int    `json:"chrt_id"`
			TrackNumber string `json:"track_number"`
			Price       int    `json:"price"`
			Rid         string `json:"rid"`
			Name        string `json:"name"`
			Sale        int    `json:"sale"`
			Size        int    `json:"size"`
			TotalPrice  int    `json:"total_price"`
			NmID        int    `json:"nm_id"`
			Brand       string `json:"brand"`
			Status      int    `json:"status"`
		} `json:"items"`
		Locale            string    `json:"locale"`
		InternalSignature string    `json:"internal_signature"`
		CustomerID        string    `json:"customer_id"`
		DeliveryService   string    `json:"delivery_service"`
		Shardkey          int       `json:"shardkey"`
		SmID              int       `json:"sm_id"`
		DateCreated       time.Time `json:"date_created"`
		OofShard          int       `json:"oof_shard"`
	}
}

func NewModel(js []byte) *Model {
	var model Model
	err := json.Unmarshal(js, &model.Json)
	if err != nil {
		log.Println(err)
	}
	return &Model{model.Json.OrderUID, model.Json}
}
