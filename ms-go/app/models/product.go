package models

import (
	"math"
	"time"
	"encoding/json"
	"ms-go/app/services/kafka"
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Product struct {
	ID          int       `json:"id" bson:"id,omitempty"`
	Name        string    `json:"name" bson:"name,omitempty"`
	Amount      int    `json:"amount"  bson:"amount,omitempty"`
	Brand       string    `json:"brand"  bson:"brand,omitempty"`
	Price       float64   `json:"price"  bson:"price,omitempty"`
	Description string    `json:"description"  bson:"description,omitempty"`
	CreatedAt   time.Time `json:"created_at"  bson:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at"  bson:"updated_at,omitempty"`
}

func (p *Product) Validate() error {
	if p.Price > 0 {
		p.Price = math.Round(p.Price*100) / 100
	}

	validate := *p
	return validation.ValidateStruct(&validate,
		validation.Field(&validate.ID, validation.Required),
		validation.Field(&validate.Name, validation.Required, validation.Length(4, 0)),
		validation.Field(&validate.Amount, validation.Required),
		validation.Field(&validate.Price, validation.Required, validation.Min(0.01), validation.Max(1000000.00)),
		validation.Field(&validate.Brand, validation.Required),
		validation.Field(&validate.Description, validation.Required),
	)
}

func (p *Product) ProduceProductEvent(ctx context.Context, eventType string, updatedProduct *Product) error {
	event := map[string]interface{}{
        "event_type": eventType,
        "product":    updatedProduct,
    }

    message, err := json.Marshal(event)
    if err != nil {
        return err
    }

    producer, err := kafka.New().Setup("go-to-rails").Write(message)

	if err != nil {
        return err
    }
    producer.Close()

    return nil
}