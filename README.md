This is a random number generator factory and ports of splitmix64, xoroshiro128+, xoshiro256++, and xoshiro256** to Go.

You need to `go get github.com/amanitaverna/rnggo` before you can use this. You don't need to `go get` the subpackages.

There are two ways to use an existing random number generator (rnggo includes three, for example, but anyone can make their own).

The first way is to `import _ "github.com/amanitaverna/rnggo/xoshiro256starstar"`, and then to create the RNG, call `rnggo.GetGenerator(rnggo.GetGeneratorType("xoshiro256**"), seed)`, where seed is zero or more uint64 parameters.

The second way is to just call `rnggo.GetGenerator(xoshiro256starstar.GenType(), seed)`, where seed is (still) zero or more uint64 parameters. This relies on the generator having a GenType function. `splitmix64`, `xoroshiro128plusplus`, `xoshiro256plusplus`, and `xoshiro256starstar` all do. If you make your own generator, you should include one too for convenience.

Note that if you don't pass any seed values, the seed will be initialized to a fixed value. To seed a `Generator` using the clock, you can for example call `gen.Seed(uint64(time.Now().UnixMicro()))`.

You can seed a `Generator` with its `Seed` method, and can re-seed an existing `Generator` as many times as you want.

To generate random numbers with a `Generator`, you can call the appropriate method on `Generator`, such as `RandNext`, which generates a uint64, `RandInt`, which takes a limit and generates a number >= 0 and < that limit, `RandBool`, which generates a random boolean value, or `RandFloat`, which generates a random float64 in the range >= 0 to < 1.

`GetGeneratorNames` returns a slice of generator names, if you want to see what is available, but things will only be available if their package has been loaded (by being imported) so that they can register their generator type and function.

To make your own generator type, see generator.go in any of the sub-packages (except splitmix64) for an example.