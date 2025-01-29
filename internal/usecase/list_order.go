package usecase

import (
	"github.com/samuelralmeida/pge-clean-architecture/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]ListOrderOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}

	dto := make([]ListOrderOutputDTO, len(orders))

	for i, order := range orders {
		dto[i] = ListOrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
	}

	return dto, nil
}
