/**
* @Author: Allen Flickinger <FuzzyStatic>
* @Date:   2017-02-22T23:25:16-05:00
* @Email:  allen.flickinger@gmail.com
* @Last modified by:   FuzzyStatic
* @Last modified time: 2017-02-23T22:35:33-05:00
 */

package search

import (
	"reflect"
	"testing"
)

var slice = &[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestBinarySearch(t *testing.T) {
	t.Logf("Test binary search of %v", *slice)
	s1, _ := New(slice)
	index := s1.BinarySearch(4)

	if !reflect.DeepEqual(index, 3) {
		t.Errorf("Expected index of %v, but it was %v instead.", 3, index)
	}

	t.Logf("Test binary search of %v", *slice)
	s2, _ := New(slice)
	index2 := s2.BinarySearch(8)

	if !reflect.DeepEqual(index2, 7) {
		t.Errorf("Expected index of %v, but it was %v instead.", 7, index2)
	}

	t.Logf("Test binary search of %v", *slice)
	s3, _ := New(slice)
	index3 := s3.BinarySearch(11)

	if !reflect.DeepEqual(index3, -1) {
		t.Errorf("Expected index of %v, but it was %v instead.", -1, index3)
	}
}

func TestLinearSearch(t *testing.T) {
	var search = []int{5}
	var search2 = []int{5, 6}
	var search3 = []int{}

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

	slice3 := &[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	t.Logf("Test linear search of %v", *slice3)
	s3, _ := New(slice3)
	indexes3 := *s3.LinearSearch(11)

	if !reflect.DeepEqual(indexes3, search3) {
		t.Errorf("Expected slice of %v, but it was %v instead.", search3, indexes3)
	}
}

func TestTernarySearch(t *testing.T) {
	t.Logf("Test binary search of %v", *slice)
	s1, _ := New(slice)
	index := s1.TernarySearch(0, len(*s1.slice)-1, 4)

	if !reflect.DeepEqual(index, 3) {
		t.Errorf("Expected index of %v, but it was %v instead.", 3, index)
	}

	t.Logf("Test binary search of %v", *slice)
	s2, _ := New(slice)
	index2 := s2.TernarySearch(0, len(*s2.slice)-1, 6)

	if !reflect.DeepEqual(index2, 5) {
		t.Errorf("Expected index of %v, but it was %v instead.", 5, index2)
	}

	t.Logf("Test binary search of %v", *slice)
	s3, _ := New(slice)
	index3 := s3.TernarySearch(0, len(*s3.slice)-1, 2)

	if !reflect.DeepEqual(index3, 1) {
		t.Errorf("Expected index of %v, but it was %v instead.", 1, index3)
	}

	t.Logf("Test binary search of %v", *slice)
	s4, _ := New(slice)
	index4 := s4.TernarySearch(0, len(*s4.slice)-1, 11)

	if !reflect.DeepEqual(index4, -1) {
		t.Errorf("Expected index of %v, but it was %v instead.", -1, index4)
	}
}
