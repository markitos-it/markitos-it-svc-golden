package grpc

import (
	"context"
	"log"

	"markitos-it-svc-acmes/internal/application/services"
	"markitos-it-svc-acmes/internal/domain/acmes"
	pb "markitos-it-svc-acmes/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AcmeServer struct {
	pb.UnimplementedAcmeServiceServer
	service *services.AcmeService
}

func NewAcmeServer(service *services.AcmeService) *AcmeServer {
	return &AcmeServer{
		service: service,
	}
}

func (s *AcmeServer) GetAllAcmes(ctx context.Context, req *pb.GetAllAcmesRequest) (*pb.GetAllAcmesResponse, error) {
	log.Println("GetAllAcmes called")

	docs, err := s.service.GetAllAcmes(ctx)
	if err != nil {
		log.Printf("Error getting all acmes: %v", err)
		return nil, status.Errorf(codes.Internal, "failed to get acmes: %v", err)
	}

	pbDocs := make([]*pb.Acme, 0, len(docs))
	for _, doc := range docs {
		pbDocs = append(pbDocs, acmeToProto(&doc))
	}

	return &pb.GetAllAcmesResponse{
		Acmes: pbDocs,
	}, nil
}

func (s *AcmeServer) GetAcmeById(ctx context.Context, req *pb.GetAcmeByIdRequest) (*pb.GetAcmeByIdResponse, error) {
	log.Printf("GetAcmeById called with id: %s", req.Id)

	doc, err := s.service.GetAcmeByID(ctx, req.Id)
	if err != nil {
		log.Printf("Error getting acme by id %s: %v", req.Id, err)
		return nil, status.Errorf(codes.NotFound, "acme not found: %v", err)
	}

	return &pb.GetAcmeByIdResponse{
		Acme: acmeToProto(doc),
	}, nil
}

func acmeToProto(doc *acmes.Acme) *pb.Acme {
	return &pb.Acme{
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
