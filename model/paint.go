package model

type RegionId int

const (
	// 都道府県の ID を定義
	Hokkaido  PrefectureId = 1
	Aomori    PrefectureId = 2
	Iwate     PrefectureId = 3
	Miyagi    PrefectureId = 4
	Akita     PrefectureId = 5
	Yamagata  PrefectureId = 6
	Fukushima PrefectureId = 7
	Ibaraki   PrefectureId = 8
	Tochigi   PrefectureId = 9
	Gunma     PrefectureId = 10
	Saitama   PrefectureId = 11
	Chiba     PrefectureId = 12
	Tokyo     PrefectureId = 13
	Kanagawa  PrefectureId = 14
	Niigata   PrefectureId = 15
	Toyama    PrefectureId = 16
	Ishikawa  PrefectureId = 17
	Fukui     PrefectureId = 18
	Yamanashi PrefectureId = 19
	Nagano    PrefectureId = 20
	Gifu      PrefectureId = 21
	Shizuoka  PrefectureId = 22
	Aichi     PrefectureId = 23
	Mie       PrefectureId = 24
	Shiga     PrefectureId = 25
	Kyoto     PrefectureId = 26
	Osaka     PrefectureId = 27
	Hyogo     PrefectureId = 28
	Nara      PrefectureId = 29
	Wakayama  PrefectureId = 30
	Tottori   PrefectureId = 31
	Shimane   PrefectureId = 32
	Okayama   PrefectureId = 33
	Hiroshima PrefectureId = 34
	Yamaguchi PrefectureId = 35
	Tokushima PrefectureId = 36
	Kagawa    PrefectureId = 37
	Ehime     PrefectureId = 38
	Kochi     PrefectureId = 39
	Fukuoka   PrefectureId = 40
	Saga      PrefectureId = 41
	Nagasaki  PrefectureId = 42
	Kumamoto  PrefectureId = 43
	Oita      PrefectureId = 44
	Miyazaki  PrefectureId = 45
	Kagoshima PrefectureId = 46
	Okinawa   PrefectureId = 47
)

var AllPrefectureIds = []PrefectureId{
	Hokkaido, Aomori, Iwate, Miyagi, Akita, Yamagata, Fukushima,
	Ibaraki, Tochigi, Gunma, Saitama, Chiba, Tokyo, Kanagawa,
	Niigata, Toyama, Ishikawa, Fukui, Yamanashi, Nagano, Gifu,
	Shizuoka, Aichi, Mie, Shiga, Kyoto, Osaka, Hyogo, Nara,
	Wakayama, Tottori, Shimane, Okayama, Hiroshima, Yamaguchi, Tokushima,
	Kagawa, Ehime, Kochi, Fukuoka, Saga, Nagasaki, Kumamoto,
	Oita, Miyazaki, Kagoshima, Okinawa,
}

var PrefectureNames = map[PrefectureId]string{
	Hokkaido:  "Hokkaido",
	Aomori:    "Aomori",
	Iwate:     "Iwate",
	Miyagi:    "Miyagi",
	Akita:     "Akita",
	Yamagata:  "Yamagata",
	Fukushima: "Fukushima",
	Ibaraki:   "Ibaraki",
	Tochigi:   "Tochigi",
	Gunma:     "Gunma",
	Saitama:   "Saitama",
	Chiba:     "Chiba",
	Tokyo:     "Tokyo",
	Kanagawa:  "Kanagawa",
	Niigata:   "Niigata",
	Toyama:    "Toyama",
	Ishikawa:  "Ishikawa",
	Fukui:     "Fukui",
	Yamanashi: "Yamanashi",
	Nagano:    "Nagano",
	Gifu:      "Gifu",
	Shizuoka:  "Shizuoka",
	Aichi:     "Aichi",
	Mie:       "Mie",
	Shiga:     "Shiga",
	Kyoto:     "Kyoto",
	Osaka:     "Osaka",
	Hyogo:     "Hyogo",
	Nara:      "Nara",
	Wakayama:  "Wakayama",
	Tottori:   "Tottori",
	Shimane:   "Shimane",
	Okayama:   "Okayama",
	Hiroshima: "Hiroshima",
	Yamaguchi: "Yamaguchi",
	Tokushima: "Tokushima",
	Kagawa:    "Kagawa",
	Ehime:     "Ehime",
	Kochi:     "Kochi",
	Fukuoka:   "Fukuoka",
	Saga:      "Saga",
	Nagasaki:  "Nagasaki",
	Kumamoto:  "Kumamoto",
	Oita:      "Oita",
	Miyazaki:  "Miyazaki",
	Kagoshima: "Kagoshima",
	Okinawa:   "Okinawa",
}

var RegionMap = map[string][]PrefectureId{
	"Hokkaido":       {Hokkaido},
	"Tohoku":         {Aomori, Iwate, Miyagi, Akita, Yamagata, Fukushima},
	"Kanto":          {Ibaraki, Tochigi, Gunma, Saitama, Chiba, Tokyo, Kanagawa},
	"Chubu":          {Niigata, Toyama, Ishikawa, Fukui, Yamanashi, Nagano, Gifu, Shizuoka, Aichi},
	"Kinki":          {Mie, Shiga, Kyoto, Osaka, Hyogo, Nara, Wakayama},
	"Chugoku":        {Tottori, Shimane, Okayama, Hiroshima, Yamaguchi},
	"Shikoku":        {Tokushima, Kagawa, Ehime, Kochi},
	"Kyushu/Okinawa": {Fukuoka, Saga, Nagasaki, Kumamoto, Oita, Miyazaki, Kagoshima, Okinawa},
}


func GetPrefectureName(prefectureId PrefectureId) string {
	name, exists := PrefectureNames[prefectureId]
	if !exists {
		return "Unknown"
	}
	return name
}

type PostsByPrefecture struct {
	PrefectureId PrefectureId
	PostIds      []PostId
}

type CountsByPrefecture struct {
	Prefecture string
	PostCount  int
}

type CountsByRegion struct {
	Region    string
	PostCount int
}
