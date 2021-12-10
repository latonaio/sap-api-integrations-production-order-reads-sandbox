package sap_api_input_reader

type EC_MC struct {
	ConnectionKey string      `json:"connection_key"`
	Result        bool        `json:"result"`
	RedisKey      string      `json:"redis_key"`
	Filepath      string      `json:"filepath"`
	ProuctionOrder struct {
		ManufacturingOrder             string      `json:"document_no"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string      `json:"quantity"`
		MfgOrderItemActualDeviationQty string      `json:"picked_quantity"`
		Price                          string      `json:"price"`
	    Batch                          string      `json:"batch"`
	} `json:"document"`
	ManufacturingOrder struct {
		ManufacturingOrder             string      `json:"document_no"`
		StatusCode                     string      `json:"status"`
		DeliverTo                      string      `json:"deliver_to"`
		MfgOrderItemPlannedTotalQty    string      `json:"quantity"`
		MfgOrderItemActualDeviationQty string      `json:"completed_quantity"`
	    PlannedStartDate               string      `json:"planned_start_date"`
	    MfgOrderItemPlndDeliveryDate   string      `json:"planned_validated_date"`
	    ActualStartDate                string      `json:"actual_start_date"`
	    MfgOrderItemActualDeliveryDate string      `json:"actual_validated_date"`
	    Batch                          string      `json:"batch"`
		BillOfOperations struct {
			OrderIntBillOfOperationsItem string      `json:"work_no"`
			OpPlannedTotalQuantity       string      `json:"quantity"`
			OpTotalConfirmedYieldQty     string      `json:"completed_quantity"`
			ErroredQuantity              string      `json:"errored_quantity"`
			Component                    string      `json:"component"`
			MaterialCompOriginalQuantity string      `json:"planned_component_quantity"`
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
	MaterialCode                   string      `json:"material_code"`
	ProductionPlant                string      `json:"plant/supplier"`
	Stock                          string      `json:"stock"`
	ManufacturingOrderType         string      `json:"document_type"`
	ManufacturingOrderNo           string      `json:"document_no"`
	MfgOrderItemPlndDeliveryDate   string      `json:"planned_date"`
	MfgOrderItemActualDeliveryDate string      `json:"validated_date"`
	Deleted                        bool        `json:"deleted"`
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
		MfgOrderItemPlannedTotalQty    string  `json:"MfgOrderItemPlannedTotalQty"`
		MfgOrderItemPlannedScrapQty    string  `json:"MfgOrderItemPlannedScrapQty"`
		MfgOrderItemGoodsReceiptQty    string  `json:"MfgOrderItemGoodsReceiptQty"`
		MfgOrderItemActualDeviationQty string  `json:"MfgOrderItemActualDeviationQty"`
		MfgOrderItemPlndDeliveryDate   string  `json:"MfgOrderItemPlndDeliveryDate"`
		MfgOrderItemActualDeliveryDate string  `json:"MfgOrderItemActualDeliveryDate"`
		IsCompletelyDelivered          bool    `json:"IsCompletelyDelivered"`
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
			RequiredQuantity              string      `json:"RequiredQuantity"`
			MaterialCompOriginalQuantity  string      `json:"MaterialCompOriginalQuantity"`
			Reservation                   string      `json:"Reservation"`
			ReservationItem               string      `json:"ReservationItem"`
			QuantityIsFixed               bool        `json:"QuantityIsFixed"`
			ComponentScrapInPercent       string      `json:"ComponentScrapInPercent"`
			WithdrawnQuantity             string      `json:"WithdrawnQuantity"`
			ConfirmedAvailableQuantity    string      `json:"ConfirmedAvailableQuantity"`
			WithdrawnQuantityAmount       string      `json:"WithdrawnQuantityAmount"`
			OrderInternalBillOfOperations string      `json:"OrderInternalBillOfOperations"`
			MatlCompRequirementDate       string      `json:"MatlCompRequirementDate"`
			ReservationIsFinallyIssued    bool        `json:"ReservationIsFinallyIssued"`
			MatlCompIsMarkedForBackflush  bool        `json:"MatlCompIsMarkedForBackflush"`
			StorageLocation               string      `json:"StorageLocation"`
			GoodsMovementType             string      `json:"GoodsMovementType"`
		} `json:"BillofMaterial"`
		BillofOperations struct {
			OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
			OrderIntBillOfOperationsItem  struct {
				OrderIntBillOfOperationsItem  string  `json:"OrderIntBillOfOperationsItem"`
				ManufacturingOrderSequence    string  `json:"ManufacturingOrderSequence"`
				MfgOrderOperationText         string  `json:"MfgOrderOperationText"`
				OperationIsReleased           string  `json:"OperationIsReleased"`
				OperationIsConfirmed          string  `json:"OperationIsConfirmed"`
				OperationIsPartiallyConfirmed string  `json:"OperationIsPartiallyConfirmed"`
				OperationIsDeleted            string  `json:"OperationIsDeleted"`
				OperationIsTechlyCompleted    string  `json:"OperationIsTechlyCompleted"`
				OperationIsClosed             string  `json:"OperationIsClosed"`
				ProductionPlant               string  `json:"ProductionPlant"`
				WorkCenterInternalID          string  `json:"WorkCenterInternalID"`
				OpErlstSchedldExecStrtDte     string  `json:"OpErlstSchedldExecStrtDte"`
				OpErlstSchedldExecStrtTme     string  `json:"OpErlstSchedldExecStrtTme"`
				OpErlstSchedldExecEndDte      string  `json:"OpErlstSchedldExecEndDte"`
				OpErlstSchedldExecEndTme      string  `json:"OpErlstSchedldExecEndTme"`
				OpActualExecutionStartDate    string  `json:"OpActualExecutionStartDate"`
				OpActualExecutionStartTime    string  `json:"OpActualExecutionStartTime"`
				OpActualExecutionEndDate      string  `json:"OpActualExecutionEndDate"`
				OpActualExecutionEndTime      string  `json:"OpActualExecutionEndTime"`
				OpActualExecutionDays         string  `json:"OpActualExecutionDays"`
				OperationUnit                 string  `json:"OperationUnit"`
				OpPlannedTotalQuantity        string  `json:"OpPlannedTotalQuantity"`
				OpTotalConfirmedYieldQty      string  `json:"OpTotalConfirmedYieldQty"`
				LastChangeDateTime            string  `json:"LastChangeDateTime"`
			} `json:"OrderIntBillOfOperationsItem"`
		} `json:"BillofOperations"`
	} `json:"ProductionOrder"`
	APISchema          string `json:"api_schema"`
	Accepter         []string `json:"accepter"`
	ProductionOrder    string `json:"production_order"`
	Deleted            bool   `json:"deleted"`
}
