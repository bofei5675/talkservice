package service

import (
    "errors"
    "fmt"
    "talkapp/biz/service/comparator"
)

const (
    BaseComparator = "base"
    EditDistance = "edit_distance"
    HammingDistance = "hamming_distance"
    JaccardDistance = "jaccard_distance"
) 

func CompareTwoString(str1, str2, method string) (float64, error) {
    var comp comparator.Comparator
    switch method {
    case BaseComparator:
        comp = comparator.GenericComparator{}
    case EditDistance:
        comp = comparator.EditDistance{}
    case JaccardDistance:
        comp = comparator.NewJaccardComparator(true)
    case HammingDistance:
        comp = comparator.NewSimHash()
    default:
        return -1, errors.New(fmt.Sprintf("unhandled comparator type %s", method))
    }
    return comp.Compare(str1, str2), nil
}
