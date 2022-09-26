package order

type deleteOrderResponse struct {
	Message string
}

func toSuccessDelete() deleteOrderResponse {
	return deleteOrderResponse{Message: "Deleted Order"}
}
