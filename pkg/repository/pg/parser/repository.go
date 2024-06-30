package parser

import (
	"context"
	"github.com/DateMine/bot-domain/pkg/models/db/parser"
	"github.com/DateMine/bot-domain/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
)

type repo struct {
	pg *pgxpool.Pool
}

func NewRepository(pg *pgxpool.Pool) repository.ParserRepository {
	return &repo{
		pg: pg,
	}
}

func (r *repo) GetById(id int, ctx context.Context) (*parser.Parser, error) {
	var parserModel parser.Parser
	err := r.pg.QueryRow(ctx, "SELECT id, name, description, created_at, start_parsing, start_report FROM parses WHERE id = $1", id).
		Scan(&parserModel.Id, &parserModel.Name, &parserModel.Description, &parserModel.CreatedAt, &parserModel.StartParsing, &parserModel.StartReport)
	if err != nil {
		return nil, err
	}
	return &parserModel, nil
}

func (r *repo) GetCompanies(ctx context.Context) ([]*parser.Parser, error) {
	rows, err := r.pg.Query(ctx, "SELECT id, name, description, created_at, start_parsing, start_report FROM parses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parsers []*parser.Parser
	for rows.Next() {
		var parserModel parser.Parser
		err := rows.Scan(&parserModel.Id, &parserModel.Name, &parserModel.Description, &parserModel.CreatedAt, &parserModel.StartParsing, &parserModel.StartReport)
		if err != nil {
			return nil, err
		}
		parsers = append(parsers, &parserModel)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return parsers, nil

}
