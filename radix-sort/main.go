package main

import "fmt"

func PhoneNumberSort(phones []string) {
	buckets := make([][]string, 10)

	for idx := 10; idx >= 0; idx-- {
		for _, v := range phones {
			bucketName := v[idx] - 48

			if buckets[bucketName] == nil {
				tmp := make([]string, 0)
				tmp = append(tmp, v)
				buckets[bucketName] = tmp
			} else {
				buckets[bucketName] = append(buckets[bucketName], v)
			}

			// fmt.Println(bucketName, buckets[bucketName], buckets, "\n")
		}

		fmt.Println(buckets, "\n\n\n")
	}

	i := 0

	for k, v := range buckets {
		if v != nil {
			fmt.Println("bucket ", v)
			for _, p := range v {
				phones[i] = p
				i++
			}
			buckets[k] = nil
		}
	}
}

func main() {
	phones := []string{
		"18671581234",
		"18971581274",
		"13671581273",
	}

	PhoneNumberSort(phones)
}
