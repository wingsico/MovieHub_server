package constants

var OrderMap = map[string]string{
	Latest: "(ifnull(pub_year * pow(10, 10), cast(release_date as unsigned)) + ifnull(cast(release_date as unsigned), 0)) desc", // 最新上映
	Highest: "(douban_rating + imdb_rating) desc", // 最高评分
	Hottest: "(douban_rating_count * douban_rating + imdb_rating_count * imdb_rating) desc", // 热度最高
	Trending: "ifnull(pub_year,0) desc, douban_rating_count * douban_rating + imdb_rating_count * imdb_rating desc", // Trending
}

var (
	Latest = "0"
	Highest = "1"
	Hottest = "2"
	Trending = "3"
)

const MinKeyCount int = 10

var (
	AccessKey = "p2_HSflxPP8l0w9-Ym_r0KS6or9nWenH_DL8QnQ5"
	SecretKey = "0jlSgVGke5-zCheC1Ht219HvoFpaMwq73GLUp6_h"
	Bucket = "images"
	Domain = "http://cdn.wingsico.org/"
)