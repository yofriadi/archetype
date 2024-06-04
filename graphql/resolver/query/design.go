package query

import (
	"context"

	"connectrpc.com/connect"

	generated "github.com/yofriadi/backend-vanilla/generated/graphql"
	designv1 "github.com/yofriadi/backend-vanilla/generated/merchandise/design/v1"
)

// Get is the resolver for the get field.
func (r *QueryResolver) Get(ctx context.Context, id string) (*generated.Design, error) {
	res, err := r.RPCClientDesign.Get(ctx, connect.NewRequest(&designv1.GetRequest{Id: id}))
	if err != nil {
		return nil, err
	}

	design := &generated.Design{
		CreatedAt:       res.Msg.Design.CreatedAt.String(),
		UpdatedAt:       res.Msg.Design.UpdatedAt.String(),
		ID:              res.Msg.Design.Id,
		Item:            res.Msg.Design.Item,
		Color:           res.Msg.Design.Color,
		ProductCategory: res.Msg.Design.ProductCategory,
		Metal:           res.Msg.Design.Metal,
		MetalRate:       res.Msg.Design.MetalRate,
		ImageURL:        res.Msg.Design.ImageUrl,
		DesignerID:      &res.Msg.Design.DesignerId,
		ParentID:        &res.Msg.Design.ParentId,
	}

	design.Stones = make([]*generated.Stone, len(res.Msg.Design.Stones))
	for i, stone := range res.Msg.Design.Stones {
		design.Stones[i] = &generated.Stone{
			Parcel: stone.Parcel,
			Pieces: int(stone.Pieces),
			Carat:  stone.Carat,
		}
	}
	design.Histories = make([]*generated.Design, len(res.Msg.Design.Histories))
	for i, history := range res.Msg.Design.Histories {
		design.Histories[i] = &generated.Design{
			CreatedAt:       history.CreatedAt.String(),
			UpdatedAt:       history.UpdatedAt.String(),
			ID:              history.Id,
			Item:            history.Item,
			Color:           history.Color,
			ProductCategory: history.ProductCategory,
			Metal:           history.Metal,
			MetalRate:       history.MetalRate,
			ImageURL:        history.ImageUrl,
			DesignerID:      &history.DesignerId,
			ParentID:        &history.ParentId,
		}
	}

	return design, nil
}

// GetAll is the resolver for the getAll field.
func (r *QueryResolver) GetAll(ctx context.Context) ([]*generated.Design, error) {
	res, err := r.RPCClientDesign.GetAll(ctx, connect.NewRequest(&designv1.GetAllRequest{}))
	if err != nil {
		return nil, err
	}

	designs := make([]*generated.Design, len(res.Msg.Designs))
	for i, design := range res.Msg.Designs {
		designs[i] = &generated.Design{
			CreatedAt:       design.CreatedAt.String(),
			UpdatedAt:       design.UpdatedAt.String(),
			ID:              design.Id,
			Item:            design.Item,
			Color:           design.Color,
			ProductCategory: design.ProductCategory,
			Metal:           design.Metal,
			MetalRate:       design.MetalRate,
			ImageURL:        design.ImageUrl,
			DesignerID:      &design.DesignerId,
			ParentID:        &design.ParentId,
		}

		stones := make([]*generated.Stone, len(design.Stones))
		for _, stone := range design.Stones {
			stones = append(stones, &generated.Stone{
				Parcel: stone.Parcel,
				Pieces: int(stone.Pieces),
				Carat:  stone.Carat,
			})
		}
		designs[i].Stones = stones

		histories := make([]*generated.Design, len(design.Histories))
		for _, history := range design.Histories {
			histories = append(histories, &generated.Design{
				CreatedAt:       history.CreatedAt.String(),
				UpdatedAt:       history.UpdatedAt.String(),
				ID:              history.Id,
				Item:            history.Item,
				Color:           history.Color,
				ProductCategory: history.ProductCategory,
				Metal:           history.Metal,
				MetalRate:       history.MetalRate,
				ImageURL:        history.ImageUrl,
				DesignerID:      &history.DesignerId,
				ParentID:        &history.ParentId,
			})
		}
		designs[i].Histories = histories
	}
	return designs, nil
}
