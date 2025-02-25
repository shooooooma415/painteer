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
	Hokkaido:  "hokkaido",
	Aomori:    "aomori",
	Iwate:     "iwate",
	Miyagi:    "miyagi",
	Akita:     "akita",
	Yamagata:  "yamagata",
	Fukushima: "fukushima",
	Ibaraki:   "ibaraki",
	Tochigi:   "tochigi",
	Gunma:     "gunma",
	Saitama:   "saitama",
	Chiba:     "chiba",
	Tokyo:     "tokyo",
	Kanagawa:  "kanagawa",
	Niigata:   "niigata",
	Toyama:    "toyama",
	Ishikawa:  "ishikawa",
	Fukui:     "fukui",
	Yamanashi: "yamanashi",
	Nagano:    "nagano",
	Gifu:      "gifu",
	Shizuoka:  "shizuoka",
	Aichi:     "aichi",
	Mie:       "mie",
	Shiga:     "shiga",
	Kyoto:     "kyoto",
	Osaka:     "osaka",
	Hyogo:     "hyogo",
	Nara:      "nara",
	Wakayama:  "wakayama",
	Tottori:   "tottori",
	Shimane:   "shimane",
	Okayama:   "okayama",
	Hiroshima: "hiroshima",
	Yamaguchi: "yamaguchi",
	Tokushima: "tokushima",
	Kagawa:    "kagawa",
	Ehime:     "ehime",
	Kochi:     "kochi",
	Fukuoka:   "fukuoka",
	Saga:      "saga",
	Nagasaki:  "nagasaki",
	Kumamoto:  "kumamoto",
	Oita:      "oita",
	Miyazaki:  "miyazaki",
	Kagoshima: "kagoshima",
	Okinawa:   "okinawa",
}


var RegionMap = map[string][]string{
	"hokkaido":       {"hokkaido"},
	"tohoku":         {"aomori", "iwate", "miyagi", "akita", "yamagata", "fukushima"},
	"kanto":          {"ibaraki", "tochigi", "gunma", "saitama", "chiba", "tokyo", "kanagawa"},
	"chubu":          {"niigata", "toyama", "ishikawa", "fukui", "yamanashi", "nagano", "gifu", "shizuoka", "aichi"},
	"kinki":          {"mie", "shiga", "kyoto", "osaka", "hyogo", "nara", "wakayama"},
	"chugoku":        {"tottori", "shimane", "okayama", "hiroshima", "yamaguchi"},
	"shikoku":        {"tokushima", "kagawa", "ehime", "kochi"},
	"kyushu_okinawa": {"fukuoka", "saga", "nagasaki", "kumamoto", "oita", "miyazaki", "kagoshima", "okinawa"},
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

var PrefectureNameToId = map[string]PrefectureId{
	"hokkaido":  Hokkaido,
	"aomori":    Aomori,
	"iwate":     Iwate,
	"miyagi":    Miyagi,
	"akita":     Akita,
	"yamagata":  Yamagata,
	"fukushima": Fukushima,
	"ibaraki":   Ibaraki,
	"tochigi":   Tochigi,
	"gunma":     Gunma,
	"saitama":   Saitama,
	"chiba":     Chiba,
	"tokyo":     Tokyo,
	"kanagawa":  Kanagawa,
	"niigata":   Niigata,
	"toyama":    Toyama,
	"ishikawa":  Ishikawa,
	"fukui":     Fukui,
	"yamanashi": Yamanashi,
	"nagano":    Nagano,
	"gifu":      Gifu,
	"shizuoka":  Shizuoka,
	"aichi":     Aichi,
	"mie":       Mie,
	"shiga":     Shiga,
	"kyoto":     Kyoto,
	"osaka":     Osaka,
	"hyogo":     Hyogo,
	"nara":      Nara,
	"wakayama":  Wakayama,
	"tottori":   Tottori,
	"shimane":   Shimane,
	"okayama":   Okayama,
	"hiroshima": Hiroshima,
	"yamaguchi": Yamaguchi,
	"tokushima": Tokushima,
	"kagawa":    Kagawa,
	"ehime":     Ehime,
	"kochi":     Kochi,
	"fukuoka":   Fukuoka,
	"saga":      Saga,
	"nagasaki":  Nagasaki,
	"kumamoto":  Kumamoto,
	"oita":      Oita,
	"miyazaki":  Miyazaki,
	"kagoshima": Kagoshima,
	"okinawa":   Okinawa,
}
