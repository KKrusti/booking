package entities

type Request struct {
	Id          string  `json:"request_id"`
	Checkin     string  `json:"check_in"`
	Nights      int     `json:"nights"`
	SellingRate float64 `json:"selling_rate"`
	Margin      float64 `json:"margin"`
}

func combinations(requests []Request) (subsets [][]Request) {
	length := uint(len(requests))

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		var subset []Request

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, requests[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}
