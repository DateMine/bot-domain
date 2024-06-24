package offer

import "time"

type Offer struct {
	Id           int
	OfferId      string
	ParserId     int
	Url          string
	CreationDate time.Time
	UpdateDate   time.Time
	OfferXML     OfferXML
	OfferJson    interface{}
}
