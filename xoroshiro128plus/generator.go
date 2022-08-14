package xoroshiro128plus

import (
	"fmt"

	"github.com/amanitaverna/rnggo"
	"github.com/amanitaverna/rnggo/splitmix64"
)

var genType rnggo.GeneratorType

func init() {
	var err error
	if genType, err = rnggo.RegisterGeneratorType(NewGenerator, "xoroshiro128+"); err != nil {
		fmt.Printf("Failed to register xoroshiro128+ with rnggo: %v\n", err.Error())
	}
}

// Returns a new IGenerator for this RNG
func NewGenerator() rnggo.IGenerator {
	return &Generator{}
}

// Returns the GeneratorType of this RNG
func GenType() rnggo.GeneratorType {
	return genType
}

const fnvOffsetBasis uint64 = 14695981039346656037
const fnvPrime uint64 = 1099511628211

type Generator struct {
	rngS0, rngS1 uint64
}

// Seed seeds the RNG using 64-bit unsigned integers.
func (gen *Generator) Seed(sx ...uint64) {
	fnv := fnvOffsetBasis
	for _, s := range sx {
		fnv = (fnv ^ s) * fnvPrime
	}
	gen.rngS0 = splitmix64.Next(&fnv)
	gen.rngS1 = splitmix64.Next(&fnv)
}

// RandNext returns a random uint64.
func (gen *Generator) RandNext() uint64 {
	return gen.next()
}

func fnv1a(ts, px, py uint64) uint64 {
	return (((((fnvOffsetBasis ^ ts) * fnvPrime) ^ px) * fnvPrime) ^ py) * fnvPrime
}
