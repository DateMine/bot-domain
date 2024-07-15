package repository

import (
	"context"
	"github.com/DateMine/bot-domain/pkg/models/db/offer"
	"github.com/DateMine/bot-domain/pkg/models/db/parser"
	"time"
)

type OfferRepository interface {
	Add(offer *offer.Offer, ctx context.Context) error
	AddRange(offers []offer.Offer, ctx context.Context) error
	GetAll(parserId int, days int, ctx context.Context) (*offer.Offer, error)
}

type ParserRepository interface {
	GetById(id int, ctx context.Context) (*parser.Parser, error)
	GetParsers(ctx context.Context) ([]*parser.Parser, error)
	GetParsersStartByTime(ctx context.Context, date time.Time) ([]*parser.Parser, error)
	GetParsersReportByTime(ctx context.Context, date time.Time) ([]*parser.Parser, error)
}
