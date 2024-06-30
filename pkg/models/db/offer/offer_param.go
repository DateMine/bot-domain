package offer

import "encoding/xml"

type OfferParam struct {
	XMLName xml.Name `xml:"param"`
	Key     string   `xml:"key"`
	Value   string   `xml:"value"`
}
