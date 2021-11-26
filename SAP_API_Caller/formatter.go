package sap_api_caller

type ManufacturingOrderReads struct {
	ConnectionKey      string `json:"connection_key"`
	Result             bool   `json:"result"`
	RedisKey           string `json:"redis_key"`
	Filepath           string `json:"filepath"`
	APISchema          string `json:"api_schema"`
	ManufacturingOrder string `json:"production_order"`
	Deleted            string `json:"deleted"`
}

type ManufacturingOrder struct {
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
}

type Status struct {
    ManufacturingOrder             string  `json:"ManufacturingOrder"`
    StatusCode                     string  `json:"StatusCode"`
    IsUserStatus                   string  `json:"IsUserStatus"`
    StatusShortName                string  `json:"StatusShortName"`
    StatusName                     string  `json:"StatusName"`
}

type BillofMaterial struct {
    ManufacturingOrder             string      `json:"ManufacturingOrder"`
    OrderInternalBillOfOperations  string      `json:"OrderInternalBillOfOperations"`  
    OrderIntBillOfOperationsItem   string      `json:"OrderIntBillOfOperationsItem"`
    BillOfMaterialItemNumber       string      `json:"BillOfMaterialItemNumber"`
    Component                      string      `json:"Component"`
    Plant                          string      `json:"Plant"`
    BaseUnit                       string      `json:"BaseUnit"`
    RequiredQuantity               float64     `json:"RequiredQuantity"`
    MaterialCompOriginalQuantity   float64     `json:"MaterialCompOriginalQuantity"`
    Reservation                    int         `json:"Reservation"`
    ReservationItem                int         `json:"ReservationItem"`
    QuantityIsFixed                string      `json:"QuantityIsFixed"`
    ComponentScrapInPercent        float64     `json:"ComponentScrapInPercent"`
    WithdrawnQuantity              float64     `json:"WithdrawnQuantity"`
    ConfirmedAvailableQuantity     float64     `json:"ConfirmedAvailableQuantity"`
    WithdrawnQuantityAmount        float64     `json:"WithdrawnQuantityAmount"`
    OrderInternalBillOfOperations  string      `json:"OrderInternalBillOfOperations"`
    MatlCompRequirementDate        string      `json:"MatlCompRequirementDate"`
    ReservationIsFinallyIssued     string      `json:"ReservationIsFinallyIssued"`
    MatlCompIsMarkedForBackflush   string      `json:"MatlCompIsMarkedForBackflush"`
    StorageLocation                string      `json:"StorageLocation"`
    GoodsMovementType              string      `json:"GoodsMovementType"`
}

type BillOfOperations struct {
    ManufacturingOrder            string  `json:"ManufacturingOrder"`
    OrderInternalBillOfOperations string  `json:"OrderInternalBillOfOperations"`
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
}
