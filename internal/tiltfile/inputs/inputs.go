package inputs

import (
	"fmt"

	"go.starlark.net/starlark"
	"go.starlark.net/starlarkstruct"

	"github.com/tilt-dev/tilt/pkg/model"

	"github.com/tilt-dev/tilt/internal/tiltfile/starkit"
	"github.com/tilt-dev/tilt/internal/tiltfile/value"
)

type Input struct {
	*starlarkstruct.Struct
	model.Input
}

var _ starlark.Value = Input{}

// Parse resource links (string or `link`) into model.Link
// Not to be confused with a LinkED List :P
type InputList struct {
	Inputs []model.Input
}

func (il *InputList) Unpack(v starlark.Value) error {
	seq := value.ValueOrSequenceToSlice(v)
	for _, val := range seq {
		switch val := val.(type) {
		case starlark.String:
			i, err := strToInput(val)
			if err != nil {
				return err
			}
			il.Inputs = append(il.Inputs, i)
		case Input:
			il.Inputs = append(il.Inputs, val.Input)
		default:
			return fmt.Errorf("`Want a string, an input, or a sequence of these; found %v (type: %T)", val, val)
		}
	}
	return nil
}

func strToInput(s starlark.String) (model.Input, error) {
	return model.Input{Key: string(s), Value: ""}, nil
}

// Implements functions for dealing with k8s secret settings.
type Plugin struct{}

var _ starkit.Plugin = Plugin{}

func NewPlugin() Plugin {
	return Plugin{}
}

func (e Plugin) OnStart(env *starkit.Environment) error {
	return env.AddBuiltin("input", e.input)
}

func (e Plugin) input(thread *starlark.Thread, fn *starlark.Builtin, args starlark.Tuple, kwargs []starlark.Tuple) (starlark.Value, error) {
	var key, value string

	if err := starkit.UnpackArgs(thread, fn.Name(), args, kwargs,
		"key", &key,
		"value", &value); err != nil {
		return nil, err
	}

	return Input{
		Struct: starlarkstruct.FromStringDict(starlark.String("input"), starlark.StringDict{
			"key":   starlark.String(key),
			"value": starlark.String(value),
		}),
		Input: model.Input{
			Key:   key,
			Value: value,
		},
	}, nil
}
