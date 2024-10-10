
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_int64_phone_number.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateInt64PhoneNumber struct{}

type GenerateInt64PhoneNumberOpts struct {
	randomizer     rng.Rand
	
}

func NewGenerateInt64PhoneNumber() *GenerateInt64PhoneNumber {
	return &GenerateInt64PhoneNumber{}
}

func NewGenerateInt64PhoneNumberOpts(
  seedArg *int64,
) (*GenerateInt64PhoneNumberOpts, error) {
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateInt64PhoneNumberOpts{
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateInt64PhoneNumber) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateInt64PhoneNumber",
		Description: "Generates a new int64 phone number with a default length of 10.",
		Example: "",
	}, nil
}

func (t *GenerateInt64PhoneNumber) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateInt64PhoneNumberOpts{}

	var seedArg *int64
	if seedValue, ok := opts["seed"].(int64); ok {
			seedArg = &seedValue
	}
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
	if err != nil {
		return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	transformerOpts.randomizer = rng.New(seed)

	return transformerOpts, nil
}
