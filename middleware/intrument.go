package middleware

// import (
// 	"github.com/go-kit/kit/metrics"
// 	"product-microservice/product"
// )

// func Metrics(requestCount metrics.Counter, requestLatency metrics.Histogram) product.Usecase {
// 	return &metricsMiddleware{
// 		requestCount,
// 		requestLatency,
// 	}
// }

// type metricsMiddleware struct {
// 	requestCount   metrics.Counter
// 	requestLatency metrics.Histogram
// }

// func (mw metricsMiddleware) CreateProduct
