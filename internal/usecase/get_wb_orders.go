package usecase

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/dto"
)

func (u *UseCase) GetWBOrders(ctx context.Context, input dto.GetWBOrdersInput) error {
	logger := log.With().
		Str("component", "GetWBOrders").
		Logger()

	orders, err := u.wb.GetWBOrders(ctx, input)
	if err != nil {
		logger.Error().Err(err).Msg("Request failed")
		return fmt.Errorf("request failed: %w", err)
	}
	logger.Info().Msgf("Successfully fetched orders, %v", orders)

	domainOrders := convertToDomainOrders(orders)

	err = u.postgres.UploadWBOrders(ctx, domainOrders)

	return nil
}


func convertToDomainOrders(dtoOrders []dto.GetWBordersOutput) []domain.WBorder {
	domainOrders := make([]domain.WBorder, len(dtoOrders))
	for i, order := range dtoOrders {
		domainOrders[i] = domain.WBorder{
			Date:            order.Date,
			LastChangeDate:  order.LastChangeDate,
			WarehouseName:   order.WarehouseName,
			WarehouseType:   order.WarehouseType,
			CountryName:     order.CountryName,
			OblastOkrugName: order.OblastOkrugName,
			RegionName:      order.RegionName,
			SupplierArticle: order.SupplierArticle,
			NmID:            order.NmID,
			Barcode:         order.Barcode,
			Category:        order.Category,
			Subject:         order.Subject,
			Brand:           order.Brand,
			TechSize:        order.TechSize,
			IncomeID:        order.IncomeID,
			IsSupply:        order.IsSupply,
			IsRealization:   order.IsRealization,
			TotalPrice:      order.TotalPrice,
			DiscountPercent: order.DiscountPercent,
			Spp:             order.Spp,
			FinishedPrice:   order.FinishedPrice,
			PriceWithDisc:   order.PriceWithDisc,
			IsCancel:        order.IsCancel,
			CancelDate:      order.CancelDate,
			OrderType:       order.OrderType,
			Sticker:         order.Sticker,
			GNumber:         order.GNumber,
			Srid:            order.Srid,
		}
	}
	return domainOrders
}
