package dto

import "time"


type GetWBOrdersInput struct {
	DateFrom time.Time `json:"dateFrom"`
	Flag     int       `json:"flag"`
}

type GetWBordersOutput struct {
	Date            string  `json:"date"`
	LastChangeDate  string  `json:"lastChangeDate"`
	WarehouseName   string  `json:"warehouseName"`
	WarehouseType   string  `json:"warehouseType"`
	CountryName     string  `json:"countryName"`
	OblastOkrugName string  `json:"oblastOkrugName"`
	RegionName      string  `json:"regionName"`
	SupplierArticle string  `json:"supplierArticle"`
	NmID            int     `json:"nmId"`
	Barcode         string  `json:"barcode"`
	Category        string  `json:"category"`
	Subject         string  `json:"subject"`
	Brand           string  `json:"brand"`
	TechSize        string  `json:"techSize"`
	IncomeID        int     `json:"incomeID"`
	IsSupply        bool    `json:"isSupply"`
	IsRealization   bool    `json:"isRealization"`
	TotalPrice      float64 `json:"totalPrice"`
	DiscountPercent int     `json:"discountPercent"`
	Spp             int     `json:"spp"`
	FinishedPrice   float64 `json:"finishedPrice"`
	PriceWithDisc   float64 `json:"priceWithDisc"`
	IsCancel        bool    `json:"isCancel"`
	CancelDate      string  `json:"cancelDate"`
	OrderType       string  `json:"orderType"`
	Sticker         string  `json:"sticker"`
	GNumber         string  `json:"gNumber"`
	Srid            string  `json:"srid"`
}


