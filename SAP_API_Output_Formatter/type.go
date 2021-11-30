package sap_api_output_formatter

type ManufacturingOrderReads struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	ManufacturingOrder string `json:"production_order"`
	Deleted            bool   `json:"deleted"`
}

type ManufacturingOrder struct {
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
}

type Status struct {
	ManufacturingOrder string `json:"ManufacturingOrder"`
	StatusCode         string `json:"StatusCode"`
	IsUserStatus       string `json:"IsUserStatus"`
	StatusShortName    string `json:"StatusShortName"`
	StatusName         string `json:"StatusName"`
}

type BillofMaterial struct {
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
}

type BillOfOperations struct {
	ManufacturingOrder            string `json:"ManufacturingOrder"`
	OrderInternalBillOfOperations string `json:"OrderInternalBillOfOperations"`
	OrderIntBillOfOperationsItem  string `json:"OrderIntBillOfOperationsItem"`
	ManufacturingOrderSequence    string `json:"ManufacturingOrderSequence"`
	MfgOrderOperationText         string `json:"MfgOrderOperationText"`
	OperationIsReleased           bool   `json:"OperationIsReleased"`
	OperationIsConfirmed          bool   `json:"OperationIsConfirmed"`
	OperationIsPartiallyConfirmed bool   `json:"OperationIsPartiallyConfirmed"`
	OperationIsDeleted            bool   `json:"OperationIsDeleted"`
	OperationIsTechlyCompleted    bool   `json:"OperationIsTechlyCompleted"`
	OperationIsClosed             bool   `json:"OperationIsClosed"`
	ProductionPlant               string `json:"ProductionPlant"`
	WorkCenterInternalID          string `json:"WorkCenterInternalID"`
	OpErlstSchedldExecStrtDte     string `json:"OpErlstSchedldExecStrtDte"`
	OpErlstSchedldExecStrtTme     string `json:"OpErlstSchedldExecStrtTme"`
	OpErlstSchedldExecEndDte      string `json:"OpErlstSchedldExecEndDte"`
	OpErlstSchedldExecEndTme      string `json:"OpErlstSchedldExecEndTme"`
	OpActualExecutionStartDate    string `json:"OpActualExecutionStartDate"`
	OpActualExecutionStartTime    string `json:"OpActualExecutionStartTime"`
	OpActualExecutionEndDate      string `json:"OpActualExecutionEndDate"`
	OpActualExecutionEndTime      string `json:"OpActualExecutionEndTime"`
	OpActualExecutionDays         string `json:"OpActualExecutionDays"`
	OperationUnit                 string `json:"OperationUnit"`
	OpPlannedTotalQuantity        string `json:"OpPlannedTotalQuantity"`
	OpTotalConfirmedYieldQty      string `json:"OpTotalConfirmedYieldQty"`
	LastChangeDateTime            string `json:"LastChangeDateTime"`
}
