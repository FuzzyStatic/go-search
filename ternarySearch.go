/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-22T23:10:25-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-23T22:34:47-05:00
 */

package search

// TernarySearch is a divide-and-conquer algorithm. It is mandatory for the array
// (in which you will search for an element) to be sorted before you begin the search.
// In this search, after each iteration it neglects ⅓ part of the array and repeats
// the same operations on the remaining ⅔.
func (s *Search) TernarySearch(low, high, value int) int {
	if high >= low {
		mid1 := low + (high-low)/3
		mid2 := high - (high-low)/3

		if (*s.slice)[mid1] == value {
			return mid1
		}

		if (*s.slice)[mid2] == value {
			return mid2
		}

		if value < (*s.slice)[mid1] {
			return s.TernarySearch(low, mid1-1, value)
		} else if value > (*s.slice)[mid2] {
			return s.TernarySearch(mid2+1, high, value)
		} else {
			return s.TernarySearch(mid1+1, mid2-1, value)
		}
	}

	return -1 // Index not found
}
