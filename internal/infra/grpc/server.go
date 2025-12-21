package grpc

import (
	"fmt"
	"net"

	grpcpkg "github.com/joseMarciano/crypto-manager/internal/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	GRPC *grpc.Server
}

func (s *Server) Start(port string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		panic(err)
	}

	return s.GRPC.Serve(lis)
}

func New() *Server {
	s := grpc.NewServer(grpc.UnaryInterceptor(grpcpkg.Interceptor()))
	reflection.Register(s)
	return &Server{GRPC: s}
}
