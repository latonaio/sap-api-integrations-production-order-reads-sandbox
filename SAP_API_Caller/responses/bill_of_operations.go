package responses

type BillOfOperations struct {
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
		} `json:"results"`
	} `json:"d"`
}
