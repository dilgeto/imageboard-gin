package reply

import {
	"context"
	"fmt"
	"github.com/jackc/pgx"
}

type Repository struct {
	DB *pgx.Conn
}