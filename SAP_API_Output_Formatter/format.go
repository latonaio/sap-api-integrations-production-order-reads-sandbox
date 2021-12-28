package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-production-order-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

func ConvertToGeneral(raw []byte, l *logger.Logger) ([]General, error) {
	pm := &responses.General{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to General. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. show the first 10 of Results array", len(pm.D.Results))
	}

	general := make([]General, 0, 10)
	for i := 0; i < 10 && i < len(pm.D.Results); i++ {
		data := pm.D.Results[i]
		general = append(general, General{
			ManufacturingOrder:            data.ManufacturingOrder,
			ManufacturingOrderCategory:    data.ManufacturingOrderCategory,
			ManufacturingOrderType:        data.ManufacturingOrderType,
			OrderIsCreated:                data.OrderIsCreated,
			OrderIsReleased:               data.OrderIsReleased,
			OrderIsPrinted:                data.OrderIsPrinted,
			OrderIsConfirmed:              data.OrderIsConfirmed,
			OrderIsPartiallyConfirmed:     data.OrderIsPartiallyConfirmed,
			OrderIsDelivered:              data.OrderIsDelivered,
			OrderIsDeleted:                data.OrderIsDeleted,
			OrderIsPreCosted:              data.OrderIsPreCosted,
			SettlementRuleIsCreated:       data.SettlementRuleIsCreated,
			OrderIsPartiallyReleased:      data.OrderIsPartiallyReleased,
			OrderIsLocked:                 data.OrderIsLocked,
			OrderIsTechnicallyCompleted:   data.OrderIsTechnicallyCompleted,
			OrderIsClosed:                 data.OrderIsClosed,
			OrderIsPartiallyDelivered:     data.OrderIsPartiallyDelivered,
			OrderIsMarkedForDeletion:      data.OrderIsMarkedForDeletion,
			OrderIsScheduled:              data.OrderIsScheduled,
			OrderHasGeneratedOperations:   data.OrderHasGeneratedOperations,
			MaterialAvailyIsNotChecked:    data.MaterialAvailyIsNotChecked,
			MfgOrderCreationDate:          data.MfgOrderCreationDate,
			MfgOrderCreationTime:          data.MfgOrderCreationTime,
			LastChangeDateTime:            data.LastChangeDateTime,
			Material:                      data.Material,
			StorageLocation:               data.StorageLocation,
			GoodsRecipientName:            data.GoodsRecipientName,
			UnloadingPointName:            data.UnloadingPointName,
			MaterialGoodsReceiptDuration:  data.MaterialGoodsReceiptDuration,
			OrderInternalBillOfOperations: data.OrderInternalBillOfOperations,
			ProductionPlant:               data.ProductionPlant,
			Plant:                         data.Plant,
			MRPArea:                       data.MRPArea,
			MRPController:                 data.MRPController,
			ProductionSupervisor:          data.ProductionSupervisor,
			ProductionVersion:             data.ProductionVersion,
			PlannedOrder:                  data.PlannedOrder,
			SalesOrder:                    data.SalesOrder,
			SalesOrderItem:                data.SalesOrderItem,
			BasicSchedulingType:           data.BasicSchedulingType,
			BusinessArea:                  data.BusinessArea,
			CompanyCode:                   data.CompanyCode,
			ProfitCenter:                  data.ProfitCenter,
			FunctionalArea:                data.FunctionalArea,
			MfgOrderPlannedStartDate:      data.MfgOrderPlannedStartDate,
			MfgOrderPlannedStartTime:      data.MfgOrderPlannedStartTime,
			MfgOrderPlannedEndDate:        data.MfgOrderPlannedEndDate,
			MfgOrderPlannedEndTime:        data.MfgOrderPlannedEndTime,
			MfgOrderScheduledStartDate:    data.MfgOrderScheduledStartDate,
			MfgOrderScheduledStartTime:    data.MfgOrderScheduledStartTime,
			MfgOrderScheduledEndDate:      data.MfgOrderScheduledEndDate,
			MfgOrderScheduledEndTime:      data.MfgOrderScheduledEndTime,
			MfgOrderActualReleaseDate:     data.MfgOrderActualReleaseDate,
			ProductionUnit:                data.ProductionUnit,
			TotalQuantity:                 data.TotalQuantity,
			MfgOrderPlannedScrapQty:       data.MfgOrderPlannedScrapQty,
			MfgOrderConfirmedYieldQty:     data.MfgOrderConfirmedYieldQty,
			WBSElementExternalID:          data.WBSElementExternalID,
			OrderLongText:                 data.OrderLongText,
			ToProductionOrderComponent:    data.ToProductionOrderComponent.Deferred.URI,
			ToProductionOrderItem:         data.ToProductionOrderItem.Deferred.URI,
			ToProductionOrderOperation:    data.ToProductionOrderOperation.Deferred.URI,
			ToProductionOrderStatus:       data.ToProductionOrderStatus.Deferred.URI,
		})
	}

	return general, nil
}

