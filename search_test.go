/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-22T23:25:16-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-23T21:27:37-05:00
 */

package search

import (
	"reflect"
	"testing"
)

var search = []int{5}
var search2 = []int{5, 6}

func TestLinearSearch(t *testing.T) {
	slice1 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test linear search of %v", *slice1)
	s1, _ := New(slice1)
	indexes := *s1.LinearSearch(4)

	if !reflect.DeepEqual(indexes, search) {
		t.Errorf("Expected slice of %v, but it was %v instead.", search, indexes)
	}

	slice2 := &[]int{6, 5, 4, 4, 8, 7, 7, 2, 1}
	t.Logf("Test linear search of %v", *slice2)
	s2, _ := New(slice2)
	indexes2 := *s2.LinearSearch(7)

	if !reflect.DeepEqual(indexes2, search2) {
		t.Errorf("Expected slice of %v, but it was %v instead.", search2, indexes2)
	}
}
