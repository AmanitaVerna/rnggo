package rnggo

/*
import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRNG(t *testing.T) {
	timeNow := uint64(time.Now().UnixMicro())
	var seeds []uint64 = []uint64{1632573360482, 1632573360493, 1632575188169, 1632575188180, timeNow, timeNow + 1}
	var results []uint64 = make([]uint64, len(seeds))
	gen := &Generator{}
	builder := strings.Builder{}
	builder.WriteString("\t\tSeeds:\t\t\t\t\t\t\t\t\tResults:\n")
	for i, seed := range seeds {
		gen.Seed(seed)
		results[i] = gen.RandNext()
		builder.WriteString(fmt.Sprintf("%16v (%013x)\t\t%21v (%016x)\n", seed, seed, results[i], results[i]))
	}

	assert.Fail(t, builder.String())
}
*/
