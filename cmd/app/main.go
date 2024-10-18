package main

import (
	"context"
	"fmt"
	blogv1 "github.com/Jurv1/blogService/proto/gen/go/blog"
	"github.com/Jurv1/blogService/proto/gen/go/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
)

type server struct {
	blogv1.UnimplementedBlogServiceServer
}

type userServer struct {
	user.UnimplementedUserServiceServer
}

func (s *server) CreateBlog(ctx context.Context, in *blogv1.CreateBlogRequest) (*blogv1.CreateBlogResponse, error) {
	return nil, nil
}

func (s *server) GetBlog(ctx context.Context, in *blogv1.GetBlogRequest) (*blogv1.GetBlogResponse, error) {
	client := getConn()

	c := make(chan *user.MakeMutationResponse)
	defer close(c)

	go getResp(ctx, client, c)

	//response, err := client.MutateUser(ctx, &user.MakeMutationRequest{
	//	Message: "Mutate",
	//})
	//if err != nil {
	//	return nil, err
	//}
	val, ok := <-c

	if ok {
		log.Println(fmt.Sprintf("AS %b", ok))
	}

	for {

	}

	log.Println(val)

	//log.Println(fmt.Sprintf("sssss %s", response))

	return &blogv1.GetBlogResponse{
		Name: val.GetMessage(),
	}, nil
}

func getConn() user.UserServiceClient {
	conn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
	}

	return user.NewUserServiceClient(conn)
}

func getResp(ctx context.Context, client user.UserServiceClient, c chan *user.MakeMutationResponse) *user.MakeMutationResponse {
	response, err := client.MutateUser(ctx, &user.MakeMutationRequest{
		Message: "Mutate",
	})
	if err != nil {
		log.Fatal(err)
	}
	c <- response

	return response
}

func (s *server) GetBlogs(ctx context.Context, in *blogv1.GetBlogsRequest) (*blogv1.GetBlogsResponse, error) {
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
