package offer

import "encoding/xml"

type OfferXML struct {
	XMLName xml.Name `xml:"Offer"`
	Params  []OfferParam
}
