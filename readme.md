[![Travis](https://travis-ci.org/etnz/permute.svg?branch=master)](https://travis-ci.org/etnz/permute.svg?branch=master)
[![GoDoc](https://godoc.org/github.com/etnz/permute?status.svg)](https://godoc.org/github.com/etnz/permute)

golang package 'permute' provides a tools to deal with 
[permutations](https://en.wikipedia.org/wiki/Permutations), 
[combinations](https://en.wikipedia.org/wiki/Combination)
.

This package provide:

- utility functions to deal with permutations, transpositions, combinations
- algorithm to generate all permutations, combinations under different constraints.

The number of generated items can get huge, hence the effort to generate permutations or combinations 
so that successives elements differ only by a 'small' change.

The definition of 'small' varies, but what remains is the need to generate all permutations, or combinations in a specific order.


# Definitions

## permutations

A n-permutation is:

- a `[]int` of length `n`
- where each values are **unique**  *and* in the interval `[0, len[`

For example, a permutation:

    permutation []int{ 2  ,  1  ,  0  }
    transforms    x= {"a" , "b" , "c" }
    into             {x[2], x[1], x[0]}
    resulting in     {"c" , "b" , "a" }

    x := []string{"a", "b", "c"}
    Strings([]int{2, 1, 0}, x)
    fmt.Println(x)
    //Output: [c b a]

## Combinations

A (n,k)-combination or (n,k)-subset is:

- a `[]int` of length `k`
- where each values are **unique**  *and* in the interval `[0, len[`
- values are sorted  in ascending order.

For example, a combination:

    combination []int{ 0  ,  2  }
    transforms    x= {"a" , "b" , "c" }
    into             {x[0], x[2]}
    resulting in     {"a" , "c" }

    x := SubStrings([]int{0, 2}, []string{"a", "b", "c"})
    fmt.Println(x)
    //Output: [a c]




# Permutation Generation

This package offers several methods to generate all permutations.

To apply successive permutation to a vector it is usual to compute the list of transposition to be applied to move from the current position to the next.

For instance in the following permutations 

    1:ABDC
    2:ACBD

moving a vector from `ABDC` to `ACBD` can be done with two transpositions: 

- `(1,2)` swapping 'B' and 'D' into `ADBC`
- `(1,3)` swapping 'D' and 'C' into `ACBD`

Therefore it is reasonable to look for a list of all permutation so that successive elements differ only by one transposition, this would be the fastest way to apply them.


## Lexicographical Order

Generates all permutation in lexicographical order. 
This is not the fastest way to generate permutations, and two successives permutation can differ by many transposition.

## Plain Change Order

There are two available implementations.

### PlainChangeNext

Implements the [Steinhaus-Johnson-Trotter](https://en.wikipedia.org/wiki/Steinhaus%E2%80%93Johnson%E2%80%93Trotter_algorithm)

Generates all permutations so that two successive elements differ by swapping two adjacent elements.

This is the slowest method to generate permutations, but it does not require extra memory.

Object that can take advantage of this property, can very quickly apply the transposition. Usually objects with linear access time.

### PlainChangeGenerator

Implements the [Even speedup](https://en.wikipedia.org/wiki/Steinhaus%E2%80%93Johnson%E2%80%93Trotter_algorithm#Even.27s_speedup) on top of the [Steinhaus-Johnson-Trotter](https://en.wikipedia.org/wiki/Steinhaus%E2%80%93Johnson%E2%80%93Trotter_algorithm).

It adds up an extra memory ( O(n) ) that greatly speeds up the generation. The generated list is identical.

## Heap Order


Implements [Heap's Algorithm](https://en.wikipedia.org/wiki/Heap%27s_algorithm).

This is the fastest way to generate the permutations.
Successive elements differ by only one swapping two elements not necessarily adjacent.

By far, Heap is the fastest of all to generate permutation. It uses O(n) extra memory.

Object with random access, can apply such permutations as quickly as in Plain Change Order.

**Caveat**: the Heap order is not cyclic. It means that the last element and the first element do not differ from 1 transposition.


## Benchmarks 

Here are some benchmarks executed on my computer (relative numbers matter)

      BenchmarkPermGenHeap 100000000    12 ns/op
      BenchmarkPermGenEven  20000000   117 ns/op
      BenchmarkPermGenLex   10000000   176 ns/op
      BenchmarkPermGenSJT    2000000   787 ns/op



# Combination Generation

The same principles applies for the combination generation. We want to generate all n,k-combinations but in the best possible order.

Usually the best possible order is when successive elements differs by one item only.

## Lexicographical Order

Generates all combinations in lexicographical order. 
This is not the fastest way to generate permutations, and two successives combinations can differ by multiple transposition.


## Minimal Change Order

Implements the [Revolving Door Algorithm](https://books.google.fr/books?id=0ArDOdcWNQcC&lpg=PA48&ots=JEsy6Hgdio&dq=revolving%20door%20algorithm&pg=PA49#v=onepage&q=revolving%20door%20algorithm&f=false) where all combinations differ by one item only.

# Still on the workbench

- generating permutation by derangements (Lynn Yarbrough)
- generating combination with strong minimal change
- generating combination with adjacent Interchange
- generate integer partitions
