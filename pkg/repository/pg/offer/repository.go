package offer

import (
	"context"
	"github.com/DateMine/bot-domain/internal/converter"
	"github.com/DateMine/bot-domain/pkg/models/db/offer"
	"github.com/DateMine/bot-domain/pkg/repository"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

type repo struct {
	pg *pgxpool.Pool
}

func NewRepository(pg *pgxpool.Pool) repository.OfferRepository {
	return &repo{pg: pg}
}
func (r *repo) Add(offer *offer.Offer, ctx context.Context) error {
	_, err := r.pg.Exec(ctx, "CALL add_offer($1, $2, $3, ...)", offer.OfferId, offer.Url, offer.OfferXML)
	return err
}

func (r *repo) AddRange(offers []offer.Offer, ctx context.Context) error {
	offersArray := make([]interface{}, len(offers))
	for i, offer := range offers {
		offerXml := converter.ToOfferXml(offer.OfferXML)
		offerJson := converter.ToOfferJson(offer.OfferJson)
		offersArray[i] = []interface{}{
			offer.OfferId,
			offer.ParserId,
			offer.Url,
			offer.CreationDate,
			offer.UpdateDate,
			offerXml,
			offerJson,
		}
	}
	_, err := r.pg.Exec(context.Background(),
		"CALL add_or_update_offers($1)",
		offersArray)
	if err != nil {
		log.Fatalf("Error calling stored procedure: %v\n", err)
	}

	log.Println("Stored procedure executed successfully!")
	return err
}
func (r *repo) GetAll(parserId int, days int, ctx context.Context) (*offer.Offer, error) {
	return nil, nil
}
