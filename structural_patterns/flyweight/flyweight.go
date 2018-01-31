package flyweight

import "time"

// Team resource struct:

// Team(id, name, players, historyData)
// player(id, name, previousTeam, photo)
// historyData(year, leagueResults)
// match(date,
//       visitorId,
//       localId,
//       LocalScore,
//       VistorScore,
//       LocalShoots,
//       VisitorShoots)

const (
	Team_A = iota
	Team_B
)

type TeamFactory struct {
	TeamCache map[int]*Team
}

func (tf *TeamFactory) GetTeam(id int) *Team {
	if tf.TeamCache[id] != nil {
		return tf.TeamCache[id]
	}

	team := getTeamFactory(id)
	tf.TeamCache[id] = team
	return team
}

func (tf *TeamFactory) GetTeamCacheLength() int {
	return len(tf.TeamCache)
}

func NewTeamFactory() *TeamFactory {
	return &TeamFactory{TeamCache: make(map[int]*Team, 0)}
}

func getTeamFactory(id int) *Team {
	switch id {
	case Team_A:
		return &Team{Id: 1, Name: "Team_A"}
	case Team_B:
		return &Team{Id: 2, Name: "Team_B"}
	default:
		return &Team{Id: 1, Name: "Team_A"}
	}
}

type Team struct {
	Id          uint32
	Name        string
	Players     []Player
	HistoryData []HistoryData
}

type Player struct {
	Id           uint32
	Name         string
	PreviousTeam []Team
	Photo        []byte
}

type HistoryData struct {
	Year          uint8
	LeagueMatches []Match
}

type Match struct { // I'm not a football guy at all...
	Date         time.Time
	VisitorId    uint64
	LocalId      uint64
	VisitorScore byte
	LcoalScore   byte
	LocalShoots  uint16
	VistorShoots uint16
}
