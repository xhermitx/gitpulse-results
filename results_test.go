package results

import (
	"reflect"
	"testing"
)

func TestTopNCandidates(t *testing.T) {
	list := []*Candidate{
		{
			TotalContributions:      1200,
			TotalFollowers:          350,
			TopRepoStars:            450,
			TopContributedRepoStars: 800,
			Languages:               []string{"Go", "JavaScript", "Python"},
			Topics:                  []string{"Distributed Systems", "Backend", "Concurrency"},
		},
		{
			TotalContributions:      950,
			TotalFollowers:          200,
			TopRepoStars:            300,
			TopContributedRepoStars: 700,
			Languages:               []string{"Python", "C++"},
			Topics:                  []string{"Compression", "Security"},
		},
		{
			TotalContributions:      1600,
			TotalFollowers:          500,
			TopRepoStars:            600,
			TopContributedRepoStars: 900,
			Languages:               []string{"Go", "Rust", "JavaScript"},
			Topics:                  []string{"DevOps", "Backend", "Data Structures"},
		},
		{
			TotalContributions:      800,
			TotalFollowers:          150,
			TopRepoStars:            250,
			TopContributedRepoStars: 650,
			Languages:               []string{"C#", "Java"},
			Topics:                  []string{"Game Development", "AI"},
		},
		{
			TotalContributions:      1400,
			TotalFollowers:          450,
			TopRepoStars:            500,
			TopContributedRepoStars: 850,
			Languages:               []string{"Go", "Java", "Python", "Ruby"},
			Topics:                  []string{"Machine Learning", "Backend", "APIs"},
		},
	}

	t.Run("check if algorithm works", func(t *testing.T) {
		want := []int{21, 24, 25}
		res := TopNCandidates(list, 3)
		got := make([]int, 3)
		for i, c := range res {
			got[i] = c.Score
		}

		if !reflect.DeepEqual(want, got) {
			t.Errorf("Got %v Want %v", got, want)
		}
	})
}
