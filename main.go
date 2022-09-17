package main

// Product is a sellable item
type Product struct {
	ID    int
	Name  string
	Price int
}

// Alternating Product Sequence (APS)

// An APS is defined as a sequence of products where the sign of the difference between
// every consecutive pair of product prices is different.

//  Examples:
//
//        Valid APS:
//        [
//              product{ID: 1, Name: "shoes", Price: 8},
//              product{ID: 2, Name: "shirt", Price: 4},
//              product{ID: 3, Name: "socks", Price: 5},
//              product{ID: 4, Name: "sweater", Price: 2},
//              product{ID: 5, Name: "jacket", Price: 6},
//              product{ID: 6, Name: "pants", Price: 1}
//        ]
//
//        Invalid APS:
//        [
//              product{ID: 1, Name: "shoes", Price: 8},
//              product{ID: 2, Name: "shirt", Price: 5},
//              product{ID: 3, Name: "socks", Price: 10},
//              product{ID: 4, Name: "sweater", Price: 12},
//              product{ID: 5, Name: "jacket", Price: 14},
//              product{ID: 6, Name: "pants", Price: 6}
//        ]
//

// isAPS returns whether or not the given array is an alternating product sequence

func IsAPS(products []*Product) bool {
	var lastResult bool
	var finalResult bool = true

	for idx := 0; idx < len(products)-1; idx++ {

		currentPrice := products[idx].Price
		nextPrice := products[idx+1].Price

		currentResult := false
		if currentPrice > nextPrice {
			currentResult = true
		}

		if currentResult == lastResult {
			finalResult = false
			break
		}

		lastResult = currentResult
	}

	return finalResult
}

// If an array of products isn't an APS, there could potentially be some subset that is an APS.
// An alternating product subsequence (APSS) is any subsequence of an array of products that is an APS itself.
//
// For example, consider the following array of products that is not an APS, as demonstrated above.
//
//      Invalid APS:
//      [
//              product{ID: 1, Name: "shoes", Price: 8},
//              product{ID: 2, Name: "shirt", Price: 5},
//              product{ID: 3, Name: "socks", Price: 10},
//              product{ID: 4, Name: "sweater", Price: 12},
//              product{ID: 5, Name: "jacket", Price: 14},
//              product{ID: 6, Name: "pants", Price: 6}
//      ]
//
// If we exclude ID's 4 and 5, we are left with:
//      [
//              product{ID: 1, Name: "shoes", Price: 8},
//              product{ID: 2, Name: "shirt", Price: 5},
//              product{ID: 3, Name: "socks", Price: 10},
//              product{ID: 6, Name: "pants", Price: 6}
//      ]
//
// which is a valid APS.
//
// Note: an array of products could have multiple APSS's.

// longestAPSSubsequenceLength returns the length of the longest APSS in the given array
func longestAPSSubsequenceLength(products []*Product) int {
	var lastResult bool
	var finalArray []Product

	for idx := 0; idx < len(products)-1; idx++ {

		currentPrice := products[idx].Price
		nextPrice := products[idx+1].Price

		currentResult := false
		if currentPrice > nextPrice {
			currentResult = true
		}

		if currentResult == lastResult {
			continue
		}

		finalArray = append(finalArray, *products[idx])

		lastResult = currentResult
	}

	return len(finalArray)
}
