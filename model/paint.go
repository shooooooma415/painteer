package model


const (
	// 都道府県の ID を定義
	Hokkaido    PrefectureId = 1
	Aomori      PrefectureId = 2
	Iwate       PrefectureId = 3
	Miyagi      PrefectureId = 4
	Akita       PrefectureId = 5
	Yamagata    PrefectureId = 6
	Fukushima   PrefectureId = 7
	Ibaraki     PrefectureId = 8
	Tochigi     PrefectureId = 9
	Gunma       PrefectureId = 10
	Saitama     PrefectureId = 11
	Chiba       PrefectureId = 12
	Tokyo       PrefectureId = 13
	Kanagawa    PrefectureId = 14
	Niigata     PrefectureId = 15
	Toyama      PrefectureId = 16
	Ishikawa    PrefectureId = 17
	Fukui       PrefectureId = 18
	Yamanashi   PrefectureId = 19
	Nagano      PrefectureId = 20
	Gifu        PrefectureId = 21
	Shizuoka    PrefectureId = 22
	Aichi       PrefectureId = 23
	Mie         PrefectureId = 24
	Shiga       PrefectureId = 25
	Kyoto       PrefectureId = 26
	Osaka       PrefectureId = 27
	Hyogo       PrefectureId = 28
	Nara        PrefectureId = 29
	Wakayama    PrefectureId = 30
	Tottori     PrefectureId = 31
	Shimane     PrefectureId = 32
	Okayama     PrefectureId = 33
	Hiroshima   PrefectureId = 34
	Yamaguchi   PrefectureId = 35
	Tokushima   PrefectureId = 36
	Kagawa      PrefectureId = 37
	Ehime       PrefectureId = 38
	Kochi       PrefectureId = 39
	Fukuoka     PrefectureId = 40
	Saga        PrefectureId = 41
	Nagasaki    PrefectureId = 42
	Kumamoto    PrefectureId = 43
	Oita        PrefectureId = 44
	Miyazaki    PrefectureId = 45
	Kagoshima   PrefectureId = 46
	Okinawa     PrefectureId = 47
)

type PostsByPrefecture struct {
	PrefectureId PrefectureId
	PostIds    []PostId
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
