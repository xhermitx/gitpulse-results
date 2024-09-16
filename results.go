package results

import (
	"container/heap"
	"fmt"
	"math"
)

func TopNCandidates(list []*Candidate, n int) []*Candidate {
	var (
		sumContributions           = 0
		sumFollowers               = 0
		sumTopRepoStars            = 0
		sumTopContributedRepoStars = 0
		sumLanguages               = 0
		sumTopics                  = 0
	)

	for _, c := range list {
		sumContributions += c.TotalContributions
		sumFollowers += c.TotalFollowers
		sumTopRepoStars += c.TopRepoStars
		sumTopContributedRepoStars += c.TopContributedRepoStars
		sumLanguages += c.Languages
		sumTopics += c.Topics
	}

	h := &MinHeap{}
	heap.Init(h)

	for _, c := range list {
		c.Score += int(math.Floor(float64(c.TotalContributions)/float64(sumContributions)*20 +
			float64(c.TotalFollowers)/float64(sumFollowers)*5 +
			float64(c.TopRepoStars)/float64(sumTopRepoStars)*25 +
			float64(c.TopContributedRepoStars)/float64(sumTopContributedRepoStars)*25 +
			float64(c.Languages)/float64(sumLanguages)*15 +
			float64(c.Topics)/float64(sumTopics)*10))

		fmt.Println("Score: ", c.Score)

		// FIXME: Add the candidate to a heap of size payload.Count
		if h.Len() < 3 {
			heap.Push(h, c)
		} else if c.Score > (*h)[0].Score {
			heap.Pop(h)
			heap.Push(h, c)
		}
	}

	res := make([]*Candidate, n)
	for i := range n {
		res[i] = heap.Pop(h).(*Candidate)
	}
	return res
}
