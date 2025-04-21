package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
)

func (postgres *Postgres) UploadWBOrders(ctx context.Context, orders []domain.WBorder) (err error) {
	fmt.Println("Postgres - upload WB Orders")

	if len(orders) == 0 {
		fmt.Println("No orders to insert.")
		return nil
	}

	// Define PostgreSQL dialect
	dialect := goqu.Dialect("postgres")

	// Convert orders to `goqu.Record` format
	var records []goqu.Record
	for _, order := range orders {
		records = append(records, goqu.Record{
			"date":             order.Date,
			"last_change_date": order.LastChangeDate,
			"warehouse_name":   order.WarehouseName,
			"warehouse_type":   order.WarehouseType,
			"country_name":     order.CountryName,
			"oblast_okrug_name": order.OblastOkrugName,
			"region_name":      order.RegionName,
			"supplier_article": order.SupplierArticle,
			"nm_id":            order.NmID,
			"barcode":          order.Barcode,
			"category":         order.Category,
			"subject":          order.Subject,
			"brand":            order.Brand,
			"tech_size":        order.TechSize,
			"income_id":        order.IncomeID,
			"is_supply":        order.IsSupply,
			"is_realization":   order.IsRealization,
			"total_price":      order.TotalPrice,
			"discount_percent": order.DiscountPercent,
			"spp":              order.Spp,
			"finished_price":   order.FinishedPrice,
			"price_with_disc":  order.PriceWithDisc,
			"is_cancel":        order.IsCancel,
			"cancel_date":      order.CancelDate,
			"order_type":       order.OrderType,
			"sticker":          order.Sticker,
			"g_number":         order.GNumber,
			"srid":             order.Srid,
		})
	}

	// Construct batch insert SQL
	sql, args, err := dialect.Insert("wb_orders").Rows(records).ToSQL()
	if err != nil {
		log.Error().Err(err).Msg("Failed to generate SQL")
		return fmt.Errorf("failed to generate SQL: %w", err)
	}

	// Execute batch insert query
	_, err = postgres.pool.Exec(ctx, sql, args...)
	if err != nil {
		log.Error().Err(err).Msg("Error inserting orders into DB")
		return fmt.Errorf("failed to insert orders: %w", err)
	}

	log.Info().Int("inserted_rows", len(orders)).Msg("Successfully inserted WB orders")
	return nil
}
