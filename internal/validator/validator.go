package validator

import (
	"log"
	"sync"

	"github.com/ferpart/hlsvalidator/domain"
)

type Validator struct {
	uri               string
	hlsValidator      domain.IHLSValidator
	waitGroup         *sync.WaitGroup
	safeValidationMap *domain.SafeValidationMap
}

func New(
	uri string,
	hlsValidator domain.IHLSValidator,
	waitGroup *sync.WaitGroup,
	safeValidationMap *domain.SafeValidationMap,
) *Validator {
	return &Validator{
		uri:               uri,
		hlsValidator:      hlsValidator,
		waitGroup:         waitGroup,
		safeValidationMap: safeValidationMap,
	}
}

func (v *Validator) Run() {
	uuid, isValid, err := v.hlsValidator.Validate(v.uri)
	if err != nil {
		log.Fatal(err)
	}

	v.safeValidationMap.Mu.Lock()
	v.safeValidationMap.ValidationMap[uuid] = isValid
	v.safeValidationMap.Mu.Unlock()

	v.waitGroup.Done()
}

func Validate(manifestURI string, hlsValidators []domain.IHLSValidator) map[string]bool {
	if err := domain.SetupOutputDirs(); err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	safeValidationMap := domain.SafeValidationMap{
		ValidationMap: make(map[string]bool),
	}

	for _, hlsValidator := range hlsValidators {
		v := New(manifestURI, hlsValidator, &wg, &safeValidationMap)
		wg.Add(1)
		go v.Run()
	}

	wg.Wait()

	return safeValidationMap.ValidationMap
}
