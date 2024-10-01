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
		"(subject, username, timestamp, comment, replycount, imagecount, "+
		"isarchived, ispinned, boardcode) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		t.Subject, t.Username, t.Timestamp, t.Comment, t.ReplyCount, t.ImageCount, t.IsArchived, t.IsPinned, t.BoardCode).
		Scan(&thread.Id, &thread.Subject, &thread.Username, &thread.Timestamp, &thread.Comment, &thread.ReplyCount, &thread.ImageCount, &thread.IsArchived, &thread.IsPinned, &thread.BoardCode)

	if err != nil {
		err = fmt.Errorf("failed query, could not save board: - %w", err)
	}
	return &thread, nil
}

func (repo *Repository) getThreadById(id uint64) (*Thread, error) {
	var thread Thread
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM thread "+
		"WHERE id = $1", id).Scan(
		&thread.Id, &thread.Subject, &thread.Username, &thread.Timestamp, &thread.Comment, &thread.ReplyCount, &thread.ImageCount, &thread.IsArchived, &thread.IsPinned, &thread.BoardCode)

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
		"SET subject = $2, username = $3, timestamp = $4, comment = $5, replycount = $6, imagecount = $7, isarchived = $8, ispinned = $9, boardcode = $10 WHERE id = $1",
		t.Id, t.Subject, t.Username, t.Timestamp, t.Comment, t.ReplyCount, t.ImageCount, t.IsArchived, t.IsPinned, t.BoardCode)

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
