package postgres

import (
	"context"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/rs/zerolog/log"
	"gitlab.golang-school.ru/potok-1/mbelogortsev/my-app/internal/domain"
)

func (postgres *Postgres) CreateProfile(ctx context.Context, p domain.Profile) (err error) {
	fmt.Println("Postgres - create profile")

	// Execute the query using the connection pool.
	sqlStr := "SELECT table_name FROM information_schema.tables"

	dialect := goqu.Dialect("postgres")



	record := goqu.Record{"name": "Аркадий", "age": 20}

	sql, _, _ := dialect.Insert("profile").Rows(record).ToSQL()

	fmt.Println(sql)
	_, err = postgres.pool.Exec(ctx, sql)
	if err != nil {
		log.Error().Err(err).Msg("Error")
	}

	// Execute the query.
	rows, err := postgres.pool.Query(ctx, sqlStr)
	
	if err != nil {
		log.Error().Err(err).Msg("Error")
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		if err := rows.Scan(&tableName); err != nil {
			log.Error().Err(err).Msg("Error2")
		}
		//fmt.Printf("Table name: %s\n", tableName)
	}

	return nil
}

func (postgres *Postgres) GetProfile(p domain.Profile) (err error) {
	fmt.Println("Postgres - get profile")
	return nil
}
