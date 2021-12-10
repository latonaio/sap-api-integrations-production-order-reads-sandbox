package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	sap_api_output_formatter "sap-api-integrations-production-order-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
	"golang.org/x/xerrors"
)

type SAPAPICaller struct {
	baseURL string
	apiKey  string
	log     *logger.Logger
}

func NewSAPAPICaller(baseUrl string, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL: baseUrl,
		apiKey:  GetApiKey(),
		log:     l,
	}
}

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

func (c *SAPAPICaller) General(manufacturingOrder string) {
	generalData, err := c.callProductionOrderSrvAPIRequirementGeneral("A_ProductionOrder_2", manufacturingOrder)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(generalData)

	componentData, err := c.callToProductionOrderComponent(generalData[0].ToProductionOrderComponent)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(componentData)

	itemData, err := c.callToProductionOrderItem(generalData[0].ToProductionOrderItem)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(itemData)

	operationData, err := c.callToProductionOrderOperation(generalData[0].ToProductionOrderOperation)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(operationData)

	statusData, err := c.callToProductionOrderStatus(generalData[0].ToProductionOrderStatus)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(statusData)
}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementGeneral(api, manufacturingOrder string) ([]sap_api_output_formatter.General, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	c.setHeaderAPIKeyAccept(req)
	c.getQueryWithGeneral(req, manufacturingOrder)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToGeneral(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductionOrderComponent(url string) (*sap_api_output_formatter.ToProductionOrderComponent, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductionOrderComponent(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductionOrderItem(url string) (*sap_api_output_formatter.ToProductionOrderItem, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductionOrderItem(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductionOrderOperation(url string) (*sap_api_output_formatter.ToProductionOrderOperation, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductionOrderOperation(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) callToProductionOrderStatus(url string) (*sap_api_output_formatter.ToProductionOrderStatus, error) {
	req, _ := http.NewRequest("GET", url, nil)
	c.setHeaderAPIKeyAccept(req)

	resp, err := new(http.Client).Do(req)
	if err != nil {
		return nil, xerrors.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToToProductionOrderStatus(byteArray, c.log)
	if err != nil {
		return nil, xerrors.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) setHeaderAPIKeyAccept(req *http.Request) {
	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")
}

func (c *SAPAPICaller) getQueryWithGeneral(req *http.Request, manufacturingOrder string) {
	params := req.URL.Query()
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s'", manufacturingOrder))
	req.URL.RawQuery = params.Encode()
}
