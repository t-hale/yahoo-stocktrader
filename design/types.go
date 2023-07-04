package design

import (
	. "goa.design/goa/v3/dsl"
)

// Foundational types
var Time = Type("Time", String, func() {
	Format(FormatDateTime)

	Meta("type:generate:force")
})

// Complex types
var Quote = Type("Quote", func() {
	Description("A stock quote object")

	Attribute("OpensAt", Time)
	Attribute("Open", Float64)
	Attribute("High", Float64)
	Attribute("Low", Float64)
	Attribute("Close", Float64)
	Attribute("Volume", Float64)
	Attribute("ClosesAt", Time)
	Attribute("Period", Int64)

	Meta("type:generate:force")
})

var TradingPeriod = Type("TradingPeriod", func() {
	Attribute("Timezone", String)
	Attribute("Start", Time)
	Attribute("End", Time)
	Attribute("GMTOffset", Int64)

	Meta("type:generate:force")
})

var CurrentTradingPeriod = Type("CurrentTradingPeriod", func() {
	Attribute("Pre", TradingPeriod)
	Attribute("Regular", TradingPeriod)
	Attribute("Post", TradingPeriod)

	Meta("type:generate:force")
})

var YahooFinanceTradingPeriod = Type("YahooFinanceTradingPeriod", func() {
	Attribute("Timezone", String)
	Attribute("Start", Int64)
	Attribute("End", Int64)
	Attribute("Start", Int64)
	Attribute("GMTOffset", Int64)

	Meta("type:generate:force")
})

var YahooFinanceCurrentTradingPeriod = Type("YahooFinanceCurrentTradingPeriod", func() {
	Attribute("Pre", YahooFinanceTradingPeriod)
	Attribute("Regular", YahooFinanceTradingPeriod)
	Attribute("Post", YahooFinanceTradingPeriod)

	Meta("type:generate:force")
})

var YahooFinanceMeta = Type("YahooFinanceMeta", func() {
	Attribute("Currency", String)
	Attribute("Symbol", String)
	Attribute("InstrumentType", String)
	Attribute("ExchangeName", String)
	Attribute("FirstTradeDate", Int64)
	Attribute("GMTOffset", Int64)
	Attribute("Timezone", String)
	Attribute("CurrentTradingPeriod", YahooFinanceCurrentTradingPeriod)
	Attribute("TradingPeriods", ArrayOf(ArrayOf(YahooFinanceTradingPeriod)))

	Meta("type:generate:force")
})

var YahooFinanceQuote = Type("YahooFinanceQuote", func() {
	Attribute("High", ArrayOf(Float64))
	Attribute("Open", ArrayOf(Float64))
	Attribute("Low", ArrayOf(Float64))
	Attribute("Close", ArrayOf(Float64))
	Attribute("Volume", ArrayOf(Float64))

	Meta("type:generate:force")
})

var YahooFinanceIndicators = Type("YahooFinanceIndicators", func() {
	Attribute("Quote", ArrayOf(YahooFinanceQuote))

	Meta("type:generate:force")
})

var YahooFinanceResult = Type("YahooFinanceResult", func() {
	Attribute("Meta", YahooFinanceMeta)
	Attribute("Timestamp", ArrayOf(Int64))
	Attribute("Indicators", YahooFinanceIndicators)
	Meta("type:generate:force")
})

var YahooFinanceResponse = Type("YahooFinanceResponse", func() {
	Attribute("Chart", YahooFinanceChart)

	Meta("type:generate:force")
})

var YahooFinanceChart = Type("YahooFinanceChart", func() {
	Description("A really cool description")
	Attribute("Result", ArrayOf(YahooFinanceResult))
	Attribute("Error", String)

	Meta("type:generate:force")
})

var ChartQueryResponse = Type("ChartQueryResponse", func() {
	Attribute("Currency", String)
	Attribute("Symbol", String)
	Attribute("ExchangeName", String)
	Attribute("InstrumentType", String)
	Attribute("FirstTradeDate", Time)
	Attribute("GMTOffset", Int64)
	Attribute("Timezone", String)
	Attribute("CurrentTradingPeriod", CurrentTradingPeriod)
	Attribute("TradingPeriods", ArrayOf(ArrayOf(TradingPeriod)))
	Attribute("Quotes", ArrayOf(Quote))

	Meta("type:generate:force")
})
