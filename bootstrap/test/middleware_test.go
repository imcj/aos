package bootstrap

type RequestIDMiddleware struct {
	context *gin.Context
}

func (m RequestIDMiddleware)HTTPMiddlewareExecute() {
	m.Header.Set("Request-ID")
}

func TestMiddleware() {
	var engine := gin.New()
	bootstram.AddMiddleware("RequestIDMiddleware", RequestIDMiddleware{engine})
}