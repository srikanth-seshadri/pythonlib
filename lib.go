/*
 * Some general purpose functions to get python like convenience.
 * Some code is lifted from https://github.com/golang/go/wiki/SliceTricks
 */
package pythonlib

import (
    "strings"
    "github.com/fatih/set"
)

// Right split a string.
func Rsplit(s string, substr string) []string {
    var result []string

    index := strings.LastIndex(s, substr)
    if index != -1 {
        result = append(result, s[:index])
        result = append(result, s[index+1:])
    }
    return result
}

// Contains tells whether s contains x.
func Contains(s []string, x string) bool {
    for _, e := range s {
        if e == x {
            return true
        }
    }
    return false
}

// Find returns the smallest index i at which x == s[i],
// or -1 if there is no such index.
func Index(s []string, x string) int {
    for i, e := range s {
        if e == x {
            return i
        }
    }
    return -1
}

// Remove the element x from the list s and return the new s.
// 2nd arguement: If x exists, returns true.
// if x doesn't exist, nil, false is returned.
func Del(s []string, x string) ([]string, bool) {
    i := Index(s, x)
    if i != -1 {
        s = append(s[:i], s[i+1:]...)
        return s, true
    } else {
        return nil, false
    }
}

// adds elements to set
func AddList(s set.Interface, l interface {}) {
    switch v := l.(type) {
    case []string:
        for _, e := range(v) {
            s.Add(e)
        }
    case []int:
        for _, e := range(v) {
            s.Add(e)
        }
    }
}

// adds elements of the integer list l to the set s.
func AddIntList(s set.Set, l []int) {
    for _, e := range(l) {
        s.Add(e)
    }
}
