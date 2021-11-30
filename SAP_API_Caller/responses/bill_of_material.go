package responses

type BillofMaterial struct {
	D struct {
		Results []struct {
			Metadata struct {
				ID   string `json:"id"`
				URI  string `json:"uri"`
				Type string `json:"type"`
			} `json:"__metadata"`
			ManufacturingOrder            string `json:"ManufacturingOrder"`
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem  string `json:"OrderIntBillOfOperationsItem"`
			BillOfMaterialItemNumber      string `json:"BillOfMaterialItemNumber"`
			Component                     string `json:"Component"`
			Plant                         string `json:"Plant"`
			BaseUnit                      string `json:"BaseUnit"`
			RequiredQuantity              string `json:"RequiredQuantity"`
			MaterialCompOriginalQuantity  string `json:"MaterialCompOriginalQuantity"`
			Reservation                   string `json:"Reservation"`
			ReservationItem               string `json:"ReservationItem"`
			QuantityIsFixed               bool   `json:"QuantityIsFixed"`
			ComponentScrapInPercent       string `json:"ComponentScrapInPercent"`
			WithdrawnQuantity             string `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity    string `json:"ConfirmedAvailableQuantity"`
			WithdrawnQuantityAmount       string `json:"WithdrawnQuantityAmount"`
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			MatlCompRequirementDate       string `json:"MatlCompRequirementDate"`
			ReservationIsFinallyIssued    bool   `json:"ReservationIsFinallyIssued"`
			MatlCompIsMarkedForBackflush  bool   `json:"MatlCompIsMarkedForBackflush"`
			StorageLocation               string `json:"StorageLocation"`
			GoodsMovementType             string `json:"GoodsMovementType"`
		} `json:"results"`
	} `json:"d"`
}
