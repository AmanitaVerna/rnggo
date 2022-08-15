This is a random number generator factory and ports of splitmix64, xoroshiro128+, xoshiro256++, and xoshiro256** to Go.

In order to use an existing generator, you first have to import the package for any generator you want to use, e.g. `import _ "github.com/amanitaverna/rnggo/xoshiro256starstar"`.
Next, call `GetGeneratorType` and pass it the name of the generator (such as "xoshiro256**"). It will return a `GeneratorType`. `GetGeneratorNames` returns a slice of generator names, if you want to see what is available, but things will only be available if their package has been loaded so that they can register their generator type and function.
With the `GeneratorType`, you can call `GetGenerator` and pass the `GeneratorType` and zero or more uint64 seed values. Note that if you don't pass any seed values, the seed will be initialized to a fixed value. 
You can also seed a `Generator` with its `Seed` method, and can re-seed an existing `Generator` at any time. To seed a `Generator` using the clock, you can for example call `gen.Seed(uint64(time.Now().UnixMicro()))`.
To generate random numbers with a `Generator`, you can call the appropriate method on `Generator`, such as `RandNext`, which generates a uint64, `RandInt`, which takes a limit and generates a number >= 0 and < that limit, `RandBool`, which generates a random boolean value, or `RandFloat`, which generates a random float64 in the range >= 0 to < 1.

To make your own generator type, see generator.go in any of the sub-packages (except splitmix64) for an example.