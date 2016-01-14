[![Travis](https://travis-ci.org/etnz/permute.svg?branch=master)](https://travis-ci.org/etnz/permute.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/etnz/permute?status.svg)](https://godoc.org/github.com/etnz/permute)

# permute


golang package 'permute' provides a tools to deal with permutations.

A permutation is exactly a '[]int', where each values are unique, *and* in the interval [0, len[

    []int{0, 1, 2} is the identity
    
A permutation

    permutation []int{ 2  ,  1  ,  0  }
    tranforms     x= {"a" , "b" , "c" }
    into             {x[2], x[1], x[0]}
    resulting in     {"c" , "b" , "a" }

In addition to applying a permutation to a collection it offers four different ways to generate all permutations :

## Lexicographical

Generates all permutation in lexicographical order. Not the fastest way to generate permutations.
Permutations in lexicographical order can imply several transposition to apply.

## SteinhausJohnsonTrotter

Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two adjacent elements.
This is the slowest method to generate permutations.
But it does not require extra memory (like with Even speed up).
The generated sequence is the fastest to apply: there is only one transposition each time, and of adjacent items.

## SteinhausJohnsonTrotterEven

Add to SteinhausJohnsonTrotter method a little extra memory (O(n)) that greatly speeds up the permutation generation.

## Heap

Generates all permutation so that each permutation in the sequence differs from the previous permutation by swapping two element (not necessarily adjacent).
By far, Heap is the fastest of all to generate permutation. If there is no huge benefits in swapping adjacent items ( random access to collection items), then
it the best choice.

Here are some benchmarks executed on my computer (relative numbers matter)

      BenchmarkPermGenHeap	100000000	        12 ns/op
      BenchmarkPermGenEven	20000000	       117 ns/op
      BenchmarkPermGenLex	10000000	       176 ns/op
      BenchmarkPermGenSJT	 2000000	       787 ns/op
