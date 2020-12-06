package comparator

import (
    "math"
    "strconv"
    "strings"
)

type SimHash struct {
    IntSimHash int64
    HashBits   int
}

func (s *SimHash) Compare(str1 string, str2 string) float64 {
    s1 := s.Simhash(str1)
    s2 := s.Simhash(str2)
    return s.Similarity(s1, s2)
}

func (s *SimHash) ComputeStrLen(str string) int {
    return len([]rune(str))
}

/**
距离 补偿
*/
func (s *SimHash) HammingDistance(hash, other int64) int {
    x := (hash ^ other) & ((1 << uint64(s.HashBits)) - 1)
    tot := 0
    for x != 0 {
        tot += 1
        x &= x - 1
    }
    return tot
}

/**
  相似度
*/
func (s *SimHash) Similarity(hash, other int64) float64 {
    a := float64(hash)
    b := float64(other)
    if a > b {
        return b / a
    }
    return a / b
}

/**
  海明距离hash
*/
func (s *SimHash) Simhash(str string) int64 {
    m := strings.Split(str, " ")
    tokenInt := make([]int, s.HashBits)
    for i := 0; i < len(m); i++ {
        temp := m[i]
        t := s.Hash(temp)
        for j := 0; j < s.HashBits; j++ {
            bitMask := int64(1 << uint(j))
            if t&bitMask != 0 {
                tokenInt[j] += 1
            } else {
                tokenInt[j] -= 1
            }
        }

    }
    var fingerprint int64 = 0
    for i := 0; i < s.HashBits; i++ {
        if tokenInt[i] >= 0 {
            fingerprint += 1 << uint64(i)
        }
    }
    return fingerprint
}

/**
  初始化
*/
func NewSimHash() (s *SimHash) {
    s = &SimHash{}
    s.HashBits = 64
    return s
}

/**
hash 值
*/
func (s *SimHash) Hash(token string) int64 {
    if token == "" {
        return 0
    } else {
        x := int64(int(token[0]) << 7)
        m := int64(1000003)
        mask := math.Pow(2, float64(s.HashBits-1))
        s := strconv.FormatFloat(mask, 'f', -1, 64)
        tsk, _ := strconv.ParseInt(s, 10, 64)
        for i := 0; i < len(token); i++ {
            tokens := int64(int(token[0]))
            x = ((x * m) ^ tokens) & tsk
        }
        x ^= int64(len(token))
        if x == -1 {
            x = -2
        }
        return x
    }
}

