package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
)

func ConvertToProductionOrder(raw []byte, l *logger.Logger) *ProductionOrder {
	pm := &responses.ProductionOrder{}
	err := json.Unmarshal(raw, pm)
	if err != nil {
		l.Error(err)
		return nil
	}
	if len(pm.D.Results) == 0 {
		l.Error("Result data is not exist.")
		return nil
	}
	if len(pm.D.Results) > 1 {
		l.Error("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ProductionOrder{
		ManufacturingOrder             data.ManufacturingOrder,
		ManufacturingOrderType         data.ManufacturingOrderType,
		Material                       data.Material,
		ProductionPlant                data.ProductionPlant,
		MfgOrderItemPlannedTotalQty    data.MfgOrderItemPlannedTotalQty,
		MfgOrderItemPlannedScrapQty    data.MfgOrderItemPlannedScrapQty,
		MfgOrderItemGoodsReceiptQty    data.MfgOrderItemGoodsReceiptQty,
		MfgOrderItemActualDeviationQty data.MfgOrderItemActualDeviationQty,
		MfgOrderItemPlndDeliveryDate   data.MfgOrderItemPlndDeliveryDate,
		MfgOrderItemActualDeliveryDate data.MfgOrderItemActualDeliveryDate,
		IsCompletelyDelivered          data.IsCompletelyDelivered,
		MRPArea                        data.MRPArea,
		StorageLocation                data.StorageLocation,
		Batch                          data.Batch,
		ProductionUnit                 data.ProductionUnit,
		StatusCode                     data.StatusCode,
		IsUserStatus                   data.IsUserStatus,
		StatusShortName                data.StatusShortName,
		StatusName                     data.StatusName,
		BillOfMaterialItemNumber       data.BillOfMaterialItemNumber,
		Component                      data.Component,
		Plant                          data.Plant,
		BaseUnit                       data.BaseUnit,
		RequiredQuantity               data.RequiredQuantity,
		MaterialCompOriginalQuantity   data.MaterialCompOriginalQuantity,
		Reservation                    data.Reservation,
		ReservationItem                data.ReservationItem,
		QuantityIsFixed                data.QuantityIsFixed,
		ComponentScrapInPercent        data.ComponentScrapInPercent,
		WithdrawnQuantity              data.WithdrawnQuantity,
		ConfirmedAvailableQuantity     data.ConfirmedAvailableQuantity,
		WithdrawnQuantityAmount        data.WithdrawnQuantityAmount,
		OrderInternalBillOfOperations  data.OrderInternalBillOfOperations,
		MatlCompRequirementDate        data.MatlCompRequirementDate,
		ReservationIsFinallyIssued     data.ReservationIsFinallyIssued,
		MatlCompIsMarkedForBackflush   data.MatlCompIsMarkedForBackflush,
		StorageLocation                data.StorageLocation,
		GoodsMovementType              data.GoodsMovementType,
		OrderInternalBillOfOperations  data.OrderInternalBillOfOperations,
		OrderIntBillOfOperationsItem   data.OrderIntBillOfOperationsItem,
		ManufacturingOrderSequence     data.ManufacturingOrderSequence,
		MfgOrderOperationText          data.MfgOrderOperationText,
		OperationIsReleased            data.OperationIsReleased,
		OperationIsConfirmed           data.OperationIsConfirmed,
		OperationIsPartiallyConfirmed  data.OperationIsPartiallyConfirmed,
		OperationIsDeleted             data.OperationIsDeleted,
		OperationIsTechlyCompleted     data.OperationIsTechlyCompleted,
		OperationIsClosed              data.OperationIsClosed,
		ProductionPlant                data.ProductionPlant,
		WorkCenterInternalID           data.WorkCenterInternalID,
		OpErlstSchedldExecStrtDte      data.OpErlstSchedldExecStrtDte,
		OpErlstSchedldExecStrtTme      data.OpErlstSchedldExecStrtTme,
		OpErlstSchedldExecEndDte       data.OpErlstSchedldExecEndDte,
		OpErlstSchedldExecEndTme       data.OpErlstSchedldExecEndTme,
		OpActualExecutionStartDate     data.OpActualExecutionStartDate,
		OpActualExecutionStartTime     data.OpActualExecutionStartTime,
		OpActualExecutionEndDate       data.OpActualExecutionEndDate,
		OpActualExecutionEndTime       data.OpActualExecutionEndTime,
		OpActualExecutionDays          data.OpActualExecutionDays,
		OperationUnit                  data.OperationUnit,
		OpPlannedTotalQuantity         data.OpPlannedTotalQuantity,
		OpTotalConfirmedYieldQty       data.OpTotalConfirmedYieldQty,
		LastChangeDateTime             data.LastChangeDateTime,
	}
}
