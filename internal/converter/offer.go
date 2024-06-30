package converter

import (
	"encoding/json"
	"encoding/xml"
	"github.com/DateMine/bot-domain/pkg/models/db/offer"
)

func ToOfferXml(offer offer.OfferXML) string {
	out, _ := xml.MarshalIndent(offer, " ", "  ")
	return string(out)
}

func ToOfferJson(offer interface{}) string {
	out, _ := json.Marshal(offer)
	return string(out)
}
