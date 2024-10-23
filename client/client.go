package main

import (
	proto "ITUServer/grpc"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Did not work")
	}

	conn, err := grpc.NewClient("localhost:5050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Not working")
	}

	client := proto.NewITUDatabaseClient(conn)

	stream, err := client.Sender(context.Background())

	stream.Send(&proto.ClientMessage{
		Message: line,
	})

	stream.CloseSend()
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		fmt.Println(res.GetMessage())
	}


	/*
		stream, err := client.ClientSend(context.Background())
		stream.Send(&proto.ClientMessage{
			Message: line,
		})
		if err != nil {
			log.Fatalf("Not working")
		}

		importStream, err := client.ServerSend(context.Background(), &proto.Empty{})

		go func() {
			for {
				value, err := importStream.Recv()
				if err == io.EOF {
					return
				}
				if err != nil {
					log.Fatalf("Not working")
				}
				fmt.Println(value.GetMessage())
			}
		}()
	*/
}
