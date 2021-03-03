package queriers

import (
	"context"
	"github.com/Pylons-tech/pylons/x/pylons/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	KeyListCookbook = "list_cookbook"
)

// ListCookbook returns a cookbook based on the cookbook id
func (querier *querierServer) ListCookbook(ctx context.Context, req *types.ListCookbookRequest) (*types.ListCookbookResponse, error) {
	if req.Address == "" {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "no address is provided in path")
	}

	accAddr, err := sdk.AccAddressFromBech32(req.Address)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	cookbook, err := querier.Keeper.GetCookbookBySender(sdk.UnwrapSDKContext(ctx), accAddr)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, err.Error())
	}

	return &types.ListCookbookResponse{
		Cookbooks: types.CookbookListToGetCookbookResponseList(cookbook),
	}, nil
}