func ConvertToToProductionOrderComponent(raw []byte, l *logger.Logger) (*ToProductionOrderComponent, error) {
	pm := &responses.ToProductionOrderComponent{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductionOrderComponent. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToProductionOrderComponent{
			Reservation:                    data.Reservation,
			ReservationItem:                data.ReservationItem,
			MaterialGroup:                  data.MaterialGroup,
			Material:                       data.Material,
			Plant:                          data.Plant,
			ManufacturingOrderCategory:     data.ManufacturingOrderCategory,
			ManufacturingOrderType:         data.ManufacturingOrderType,
			ManufacturingOrder:             data.ManufacturingOrder,
			ManufacturingOrderSequence:     data.ManufacturingOrderSequence,
			ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
			ProductionPlant:                data.ProductionPlant,
			OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
			MatlCompRequirementDate:        data.MatlCompRequirementDate,
			MatlCompRequirementTime:        data.MatlCompRequirementTime,
			ReservationIsFinallyIssued:     data.ReservationIsFinallyIssued,
			IsBulkMaterialComponent:        data.IsBulkMaterialComponent,
			MatlCompIsMarkedForBackflush:   data.MatlCompIsMarkedForBackflush,
			MaterialCompIsCostRelevant:     data.MaterialCompIsCostRelevant,
			SalesOrder:                     data.SalesOrder,
			SalesOrderItem:                 data.SalesOrderItem,
			SortField:                      data.SortField,
			BillOfMaterialCategory:         data.BillOfMaterialCategory,
			BOMItem:                        data.BOMItem,
			BOMItemCategory:                data.BOMItemCategory,
			BillOfMaterialItemNumber:       data.BillOfMaterialItemNumber,
			BOMItemDescription:             data.BOMItemDescription,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			BatchSplitType:                 data.BatchSplitType,
			GoodsMovementType:              data.GoodsMovementType,
			SupplyArea:                     data.SupplyArea,
			GoodsRecipientName:             data.GoodsRecipientName,
			UnloadingPointName:             data.UnloadingPointName,
			MaterialCompIsAlternativeItem:  data.MaterialCompIsAlternativeItem,
			AlternativeItemGroup:           data.AlternativeItemGroup,
			AlternativeItemStrategy:        data.AlternativeItemStrategy,
			AlternativeItemPriority:        data.AlternativeItemPriority,
			UsageProbabilityPercent:        data.UsageProbabilityPercent,
			MaterialComponentIsPhantomItem: data.MaterialComponentIsPhantomItem,
			LeadTimeOffset:                 data.LeadTimeOffset,
			QuantityIsFixed:                data.QuantityIsFixed,
			IsNetScrap:                     data.IsNetScrap,
			ComponentScrapInPercent:        data.ComponentScrapInPercent,
			OperationScrapInPercent:        data.OperationScrapInPercent,
			BaseUnit:                       data.BaseUnit,
			RequiredQuantity:               data.RequiredQuantity,
			WithdrawnQuantity:              data.WithdrawnQuantity,
			ConfirmedAvailableQuantity:     data.ConfirmedAvailableQuantity,
			MaterialCompOriginalQuantity:   data.MaterialCompOriginalQuantity,
			Currency:                       data.Currency,
			WithdrawnQuantityAmount:        data.WithdrawnQuantityAmount,
	}, nil
}

func ConvertToToProductionOrderItem(raw []byte, l *logger.Logger) (*ToProductionOrderItem, error) {
	pm := &responses.ToProductionOrderItem{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductionOrderItem. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToProductionOrderItem{
			ManufacturingOrder:             data.ManufacturingOrder,
			ManufacturingOrderItem:         data.ManufacturingOrderItem,
			ManufacturingOrderCategory:     data.ManufacturingOrderCategory,
			ManufacturingOrderType:         data.ManufacturingOrderType,
			IsCompletelyDelivered:          data.IsCompletelyDelivered,
			Material:                       data.Material,
			ProductionPlant:                data.ProductionPlant,
			Plant:                          data.Plant,
			MRPArea:                        data.MRPArea,
			QuantityDistributionKey:        data.QuantityDistributionKey,
			MaterialGoodsReceiptDuration:   data.MaterialGoodsReceiptDuration,
			StorageLocation:                data.StorageLocation,
			Batch:                          data.Batch,
			InventoryUsabilityCode:         data.InventoryUsabilityCode,
			GoodsRecipientName:             data.GoodsRecipientName,
			UnloadingPointName:             data.UnloadingPointName,
			MfgOrderItemPlndDeliveryDate:   data.MfgOrderItemPlndDeliveryDate,
			MfgOrderItemActualDeliveryDate: data.MfgOrderItemActualDeliveryDate,
			ProductionUnit:                 data.ProductionUnit,
			MfgOrderItemPlannedTotalQty:    data.MfgOrderItemPlannedTotalQty,
			MfgOrderItemPlannedScrapQty:    data.MfgOrderItemPlannedScrapQty,
			MfgOrderItemGoodsReceiptQty:    data.MfgOrderItemGoodsReceiptQty,
			MfgOrderItemActualDeviationQty: data.MfgOrderItemActualDeviationQty,
	}, nil
}

func ConvertToToProductionOrderOperation(raw []byte, l *logger.Logger) (*ToProductionOrderOperation, error) {
	pm := &responses.ToProductionOrderOperation{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductionOrderOperation. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToProductionOrderOperation{
			OrderInternalBillOfOperations:  data.OrderInternalBillOfOperations,
			OrderIntBillOfOperationsItem:   data.OrderIntBillOfOperationsItem,
			ManufacturingOrder:             data.ManufacturingOrder,
			ManufacturingOrderSequence:     data.ManufacturingOrderSequence,
			MfgOrderSequenceText:           data.MfgOrderSequenceText,
			ManufacturingOrderOperation:    data.ManufacturingOrderOperation,
			ManufacturingOrderSubOperation: data.ManufacturingOrderSubOperation,
			ManufacturingOrderCategory:     data.ManufacturingOrderCategory,
			ManufacturingOrderType:         data.ManufacturingOrderType,
			MfgOrderOperationText:          data.MfgOrderOperationText,
			OperationIsCreated:             data.OperationIsCreated,
			OperationIsReleased:            data.OperationIsReleased,
			OperationIsPrinted:             data.OperationIsPrinted,
			OperationIsConfirmed:           data.OperationIsConfirmed,
			OperationIsPartiallyConfirmed:  data.OperationIsPartiallyConfirmed,
			OperationIsDeleted:             data.OperationIsDeleted,
			OperationIsTechlyCompleted:     data.OperationIsTechlyCompleted,
			OperationIsClosed:              data.OperationIsClosed,
			OperationIsScheduled:           data.OperationIsScheduled,
			OperationIsPartiallyDelivered:  data.OperationIsPartiallyDelivered,
			OperationIsDelivered:           data.OperationIsDelivered,
			ProductionPlant:                data.ProductionPlant,
			WorkCenterInternalID:           data.WorkCenterInternalID,
			WorkCenterTypeCode:             data.WorkCenterTypeCode,
			WorkCenter:                     data.WorkCenter,
			OperationControlProfile:        data.OperationControlProfile,
			OpErlstSchedldExecStrtDte:      data.OpErlstSchedldExecStrtDte,
			OpErlstSchedldExecStrtTme:      data.OpErlstSchedldExecStrtTme,
			OpErlstSchedldExecEndDte:       data.OpErlstSchedldExecEndDte,
			OpErlstSchedldExecEndTme:       data.OpErlstSchedldExecEndTme,
			OpActualExecutionStartDate:     data.OpActualExecutionStartDate,
			OpActualExecutionStartTime:     data.OpActualExecutionStartTime,
			OpActualExecutionEndDate:       data.OpActualExecutionEndDate,
			OpActualExecutionEndTime:       data.OpActualExecutionEndTime,
			ErlstSchedldExecDurnInWorkdays: data.ErlstSchedldExecDurnInWorkdays,
			OpActualExecutionDays:          data.OpActualExecutionDays,
			OperationUnit:                  data.OperationUnit,
			OpPlannedTotalQuantity:         data.OpPlannedTotalQuantity,
			OpTotalConfirmedYieldQty:       data.OpTotalConfirmedYieldQty,
			LastChangeDateTime:             data.LastChangeDateTime,
	}, nil
}

func ConvertToToProductionOrderStatus(raw []byte, l *logger.Logger) (*ToProductionOrderStatus, error) {
	pm := &responses.ToProductionOrderStatus{}

	err := json.Unmarshal(raw, pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to ToProductionOrderStatus. unmarshal error: %w", err)
	}
	if len(pm.D.Results) == 0 {
		return nil, xerrors.New("Result data is not exist")
	}
	if len(pm.D.Results) > 10 {
		l.Info("raw data has too many Results. %d Results exist. expected only 1 Result. Use the first of Results array", len(pm.D.Results))
	}
	data := pm.D.Results[0]

	return &ToProductionOrderStatus{
			ManufacturingOrder: data.ManufacturingOrder,
			StatusCode:         data.StatusCode,
			IsUserStatus:       data.IsUserStatus,
			StatusShortName:    data.StatusShortName,
			StatusName:         data.StatusName,
	}, nil
}
