package repository

import (
	"context"
	"github.com/DateMine/bot-domain/pkg/models/offer"
)

type OfferRepository interface {
	Add(offer *offer.Offer, ctx context.Context) error
	AddRange(offers []offer.Offer, ctx context.Context) error
	GetAll(parserId int, days int, ctx context.Context) (*offer.Offer, error)
}
