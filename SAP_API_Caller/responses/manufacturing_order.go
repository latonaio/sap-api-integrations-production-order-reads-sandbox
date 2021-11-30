package responses

type ManufacturingOrder struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderType         string `json:"ManufacturingOrderType"`
			Material                       string `json:"Material"`
			ProductionPlant                string `json:"ProductionPlant"`
			MfgOrderItemPlannedTotalQty    string `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    string `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    string `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty string `json:"MfgOrderItemActualDeviationQty"`
			MfgOrderItemPlndDeliveryDate   string `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate string `json:"MfgOrderItemActualDeliveryDate"`
			IsCompletelyDelivered          bool   `json:"IsCompletelyDelivered"`
			MRPArea                        string `json:"MRPArea"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			ProductionUnit                 string `json:"ProductionUnit"`
		} `json:"results"`
	} `json:"d"`
}
