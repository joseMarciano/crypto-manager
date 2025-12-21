package usecase

type (
	Input struct {
		StartDate string `validate:"required"`
		EndDate   string `validate:"required"`
	}

	Output struct {
		Data []Exchange
	}
	Exchange struct {
		ExchangeID string
		Dates      []Date
	}

	Date struct {
		Date   string
		Totals []Total
	}

	Total struct {
		Amount float64
		Type   string
	}
)
