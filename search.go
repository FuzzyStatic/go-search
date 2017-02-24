/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-18T12:56:00-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-22T23:26:44-05:00
 */

package search

// Search slice
type Search struct {
	slice *[]int
}

// New create new Search object with slice
func New(slice *[]int) (*Search, error) {
	var s = Search{slice: slice}
	return &s, nil
}
