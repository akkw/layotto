package sequencer

type Config struct {
	BiggerThan map[string]int64  `json:"biggerThan"`
	Metadata   map[string]string `json:"metadata"`
}
