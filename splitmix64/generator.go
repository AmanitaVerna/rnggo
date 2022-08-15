package splitmix64

import (
	"fmt"

	"github.com/amanitaverna/rnggo"
)

var genType rnggo.GeneratorType

func init() {
	var err error
	if genType, err = rnggo.RegisterGeneratorType(NewGenerator, "splitmix64"); err != nil {
		fmt.Printf("Failed to register Splitmix64 with rnggo: %v\n", err.Error())
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
	seed uint64
}

// Seed seeds the RNG using 64-bit unsigned integers.
func (gen *Generator) Seed(sx ...uint64) {
	fnv := fnvOffsetBasis
	for _, s := range sx {
		fnv = (fnv ^ s) * fnvPrime
	}
	gen.seed = Next(&fnv)
}

// RandNext returns a random uint64.
func (gen *Generator) RandNext() uint64 {
	return Next(&gen.seed)
}

// not used
func fnv1a(ts, px, py uint64) uint64 {
	return (((((fnvOffsetBasis ^ ts) * fnvPrime) ^ px) * fnvPrime) ^ py) * fnvPrime
}
