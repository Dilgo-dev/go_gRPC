package grpc

import (
	"context"
	"go_server/internal/proto"
	"go_server/internal/utils"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type LogServer struct {
	proto.UnimplementedLogServiceServer
}

func (s *LogServer) CreateLog(ctx context.Context, req *proto.LogRequest) (*proto.LogResponse, error) {
	db := utils.Database()
	
	log := utils.Log{
		Message: req.Message,
	}
	
	result := db.Create(&log)
	if result.Error != nil {
		return &proto.LogResponse{
			Success: false,
			Message: "Failed to create log: " + result.Error.Error(),
		}, nil
	}
	
	return &proto.LogResponse{
		Success: true,
		Message: "Log created successfully",
	}, nil
}

func (s *LogServer) GetLogs(ctx context.Context, req *proto.EmptyRequest) (*proto.LogsResponse, error) {
	db := utils.Database()
	var logs []utils.Log
	
	db.Find(&logs)
	
	var protoLogs []*proto.LogEntry
	for _, log := range logs {
		protoLogs = append(protoLogs, &proto.LogEntry{
			Id:        uint64(log.ID),
			Message:   log.Message,
			CreatedAt: log.CreatedAt.Format(time.RFC3339),
		})
	}
	
	return &proto.LogsResponse{
		Logs: protoLogs,
	}, nil
}

func RegisterGrpcServer(s *grpc.Server) {
	proto.RegisterLogServiceServer(s, &LogServer{})
	reflection.Register(s)
}