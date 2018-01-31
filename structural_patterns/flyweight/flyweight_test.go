package flyweight

import (
	"log"
	"testing"
)

func TestFlyweight(t *testing.T) {
	t.Run("get team", func(t *testing.T) {
		factory := NewTeamFactory()
		if factory.TeamCache == nil {
			t.Errorf("team factory cache should not be nil\n")
		}

		team1 := factory.GetTeam(Team_A)
		if factory.GetTeamCacheLength() != 1 {
			t.Errorf("the length of team cache should be 1 but got: %d\n",
				factory.GetTeamCacheLength())
		}

		team2 := factory.GetTeam(Team_A)
		if factory.GetTeamCacheLength() != 1 {
			t.Errorf("cache length after the same team should not change, but got: %d\n",
				factory.GetTeamCacheLength())
		}

		if team1 != team2 {
			t.Errorf("query team with same id should return the same team pointer but not\n")
		}

		team3 := factory.GetTeam(Team_B)
		if factory.GetTeamCacheLength() != 2 {
			t.Errorf("cache length after the diff team should change to 2, but got: %d\n",
				factory.GetTeamCacheLength())
		}

		if team3 == team2 || team3 == team1 {
			t.Errorf("team with diff id should return diff team, but not\n")
		}

	})

	t.Run("get many team", func(t *testing.T) {
		factory := NewTeamFactory()
		teamBuffer := make([]*Team, 5000*2)
		for i := 0; i < 5000; i++ {
			teamBuffer[i] = factory.GetTeam(Team_A)
		}
		for i := 0; i < 2*5000; i++ {
			teamBuffer[i] = factory.GetTeam(Team_B)
		}

		if factory.GetTeamCacheLength() != 2 {
			t.Errorf("cache length after 2 diff team query should return 2 but got: %d\n",
				factory.GetTeamCacheLength())
		}

		for i := 0; i < 3; i++ {
			// addr should be equal
			log.Printf("addr: %p pointers locate at: %p\n", teamBuffer[i], &teamBuffer[i])
		}
	})
}
