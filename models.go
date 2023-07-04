package stocktrader

import (
	"google.golang.org/protobuf/proto"
	sop "stocktrader/gen/sop"
	"time"
)

func GetFormattedOutput(resp *sop.YahooFinanceResponse) (*sop.YahooFinanceResponse, error) {
	if len(resp.Chart.Result) < 1 {
		return nil, InvalidYahooFinanceResponseLengthError
	}

	out := sop.ChartQueryResponse{
		Currency:       resp.Chart.Result[0].Meta.Currency,
		Symbol:         resp.Chart.Result[0].Meta.Symbol,
		ExchangeName:   resp.Chart.Result[0].Meta.ExchangeName,
		InstrumentType: resp.Chart.Result[0].Meta.InstrumentType,
		FirstTradeDate: time.Unix(resp.Chart.Result[0].Meta.FirstTradeDate, 0),
		GMTOffset:      time.Duration(*resp.Chart.Result[0].Meta.GMTOffset) * time.Second,
		Timezone:       resp.Chart.Result[0].Meta.Timezone,
	}

	out.Quotes = make([]*sop.Quote, len(resp.Chart.Result[0].Timestamp))
	if len(resp.Chart.Result[0].Timestamp) < 2 {
		return nil, InvalidYahooFinanceResponseNotEnoughDataError
	}

	periodSeconds := sop.Duration(time.Second * time.Duration(resp.Chart.Result[0].Timestamp[1]-resp.Chart.Result[0].Timestamp[0]))
	for ind := range resp.Chart.Result[0].Timestamp {

		opensAt := sop.Time(resp.Chart.Result[0].Timestamp[ind])
		closesAt := sop.Time(resp.Chart.Result[0].Timestamp[ind] + int64(periodSeconds))

		out.Quotes[ind] = &sop.Quote{
			OpensAt:  &opensAt,
			ClosesAt: &closesAt,
			Period:   &periodSeconds,
			Open:     proto.Float64(resp.Chart.Result[0].Indicators.Quote[0].Open[ind]),
			High:     proto.Float64(resp.Chart.Result[0].Indicators.Quote[0].High[ind]),
			Low:      proto.Float64(resp.Chart.Result[0].Indicators.Quote[0].Low[ind]),
			Close:    proto.Float64(resp.Chart.Result[0].Indicators.Quote[0].Close[ind]),
			Volume:   proto.Float64(resp.Chart.Result[0].Indicators.Quote[0].Volume[ind]),
		}
	}

	out.CurrentTradingPeriod.Pre = &sop.TradingPeriod{
		Timezone:  proto.String(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.Timezone),
		GMTOffset: time.Second * resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.GMTOffset,
		Start:     time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.Start, 0),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Pre.End, 0),
	}
	out.CurrentTradingPeriod.Regular = &sop.TradingPeriod{
		Timezone:  proto.String(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.Timezone),
		GMTOffset: time.Second * resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.GMTOffset,
		Start:     proto.String(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.Start),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Regular.End, 0),
	}
	out.CurrentTradingPeriod.Post = &sop.TradingPeriod{
		Timezone:  proto.String(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.Timezone),
		GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.GMTOffset),
		Start:     time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.Start, 0),
		End:       time.Unix(resp.Chart.Result[0].Meta.CurrentTradingPeriod.Post.End, 0),
	}

	out.TradingPeriods = make([][]TradingPeriod, len(resp.Chart.Result[0].Meta.TradingPeriods))
	for dayInd := range resp.Chart.Result[0].Meta.TradingPeriods {
		out.TradingPeriods[dayInd] = make([]TradingPeriod, len(resp.Chart.Result[0].Meta.TradingPeriods[dayInd]))
		for periodInd := range resp.Chart.Result[0].Meta.TradingPeriods[dayInd] {
			out.TradingPeriods[dayInd][periodInd] = &sop.TradingPeriod{
				Timezone:  resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].Timezone,
				GMTOffset: time.Second * time.Duration(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].GMTOffset),
				Start:     time.Unix(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].Start, 0),
				End:       time.Unix(resp.Chart.Result[0].Meta.TradingPeriods[dayInd][periodInd].End, 0),
			}
		}
	}

	return &out, nil
}
