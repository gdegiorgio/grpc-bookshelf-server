package server

import (
	"context"
	pb "github.com/gdegiorgio/grpc-bookshelf-server/internal/proto/book"
)

type GRPCBookServer = pb.BookServer

type grpcBookServer struct {
	pb.UnimplementedBookServer
}

func (s grpcBookServer) GetBooks(ctx context.Context, request *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	return &pb.GetBookResponse{
		Id:        "1",
		Author:    "J.R.R Tolkien",
		Title:     "The Lord Of The Rings - The Return of The King",
		Pagecount: 434,
	}, nil
}

func NewBookServer() GRPCBookServer {
	return &grpcBookServer{}
}
