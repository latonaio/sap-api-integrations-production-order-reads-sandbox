package file_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	ManufacturingOrder    struct {
		ManufacturingOrder             string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    float64     `json:"quantity"`
		MfgOrderItemActualDeviationQty float64     `json:"picked_quantity"`
		Price                          float64     `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ManufacturingOrder struct {
		ManufacturingOrder             string      `json:"document_no"`
		StatusCode                     string      `json:"status"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    float64     `json:"quantity"`
		MfgOrderItemActualDeviationQty float64     `json:"completed_quantity"`
	    PlannedStartDate               string      `json:"planned_start_date"`
	    MfgOrderItemPlndDeliveryDate   string      `json:"planned_validated_date"`
	    ActualStartDate                string      `json:"actual_start_date"`
	    MfgOrderItemActualDeliveryDate string      `json:"actual_validated_date"`
	    Batch                          string      `json:"batch"`
		BillOfOperations struct {
			OrderIntBillOfOperationsItem string      `json:"work_no"`
			OpPlannedTotalQuantity       float64     `json:"quantity"`
			OpTotalConfirmedYieldQty     float64     `json:"completed_quantity"`
			ErroredQuantity              float64     `json:"errored_quantity"`
			Component                    string      `json:"component"`
			MaterialCompOriginalQuantity float64     `json:"planned_component_quantity"`
			OpErlstSchedldExecStrtDte    string      `json:"planned_start_date"`
			OpErlstSchedldExecStrtTme    string      `json:"planned_start_time"`
			OpErlstSchedldExecEndDte     string      `json:"planned_validated_date"`
			OpErlstSchedldExecEndTme     string      `json:"planned_validated_time"`
			OpActualExecutionStartDate   string      `json:"actual_start_date"`
			OpActualExecutionStartTime   string      `json:"actual_start_time"`
			OpActualExecutionEndDate     string      `json:"actual_validated_date"`
			OpActualExecutionEndTime     string      `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"production_order"`
	APISchema                      string      `json:"api_schema"`
	Material                       string      `json:"material_code"`
	ProductionPlant                string      `json:"plant/supplier"`
	Stock                          float64     `json:"stock"`
	ManufacturingOrderType         string      `json:"document_type"`
	ManufacturingOrder             string      `json:"document_no"`
	MfgOrderItemPlndDeliveryDate   string      `json:"planned_date"`
	MfgOrderItemActualDeliveryDate string      `json:"validated_date"`
	Deleted                        string      `json:"deleted"`
}

type SDC struct {
	ConnectionKey   string `json:"connection_key"`
	Result          bool   `json:"result"`
	RedisKey        string `json:"redis_key"`
	Filepath        string `json:"filepath"`
	ManufacturingOrder struct {
		ManufacturingOrder             string  `json:"ManufacturingOrder"`
		ManufacturingOrderType         string  `json:"ManufacturingOrderType"`
		Material                       string  `json:"Material"`
		ProductionPlant                string  `json:"ProductionPlant"`
		MfgOrderItemPlannedTotalQty    float64 `json:"MfgOrderItemPlannedTotalQty"`
		MfgOrderItemPlannedScrapQty    float64 `json:"MfgOrderItemPlannedScrapQty"`
		MfgOrderItemGoodsReceiptQty    float64 `json:"MfgOrderItemGoodsReceiptQty"`
		MfgOrderItemActualDeviationQty float64 `json:"MfgOrderItemActualDeviationQty"`
		MfgOrderItemPlndDeliveryDate   string  `json:"MfgOrderItemPlndDeliveryDate"`
		MfgOrderItemActualDeliveryDate string  `json:"MfgOrderItemActualDeliveryDate"`
		IsCompletelyDelivered          string  `json:"IsCompletelyDelivered"`
		MRPArea                        string  `json:"MRPArea"`
		StorageLocation                string  `json:"StorageLocation"`
		Batch                          string  `json:"Batch"`
		ProductionUnit                 string  `json:"ProductionUnit"`
		Status                         struct {
			StatusCode      string `json:"StatusCode"`
			IsUserStatus    string `json:"IsUserStatus"`
			StatusShortName string `json:"StatusShortName"`
			StatusName      string `json:"StatusName"`
		} `json:"Status"`
		BillofMaterial struct {
			BillOfMaterialItemNumber      string      `json:"BillOfMaterialItemNumber"`
			Component                     string      `json:"Component"`
			Plant                         string      `json:"Plant"`
			BaseUnit                      string      `json:"BaseUnit"`
			RequiredQuantity              float64     `json:"RequiredQuantity"`
			MaterialCompOriginalQuantity  float64     `json:"MaterialCompOriginalQuantity"`
			Reservation                   int         `json:"Reservation"`
			ReservationItem               int         `json:"ReservationItem"`
			QuantityIsFixed               string      `json:"QuantityIsFixed"`
			ComponentScrapInPercent       float64     `json:"ComponentScrapInPercent"`
			WithdrawnQuantity             float64     `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity    float64     `json:"ConfirmedAvailableQuantity"`
			WithdrawnQuantityAmount       float64     `json:"WithdrawnQuantityAmount"`
			OrderInternalBillOfOperations string      `json:"OrderInternalBillOfOperations"`
			MatlCompRequirementDate       string      `json:"MatlCompRequirementDate"`
			ReservationIsFinallyIssued    string      `json:"ReservationIsFinallyIssued"`
			MatlCompIsMarkedForBackflush  string      `json:"MatlCompIsMarkedForBackflush"`
			StorageLocation               string      `json:"StorageLocation"`
			GoodsMovementType             string      `json:"GoodsMovementType"`
		} `json:"BillofMaterial"`
		BillofOperations struct {
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem  struct {
				OrderIntBillOfOperationsItem  string  `json:"OrderIntBillOfOperationsItem"`
				ManufacturingOrderSequence    int     `json:"ManufacturingOrderSequence"`
				MfgOrderOperationText         string  `json:"MfgOrderOperationText"`
				OperationIsReleased           string  `json:"OperationIsReleased"`
				OperationIsConfirmed          string  `json:"OperationIsConfirmed"`
				OperationIsPartiallyConfirmed string  `json:"OperationIsPartiallyConfirmed"`
				OperationIsDeleted            string  `json:"OperationIsDeleted"`
				OperationIsTechlyCompleted    string  `json:"OperationIsTechlyCompleted"`
				OperationIsClosed             string  `json:"OperationIsClosed"`
				ProductionPlant               string  `json:"ProductionPlant"`
				WorkCenterInternalID          int     `json:"WorkCenterInternalID"`
				OpErlstSchedldExecStrtDte     string  `json:"OpErlstSchedldExecStrtDte"`
				OpErlstSchedldExecStrtTme     string  `json:"OpErlstSchedldExecStrtTme"`
				OpErlstSchedldExecEndDte      string  `json:"OpErlstSchedldExecEndDte"`
				OpErlstSchedldExecEndTme      string  `json:"OpErlstSchedldExecEndTme"`
				OpActualExecutionStartDate    string  `json:"OpActualExecutionStartDate"`
				OpActualExecutionStartTime    string  `json:"OpActualExecutionStartTime"`
				OpActualExecutionEndDate      string  `json:"OpActualExecutionEndDate"`
				OpActualExecutionEndTime      string  `json:"OpActualExecutionEndTime"`
				OpActualExecutionDays         int     `json:"OpActualExecutionDays"`
				OperationUnit                 string  `json:"OperationUnit"`
				OpPlannedTotalQuantity        float64 `json:"OpPlannedTotalQuantity"`
				OpTotalConfirmedYieldQty      float64 `json:"OpTotalConfirmedYieldQty"`
				LastChangeDateTime            string  `json:"LastChangeDateTime"`
			} `json:"OrderIntBillOfOperationsItem"`
		} `json:"BillofOperations"`
	} `json:"ProductionOrder"`
	APISchema          string `json:"api_schema"`
	ManufacturingOrder string `json:"production_order"`
	Deleted            string `json:"deleted"`
}