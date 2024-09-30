package thread

import (
	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}
