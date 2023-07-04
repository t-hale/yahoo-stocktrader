package stocktrader

import "errors"

var (
	InvalidYahooFinanceResponseLengthError        error = errors.New("invalid yahoo finance query result array length")
	InvalidYahooFinanceResponseNotEnoughDataError error = errors.New("invalid yahoo finance query result not enough data returned")

	YahooDomain = "finance.yahoo.com"
)
