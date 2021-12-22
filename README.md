# sap-api-integrations-production-order-reads  
sap-api-integrations-production-order-reads は、外部システム(特にエッジコンピューティング環境)をSAPと統合することを目的に、SAP API で製造指図データを取得するマイクロサービスです。    
sap-api-integrations-production-order-reads には、サンプルのAPI Json フォーマットが含まれています。   
sap-api-integrations-production-order-reads は、オンプレミス版である（＝クラウド版ではない）SAPS4HANA API の利用を前提としています。クラウド版APIを利用する場合は、ご注意ください。   
https://api.sap.com/api/OP_API_PRODUCTION_ORDER_2_SRV_0001/overview    

## 動作環境  

sap-api-integrations-production-order-reads は、主にエッジコンピューティング環境における動作にフォーカスしています。  
使用する際は、事前に下記の通り エッジコンピューティングの動作環境（推奨/必須）を用意してください。  
・ エッジ Kubernetes （推奨）    
・ AION のリソース （推奨)    
・ OS: LinuxOS （必須）    
・ CPU: ARM/AMD/Intel（いずれか必須）    

## クラウド環境での利用

sap-api-integrations-production-order-reads は、外部システムがクラウド環境である場合にSAPと統合するときにおいても、利用可能なように設計されています。  

## 本レポジトリ が 対応する API サービス
sap-api-integrations-production-order-reads が対応する APIサービス は、次のものです。

* APIサービス概要説明 URL: https://api.sap.com/api/OP_API_PRODUCTION_ORDER_2_SRV_0001/overview  
* APIサービス名(=baseURL): API_PRODUCTION_ORDER_2_SRV

## 本レポジトリ に 含まれる API名
sap-api-integrations-production-order-reads には、次の API をコールするためのリソースが含まれています。  

* A_ProductionOrder_2（製造指図 - 一般）※製造指図の詳細データを取得するために、ToProductionOrderComponent、ToProductionOrderItem、ToProductionOrderOperation、ToProductionOrderStatus、と合わせて利用されます。
* ToProductionOrderComponent（製造指図 - 構成品目）
* ToProductionOrderItem（製造指図 - 明細）
* ToProductionOrderOperation（製造指図 - 作業手順）
* ToProductionOrderStatus（製造指図 - ステータス）

## API への 値入力条件 の 初期値
sap-api-integrations-production-order-reads において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

### SDC レイアウト

* inoutSDC.ManufacturingOrder（製造指図）

## SAP API Bussiness Hub の API の選択的コール

Latona および AION の SAP 関連リソースでは、Inputs フォルダ下の sample.json の accepter に取得したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて取得することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "sap.s4.beh.productionorder.v1.ProductionOrder.Created.v1",
	"accepter": ["General"],
	"production_order": "1000000",
	"deleted": false
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "sap.s4.beh.productionorder.v1.ProductionOrder.Created.v1",
	"accepter": ["All"],
	"production_order": "1000000",
	"deleted": false
```



## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて SAP_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *SAPAPICaller) AsyncGetProductionOrder(manufacturingOrder string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "General":
			func() {
				c.General(manufacturingOrder)
				wg.Done()
			}()
		default:
			wg.Done()
		}
	}

	wg.Wait()
}
```

## Output  
本マイクロサービスでは、[golang-logging-library](https://github.com/latonaio/golang-logging-library) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は、SAP 製造指図 の 一般データ が取得された結果の JSON の例です。  
以下の項目のうち、"ManufacturingOrder" ～ "to_ProductionOrderStatus" は、/SAP_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
{
	"cursor": "/Users/latona2/bitbucket/sap-api-integrations-production-order-reads/SAP_API_Caller/caller.go#L46",
	"function": "sap-api-integrations-production-order-reads/SAP_API_Caller.(*SAPAPICaller).General",
	"level": "INFO",
	"message": [
		{
			"ManufacturingOrder": "1000000",
			"ManufacturingOrderCategory": "10",
			"ManufacturingOrderType": "YBM1",
			"OrderIsCreated": "",
			"OrderIsReleased": "",
			"OrderIsPrinted": "",
			"OrderIsConfirmed": "",
			"OrderIsPartiallyConfirmed": "",
			"OrderIsDelivered": "",
			"OrderIsDeleted": "",
			"OrderIsPreCosted": "X",
			"SettlementRuleIsCreated": "X",
			"OrderIsPartiallyReleased": "",
			"OrderIsLocked": "",
			"OrderIsTechnicallyCompleted": "X",
			"OrderIsClosed": "",
			"OrderIsPartiallyDelivered": "",
			"OrderIsMarkedForDeletion": "",
			"OrderIsScheduled": "",
			"OrderHasGeneratedOperations": "",
			"MaterialAvailyIsNotChecked": "",
			"MfgOrderCreationDate": "/Date(1470268800000)/",
			"MfgOrderCreationTime": "PT12H26M00S",
			"LastChangeDateTime": "20200723125718",
			"Material": "MZ-FG-R10",
			"StorageLocation": "171A",
			"GoodsRecipientName": "",
			"UnloadingPointName": "",
			"MaterialGoodsReceiptDuration": "0",
			"OrderInternalBillOfOperations": "1",
			"ProductionPlant": "1710",
			"Plant": "1710",
			"MRPArea": "1710",
			"MRPController": "MZ1",
			"ProductionSupervisor": "",
			"ProductionVersion": "0001",
			"PlannedOrder": "",
			"SalesOrder": "",
			"SalesOrderItem": "0",
			"BasicSchedulingType": "1",
			"BusinessArea": "",
			"CompanyCode": "1710",
			"ProfitCenter": "Z_LOG_PC7",
			"FunctionalArea": "YB20",
			"MfgOrderPlannedStartDate": "/Date(1496707200000)/",
			"MfgOrderPlannedStartTime": "PT00H00M00S",
			"MfgOrderPlannedEndDate": "/Date(1496880000000)/",
			"MfgOrderPlannedEndTime": "PT00H00M00S",
			"MfgOrderScheduledStartDate": "/Date(1496707200000)/",
			"MfgOrderScheduledStartTime": "PT07H00M00S",
			"MfgOrderScheduledEndDate": "/Date(1496793600000)/",
			"MfgOrderScheduledEndTime": "PT10H22M30S",
			"MfgOrderActualReleaseDate": "/Date(1470614400000)/",
			"ProductionUnit": "PC",
			"TotalQuantity": "500",
			"MfgOrderPlannedScrapQty": "0",
			"MfgOrderConfirmedYieldQty": "0",
			"WBSElementExternalID": "",
			"OrderLongText": "",
			"to_ProductionOrderComponent": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ORDER_2_SRV/A_ProductionOrder_2('1000000')/to_ProductionOrderComponent",
			"to_ProductionOrderItem": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ORDER_2_SRV/A_ProductionOrder_2('1000000')/to_ProductionOrderItem",
			"to_ProductionOrderOperation": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ORDER_2_SRV/A_ProductionOrder_2('1000000')/to_ProductionOrderOperation",
			"to_ProductionOrderStatus": "https://sandbox.api.sap.com/s4hanacloud/sap/opu/odata/sap/API_PRODUCTION_ORDER_2_SRV/A_ProductionOrder_2('1000000')/to_ProductionOrderStatus"
		}
	],
	"time": "2021-12-10T11:28:38.6947+09:00"
}
```
