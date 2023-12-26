package itemservice

import (
	"context"

	"github.com/DanilaNik/IU5_RIP2023/internal/ds"
	"github.com/DanilaNik/IU5_RIP2023/internal/httpmodels"
	"github.com/DanilaNik/IU5_RIP2023/internal/repository"
	"github.com/pkg/errors"
)

type ItemService struct {
	Repository *repository.Repository
}

func NewItemService(repo *repository.Repository) *ItemService {
	return &ItemService{
		Repository: repo,
	}
}

func (i *ItemService) GetItems(ctx context.Context, req *httpmodels.TestingGetItemsRequest) (*httpmodels.TestingGetItemsResponse, error) {
	items, err := i.Repository.GetItems(req.SearchText)
	if err != nil {
		return nil, errors.Wrap(err, "get online items")
	}
	res := convertToResponse(items)
	resp := &httpmodels.TestingGetItemsResponse{
		Items: res,
	}

	return resp, nil
}

func (i *ItemService) GetItemByID(ctx context.Context, req *httpmodels.TestingGetItemByIDRequest) (*httpmodels.TestingGetItemByIDResponse, error) {
	item, err := i.Repository.GetItemByID(int(req.ID))
	if err != nil {
		return nil, errors.Wrap(err, "get item by id")
	}

	resp := &httpmodels.TestingGetItemByIDResponse{
		Item: httpmodels.Item{
			ID:       item.ID,
			Name:     item.Name,
			ImageURL: item.ImageURL,
			Status:   item.Status,
			Quantity: item.Quantity,
			Height:   item.Height,
			Width:    item.Width,
			Depth:    item.Depth,
			Barcode:  item.Barcode,
		},
	}

	return resp, nil
}

// func (i *ItemService) DeleteItem(ctx context.Context, req *httpmodels.TestingDeleteItemRequest) error {
// 	err := i.Repository.DeleteItem(int(req.ID))
// 	if err != nil {
// 		errors.Wrap(err, "delete item")
// 	}

// 	return nil
// }

func convertToResponse(items []*ds.Item) []*httpmodels.Item {
	result := make([]*httpmodels.Item, 0)
	for _, item := range items {
		result = append(result, &httpmodels.Item{
			ID:       item.ID,
			Name:     item.Name,
			ImageURL: item.ImageURL,
			Status:   item.Status,
			Quantity: item.Quantity,
			Height:   item.Height,
			Width:    item.Width,
			Depth:    item.Depth,
			Barcode:  item.Barcode,
		})
	}
	return result
}
