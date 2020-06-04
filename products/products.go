package products

import (
	"encoding/xml"
	"fmt"

	"github.com/nikolasmelui/go-xml2json-mapper/cconfig"
)

// ProductsURL ...
var ProductsURL = cconfig.Config.BaseURL + "/products"

type stock struct {
	XMLName     xml.Name `xml:"Склад"`
	ID          string   `xml:"Ид,attr"`
	ReceiptDate string   `xml:"ДатаПоступления,attr"`
	Quantity    string   `xml:"Количество,attr"`
}

type admission struct {
	XMLName xml.Name `xml:"Поступление"`
	Stock   stock    `xml:"Склад"`
}

type availability struct {
	XMLName xml.Name `xml:"Наличие"`
	Stock   stock    `xml:"Склад"`
}

type price struct {
	XMLName xml.Name `xml:"Цена"`
	ID      string   `xml:"Ид,attr"`
	Title   string   `xml:"Наименование,attr"`
	Price   string   `xml:"Цена,attr"`
}

type prices struct {
	XMLName xml.Name `xml:"Цены"`
	Prices  []price  `xml:"Цена"`
}

type description struct {
	XMLName xml.Name `xml:"Описание"`
}

// Product ...
type Product struct {
	XMLName      xml.Name     `xml:"Товар"`
	ID           string       `xml:"Ид,attr"`
	Title        string       `xml:"Наименование,attr"`
	VendorCode   string       `xml:"Артикул,attr"`
	Multiplicity string       `xml:"Кратность,attr"`
	Reservation  string       `xml:"Зарезервировано,attr"`
	Description  description  `xml:"Описание"`
	Prices       prices       `xml:"Цены"`
	Availability availability `xml:"Наличие"`
	Admission    admission    `xml:"Поступление"`
}

// Products ...
type Products struct {
	XMLName  xml.Name  `xml:"Товары"`
	XMLNS    string    `xml:"xmlns"`
	XMLNSXS  string    `xml:"xmlns:xs"`
	XMLNSXSI string    `xml:"xmlns:xsi"`
	Version  string    `xml:"Версия"`
	Products []Product `xml:"Товар"`
}

// BeautyPrint ...
func (p *Products) BeautyPrint() {
	for _, product := range p.Products {
		fmt.Printf("Товар:\n")
		fmt.Printf("  Ид: %s\n", product.ID)
		fmt.Printf("  Наименование: %s\n", product.Title)
		fmt.Printf("  Артикул: %s\n", product.VendorCode)
		fmt.Printf("  Кратность: %s\n", product.Multiplicity)
		fmt.Printf("  Зарезервировано: %s\n", product.Reservation)
		fmt.Printf("    Описание: %s\n", product.Description)
		fmt.Printf("    Цены:\n")
		for _, price := range product.Prices.Prices {
			fmt.Printf("      Ценa:\n")
			fmt.Printf("        Ид: %s\n", price.ID)
			fmt.Printf("        Наименование: %s\n", price.Title)
			fmt.Printf("        Цена: %s\n", price.Price)
		}
		fmt.Printf("    Наличие:\n")
		fmt.Printf("      Cклад:\n")
		fmt.Printf("        Ид: %s\n", product.Availability.Stock.ID)
		fmt.Printf("        ДатаПоступления: %s\n", product.Availability.Stock.ReceiptDate)
		fmt.Printf("        Количество: %s\n", product.Availability.Stock.Quantity)

		fmt.Printf("    Поступление:\n")
		fmt.Printf("      Cклад:\n")
		fmt.Printf("        Ид: %s\n", product.Admission.Stock.ID)
		fmt.Printf("        ДатаПоступления: %s\n", product.Admission.Stock.ReceiptDate)
		fmt.Printf("        Количество: %s\n", product.Admission.Stock.Quantity)
		fmt.Println()
	}
}
