package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/rs/cors"

	greetv1 "github.com/hokaccha/connect-example/backend/gen/greet/v1"
	"github.com/hokaccha/connect-example/backend/gen/greet/v1/greetv1connect"
)

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetv1.GreetRequest],
) (*connect.Response[greetv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())
	log.Println("req.Msg.Name: ", req.Msg.Name)
	res := connect.NewResponse(&greetv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreetServiceHandler(greeter)
	mux.Handle(path, handler)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{"POST"},
		AllowedHeaders: []string{"*"},
	})
	http.ListenAndServe(
		"localhost:8080",
		c.Handler(h2c.NewHandler(mux, &http2.Server{})),
	)
}
