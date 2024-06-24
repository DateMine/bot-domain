package offer_storage

import (
	"context"
	"errors"
	"github.com/DateMine/bot-domain/pkg/models/offer"
	"github.com/DateMine/bot-domain/pkg/repository"
	"sync"
)

type OfferStorage struct {
	mx              sync.RWMutex
	Offers          map[string]offer.Offer
	Portion         int
	offerRepository repository.OfferRepository
}

func NewOfferStorage(offerRepository repository.OfferRepository) *OfferStorage {
	return &OfferStorage{
		Offers:          make(map[string]offer.Offer),
		offerRepository: offerRepository,
	}
}

func (s *OfferStorage) Add(offer offer.Offer, ctx context.Context) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	if s.Offers == nil {
		return errors.New("offers is nil")
	}
	s.Offers[offer.OfferId] = offer
	if len(s.Offers) == s.Portion {
		s.commit(ctx)
	}

	return nil
}

func (s *OfferStorage) commit(ctx context.Context) error {
	s.mx.Lock()
	defer s.mx.Unlock()
	offers := make([]offer.Offer, 0, len(s.Offers))
	for _, offer := range s.Offers {
		offers = append(offers, offer)
	}
	s.offerRepository.AddRange(offers, ctx)
	s.Clear()
	return nil
}

func (s *OfferStorage) Clear() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.Offers = make(map[string]offer.Offer)
}

func (s *OfferStorage) Count() int {
	return len(s.Offers)
}
