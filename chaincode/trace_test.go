package main

import (
	"testing"
	"encoding/json"
)

func TestAddTrace(t *testing.T) {
	var traceInfo Trace
	traceInfo.UniqueCode = "02130041810053096512"
	traceInfo.ProductInfo.Barcode = "7613035260160"
	traceInfo.ProductInfo.Name = "启赋3段@6x900g N2 CN"
	traceInfo.ProductInfo.Category = "3 Stage"
	traceInfo.ProductInfo.Package = "听装"
	traceInfo.ProductInfo.Spec = "900"
	traceInfo.ProductInfo.BatchNum = "8301755841"
	traceInfo.ProductInfo.ProduceDate = "2018-10-28"
	traceInfo.ProductInfo.ValidDate = "2020-10-27"
	traceInfo.LogisticsInfo.DepartDate = "2018-11-05"
	traceInfo.LogisticsInfo.ArrivalDate = "2018-12-08"
	traceInfo.ManufactInfo.Origin = "爱尔兰"
	traceInfo.ManufactInfo.Factory = "Wyeth Nutritionals Ireland Ltd."
	traceInfo.ManufactInfo.Address = "Askeaton,Co.Limerick,Ireland"

	b, err := json.Marshal(traceInfo)
	if err != nil {
		t.Errorf("Json marshal failed: %v", err)
	}

	t.Log(string(b))
}