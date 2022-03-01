package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"githib.com/codergab/go-grpc-livescore/api/livescore"
	"google.golang.org/grpc"
)

var matches livescore.ListMatchesResponse

type liveScoreServer struct {
	livescore.UnimplementedScoreServiceServer
}

func (lss *liveScoreServer) ListMatches(ctx context.Context, req *livescore.ListMatchesRequest) (*livescore.ListMatchesResponse, error) {
	match := &livescore.MatchScoreResponse{
		Score: "4:1",
		Live:  true,
	}

	matches.Scores = append(matches.Scores, match)
	return &matches, nil
}

const addr = "localhost:50004"

func main() {
	conn, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("tcp connection error: ", err.Error())
	}

	// Create GRPC Server
	grpcServer := grpc.NewServer()
	server := liveScoreServer{}

	// Register livescore server
	livescore.RegisterScoreServiceServer(grpcServer, &server)
	fmt.Println("Starting GRPC server at :", addr)

	if err := grpcServer.Serve(conn); err != nil {
		log.Fatal(err)
	}
}
