package sap_api_output_formatter

type ProductionOrderReads struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	ManufacturingOrder string `json:"production_order"`
	Deleted            bool   `json:"deleted"`
}

type General struct {
			ManufacturingOrder            string `json:"ManufacturingOrder"`
			ManufacturingOrderCategory    string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType        string `json:"ManufacturingOrderType"`
			OrderIsCreated                string `json:"OrderIsCreated"`
			OrderIsReleased               string `json:"OrderIsReleased"`
			OrderIsPrinted                string `json:"OrderIsPrinted"`
			OrderIsConfirmed              string `json:"OrderIsConfirmed"`
			OrderIsPartiallyConfirmed     string `json:"OrderIsPartiallyConfirmed"`
			OrderIsDelivered              string `json:"OrderIsDelivered"`
			OrderIsDeleted                string `json:"OrderIsDeleted"`
			OrderIsPreCosted              string `json:"OrderIsPreCosted"`
			SettlementRuleIsCreated       string `json:"SettlementRuleIsCreated"`
			OrderIsPartiallyReleased      string `json:"OrderIsPartiallyReleased"`
			OrderIsLocked                 string `json:"OrderIsLocked"`
			OrderIsTechnicallyCompleted   string `json:"OrderIsTechnicallyCompleted"`
			OrderIsClosed                 string `json:"OrderIsClosed"`
			OrderIsPartiallyDelivered     string `json:"OrderIsPartiallyDelivered"`
			OrderIsMarkedForDeletion      string `json:"OrderIsMarkedForDeletion"`
			OrderIsScheduled              string `json:"OrderIsScheduled"`
			OrderHasGeneratedOperations   string `json:"OrderHasGeneratedOperations"`
			MaterialAvailyIsNotChecked    string `json:"MaterialAvailyIsNotChecked"`
			MfgOrderCreationDate          string `json:"MfgOrderCreationDate"`
			MfgOrderCreationTime          string `json:"MfgOrderCreationTime"`
			LastChangeDateTime            string `json:"LastChangeDateTime"`
			Material                      string `json:"Material"`
			StorageLocation               string `json:"StorageLocation"`
			GoodsRecipientName            string `json:"GoodsRecipientName"`
			UnloadingPointName            string `json:"UnloadingPointName"`
			MaterialGoodsReceiptDuration  string `json:"MaterialGoodsReceiptDuration"`
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			ProductionPlant               string `json:"ProductionPlant"`
			Plant                         string `json:"Plant"`
			MRPArea                       string `json:"MRPArea"`
			MRPController                 string `json:"MRPController"`
			ProductionSupervisor          string `json:"ProductionSupervisor"`
			ProductionVersion             string `json:"ProductionVersion"`
			PlannedOrder                  string `json:"PlannedOrder"`
			SalesOrder                    string `json:"SalesOrder"`
			SalesOrderItem                string `json:"SalesOrderItem"`
			BasicSchedulingType           string `json:"BasicSchedulingType"`
			BusinessArea                  string `json:"BusinessArea"`
			CompanyCode                   string `json:"CompanyCode"`
			ProfitCenter                  string `json:"ProfitCenter"`
			FunctionalArea                string `json:"FunctionalArea"`
			MfgOrderPlannedStartDate      string `json:"MfgOrderPlannedStartDate"`
			MfgOrderPlannedStartTime      string `json:"MfgOrderPlannedStartTime"`
			MfgOrderPlannedEndDate        string `json:"MfgOrderPlannedEndDate"`
			MfgOrderPlannedEndTime        string `json:"MfgOrderPlannedEndTime"`
			MfgOrderScheduledStartDate    string `json:"MfgOrderScheduledStartDate"`
			MfgOrderScheduledStartTime    string `json:"MfgOrderScheduledStartTime"`
			MfgOrderScheduledEndDate      string `json:"MfgOrderScheduledEndDate"`
			MfgOrderScheduledEndTime      string `json:"MfgOrderScheduledEndTime"`
			MfgOrderActualReleaseDate     string `json:"MfgOrderActualReleaseDate"`
			ProductionUnit                string `json:"ProductionUnit"`
			TotalQuantity                 string `json:"TotalQuantity"`
			MfgOrderPlannedScrapQty       string `json:"MfgOrderPlannedScrapQty"`
			MfgOrderConfirmedYieldQty     string `json:"MfgOrderConfirmedYieldQty"`
			WBSElementExternalID          string `json:"WBSElementExternalID"`
			OrderLongText                 string `json:"OrderLongText"`
			ToProductionOrderComponent    string `json:"to_ProductionOrderComponent"`
			ToProductionOrderItem         string `json:"to_ProductionOrderItem"`
			ToProductionOrderOperation    string `json:"to_ProductionOrderOperation"`
			ToProductionOrderStatus       string `json:"to_ProductionOrderStatus"`
}

