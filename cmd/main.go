package main

import (
	"flag"
	"fmt"
	pb "github.com/gdegiorgio/grpc-bookshelf-server/internal/proto/book"
	"github.com/gdegiorgio/grpc-bookshelf-server/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

var (
	host   = flag.String("host", "localhost", "Set the GRPC Server host")
	port   = flag.String("port", "4500", "Set the GRPC Server port")
	pretty = flag.Bool("pretty", false, "Prettify logs")
)

func main() {
	flag.Parse()
	if *pretty {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", *host, *port))
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error listening on  %s:%s : %+v", *host, *port, err))
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterBookServer(grpcServer, server.NewBookServer())
	log.Info().Msg(fmt.Sprintf("Starting Book GRPC Server on  %s:%s", *host, *port))
	err = grpcServer.Serve(ln)
	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Error starting grpc server : %+v", err))
	}
}
