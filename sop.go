package sopapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	sop "stocktrader/gen/sop"
)

// sop service example implementation.
// The example methods log the requests and return zero values.
type sopsrvc struct {
	logger *log.Logger
}

// NewSop returns the sop service implementation.
func NewSop(logger *log.Logger) sop.Service {
	return &sopsrvc{logger}
}

// Plan implements plan.
func (s *sopsrvc) Plan(ctx context.Context, p *sop.PlanPayload) (res int, err error) {
	s.logger.Printf("sop.plan called with %+v\n", p)

	tickerData, err := GetTickerData(p.Symbol, OneMonth, OneDay, "quote", true, false)

	if err != nil {
		s.logger.Printf("Error retrieving ticker data for %q : %s", p.Symbol, err)
	}

	prettyData, err := json.MarshalIndent(tickerData, "", "    ")
	if err != nil {
		s.logger.Printf("tickerData : %+v\n", tickerData)
	} else {
		s.logger.Printf("tickerData : %s\n", prettyData)
	}
	return
}

func GetTickerData(ticker, rangeStr, intervalStr, indicators string, includeTimestamps, includePrePostTradingPeriods bool) (ChartQueryResponse, error) {
	data := YahooFinanceResponse{}

	values := url.Values{}
	values.Add("range", rangeStr)
	values.Add("interval", intervalStr)
	values.Add("indicators", indicators)
	values.Add("includeTimestamps", fmt.Sprint(includeTimestamps))
	values.Add("includePrePost", fmt.Sprint(includePrePostTradingPeriods))
	values.Add("corsDomain", YahooDomain)

	url := fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/chart/%s?%s", ticker, values.Encode())
	err := getJson(url, &data)
	if err != nil {
		return ChartQueryResponse{}, err
	}

	return data.GetFormattedOutput()
}
