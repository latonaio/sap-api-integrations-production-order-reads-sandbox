package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/latonaio/golang-logging-library/logger"
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


func (c *SAPAPICaller) AsyncGetProductionOrder(ManufacturingOrder, StatusCode string) {
	wg := &sync.WaitGroup{}

	wg.Add(5)
	go func() {
		c.Header(ManufacturingOrder)
		wg.Done()
	}()
	go func() {
		c.Status(ManufacturingOrder, StatusCode)
		wg.Done()
	}()
	go func() {
		c.Items(ManufacturingOrder)
		wg.Done()
	}()
	go func() {
		c.Components(ManufacturingOrder)
		wg.Done()
	}()
	go func() {
		c.Operations(ManufacturingOrder)
		wg.Done()
	}()
	wg.Wait()
}

func (c *SAPAPICaller) Header(ManufacturingOrder string) {
	res, err := c.callProductionOrderSrvAPIRequirementHeader("A_ProductionOrder_2", ManufacturingOrder)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Status(ManufacturingOrder, StatusCode string) {
	res, err := c.callProductionOrderSrvAPIRequirementStatus("A_ProductionOrderStatus_2", ManufacturingOrder, StatusCode)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Items(ManufacturingOrder string) {
	res, err := c.callProductionOrderSrvAPIRequirementItems("A_ProductionOrderItem_2(ManufacturingOrder='{ManufacturingOrder}',ManufacturingOrderItem='{ManufacturingOrderItem}')", ManufacturingOrder)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Components(ManufacturingOrder string) {
	res, err := c.callProductionOrderSrvAPIRequirementComponents("A_ProductionOrderComponent_2", ManufacturingOrder)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) Operations(ManufacturingOrder string) {
	res, err := c.callProductionOrderSrvAPIRequirementOperations("A_ProductionOrderOperation_2", ManufacturingOrder)
	if err != nil {
		c.log.Error(err)
		return
	}

	c.log.Info(res)

}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementHeader(api, ManufacturingOrder string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "ManufacturingOrder")
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s'", ManufacturingOrder))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementStatus(api, ManufacturingOrder, StatusCode string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "ManufacturingOrder, StatusCode")
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s' and StatusCode eq '%s'", ManufacturingOrder, StatusCode))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementItems(api, ManufacturingOrder string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "ManufacturingOrder")
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s'", ManufacturingOrder))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementComponents(api, ManufacturingOrder string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "ManufacturingOrder")
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s'", ManufacturingOrder))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}

func (c *SAPAPICaller) callProductionOrderSrvAPIRequirementOperations(api, ManufacturingOrder string) ([]byte, error) {
	url := strings.Join([]string{c.baseURL, "API_PRODUCTION_ORDER_2_SRV", api}, "/")
	req, _ := http.NewRequest("GET", url, nil)

	params := req.URL.Query()
	// params.Add("$select", "ManufacturingOrder")
	params.Add("$filter", fmt.Sprintf("ManufacturingOrder eq '%s'", ManufacturingOrder))
	req.URL.RawQuery = params.Encode()

	req.Header.Set("APIKey", c.apiKey)
	req.Header.Set("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	return byteArray, nil
}