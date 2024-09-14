package results

type CandidateMeta struct {
	Name       string
	Username   string
	AvatarURL  string
	Bio        string
	Email      string
	WebsiteURL string
}

type Candidate struct {
	CandidateMeta
	CandidateId             string
	TotalContributions      int
	TotalFollowers          int
	TopRepo                 string
	TopRepoStars            int
	TopContributedRepo      string
	TopContributedRepoStars int
	Languages               []string
	Topics                  []string
	JobId                   string
	Score                   int
}

// Minheap to maintaing the top n candidates with highest scores
type MinHeap []*Candidate

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Score < h[j].Score }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Candidate))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
