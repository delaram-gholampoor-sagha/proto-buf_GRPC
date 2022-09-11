package main


// protoc --go_out=. --go_opt=paths=source_relative     --go-grpc_out=. --go-grpc_opt=paths=source_relative     greet/greetpb/greet.proto

// https://github.com/techschool/pcbook-go/issues/3

// func main() {
// 	fmt.Println(" Hello I'm a client")

// 	// cc = client connection
// 	cc ,err :=  grpc.Dial("localhost:50051" , grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatal("could not connect %v" , err)
// 	}

// 	defer cc.Close()
//     c := greetpb.NewGreetServiceClient(cc)
// 	fmt.Println("Created CLient %f" , c)

// }