package usecase

type (
	Input struct {
		UserID string `validate:"required"`
	}

	Output struct {
		TotalBalance float64
		UserID       string
		Balances     []Balance
	}
	Balance struct {
		ExchangeID string
		Balance    float64
	}
)
