package main

import (
	"context"
)

//our twirp file has the haberdasher interface whoch has the
//MakeHat fuc=nction. Any type that implements the MakeHat
//function implements the haberdasher twirp interface
type Server struct{}

func (s *Server) MakeHat(ctx context.Context, size *pb.Size) (hat *pb.Hat, err error) {

}
