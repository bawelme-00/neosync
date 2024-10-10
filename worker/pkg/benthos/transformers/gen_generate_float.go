
// Code generated by Neosync neosync_transformer_generator.go. DO NOT EDIT.
// source: generate_float.go

package transformers

import (
	"fmt"
	
	transformer_utils "github.com/nucleuscloud/neosync/worker/pkg/benthos/transformers/utils"
	"github.com/nucleuscloud/neosync/worker/pkg/rng"
	
)

type GenerateFloat64 struct{}

type GenerateFloat64Opts struct {
	randomizer     rng.Rand
	
	randomizeSign bool
	min float64
	max float64
	precision *int64
	scale *int64
}

func NewGenerateFloat64() *GenerateFloat64 {
	return &GenerateFloat64{}
}

func NewGenerateFloat64Opts(
	randomizeSignArg *bool,
	min float64,
	max float64,
	precision *int64,
	scale *int64,
  seedArg *int64,
) (*GenerateFloat64Opts, error) {
	randomizeSign := bool(false) 
	if randomizeSignArg != nil {
		randomizeSign = *randomizeSignArg
	}
	
	seed, err := transformer_utils.GetSeedOrDefault(seedArg)
  if err != nil {
    return nil, fmt.Errorf("unable to generate seed: %w", err)
	}
	
	return &GenerateFloat64Opts{
		randomizeSign: randomizeSign,
		min: min,
		max: max,
		precision: precision,
		scale: scale,
		randomizer: rng.New(seed),	
	}, nil
}

func (t *GenerateFloat64) GetJsTemplateData() (*TemplateData, error) {
	return &TemplateData{
		Name: "generateFloat64",
		Description: "Generates a random floating point number with a max precision of 17. Go float64 adheres to the IEEE 754 standard for double-precision floating-point numbers.",
		Example: "",
	}, nil
}

func (t *GenerateFloat64) ParseOptions(opts map[string]any) (any, error) {
	transformerOpts := &GenerateFloat64Opts{}

	randomizeSign, ok := opts["randomizeSign"].(bool)
	if !ok {
		randomizeSign = false
	}
	transformerOpts.randomizeSign = randomizeSign

	if _, ok := opts["min"].(float64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "generateFloat64", "min")
	}
	min := opts["min"].(float64)
	transformerOpts.min = min

	if _, ok := opts["max"].(float64); !ok {
		return nil, fmt.Errorf("missing required argument. function: %s argument: %s", "generateFloat64", "max")
	}
	max := opts["max"].(float64)
	transformerOpts.max = max

	var precision *int64
	if arg, ok := opts["precision"].(int64); ok {
		precision = &arg
	}
	transformerOpts.precision = precision

	var scale *int64
	if arg, ok := opts["scale"].(int64); ok {
		scale = &arg
	}
	transformerOpts.scale = scale

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
