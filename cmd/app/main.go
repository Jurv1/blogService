package main

import (
	"context"
	blogv1 "github.com/Jurv1/blogService/proto/gen/go/blog"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	blogv1.UnimplementedBlogServiceServer
}

func (s *server) CreateBlog(ctx context.Context, in *blogv1.CreateBlogRequest) (*blogv1.CreateBlogResponse, error) {
	return nil, nil
}

func (s *server) GetBlog(ctx context.Context, in *blogv1.GetBlogRequest) (*blogv1.GetBlogResponse, error) {

	return &blogv1.GetBlogResponse{
		Name: "new Blog",
	}, nil
}

func (s *server) GetBlogs(ctx context.Context, in *blogv1.GetBlogsRequest) (*blogv1.GetBlogsResponse, error) {
	return nil, nil
}

func (s *server) MakeMutation(ctx context.Context, in *blogv1.MakeMutationRequest) (*blogv1.MakeMutationResponse, error) {
	return nil, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8070")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	blogv1.RegisterBlogServiceServer(grpcServer, &server{})

	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}
