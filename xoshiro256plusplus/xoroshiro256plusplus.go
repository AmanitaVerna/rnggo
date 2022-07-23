package xoshiro256plusplus

// Ported to Go by Amy Snow in 2021.

/*  Written in 2019 by David Blackman and Sebastiano Vigna (vigna@acm.org)

To the extent possible under law, the author has dedicated all copyright
and related and neighboring rights to this software to the public domain
worldwide. This software is distributed without any warranty.

See <http://creativecommons.org/publicdomain/zero/1.0/>. */

/* This is xoshiro256++ 1.0, one of our all-purpose, rock-solid generators.
   It has excellent (sub-ns) speed, a state (256 bits) that is large
   enough for any parallel application, and it passes all tests we are
   aware of.

   For generating just floating-point numbers, xoshiro256+ is even faster.

   The state must be seeded so that it is not everywhere zero. If you have
   a 64-bit seed, we suggest to seed a splitmix64 generator and use its
   output to fill s. */

var jumpConst []uint64 = []uint64{0x180ec6d33cfd0aba, 0xd5a61266f0c9392c, 0xa9582618e03fc9aa, 0x39abdc4529b1661c}
var longJumpConst []uint64 = []uint64{0x76e15d3efefdcbbf, 0xc5004e441c522fb3, 0x77710069854ee241, 0x39109bb02acbe635}

func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

func (gen *Generator) next() (result uint64) {
	result = rotl(gen.seed[0]+gen.seed[3], 23) + gen.seed[0]
	t := gen.seed[1] << 17
	gen.seed[2] ^= gen.seed[0]
	gen.seed[3] ^= gen.seed[1]
	gen.seed[1] ^= gen.seed[2]
	gen.seed[0] ^= gen.seed[3]

	gen.seed[2] ^= t

	gen.seed[3] = rotl(gen.seed[3], 45)

	return result
}

/* This is the jump function for the generator. It is equivalent
   to 2^128 calls to next(); it can be used to generate 2^128
   non-overlapping subsequences for parallel computations. */
func (gen *Generator) Jump() {
	var s0, s1, s2, s3 uint64
	var b uint64
	for i := 0; i < len(jumpConst); i++ {
		for b = 0; b < 64; b++ {
			if (jumpConst[i] & (1 << b)) != 0 {
				s0 ^= gen.seed[0]
				s1 ^= gen.seed[1]
				s2 ^= gen.seed[2]
				s3 ^= gen.seed[3]
			}
			gen.next()
		}
	}
	gen.seed[0] = s0
	gen.seed[1] = s1
	gen.seed[2] = s1
	gen.seed[3] = s1
}

/* This is the long-jump function for the generator. It is equivalent to
2^192 calls to next(); it can be used to generate 2^64 starting points,
from each of which jump() will generate 2^64 non-overlapping
subsequences for parallel distributed computations. */
func (gen *Generator) LongJump() {
	var s0, s1, s2, s3 uint64
	var b uint64
	for i := 0; i < len(longJumpConst); i++ {
		for b = 0; b < 64; b++ {
			if (longJumpConst[i] & (1 << b)) != 0 {
				s0 ^= gen.seed[0]
				s1 ^= gen.seed[1]
				s2 ^= gen.seed[2]
				s3 ^= gen.seed[3]
			}
			gen.next()
		}
	}
	gen.seed[0] = s0
	gen.seed[1] = s1
	gen.seed[2] = s1
	gen.seed[3] = s1
}
