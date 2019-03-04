package main

type TraceProduct struct {
	Barcode string `json: "Barcode"`
	Name string `json: "Name"` 
	Category string `json: "Category"`
	Package string `json: "Package"`
	Spec string `json: "Spec"`
	BatchNum string `json: "BatchNum"`
	ProduceDate string `json: "ProduceDate"`
	ValidDate string `json: "ValidDate"`
}

type TraceLogistics struct {
	DepartDate string `json: "DepartDate"`
	ArrivalDate string `json: "ArrivalDate"`
}

type TraceManufact struct {
	Origin string `json: "Origin"`
	Factory string `json: "Factory"`
	Address string `json: "Address"`
}

type Trace struct {
	UniqueCode string `json: "UniqueCode"`
	ProductInfo TraceProduct `json: "ProductInfo"`
	LogisticsInfo TraceLogistics `json: "LogisticsInfo"`
	ManufactInfo TraceManufact `json: "ManufactInfo"`
	Historys []HistoryItem
}

type HistoryItem struct {
	TxId	string
	TraceInfo	Trace
}