package design

import (
	"context"
	"log"

	"connectrpc.com/connect"

	designv1 "github.com/yofriadi/backend-vanilla/generated/merchandise/design/v1"
)

func (s *DesignServer) GetAll(
	ctx context.Context,
	req *connect.Request[designv1.GetAllRequest],
) (*connect.Response[designv1.GetAllResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&designv1.GetAllResponse{
		Designs: []*designv1.Design{
			{
				Id:              "1",
				Item:            "Item 1",
				Color:           "Color 1",
				ProductCategory: "Product Category 1",
				Metal:           "Metal 1",
				MetalRate:       1.0,
				ImageUrl:        "Image Url 1",
				DesignerId:      "Designer Id 1",
				ParentId:        "Parent Id 1",
				Stones: []*designv1.Stone{
					{
						Parcel: "Parcel 1",
						Pieces: 1,
						Carat:  1.0,
					},
					{
						Parcel: "Parcel 2",
						Pieces: 2,
						Carat:  2.0,
					},
				},
			},
		},
	})
	res.Header().Set("Merchandise-Design-Version", "v1")
	return res, nil
}
