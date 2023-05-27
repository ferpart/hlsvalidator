package domain

import (
	"bytes"
	"errors"
	"os"
)

const (
	OutputDir    = "output"
	ValidatedDir = OutputDir + "/validated"
	ReportsDir   = OutputDir + "/reports"

	ValidatedPath = ValidatedDir + "/%s.json"
	ReportsPath   = ReportsDir + "/%s.html"
)

func SetupOutputDirs() error {
	if err := os.RemoveAll(OutputDir); err != nil {
		return err
	}

	if err := os.MkdirAll(ValidatedDir, 0755); err != nil {
		return err
	}

	if err := os.MkdirAll(ReportsDir, 0755); err != nil {
		return err
	}

	return nil
}

func StdErrToErr(err error, stdErr *bytes.Buffer) error {
	return errors.New(err.Error() + ": " + stdErr.String())
}
