package interfaces

// TODO: 
// type Request struct {

// }

// type Response struct {

// }

// type Middleware interface {
// 	Next(request, response)func(request *Request, response *Response)
// }

// type HTTPMiddleware interface {
// 	func Execute()
// }

// var middlewares map[string]HTTPMiddleware


// // Add command to commands list
// func AddHTTPMiddleware(name string, middleware HTTPMiddleware) {
// 	if nil == middlewares {
// 		middlewares = make(map[string]Middleware)
// 	}
// 	middlewares[name] = middleware
// }

// func ExecuteMiddlewares(aMiddleware []string) {
// 	for _, name := range middlewares {
// 		middlewares[name].Execute()
// 	}
// }