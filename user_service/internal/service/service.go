package service

import (
	"context"
	"google.golang.org/grpc/metadata"
	"log"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
	"net"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Service Struct of server
type Service struct {
	cfg *config.ServerConfig
	db  db.StoreQuerier
	// embedded unimplemented service
	api.UnimplementedUserServiceServer
}

// 2 below functions were made via this post https://fale.io/blog/2021/07/28/cors-headers-with-grpc-gateway
func allowedOrigin(origin string) bool {
	// allow all origin
	// TODO: add regex to filter origin
	return true
}

func cors(handler http.Handler, writer http.ResponseWriter, request *http.Request) {
	if allowedOrigin(request.Header.Get("Origin")) {
		writer.Header().Set("Access-Control-Allow-Origin", request.Header.Get("Origin"))
		writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType, _session")
	}
	if request.Method == "OPTIONS" {
		return
	}
	handler.ServeHTTP(writer, request)
}

func withLogger(request *http.Request) {
	log.Printf("[%s] -- %s", request.Method, request.RequestURI)
}

// This function help to apply multiple handler to final one
func httpHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		// Enable traffic log
		withLogger(request)
		// Enable cors
		cors(handler, writer, request)
	})
}

var allowedHeaders = map[string]struct{}{
	"x-request-id": {},
}

func isHeaderAllowed(s string) (string, bool) {
	// check if allowedHeaders contain the header
	if _, isAllowed := allowedHeaders[s]; isAllowed {
		// send uppercase header
		return strings.ToUpper(s), true
	}
	// if not in the allowed header, don't send the header
	return s, false
}

func CreateServer(cfg *config.ServerConfig, store db.StoreQuerier) error {
	// Create GRPC listener
	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalln("Create GPRC listener fail", err)
		return err
	}

	s := grpc.NewServer()
	api.RegisterUserServiceServer(s, &Service{
		cfg: cfg,
		db:  store,
	})
	go func() {
		log.Println("Serving GRPC on port: ", cfg.GRPCPort)
		log.Fatalln(s.Serve(lis))
	}()

	// Create GRPC Gateway for http
	conn, err := grpc.DialContext(
		context.Background(),
		cfg.Host+":"+cfg.GRPCPort,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Fail to dial grpc server", err)
		return err
	}

	gwMux := runtime.NewServeMux(
		runtime.WithOutgoingHeaderMatcher(isHeaderAllowed),
		// Get your custom headers of request and convert to grpc metadata
		runtime.WithMetadata(func(ctx context.Context, r *http.Request) metadata.MD {
			sessionKey := r.Header.Get("_session")
			md := metadata.Pairs("_session", sessionKey)
			return md
		}),
	)
	_ = api.RegisterUserServiceHandler(context.Background(), gwMux, conn)
	gwServer := &http.Server{
		Addr: cfg.Host + ":" + cfg.HttpPort,
		// Enable CORS
		Handler: httpHandler(gwMux),
	}
	log.Println("Serving http on port: ", cfg.HttpPort)
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
