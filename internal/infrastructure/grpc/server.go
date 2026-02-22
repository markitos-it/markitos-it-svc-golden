package grpc

import (
	"context"
	"log"

	"markitos-it-svc-goldens/internal/application/services"
	"markitos-it-svc-goldens/internal/domain"
	pb "markitos-it-svc-goldens/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GoldenServer struct {
	pb.UnimplementedGoldenServiceServer
	service *services.GoldenService
}

func NewGoldenServer(service *services.GoldenService) *GoldenServer {
	return &GoldenServer{
		service: service,
	}
}

func (s *GoldenServer) GetAllGoldens(ctx context.Context, req *pb.GetAllGoldensRequest) (*pb.GetAllGoldensResponse, error) {
	log.Println("GetAllGoldens called")

	docs, err := s.service.GetAllGoldens(ctx)
	if err != nil {
		log.Printf("Error getting all goldens: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get goldens: %v", err)
	}

	pbDocs := make([]*pb.Golden, 0, len(docs))
	for _, doc := range docs {
		pbDocs = append(pbDocs, goldenToProto(&doc))
	}

	return &pb.GetAllGoldensResponse{
		Goldens: pbDocs,
	}, nil
}

func (s *GoldenServer) GetGoldenById(ctx context.Context, req *pb.GetGoldenByIdRequest) (*pb.GetGoldenByIdResponse, error) {
	log.Printf("GetGoldenById called with id: %s", req.Id)

	doc, err := s.service.GetGoldenByID(ctx, req.Id)
	if err != nil {
		log.Printf("Error getting golden by id %s: %v", req.Id, err)
		return nil, status.Errorf(codes.NotFound, "golden not found: %v", err)
	}

	return &pb.GetGoldenByIdResponse{
		Golden: goldenToProto(doc),
	}, nil
}

func goldenToProto(doc *domain.Golden) *pb.Golden {
	return &pb.Golden{
		Id:          doc.ID,
		Title:       doc.Title,
		Description: doc.Description,
		Category:    doc.Category,
		Tags:        doc.Tags,
		UpdatedAt:   timestamppb.New(doc.UpdatedAt),
		ContentB64:  doc.ContentB64,
		CoverImage:  doc.CoverImage,
	}
}
