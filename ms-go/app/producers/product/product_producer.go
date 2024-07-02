package product

import (
    "context"
    "encoding/json"
    "ms-go/app/models"
    "ms-go/app/services/kafka"
)

func ProduceProductEvent(ctx context.Context, eventType string, data *models.Product) error {
    message, err := json.Marshal(struct {
        EventType string          `json:"event_type"`
        Product   *models.Product `json:"product"`
    }{
        EventType: eventType,
        Product:   data,
    })
    if err != nil {
        return err
    }

    return kafka.Produce(ctx, "go-to-rails", message)
}
