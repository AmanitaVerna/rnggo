package rnggo

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

// The type of random number generator. To register a new one, implement IGenerator, and call RegisterGeneratorType and pass your IGenerator to it from a func init().
type GeneratorType int

const (
	Nonexistent GeneratorType = -1
)

// A function which returns a new Generator
type NewGeneratorFunc func() IGenerator

var registeredGeneratorNames map[string]GeneratorType = make(map[string]GeneratorType)
var registeredGeneratorFunctions map[GeneratorType]NewGeneratorFunc = make(map[GeneratorType]NewGeneratorFunc)
var numRegisteredGenerators GeneratorType = 0

// Create a Generator using NewGenerator.
type Generator struct {
	iGen IGenerator
}

// All random number generators that are registered with rnggo must implement IGenerator.
type IGenerator interface {
	// Seed the generator using an arbitrary amount of uint64s.
	Seed(sx ...uint64)
	// Return the generator's next random uint64.
	RandNext() uint64
}

// RegisterGeneratorType is for registering an RNG with an implementation of IGenerator.
// Pass RegisterGeneratorType the function that should be called to create a new instance of the RNG, along with the name of the RNG,
// and this will return a new GeneratorType for it.
// Returns an error if the name is "" or if the name is already registered. If an error is returned, genType will be Nonexistent (-1).
// Names are case-insensitive.
func RegisterGeneratorType(fn NewGeneratorFunc, name string) (genType GeneratorType, err error) {
	genType = Nonexistent
	err = nil
	if len(name) == 0 {
		err = errors.New("Name passed to RegisterGenerator must not be \"\"")
	} else {
		nameLower := strings.ToLower(name)
		if existingGenType, exists := registeredGeneratorNames[nameLower]; exists {
			err = errors.New(fmt.Sprintf("Name (%v) is already registered as GeneratorType %v.", nameLower, existingGenType))
		} else {
			genType = numRegisteredGenerators
			registeredGeneratorNames[nameLower] = genType
			registeredGeneratorFunctions[genType] = fn
			numRegisteredGenerators++
		}
	}
	return
}

// GetGeneratorType returns the GeneratorType for the specified name, or Nonexistent (-1) if no generator is registered with that name.
// Names are case-insensitive.
func GetGeneratorType(name string) (genType GeneratorType) {
	nameLower := strings.ToLower(name)
	var exists bool
	if genType, exists = registeredGeneratorNames[nameLower]; !exists {
		genType = -1
	}
	return
}

// NewGenerator creates a new Generator and seeds it with sx.
func NewGenerator(genType GeneratorType, sx ...uint64) (gen *Generator) {
	if genType < numRegisteredGenerators {
		genFunc := registeredGeneratorFunctions[genType]
		gen = &Generator{iGen: genFunc()}
	} else {
		return nil
	}
	gen.Seed(sx...)
	return
}

// Seed seeds the Generator.
func (gen *Generator) Seed(sx ...uint64) {
	gen.iGen.Seed(sx...)
}

// RandNext returns the Generator's next uint64.
func (gen *Generator) RandNext() uint64 {
	return gen.iGen.RandNext()
}

// RandInt returns a number from 0 to limit-1.
func (gen *Generator) RandInt(limit uint64) uint64 {
	return gen.RandNext() % limit
}

// RandBool returns a random boolean value.
func (gen *Generator) RandBool() bool {
	return gen.RandNext()&0x8000000000000000 > 0
}

// RandFloat returns a random 64-bit floating-point number.
func (gen *Generator) RandFloat() float64 {
	return math.Float64frombits(gen.RandNext())
}
