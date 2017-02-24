/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-22T23:10:25-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-23T22:15:47-05:00
 */

package search

// BinarySearch is used to perform operations on a sorted set, the number of iterations
// can always be reduced on the basis of the value that is being searched.
func (s *Search) BinarySearch(value int) int {
	low := 0
	high := len(*s.slice) - 1

	for low <= high {
		mid := (low + high) / 2

		if (*s.slice)[mid] < value {
			low = mid + 1
		} else if (*s.slice)[mid] > value {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1 // Index not found
}
