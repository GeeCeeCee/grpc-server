package persons

import (
	"log"

	"github.com/GeeCeeCee/grpc-proto/person"
	"golang.org/x/net/context" // indirect
)

type Server struct {
	person.UnimplementedSayHelloServiceServer
}

func (s *Server) SayHello(ctx context.Context, in *person.Message) (*person.Message, error) {
	
	log.Printf("Receive message body from client: %s", in.Body)
	return &person.Message{Body: "Hello From the Server!"}, nil
}
