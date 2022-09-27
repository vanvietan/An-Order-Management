package order

import (
	"errors"
	"gorm.io/gorm"
	"math"
	"order-mg/internal/model"
	"time"
)

// CreateOrderInput input from clients
type CreateOrderInput struct {
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	TotalPrice    int32          `json:"total_price"`
	Quantity      int            `json:"quantity"`
	Discount      int8           `json:"discount"`
	Shipping      string         `json:"shipping"`
	Status        model.Status   `json:"status"` // Enums Status
	UserId        int64          `json:"userID"`
	DatePurchased time.Time      `json:"date_purchased"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

// AnOrderResponse output from client
type AnOrderResponse struct {
	Id            int64           `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	TotalPrice    int32           `json:"total_price"`
	Quantity      int             `json:"quantity"`
	Discount      int8            `json:"discount"`
	Shipping      string          `json:"shipping"`
	Status        model.Status    `json:"status"` // Enums Status
	UserId        int64           `json:"userID"`
	DatePurchased time.Time       `json:"date_purchased"`
	CreatedAt     time.Time       `json:"created_at" `
	UpdatedAt     time.Time       `json:"updated_at"`
	Histories     []model.History `json:"histories,omitempty"`
}

func toGetAnOrderResponse(order model.Order) AnOrderResponse {
	return AnOrderResponse{
		Id:            order.Id,
		Name:          order.Name,
		Description:   order.Description,
		TotalPrice:    order.TotalPrice,
		Quantity:      order.Quantity,
		Discount:      order.Discount,
		Shipping:      order.Shipping,
		Status:        order.Status,
		UserId:        order.UserId,
		DatePurchased: order.DatePurchased,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		Histories:     order.Histories,
	}
}

func (c CreateOrderInput) validateAndMap() (model.Order, error) {
	if c.Name == "" {
		return model.Order{}, errors.New("invalid name")
	}
	if c.Description == "" {
		return model.Order{}, errors.New("invalid description")
	}
	if c.TotalPrice <= 0 || c.TotalPrice > math.MaxInt32 {
		return model.Order{}, errors.New("invalid price")
	}
	if c.Quantity <= 0 || c.Quantity > math.MaxInt32 {
		return model.Order{}, errors.New("invalid quantity")
	}
	if c.Discount <= 0 || c.Discount > 120 {
		return model.Order{}, errors.New("invalid discount")
	}
	if c.Shipping == "" {
		return model.Order{}, errors.New("invalid shipping method")
	}
	if c.Status == "" {
		return model.Order{}, errors.New("invalid status")
	}
	err := validStatus(c.Status)
	if err != nil {
		return model.Order{}, err
	}

	if c.UserId <= 0 || c.UserId > math.MaxInt64 {
		return model.Order{}, errors.New("invalid userID")
	}

	return model.Order{
		Name:          c.Name,
		Description:   c.Description,
		TotalPrice:    c.TotalPrice,
		Quantity:      c.Quantity,
		Discount:      c.Discount,
		Shipping:      c.Shipping,
		Status:        c.Status,
		UserId:        c.UserId,
		DatePurchased: c.DatePurchased,
	}, nil
}
func validStatus(s model.Status) error {
	if s == model.StatusApproved {
		return nil
	}
	if s == model.StatusApprovalPending {
		return nil
	}
	if s == model.StatusShipping {
		return nil
	}
	if s == model.StatusShipped {
		return nil
	}

	return errors.New("invalid status")
}
