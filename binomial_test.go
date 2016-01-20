package permute

import "fmt"

func ExampleBinomial() {
	for n := 0; n <= 5; n++ {
		for k := 0; k <= n; k++ {
			fmt.Printf("%v ", Binomial(n, k))
		}
		fmt.Println(",")
	}

	//Output:
	// 1 ,
	// 1 1 ,
	// 1 2 1 ,
	// 1 3 3 1 ,
	// 1 4 6 4 1 ,
	// 1 5 10 10 5 1 ,
}
