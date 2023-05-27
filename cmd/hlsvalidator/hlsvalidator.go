package main

import (
	"fmt"

	flag "github.com/spf13/pflag"

	"github.com/ferpart/hlsvalidator/domain"
	"github.com/ferpart/hlsvalidator/internal/reporter"
	"github.com/ferpart/hlsvalidator/internal/validator"
	"github.com/ferpart/hlsvalidator/pkg/appleHLSReport"
	"github.com/ferpart/hlsvalidator/pkg/appleMediaStreamValidator"
)

var (
	manifestURI string
	times       int
	saveValid   bool
)

func init() {
	flag.StringVarP(
		&manifestURI,
		"manifest",
		"m",
		"https://example.com/playlist.m3u8",
		"manifest url to verify with HLS spec",
	)

	flag.IntVarP(
		&times,
		"times",
		"t",
		1,
		"number of times to validate url",
	)

	flag.BoolVarP(
		&saveValid,
		"save-valid",
		"s",
		false,
		"save reports for valid HLS manifests. validated jsons will always be generated",
	)
	flag.Lookup("save-valid").NoOptDefVal = "true"
}

func main() {
	flag.Parse()

	HLSValidators := generateValidators()
	validationMap := validator.Validate(manifestURI, HLSValidators)

	HLSReporters := generateReporters(validationMap)
	reporter.Report(HLSReporters)

	fmt.Println("HLS spec validation done.")
}

func generateValidators() []domain.IHLSValidator {
	var HLSValidators []domain.IHLSValidator
	for i := 0; i < times; i++ {
		HLSValidators = append(HLSValidators, appleMediaStreamValidator.New())
	}
	return HLSValidators
}

func generateReporters(validationMap map[string]bool) []domain.IHLSReporter {
	var HLSReporters []domain.IHLSReporter
	for uuid, isValid := range validationMap {
		if !saveValid && isValid {
			continue
		}

		HLSReporters = append(HLSReporters, appleHLSReport.New(uuid))
	}
	return HLSReporters
}
