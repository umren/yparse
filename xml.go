package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
)

func getXMLDocument(products []map[string]string) {
	productsXML := ""

	for _, product := range products {
		productsXML += "\n" + getXMLProduct(product)
	}

	fmt.Println(productsXML)
}

func getXMLProduct(product map[string]string) string {
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

	var b bytes.Buffer
	enc := xml.NewEncoder(&b)
	enc.Indent(" ", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	return b.String()
}
