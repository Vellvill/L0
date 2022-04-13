package model

import (
	"encoding/json"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"time"
)

type Model struct {
	OrderUID string
	Json     struct {
		OrderUID    string `json:"order_uid" validate:"required"`
		TrackNumber string `json:"track_number" validate:"required"`
		Entry       string `json:"entry" validate:"required"`
		Delivery    struct {
			Name    string `json:"name" validate:"required"`
			Phone   string `json:"phone" validate:"required"`
			Zip     string `json:"zip" validate:"required"`
			City    string `json:"city" validate:"required"`
			Address string `json:"address" validate:"required"`
			Region  string `json:"region" validate:"required"`
			Email   string `json:"email" validate:"required"`
		} `json:"delivery" validate:"required"`
		Payment struct {
			Transaction  string `json:"transaction" validate:"required"`
			RequestID    string `json:"request_id" validate:"required"`
			Currency     string `json:"currency" validate:"required"`
			Provider     string `json:"provider" validate:"required"`
			Amount       int    `json:"amount" validate:"required,numeric"`
			PaymentDt    int    `json:"payment_dt" validate:"required"`
			Bank         string `json:"bank" validate:"required"`
			DeliveryCost int    `json:"delivery_cost" validate:"required"`
			GoodsTotal   int    `json:"goods_total" validate:"required,numeric"`
			CustomFee    int    `json:"custom_fee" validate:"required,numeric"`
		} `json:"payment" validate:"required"`
		Items []struct {
			ChrtID      int    `json:"chrt_id" validate:"required"`
			TrackNumber string `json:"track_number" validate:"required"`
			Price       int    `json:"price" validate:"required"`
			Rid         string `json:"rid" validate:"required"`
			Name        string `json:"name" validate:"required"`
			Sale        int    `json:"sale" validate:"required"`
			Size        string `json:"size" validate:"required"`
			TotalPrice  int    `json:"total_price" validate:"required"`
			NmID        int    `json:"nm_id" validate:"required"`
			Brand       string `json:"brand" validate:"required"`
			Status      int    `json:"status" validate:"required"`
		} `json:"items" validate:"required"`
		Locale            string    `json:"locale" validate:"required"`
		InternalSignature string    `json:"internal_signature" validate:"required"`
		CustomerID        string    `json:"customer_id" validate:"required"`
		DeliveryService   string    `json:"delivery_service" validate:"required"`
		Shardkey          string    `json:"shardkey" validate:"required"`
		SmID              int       `json:"sm_id" validate:"required,numeric"`
		DateCreated       time.Time `json:"date_created" validate:"required"`
		OofShard          string    `json:"oof_shard" validate:"required"`
	}
}

func (m *Model) Validate() error {

	err := validation.Validate(m.Json, is.JSON)
	if err != nil {
		return err
	}

	return validation.ValidateStruct(m.Json,
		validation.Field(&m.Json.OrderUID, validation.Required),
		validation.Field(&m.Json.CustomerID, validation.Required),
		validation.Field(&m.Json.DateCreated, validation.Required),
	)
}

func NewModel(js []byte) (*Model, error) {
	var model Model
	err := model.Validate()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(js, &model.Json)
	if err != nil {
		return nil, err
	}
	return &Model{model.Json.OrderUID, model.Json}, nil
}
