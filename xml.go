package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func getXMLDocument(products []map[string]string) {

}

func getXMLProduct(product map[string]string) {
	type Product struct {
		XMLName     xml.Name `xml:"offer"`
		ID          string   `xml:"id,attr"`
		Available   string   `xml:"available,attr"`
		URL         string   `xml:"url"`
		Price       string   `xml:"price"`
		OldPrice    string   `xml:"oldprice"`
		CurrencyID  string   `xml:"currencyId"`
		CategoryID  string   `xml:"categoryId"`
		Picture     string   `xml:"picture"`
		Name        string   `xml:"name"`
		Vendor      string   `xml:"vendor"`
		Description string   `xml:"description"`
	}

	v := &Product{
		ID:          product["id"],
		Available:   "true",
		URL:         product["url"],
		Price:       product["price"],
		OldPrice:    product["oldprice"],
		CurrencyID:  "RUB",
		CategoryID:  "1337",
		Picture:     product["image"],
		Name:        product["title"],
		Vendor:      "ИВ РОШЕ",
		Description: product["desc"],
	}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent(" ", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}
}
