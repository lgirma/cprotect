package cprotect

import "errors"

type FingerprintService interface {	
	GetMachineId(elevatedPrivilege bool) (string, error)
}

type PCFingerprintService struct {
}

func GetFingerprintService() FingerprintService {
	return &PCFingerprintService{}
}

func (service *PCFingerprintService) GetMachineId (elevatedPrivilege bool) (string, error) {
	errs := make([]error, 0)
	hddId, err := getDiskDriveId(elevatedPrivilege)
	if err != nil {
		errs = append(errs, err)
	}
	mbId, err := getMotherboardId(elevatedPrivilege)
	if err != nil {
		errs = append(errs, err)
	}

	hardwareId := hddId + mbId
	if len(hardwareId) == 0 {
		if len(errs) > 0 {
			return "", errors.New(ErrorHardwareIdExecutionFailure)
		}
		return "", errors.New(ErrorHardwareIdEmpty)
	}

	return hardwareId, nil
}
