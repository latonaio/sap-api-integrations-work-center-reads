package sap_api_output_formatter

import (
	"encoding/json"
	"sap-api-integrations-work-center-reads/SAP_API_Caller/responses"

	"github.com/latonaio/golang-logging-library-for-sap/logger"
	"golang.org/x/xerrors"
)

func ConvertToWorkCenter(raw []byte, l *logger.Logger) (*WorkCenter, error) {
	pm := &responses.WorkCenter{}
	err := json.Unmarshal(raw, &pm)
	if err != nil {
		return nil, xerrors.Errorf("cannot convert to WorkCenter. unmarshal error: %w", err)
	}

	return &WorkCenter{
		WorkCenterInternalID:         pm.WorkCenterInternalID,
		WorkCenterTypeCode:           pm.WorkCenterTypeCode,
		WorkCenter:                   pm.WorkCenter,
		WorkCenterDesc:               pm.WorkCenterDesc,
		Plant:                        pm.Plant,
		WorkCenterCategoryCode:       pm.WorkCenterCategoryCode,
		WorkCenterResponsible:        pm.WorkCenterResponsible,
		SupplyArea:                   pm.SupplyArea,
		WorkCenterUsage:              pm.WorkCenterUsage,
		MatlCompIsMarkedForBackflush: pm.MatlCompIsMarkedForBackflush,
		WorkCenterLocation:           pm.WorkCenterLocation,
		CapacityInternalID:           pm.CapacityInternalID,
		CapacityCategoryCode:         pm.CapacityCategoryCode,
		ValidityStartDate:            pm.ValidityStartDate,
		ValidityEndDate:              pm.ValidityEndDate,
		WorkCenterIsToBeDeleted:      pm.WorkCenterIsToBeDeleted,
	}, nil
}
