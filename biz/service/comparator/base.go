package comparator

import (
    "github.com/yanyiwu/gojieba"
)

type Comparator interface {
    Compare(str1 string, str2 string) float64
    ComputeStrLen(str string) int
}
// generic type just do direct comparison
type GenericComparator struct {
}

func (comp GenericComparator) ComputeStrLen(str string) int {
    return len(str)
}

func (comp GenericComparator) Compare(str1 string, str2 string) float64 {
    if str1 == str2 {
        return 0
    } else {
        return 1
    }
}

// 编辑距离
type EditDistance struct {
    gene GenericComparator
}

func (edit EditDistance) Compare(str1 string, str2 string) float64 {
    return minDistance(str1, str2)
}

func (edit EditDistance) ComputeStrLen(str string) int {
    // a string can have either english or chinese
    out := []rune(str)
    return len(out)
}
func minDistance(word1 string, word2 string) float64 {
    word1Rune := []rune(word1)
    word2Rune := []rune(word2)
    n1, n2 := len(word1Rune), len(word2Rune)
    dp := make([][]int, n1+1)
    for i := range dp {
        dp[i] = make([]int, n2+1)
    }
    for i := 0; i <= n1; i++ {
        dp[i][0] = i
    }
    for i := 0; i <= n2; i++ {
        dp[0][i] = i
    }

    for i := 1; i <= n1; i++ {
        for j := 1; j <= n2; j++ {
            if word1Rune[i-1] == word2Rune[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else {
                dp[i][j] = Min(
                    dp[i-1][j-1]+1,
                    dp[i-1][j]+1,
                    dp[i][j-1]+1)
            }
        }
    }
    return float64(dp[n1][n2])
}

func Min(args ...int) int {
    min := args[0]
    for _, item := range args {
        if item < min {
            min = item
        }
    }
    return min
}

type JaccardSimilarity struct {
    gene   GenericComparator
    UseHmm bool
    jieba  *gojieba.Jieba
}

func NewJaccardComparator(useHmm bool) *JaccardSimilarity {
    return &JaccardSimilarity{jieba: gojieba.NewJieba(), UseHmm: useHmm}
}

func (jac *JaccardSimilarity) Compare(str1 string, str2 string) float64 {
    // word segmentation
    words1 := jac.Cut(str1)
    words2 := jac.Cut(str2)
    inter := jac.intersection(words1, words2)
    union := jac.union(words1, words2)
    return float64(inter) / float64(union)
}

func (jac *JaccardSimilarity) ComputeStrLen(str string) int {
    // a string can have either english or chinese
    out := []rune(str)
    return len(out)
}

func (jac *JaccardSimilarity) intersection(words1 []string, words2 []string) int {

    wordSet := make(map[string]bool)
    for _, w1 := range words1 {
        for _, w2 := range words2 {
            _, found := wordSet[w1]
            if w1 == w2 && !found {
                wordSet[w1] = true
            }
        }
    }
    return len(wordSet)
}

func (jac *JaccardSimilarity) union(words1 []string, words2 []string) int {
    wordSet := make(map[string]bool)
    for _, w1 := range words1 {
        for _, w2 := range words2 {
            wordSet[w1] = true
            wordSet[w2] = true
        }
    }
    return len(wordSet)
}

func (jac *JaccardSimilarity) Cut(str string) []string {
    // a string can have either english or chinese
    if !jac.UseHmm {
        return jac.jieba.CutAll(str)
    } else {
        return jac.jieba.Cut(str, jac.UseHmm)
    }
}

func (jac *JaccardSimilarity) Close() {
    jac.jieba.Free()
}
