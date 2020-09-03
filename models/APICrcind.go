package models

import "encoding/xml"

// CrcindAPI .
type CrcindAPI struct {
	XMLName xml.Name `xml:"Envelope"`
	Text    string   `xml:",chardata"`
	Body    struct {
		Text                  string `xml:",chardata"`
		GetListByNameResponse struct {
			Text                string `xml:",chardata"`
			GetListByNameResult struct {
				Text                 string `xml:",chardata"`
				PersonIdentification []struct {
					Text string `xml:",chardata"`
					ID   string `xml:"ID"`
					Name string `xml:"Name"`
					SSN  string `xml:"SSN"`
					DOB  string `xml:"DOB"`
				} `xml:"PersonIdentification"`
			} `xml:"GetListByNameResult"`
		} `xml:"GetListByNameResponse"`
	} `xml:"Body"`
}
