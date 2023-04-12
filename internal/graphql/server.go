package graphql

import (
	"github.com/kubeshop/testkube/internal/graphql/services"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"

	executorsclientv1 "github.com/kubeshop/testkube-operator/client/executors/v1"
	"github.com/kubeshop/testkube/internal/graphql/gen"
	"github.com/kubeshop/testkube/internal/graphql/resolvers"
	"github.com/kubeshop/testkube/pkg/event/bus"
	"github.com/kubeshop/testkube/pkg/log"
)

func GetServer(eventBus bus.Bus, executorsClient *executorsclientv1.ExecutorsClient) *handler.Server {
	service := &services.Service{
		Logger: log.DefaultLogger,
		Bus:    eventBus,
	}
	resolver := &resolvers.Resolver{
		ExecutorsService: &services.ExecutorsService{
			Service: service,
			Client:  executorsClient,
		},
	}
	srv := handler.New(gen.NewExecutableSchema(gen.Config{Resolvers: resolver}))
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.Use(extension.Introspection{})

	return srv

}
