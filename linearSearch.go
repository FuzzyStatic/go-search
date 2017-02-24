/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-22T23:10:25-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-23T21:26:22-05:00
 */

package search

// LinearSearch  is used on a collections of items. It relies on the technique of
// traversing a list from start to end by exploring properties of all the elements
// that are found on the way.
func (s *Search) LinearSearch(value int) *[]int {
	var slice []int

	for i, v := range *s.slice {
		if v == value {
			slice = append(slice, i)
		}
	}

	return &slice
}
