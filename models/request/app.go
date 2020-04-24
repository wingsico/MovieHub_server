package request

type TypeInfo struct {
	Start    int    `form:"start"`
	Limit    int    `form:"limit"`
	Genres   string `form:"genres"`
	Regions  string `form:"regions"`
	PubYears string `form:"pub_years"`
	Sort     string `form:"sort"`
}

type RandomInfo struct {
	Start int    `form:"start"`
	Limit int    `form:"limit"`
	Seed  string `form:"seed"`
}

type IdsInfo struct {
	Ids []int `form:"ids"`
}

type SearchInfo struct {
	Start int    `form:"start"`
	Limit int    `form:"limit"`
	Title string `form:"q"`
}
