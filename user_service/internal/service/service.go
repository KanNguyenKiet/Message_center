package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"message-server/user_service/api"
	"message-server/user_service/config"
	"message-server/user_service/internal/db"
	"net"
	"net/http"
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

func cors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if allowedOrigin(r.Header.Get("Origin")) {
			w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, ResponseType")
		}
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
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

	gwMux := runtime.NewServeMux()
	err = api.RegisterUserServiceHandler(context.Background(), gwMux, conn)
	gwServer := &http.Server{
		Addr: cfg.Host + ":" + cfg.HttpPort,
		// Enable CORS
		Handler: cors(gwMux),
	}
	log.Println("Serving http on port: ", cfg.HttpPort)
	if err = gwServer.ListenAndServe(); err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}
