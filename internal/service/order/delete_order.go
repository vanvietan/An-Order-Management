package order

import (
	"context"
	log "github.com/sirupsen/logrus"
)

// DeleteOrder delete an order
func (i impl) DeleteOrder(ctx context.Context, orderID int64) error {
	err := i.orderRepo.DeleteOrder(ctx, orderID)
	if err != nil {
		log.Printf("error when deleting an order :  %v", err)
		return err
	}
	return nil
}
