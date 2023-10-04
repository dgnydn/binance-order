package app

import (
	"context"
	orderv1 "github.com/dgnydn/binance-order/api/binance-order/v1"
	"github.com/dgnydn/binance-order/internal/app/order"
	"github.com/dgnydn/binance-order/internal/app/order/orderdriver"
	"github.com/dgnydn/binance-order/internal/app/order/store"
	"github.com/dgnydn/binance-order/internal/common"
	"github.com/dgnydn/binance-order/internal/platform/database"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/tracing/opencensus"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"github.com/gorilla/mux"
	appkitendpoint "github.com/sagikazarmark/appkit/endpoint"
	"github.com/sagikazarmark/kitx/correlation"
	kitxendpoint "github.com/sagikazarmark/kitx/endpoint"
	kitxtransport "github.com/sagikazarmark/kitx/transport"
	kitxgrpc "github.com/sagikazarmark/kitx/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitializeApp initializes a new HTTP and a new gRPC application.
func InitializeApp(
	grpcServer *grpc.Server,
	httpRouter *mux.Router,
	logger common.Logger,
	errorHandler common.ErrorHandler, // nolint: interfacer
) {
	endpointMiddleware := []endpoint.Middleware{
		correlation.Middleware(),
		opencensus.TraceEndpoint("", opencensus.WithSpanName(func(ctx context.Context, _ string) string {
			name, _ := kitxendpoint.OperationName(ctx)
			return name
		})),
		appkitendpoint.LoggingMiddleware(logger),
	}

	transportErrorHandler := kitxtransport.NewErrorHandler(errorHandler)

	grpcServerOptions := []kitgrpc.ServerOption{
		kitgrpc.ServerErrorHandler(transportErrorHandler),
		kitgrpc.ServerBefore(correlation.GRPCToContext()),
	}

	{
		dbConfig := database.Config{
			Host:   "localhost",
			Port:   3306,
			User:   "root",
			Pass:   "root",
			Name:   "binance-order",
			Params: nil,
		}

		gormDb, err := gorm.Open(mysql.Open(dbConfig.DSN()), &gorm.Config{})
		if err != nil {
			return
		}

		gormStore := store.New(store.WithGorm(gormDb))

		service := order.NewService(gormStore, logger)
		//service = orderdriver.ValidatorMiddleware(commonValidator)(service)
		//service = orderdriver.LoggingMiddleware(logger)(service)
		//service = orderdriver.InstrumentationMiddleware()(service)

		endpoints := orderdriver.MakeEndpoints(
			service,
			kitxendpoint.Combine(endpointMiddleware...),
		)

		orderv1.RegisterOrderServiceServer(
			grpcServer,
			orderdriver.MakeGRPCServer(endpoints, kitxgrpc.ServerOptions(grpcServerOptions)),
		)

		subRouter := httpRouter.PathPrefix("/v1/convert").Subrouter()
		orderdriver.RegisterHTTPHandler(endpoints, subRouter)

		reflection.Register(grpcServer)
	}
}
