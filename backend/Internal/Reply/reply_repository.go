package reply

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) saveReply(r Reply) (*Reply, error) {
	var reply Reply
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO reply "+
		"(file, username, timestamp, comment, id_thread)"+
		"VALUES ($1, $2, $3, $4, $5)",
		r.File, r.Username, r.Timestamp, r.Comment, r.Id_thread).
		Scan(&reply.Id, &reply.Username, &reply.Timestamp, &reply.Comment, &reply.Id_thread)

	if err != nil {
		err = fmt.Errorf("failed query, could not save reply: - %w", err)
	}
	return &reply, nil
}

func (repo *Repository) getReplyById(id uint64) (*Reply, error) {
	var reply Reply
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM reply "+
		"WHERE id = $1", id).Scan(
		&reply.Id, &reply.Username, &reply.Timestamp, &reply.Comment, &reply.Id_thread)

	if err != nil {
		err = fmt.Errorf("failed query, could not get reply with ID: - %w", err)
		return nil, err
	}

	return &reply, nil
}

func (repo *Repository) getAllReplies() ([]Reply, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT * FROM reply")

	if err != nil {
		err = fmt.Errorf("failed query, couldn't get all replies: - %w", err)
		return nil, err
	}

	replies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Reply])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, couldn't get rows or parse them: - %w", err)
		return nil, err
	}

	return replies, nil
}

func (repo *Repository) updateReply(r Reply) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE reply "+
		"SET subject = $2, username = $3, timestamp = $4, comment = $5, id_thread = $6 WHERE id = $1",
		r.Id, r.File, r.Username, r.Timestamp, r.Comment, r.Id_thread)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't update reply: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) deleteReplyById(id uint64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM reply WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't delete reply: - %w", err)
		return err
	}

	return nil
}
