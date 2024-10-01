package thread

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) saveThread(t Thread) (*Thread, error) {
	var thread Thread
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO thread "+
		"(filepath, subject, username, timestampp, comment, reply_count, image_count, "+
		"is_archived, is_pinned, code) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		t.Filepath, t.Subject, t.Username, t.Timestampp, t.Commenta, t.Reply_count, t.Image_count, t.Is_archived, t.Is_pinned, t.Code).
		Scan(&thread.Id_thread, &thread.Filepath, &thread.Subject, &thread.Username, &thread.Timestampp, &thread.Commenta, &thread.Reply_count, &thread.Image_count, &thread.Is_archived, &thread.Is_pinned, &thread.Code)

	if err != nil {
		err = fmt.Errorf("failed query, could not save board: - %w", err)
	}
	return &thread, nil
}

func (repo *Repository) getThreadById(id uint64) (*Thread, error) {
	var thread Thread
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM thread "+
		"WHERE id = $1", id).Scan(
		&thread.Id_thread, &thread.Filepath, &thread.Subject, &thread.Username, &thread.Timestampp, &thread.Commenta, &thread.Reply_count, &thread.Image_count, &thread.Is_archived, &thread.Is_pinned, &thread.Code)

	if err != nil {
		err = fmt.Errorf("failed query, could not get thread with ID: - %w", err)
		return nil, err
	}

	return &thread, nil
}

func (repo *Repository) getAllThreads() ([]Thread, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT * FROM thread")

	if err != nil {
		err = fmt.Errorf("failed query, couldn't get all threads: - %w", err)
		return nil, err
	}

	threads, err := pgx.CollectRows(rows, pgx.RowToStructByName[Thread])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, couldn't get rows or parse them: - %w", err)
		return nil, err
	}

	return threads, nil
}

func (repo *Repository) updateThread(t Thread) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE thread "+
		"SET filepath = $2, subject = $3, username = $4, timestampp = $5, comment = $6, reply_count = $7, image_count = $8, is_archived = $9, is_pinned = $10, boardcode = $11 WHERE id = $1",
		t.Id_thread, t.Filepath, t.Subject, t.Username, t.Timestampp, t.Commenta, t.Reply_count, t.Image_count, t.Is_archived, t.Is_pinned, t.Code)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't update thread: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) deleteThreadById(id uint64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM thread WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't delete thread: - %w", err)
		return err
	}

	return nil
}
