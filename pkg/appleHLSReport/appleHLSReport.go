package appleHLSReport

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"

	"github.com/ferpart/hlsvalidator/domain"
)

const hlsReport = "hlsreport"

type AppleHLSReport struct {
	validatedUUID string
	generator     *exec.Cmd
}

func New(validatedUUID string) *AppleHLSReport {
	return &AppleHLSReport{
		validatedUUID: validatedUUID,
		generator:     exec.Command(hlsReport),
	}
}

func (a *AppleHLSReport) Report() {
	var stdErr bytes.Buffer

	jsonFilePath := fmt.Sprintf(domain.ValidatedPath, a.validatedUUID)
	outputFilePath := fmt.Sprintf(domain.ReportsPath, a.validatedUUID)

	a.generator.Args = append(a.generator.Args, []string{
		"-o", outputFilePath,
		jsonFilePath,
	}...)

	a.generator.Stderr = &stdErr

	if err := a.generator.Run(); err != nil {
		log.Fatal(domain.StdErrToErr(err, &stdErr))
	}
}
