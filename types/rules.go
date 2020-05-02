package types

type TypeRules struct {
	Start    int
	Limit    int
	Genres   []int
	PubYears []int
	Regions  []int
	OrderBy  string
}

type SearchRules struct {
	Start int
	Limit int
	Title string
	Genre int
	Region int
	PubYears []int
}
