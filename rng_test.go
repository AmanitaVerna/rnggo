package rnggo_test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/amanitaverna/rnggo"
	_ "github.com/amanitaverna/rnggo/xoroshiro128plus"
	_ "github.com/amanitaverna/rnggo/xoshiro256plusplus"
	_ "github.com/amanitaverna/rnggo/xoshiro256starstar"
	"github.com/stretchr/testify/assert"
)

func TestRNG(t *testing.T) {
	timeNow := uint64(time.Now().UnixMicro())
	numGenerated := 8
	var seeds []uint64 = []uint64{1632573360482, 1632573360493, 1632575188169, 1632575188180, timeNow, timeNow + 1}
	genNames := []string{"xoroshiro128+", "xoshiro256++", "xoshiro256**"}
	// map the generator names through GetGeneratorType to get the corresponding GeneratorTypes
	builder := strings.Builder{}
	genTypes := mapFunc(rnggo.GetGeneratorType, genNames)

	// if we wanted to use all the seeds to seed a single generator per type, we could have done this:
	// generators := mapFunc(func(x rnggo.GeneratorType) *rnggo.Generator { return rnggo.NewGenerator(x, seeds...) }, genTypes)
	// instead we're doing this:
	for i, genType := range genTypes {
		if genType == rnggo.Nonexistent {
			assert.Fail(t, fmt.Sprintf("Failed to find generator type '%v'", genNames[i]))
		} else {
			builder.WriteString(fmt.Sprintf("\t\t\t\t=== %v ===\n", genNames[i]))
			builder.WriteString("\t\tSeeds:\t\t\t\t\t\t\t\t\tResults:\n")
			for _, seed := range seeds {
				gen := rnggo.NewGenerator(genType, seed)
				builder.WriteString(fmt.Sprintf("%16v (%013x)\t\t", seed, seed))
				for ri := 0; ri < numGenerated; ri++ {
					result := gen.RandNext()
					builder.WriteString(fmt.Sprintf("%016x ", result))
				}
				builder.WriteString("\n")
			}
			builder.WriteString("\n")
		}
	}
	fmt.Printf(builder.String())
}

// mapFunc works like the map function in haskell or other functional languages.
// It returns the slice "obtained by applying f to every element of xs" (to quote the Haskell documentation).
// It almost certainly isn't as efficient as the function in Haskell, but, well.
// It also can't do partial application, of course, so that has to be done another way.
func mapFunc[In comparable, Out comparable](f func(In) Out, xs []In) (ret []Out) {
	if f != nil && xs != nil && len(xs) > 0 {
		ret = make([]Out, len(xs))
		for i, x := range xs {
			ret[i] = f(x)
		}
	}
	return
}
