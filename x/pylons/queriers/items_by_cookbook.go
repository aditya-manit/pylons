package queriers

import (
	"context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/Pylons-tech/pylons/x/pylons/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	KeyItemsByCookbook = "items_by_cookbook"
)

// ItemResp is the response for Items
type ItemResp struct {
	Items []types.Item
}

func (ir ItemResp) String() string {
	output := "ItemResp{"
	for _, it := range ir.Items {
		output += it.String()
		output += ",\n"
	}
	output += "}"
	return output
}

// ItemsByCookbook returns a cookbook based on the cookbook id
func (querier *querierServer) ItemsByCookbook(ctx context.Context, req *types.ItemsByCookbookRequest) (*types.ItemsByCookbookResponse, error) {
	if req.CookbookID == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no cookbook id is provided in path")
	}

	items, err := querier.Keeper.ItemsByCookbook(sdk.UnwrapSDKContext(ctx), req.CookbookID)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &types.ItemsByCookbookResponse{
		Items: types.ItemInputsToProto(items),
	}, nil
}
