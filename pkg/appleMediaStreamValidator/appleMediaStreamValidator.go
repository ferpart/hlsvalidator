package appleMediaStreamValidator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/google/uuid"

	"github.com/ferpart/hlsvalidator/domain"
)

const mediaStreamValidator = "mediastreamvalidator"

type AppleMediaStreamValidator struct {
	validator *exec.Cmd
}

func New(parsePlaylist bool) *AppleMediaStreamValidator {
	cmd := exec.Command(mediaStreamValidator)
	if parsePlaylist {
		cmd.Args = append(cmd.Args, "-p")
	}

	return &AppleMediaStreamValidator{
		validator: cmd,
	}
}

func (a *AppleMediaStreamValidator) Validate(uri string) (string, bool, error) {
	var stdErr bytes.Buffer

	validatedUUID := uuid.NewString()
	outputFilePath := fmt.Sprintf(domain.ValidatedPath, validatedUUID)

	a.validator.Args = append(a.validator.Args, []string{
		"-p",
		"-O", outputFilePath,
		uri,
	}...)

	a.validator.Stderr = &stdErr

	if err := a.validator.Run(); err != nil {
		return validatedUUID, false, domain.StdErrToErr(err, &stdErr)
	}

	validationDataBytes, err := os.ReadFile(outputFilePath)
	if err != nil {
		return validatedUUID, false, err
	}

	validationData := &domain.ValidationData{}

	if err = json.Unmarshal(validationDataBytes, validationData); err != nil {
		return validatedUUID, false, err
	}

	return validatedUUID, isValid(validationData), nil
}

func isValid(validationData *domain.ValidationData) bool {
	if len(validationData.Messages) == 0 {
		return true
	}

	for _, message := range validationData.Messages {
		if _, ok := domain.ValidErrors[message.ErrorStatusCode]; ok {
			continue
		}

		return false
	}

	return true
}