type ToProductionOrderComponent struct {
			Reservation                    string `json:"Reservation"`
			ReservationItem                string `json:"ReservationItem"`
			MaterialGroup                  string `json:"MaterialGroup"`
			Material                       string `json:"Material"`
			Plant                          string `json:"Plant"`
			ManufacturingOrderCategory     string `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string `json:"ManufacturingOrderType"`
			ManufacturingOrder             string `json:"ManufacturingOrder"`
			ManufacturingOrderSequence     string `json:"ManufacturingOrderSequence"`
			ManufacturingOrderOperation    string `json:"ManufacturingOrderOperation"`
			ProductionPlant                string `json:"ProductionPlant"`
			OrderInternalBillOfOperations  string `json:"OrderInternalBillOfOperations"`
			MatlCompRequirementDate        string `json:"MatlCompRequirementDate"`
			MatlCompRequirementTime        string `json:"MatlCompRequirementTime"`
			ReservationIsFinallyIssued     bool   `json:"ReservationIsFinallyIssued"`
			IsBulkMaterialComponent        bool   `json:"IsBulkMaterialComponent"`
			MatlCompIsMarkedForBackflush   bool   `json:"MatlCompIsMarkedForBackflush"`
			MaterialCompIsCostRelevant     string `json:"MaterialCompIsCostRelevant"`
			SalesOrder                     string `json:"SalesOrder"`
			SalesOrderItem                 string `json:"SalesOrderItem"`
			SortField                      string `json:"SortField"`
			BillOfMaterialCategory         string `json:"BillOfMaterialCategory"`
			BOMItem                        string `json:"BOMItem"`
			BOMItemCategory                string `json:"BOMItemCategory"`
			BillOfMaterialItemNumber       string `json:"BillOfMaterialItemNumber"`
			BOMItemDescription             string `json:"BOMItemDescription"`
			StorageLocation                string `json:"StorageLocation"`
			Batch                          string `json:"Batch"`
			BatchSplitType                 string `json:"BatchSplitType"`
			GoodsMovementType              string `json:"GoodsMovementType"`
			SupplyArea                     string `json:"SupplyArea"`
			GoodsRecipientName             string `json:"GoodsRecipientName"`
			UnloadingPointName             string `json:"UnloadingPointName"`
			MaterialCompIsAlternativeItem  bool   `json:"MaterialCompIsAlternativeItem"`
			AlternativeItemGroup           string `json:"AlternativeItemGroup"`
			AlternativeItemStrategy        string `json:"AlternativeItemStrategy"`
			AlternativeItemPriority        string `json:"AlternativeItemPriority"`
			UsageProbabilityPercent        string `json:"UsageProbabilityPercent"`
			MaterialComponentIsPhantomItem bool   `json:"MaterialComponentIsPhantomItem"`
			LeadTimeOffset                 string `json:"LeadTimeOffset"`
			QuantityIsFixed                bool   `json:"QuantityIsFixed"`
			IsNetScrap                     bool   `json:"IsNetScrap"`
			ComponentScrapInPercent        string `json:"ComponentScrapInPercent"`
			OperationScrapInPercent        string `json:"OperationScrapInPercent"`
			BaseUnit                       string `json:"BaseUnit"`
			RequiredQuantity               string `json:"RequiredQuantity"`
			WithdrawnQuantity              string `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity     string `json:"ConfirmedAvailableQuantity"`
			MaterialCompOriginalQuantity   string `json:"MaterialCompOriginalQuantity"`
			Currency                       string `json:"Currency"`
			WithdrawnQuantityAmount        string `json:"WithdrawnQuantityAmount"`
}

