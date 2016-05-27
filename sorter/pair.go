package sorter

import "sort"

// Sort the words
func Sort(wordFrequencies map[string]int) PairList{
  pl := make(PairList, len(wordFrequencies))
  i := 0
  for k, v := range wordFrequencies {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

type Pair struct {
  Word string
  Occurances int
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Occurances < p[j].Occurances }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
