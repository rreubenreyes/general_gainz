package training

type Tier string

const (
	T1 Tier = "t1"
	T2 Tier = "t2"
	T3 Tier = "t3"
)

type Decision string

const (
	Hold Decision = "hold"
	Push Decision = "push"
)

type Exercise struct {
	Name   string `json:"name"`
	Tier   Tier   `json:"tier"`
	Weight int    `json:"weight"`
	Reps   int    `json:"reps"`
	Flavor Flavor `json:"flavor"`
}

type FollowUp struct {
	Sets int
}

type Flavor struct {
	FollowUp  FollowUp
	Decision  Decision
	Extension int
}

