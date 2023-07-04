package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = API("sop", func() {
	Title("Stock Option Planner Service")
	Description("Service for planning stock options including vesting schedules")
	Server("server", func() {
		Host("localhost", func() {
			URI("http://localhost:8000")
			//URI("grpc://localhost:8080")
		})
	})
})

var _ = Service("sop", func() {
	Description("The sop service provides advisors with a comprehensive view of a particular stock schedule.")

	Method("plan", func() {
		Payload(func() {
			Field(1, "symbol", String, "stock symbol to retrieve plan for")
			Required("symbol")
		})

		Result(YahooFinanceResponse)

		HTTP(func() {
			GET("/plan/{symbol}")
		})

		//GRPC(func() {
		//})
	})

	Files("/openapi.json", "./gen/http/openapi.json")
})
