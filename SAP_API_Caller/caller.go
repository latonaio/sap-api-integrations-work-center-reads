package sap_api_caller

import (
	"fmt"
	"io/ioutil"
	sap_api_output_formatter "sap-api-integrations-work-center-reads/SAP_API_Output_Formatter"
	"strings"
	"sync"

	sap_api_request_client_header_setup "github.com/latonaio/sap-api-request-client-header-setup"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
)

type SAPAPICaller struct {
	baseURL         string
	sapClientNumber string
	requestClient   *sap_api_request_client_header_setup.SAPRequestClient
	log             *logger.Logger
}

func NewSAPAPICaller(baseUrl, sapClientNumber string, requestClient *sap_api_request_client_header_setup.SAPRequestClient, l *logger.Logger) *SAPAPICaller {
	return &SAPAPICaller{
		baseURL:         baseUrl,
		requestClient:   requestClient,
		sapClientNumber: sapClientNumber,
		log:             l,
	}
}

func (c *SAPAPICaller) AsyncGetWorkCenter(workCenterInternalID, workCenterTypeCode string, accepter []string) {
	wg := &sync.WaitGroup{}
	wg.Add(len(accepter))
	for _, fn := range accepter {
		switch fn {
		case "WorkCenter":
			func() {
				c.WorkCenter(workCenterInternalID, workCenterTypeCode)
				wg.Done()
			}()

		default:
			wg.Done()
		}
	}

	wg.Wait()
}

func (c *SAPAPICaller) WorkCenter(workCenterInternalID, workCenterTypeCode string) {
	data, err := c.callWorkCenterSrvAPIRequirementWorkCenter("A_WorkCenter", workCenterInternalID, workCenterTypeCode)
	if err != nil {
		c.log.Error(err)
		return
	}
	c.log.Info(data)
}

func (c *SAPAPICaller) callWorkCenterSrvAPIRequirementWorkCenter(api, workCenterInternalID, workCenterTypeCode string) (*sap_api_output_formatter.WorkCenter, error) {
	url := strings.Join([]string{c.baseURL, "api_work_center/srvd_a2x/sap/workcenter/0001", api}, "/")
	param := c.getQueryWithWorkCenter(map[string]string{}, workCenterInternalID, workCenterTypeCode)

	resp, err := c.requestClient.Request("GET", url, param, "")
	if err != nil {
		return nil, fmt.Errorf("API request error: %w", err)
	}
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)
	data, err := sap_api_output_formatter.ConvertToWorkCenter(byteArray, c.log)
	if err != nil {
		return nil, fmt.Errorf("convert error: %w", err)
	}
	return data, nil
}

func (c *SAPAPICaller) getQueryWithWorkCenter(params map[string]string, workCenterInternalID, workCenterTypeCode string) map[string]string {
	if len(params) == 0 {
		params = make(map[string]string, 1)
	}
	params["$filter"] = fmt.Sprintf("WorkCenterInternalID eq '%s' and WorkCenterTypeCode eq '%s'", workCenterInternalID, workCenterTypeCode)
	return params
}
