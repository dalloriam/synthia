package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/dalloriam/synthia/core"
	"github.com/dalloriam/synthia/modular"
)

type ModuleDef map[string]interface{}

type MixerChanDef struct {
	Input string `json:"input"`
}

type PatchDefinition struct {
	Modules       map[string]ModuleDef `json:"modules"`
	MixerChannels []MixerChanDef       `json:"mixer"`
}

func LoadPatchFile(patchFile string) (PatchDefinition, error) {
	data, err := ioutil.ReadFile(patchFile)
	if err != nil {
		return PatchDefinition{}, err
	}

	var patch PatchDefinition
	if err := json.Unmarshal(data, &patch); err != nil {
		return PatchDefinition{}, err
	}

	fmt.Println("Loaded patch: \n\t", patchFile)
	return patch, nil
}

func buildModule(def ModuleDef) (interface{}, error) {
	moduleType, ok := def["_type"].(string)
	if !ok {
		return nil, errors.New("no type defined for module")
	}
	fmt.Print("\t- Building [", moduleType, "]... ")

	var m interface{}
	switch moduleType {
	case "Oscillator":
		m = modular.NewOscillator()
	case "Clock":
		m = modular.NewClock()
	case "Sequencer":
		m = modular.NewSequencer()
	case "VCA":
		m = modular.NewVCA()
	default:
		return nil, fmt.Errorf("unknown module type: %s", moduleType)
	}
	fmt.Println("OK.")
	return m, nil
}

func extractModuleDependencies(def ModuleDef) map[string]string {
	dependencies := make(map[string]string)

	for k, v := range def {
		if x, ok := v.(string); ok && strings.HasPrefix(x, "@") {
			dependencies[k] = x[1:]
		}
	}
	return dependencies
}

// Initializes all modules in the patch to their default values.
func buildAllModules(patch PatchDefinition) (map[string]interface{}, error) {
	fmt.Println("Building modules...")
	builtModules := make(map[string]interface{})
	for moduleName, moduleToBuild := range patch.Modules {
		newModule, err := buildModule(moduleToBuild)
		if err != nil {
			return nil, err
		}
		builtModules[moduleName] = newModule
	}
	return builtModules, nil
}

func patchSingleModule(targetField, destObject reflect.Value) {
	if knob, ok := targetField.Interface().(*core.Knob); ok {
		knob.Line = destObject.Interface().(core.Signal)
	} else {
		targetField.Set(destObject)
	}
}

func patchAllModules(modules map[string]interface{}, patch PatchDefinition) error {
	fmt.Println("Applying patch to initialized modules...")

	for moduleName, moduleDefinition := range patch.Modules {
		deps := extractModuleDependencies(moduleDefinition)
		module := modules[moduleName]

		for attribute, targetModule := range deps {
			ps := reflect.ValueOf(module)
			s := ps.Elem() // Get the struct behind the module pointer
			targetField := s.FieldByName(attribute)
			if !targetField.IsValid() || !targetField.CanSet() {
				panic(fmt.Sprintf("invalid field: %s", attribute))
			}

			splitted := strings.Split(targetModule, ".")
			if len(splitted) == 1 {
				// Patch target is other module.
				// Field is settable. Obtain pointer to the target module and patch
				patchSingleModule(targetField, reflect.ValueOf(modules[targetModule]))
				fmt.Printf("\t- Patched %s -> %s.%s\n", targetModule, moduleName, attribute)
			} else if len(splitted) == 2 {
				// Patch target is an attribute of the module.
				targetModule = splitted[0]
				targetAttr := splitted[1]
				targetAttrVal := reflect.ValueOf(modules[targetModule]).Elem().FieldByName(targetAttr)
				patchSingleModule(targetField, targetAttrVal)
				fmt.Printf("\t- Patched %s.%s -> %s.%s\n", targetModule, targetAttr, moduleName, attribute)
			} else {
				return errors.New("arbitrary nesting not supported")
			}
		}
	}
	return nil
}

// Patch applies a patch to the synthesizer.
func Patch(patch PatchDefinition, synth *core.Synthia) error {
	initializedModules, err := buildAllModules(patch)
	if err != nil {
		return err
	}

	if err := patchAllModules(initializedModules, patch); err != nil {
		return err
	}

	// Hook up to mixer
	for i := 0; i < len(patch.MixerChannels); i++ {
		synth.Mixer.Channels[i].Input = initializedModules[patch.MixerChannels[i].Input[1:]].(core.Signal)
	}
	return nil
}
