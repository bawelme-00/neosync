
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_full_address.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateFullAddress struct{}

type GenerateFullAddressOpts struct {
	randomizer     rng.Rand
	
	maxLength int64
}

func NewGenerateFullAddress() *GenerateFullAddress {
	return &GenerateFullAddress{}
}

func NewGenerateFullAddressOpts(
	maxLength int64,
  seedArg *int64,
) (*GenerateFullAddressOpts, error) {
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateFullAddressOpts{
		maxLength: maxLength,
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateFullAddress) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateFullAddress",
		Description: "Generates a randomly selected real full address that exists in the United States.",
		Example: "",
	}, nil
}

func (t *GenerateFullAddress) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateFullAddressOpts{}

	if _, ok := opts["maxLength"].(int64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "generateFullAddress", "maxLength")
	}
	maxLength := opts["maxLength"].(int64)
	transformerOpts.maxLength = maxLength

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
