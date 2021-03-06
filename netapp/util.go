package netapp

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/candidpartners/occm-sdk-go/api/workenv"
)

// number of attempts to try to resolve a request
const RequestResolutionRetryCount = 60

// amount of time to wait between retrying request resolution
const RequestResolutionWaitTime = 2 * time.Second

// Volume identifier
type VolumeID struct {
	VolumeType string
	WorkEnvId  string
	SvmName    string
	VolumeName string
	IsHA       bool
}

func ParseVolumeID(id string) (*VolumeID, error) {
	parts := strings.Split(id, "/")
	if len(parts) != 4 {
		return nil, fmt.Errorf("Invalid volume ID format: %s", id)
	}

	volumeType := parts[0]
	workenvId := parts[1]
	svmName := parts[2]
	volumeName := parts[3]
	isHA := false
	if volumeType == "ha" {
		isHA = true
	}

	result := VolumeID{
		VolumeType: volumeType,
		WorkEnvId:  workenvId,
		SvmName:    svmName,
		VolumeName: volumeName,
		IsHA:       isHA,
	}

	return &result, nil
}

func GetWorkingEnvironments(apis *APIs) ([]workenv.VsaWorkingEnvironment, error) {
	resp, err := apis.WorkingEnvironmentAPI.GetWorkingEnvironments()
	if err != nil {
		return nil, fmt.Errorf("Error retrieving working environments: %s", err)
	}

	return resp.VSA, nil
}

func GetWorkingEnvironmentByName(apis *APIs, workEnvName string) (*workenv.VsaWorkingEnvironment, error) {
	workEnvs, err := GetWorkingEnvironments(apis)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Reading working environment %s", workEnvName)

	var found *workenv.VsaWorkingEnvironment

	for _, workenv := range workEnvs {
		if workenv.Name == workEnvName {
			found = &workenv
			break
		}
	}

	if found == nil {
		return nil, fmt.Errorf("Working environment %s not found", workEnvName)
	}

	log.Printf("[DEBUG] Found working environment %s", workEnvName)

	return found, nil
}

func GetWorkingEnvironmentById(apis *APIs, workEnvId string) (*workenv.VsaWorkingEnvironment, error) {
	workEnvs, err := GetWorkingEnvironments(apis)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Reading working environment for ID %s", workEnvId)

	var found *workenv.VsaWorkingEnvironment

	for _, workEnv := range workEnvs {
		if workEnv.PublicId == workEnvId {
			found = &workEnv
			break
		}
	}

	if found == nil {
		return nil, fmt.Errorf("Working environment with ID %s not found", workEnvId)
	}

	return found, nil
}

func WaitForRequest(apis *APIs, requestId string) error {
	log.Printf("[DEBUG] Waiting for completion of request %s", requestId)

	for i := 0; i < RequestResolutionRetryCount; i++ {
		summary, err := apis.AuditAPI.GetAuditSummary(requestId)
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] Received status for request %s: %s", requestId, summary.Status)

		if summary.Status == "Failed" {
			log.Printf("[DEBUG] Failure detected, breaking wait loop")
			return fmt.Errorf(summary.ErrorMessage)
		}

		if summary.Status == "Success" {
			log.Printf("[DEBUG] Request completion detected, breaking wait loop")
			return nil
		}

		time.Sleep(RequestResolutionWaitTime)
	}

	return fmt.Errorf("Timed out waiting for request completion")
}

func FormatString(format string, a ...interface{}) string {
	return fmt.Sprintf(format, a)
}
