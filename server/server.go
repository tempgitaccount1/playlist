package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/tempgitaccount1/playlist/proto"

	"google.golang.org/grpc"
)

type Playlist struct {
}

func New() *Playlist {
	return &Playlist{}
}

func (p *Playlist) Start() error {
	lis, err := net.Listen("tcp", "localhost:8080") // this should come from a database
	if err != nil {
		return fmt.Errorf("failed to start playlist service:\n%w", err)
	}

	server := grpc.NewServer()
	pb.RegisterPlaylistManagerService(server, p)
	server.Serve(lis)
	return nil
}

func (p *Playlist) NewPlaylist(ctx context.Context, r *pb.NewPlaylistRequest) (*pb.NewPlaylistReply, error) {
	return &pb.NewPlaylistReply{
		Message: fmt.Sprintf("make playlist %v", r.Name),
	}, nil
}
