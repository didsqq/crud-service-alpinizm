package grpc

import (
	"context"

	"github.com/didsqq/crud-service-alpinizm/internal/domain"
	crudv1 "github.com/didsqq/protos/gen/go/crud"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Climbs interface {
	GetAll() ([]domain.Climb, error)
	GetById(climbID int64) (domain.Climb, error)
}

type serverAPI struct {
	crudv1.UnimplementedCrudServer
	climb Climbs
}

func Register(gRPC *grpc.Server, climb Climbs) {
	crudv1.RegisterCrudServer(gRPC, &serverAPI{climb: climb})
}

func (s *serverAPI) Climbs(ctx context.Context, req *crudv1.Empty) (*crudv1.ClimbsResponse, error) {
	climbs, err := s.climb.GetAll()
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	var climbsResponse []*crudv1.ClimbResponse
	for _, climb := range climbs {
		climbsResponse = append(climbsResponse, &crudv1.ClimbResponse{
			IdMountainClimbs: climb.ID,
			IdGroups:         climb.ID_group,
			IdMountain:       climb.ID_mountain,
			IdCategory:       climb.ID_category,
			StartDate:        timestamppb.New(climb.Start_date),
			EndDate:          timestamppb.New(climb.End_date),
			Total:            climb.Total,
		})
	}

	return &crudv1.ClimbsResponse{
		Climbs: climbsResponse,
	}, nil
}

func (s *serverAPI) Climb(ctx context.Context, req *crudv1.ClimbRequest) (*crudv1.ClimbResponse, error) {
	return &crudv1.ClimbResponse{}, nil
}
