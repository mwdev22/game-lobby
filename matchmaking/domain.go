package matchmaking

type GameMode int

const (
	Deathmatch GameMode = iota
	CaptureTheFlag
	TeamDeathmatch
)

type Region string

const (
	RegionNA Region = "North America"
	RegionEU Region = "Europe"
	RegionAS Region = "Asia"
)

// here is the place where u define all domain objects and interfaces used by application

type Player struct {
	ID        string
	Name      string
	SkillRank int
}

type Queue interface {
	GetPlayers() []*Player
	AddPlayer(player *Player) error
	RemovePlayer(playerID string) error
	PlayerCount() int
	Capacity() int
	OpenSlots() int
}

type Match struct {
	ID           string
	Players      []*Player
	GameMode     GameMode
	Region       Region
	SessionToken SessionToken
}

type SessionToken struct {
	Token     string
	ExpiresAt int64
}

type Notifier interface {
	MatchCreated(match *Match) error
}

type Publisher interface {
	Publish(topic string, message []byte) error
}

type MatchMaker interface {
	Create(players []*Player, mode GameMode, region Region) (*Match, error)
	FindForPlayer(player *Player, mode GameMode, region Region) (*Match, error)
	NewMatchNotification(match *Match) error
}
