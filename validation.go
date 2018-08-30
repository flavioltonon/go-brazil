package brazil

type Validation struct {
	Valid  bool  `json:"valid"`
	Reason error `json:"reason"`
}
