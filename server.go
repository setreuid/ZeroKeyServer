package main

import (
    "context"
    "log"
    "net"

    pb "ZeroKey/data"
    "google.golang.org/grpc"

    "encoding/json"
    "io/ioutil"
)

type Server struct {}


func (self *Server) Start() {
    listener, err := net.Listen("tcp", PORT)
    if (err != nil) {
        log.Fatalf("Tcp listen error: %v", err)
    }

    server := grpc.NewServer()
    pb.RegisterCommandServer(server, &Server{})
    if err := server.Serve(listener); err != nil {
        log.Fatalf("Fail to serve: %v", err)
    }
}


// Client -> Server(Raspi)
func (self *Server) Save(ctx context.Context, in *pb.KeysRequest) (*pb.Noting, error) {
    // log.Printf("RECV: %+v", in)
    Preset = *in

    file, _ := json.MarshalIndent(in, "", " ")
    if err := ioutil.WriteFile("preset.json", file, 0644); err == nil {
        _P("Save preset to preset.json")
    }

    return &pb.Noting{Status: 200}, nil
}


// Server(Raspi) -> Client
func (self *Server) Load(ctx context.Context, in *pb.Noting) (*pb.KeysRequest, error) {
    // log.Printf("RECV: %+v", in.Status)
    return &Preset, nil
}


// Client -> Server(Raspi)
func (self *Server) Run(ctx context.Context, in *pb.Noting) (*pb.Noting, error) {
    log.Printf("RECV: %+v", in.Status)
    MainKeyboard.Run(int(in.Status))

    return &pb.Noting{Status: 200}, nil
}


// Client -> Server(Raspi)
func (self *Server) Stop(ctx context.Context, in *pb.Noting) (*pb.Noting, error) {
    log.Printf("RECV: %+v", in.Status)
    MainKeyboard.Stop(int(in.Status))

    return &pb.Noting{Status: 200}, nil
}
