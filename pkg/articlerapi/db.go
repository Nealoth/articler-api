package articlerapi

import (
	"context"
	"fmt"
	"github.com/Nealoth/articler-api/pkg/articlerapi/configuration"
	"github.com/jmoiron/sqlx"
)

func InitPostgresConnection(ctx context.Context, conf configuration.DatabaseConfiguration) (*sqlx.DB, error) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.DbName,
		conf.Password,
		conf.SslMode,
	)

	db, err := sqlx.ConnectContext(ctx, "postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
