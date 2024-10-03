package board

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Repository struct {
	DB *pgx.Conn
}

func (repo *Repository) saveBoard(b Board) (*Board, error) {
	var board Board
	err := repo.DB.QueryRow(context.Background(), "INSERT INTO board "+
		"(code, name, id_category) "+
		"VALUES ($1, $2, $3)",
		b.Code, b.Name, b.Id_category).
		Scan(&board.Id_board, &board.Code, &board.Name, &board.Id_category)

	if err != nil {
		err = fmt.Errorf("failed query, could not save board: - %w", err)
	}
	return &board, nil
}

func (repo *Repository) getBoardById(id uint64) (*Board, error) {
	var board Board
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM board "+
		"WHERE id = $1", id).Scan(
		&board.Id_board, &board.Code, &board.Name, &board.Id_category)

	if err != nil {
		err = fmt.Errorf("failed query, could not get board with ID: - %w", err)
		return nil, err
	}

	return &board, nil
}

func (repo *Repository) getAllBoards() ([]Board, error) {
	rows, err := repo.DB.Query(context.Background(), "SELECT * FROM board")

	if err != nil {
		err = fmt.Errorf("failed query, couldn't get all boards: - %w", err)
		return nil, err
	}

	boards, err := pgx.CollectRows(rows, pgx.RowToStructByName[Board])
	if err != nil {
		err = fmt.Errorf("failed Row Collection, couldn't get rows or parse them: - %w", err)
		return nil, err
	}

	return boards, nil
}

func (repo *Repository) updateBoard(b Board) error {
	_, err := repo.DB.Exec(context.Background(), "UPDATE board "+
		"SET code = $2, name = $3, nsfw = $4 WHERE id = $1", b.Id_board, b.Code, b.Name, b.Id_category)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't update board: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) deleteBoardById(id uint64) error {
	_, err := repo.DB.Exec(context.Background(), "DELETE FROM board WHERE id = $1", id)

	if err != nil {
		err = fmt.Errorf("failed query, couldn't delete board: - %w", err)
		return err
	}

	return nil
}

func (repo *Repository) getBoardByCode(code string) (*Board, error) {
	var board Board
	err := repo.DB.QueryRow(context.Background(), "SELECT * FROM board WHERE code = $1",
		code).Scan(&board.Id_board, &board.Code, &board.Name, &board.Id_category)

	if err != nil {
		err = fmt.Errorf("failed to get board by code: - %w", err)
		return nil, err
	}

	return &board, nil
}

// TODO: Hacer una query que verifique la existencia del código,