type ToProductionOrderItem struct {
			ManufacturingOrder             string      `json:"ManufacturingOrder"`
			ManufacturingOrderItem         string      `json:"ManufacturingOrderItem"`
			ManufacturingOrderCategory     string      `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string      `json:"ManufacturingOrderType"`
			IsCompletelyDelivered          bool        `json:"IsCompletelyDelivered"`
			Material                       string      `json:"Material"`
			ProductionPlant                string      `json:"ProductionPlant"`
			Plant                          string      `json:"Plant"`
			MRPArea                        string      `json:"MRPArea"`
			QuantityDistributionKey        string      `json:"QuantityDistributionKey"`
			MaterialGoodsReceiptDuration   string      `json:"MaterialGoodsReceiptDuration"`
			StorageLocation                string      `json:"StorageLocation"`
			Batch                          string      `json:"Batch"`
			InventoryUsabilityCode         string      `json:"InventoryUsabilityCode"`
			GoodsRecipientName             string      `json:"GoodsRecipientName"`
			UnloadingPointName             string      `json:"UnloadingPointName"`
			MfgOrderItemPlndDeliveryDate   string      `json:"MfgOrderItemPlndDeliveryDate"`
			MfgOrderItemActualDeliveryDate string      `json:"MfgOrderItemActualDeliveryDate"`
			ProductionUnit                 string      `json:"ProductionUnit"`
			MfgOrderItemPlannedTotalQty    string      `json:"MfgOrderItemPlannedTotalQty"`
			MfgOrderItemPlannedScrapQty    string      `json:"MfgOrderItemPlannedScrapQty"`
			MfgOrderItemGoodsReceiptQty    string      `json:"MfgOrderItemGoodsReceiptQty"`
			MfgOrderItemActualDeviationQty string      `json:"MfgOrderItemActualDeviationQty"`
}

type ToProductionOrderOperation struct {
			OrderInternalBillOfOperations  string      `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem   string      `json:"OrderIntBillOfOperationsItem"`
			ManufacturingOrder             string      `json:"ManufacturingOrder"`
			ManufacturingOrderSequence     string      `json:"ManufacturingOrderSequence"`
			MfgOrderSequenceText           string      `json:"MfgOrderSequenceText"`
			ManufacturingOrderOperation    string      `json:"ManufacturingOrderOperation"`
			ManufacturingOrderSubOperation string      `json:"ManufacturingOrderSubOperation"`
			ManufacturingOrderCategory     string      `json:"ManufacturingOrderCategory"`
			ManufacturingOrderType         string      `json:"ManufacturingOrderType"`
			MfgOrderOperationText          string      `json:"MfgOrderOperationText"`
			OperationIsCreated             string      `json:"OperationIsCreated"`
			OperationIsReleased            string      `json:"OperationIsReleased"`
			OperationIsPrinted             string      `json:"OperationIsPrinted"`
			OperationIsConfirmed           string      `json:"OperationIsConfirmed"`
			OperationIsPartiallyConfirmed  string      `json:"OperationIsPartiallyConfirmed"`
			OperationIsDeleted             string      `json:"OperationIsDeleted"`
			OperationIsTechlyCompleted     string      `json:"OperationIsTechlyCompleted"`
			OperationIsClosed              string      `json:"OperationIsClosed"`
			OperationIsScheduled           string      `json:"OperationIsScheduled"`
			OperationIsPartiallyDelivered  string      `json:"OperationIsPartiallyDelivered"`
			OperationIsDelivered           string      `json:"OperationIsDelivered"`
			ProductionPlant                string      `json:"ProductionPlant"`
			WorkCenterInternalID           string      `json:"WorkCenterInternalID"`
			WorkCenterTypeCode             string      `json:"WorkCenterTypeCode"`
			WorkCenter                     string      `json:"WorkCenter"`
			OperationControlProfile        string      `json:"OperationControlProfile"`
			OpErlstSchedldExecStrtDte      string      `json:"OpErlstSchedldExecStrtDte"`
			OpErlstSchedldExecStrtTme      string      `json:"OpErlstSchedldExecStrtTme"`
			OpErlstSchedldExecEndDte       string      `json:"OpErlstSchedldExecEndDte"`
			OpErlstSchedldExecEndTme       string      `json:"OpErlstSchedldExecEndTme"`
			OpActualExecutionStartDate     string      `json:"OpActualExecutionStartDate"`
			OpActualExecutionStartTime     string      `json:"OpActualExecutionStartTime"`
			OpActualExecutionEndDate       string      `json:"OpActualExecutionEndDate"`
			OpActualExecutionEndTime       string      `json:"OpActualExecutionEndTime"`
			ErlstSchedldExecDurnInWorkdays int         `json:"ErlstSchedldExecDurnInWorkdays"`
			OpActualExecutionDays          int         `json:"OpActualExecutionDays"`
			OperationUnit                  string      `json:"OperationUnit"`
			OpPlannedTotalQuantity         string      `json:"OpPlannedTotalQuantity"`
			OpTotalConfirmedYieldQty       string      `json:"OpTotalConfirmedYieldQty"`
			LastChangeDateTime             string      `json:"LastChangeDateTime"`
}

type ToProductionOrderStatus struct {
			ManufacturingOrder string `json:"ManufacturingOrder"`
			StatusCode         string `json:"StatusCode"`
			IsUserStatus       bool   `json:"IsUserStatus"`
			StatusShortName    string `json:"StatusShortName"`
			StatusName         string `json:"StatusName"`
}
