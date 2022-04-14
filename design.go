package design

import (
	//lint:ignore ST1001 Definitions are overly verbose without dot importing
	. "goa.design/goa/v3/dsl"
)

var _ = API("Test", func() {
	Title("Experiment")
})

var _ = Service("Test", func() {
	Method("CreateEndpoint", func() {
		Payload(Endpoint)
		Result(func() {
			Field(1, "result", String)
			Required("result")
		})
		HTTP(func() {
			PUT("/manage/v1/endpoint")
		})
	})

	Method("ListEndpoints", func() {
		Result(func() {
			Field(1, "result", String)
			Field(2, "endpoints", ArrayOf(Endpoint))
			Required("result", "endpoints")
		})
		HTTP(func() {
			GET("/manage/v1/endpoints")
		})
	})
})

var Endpoint = Type("application/x-endpoint", func() {
	TypeName("Endpoint")
	Attribute("config", ArrayOf(EndpointStepConfig))
	Required("config")
})

var EndpointStepConfig = Type("application/x-endpointstepconfig", func() {
	TypeName("EndpointStepConfig")
	Attribute("name")
	Required("name")
})
