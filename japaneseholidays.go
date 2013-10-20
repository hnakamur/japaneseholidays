package japaneseholidays

import (
	"time"
)

// Japan Standard Time (UTC+9)
var JST = time.FixedZone("Asia/Tokyo", 9*60*60)

type Holiday struct {
	Date time.Time // Holiday date in JST timezone
	Name string    // Holiday name in Kanji
}

func (h Holiday) String() string {
	return h.Date.Format("2006-01-02") + ":" + h.Name
}

// Compare the year, month, and day of the date to holidays and returns
// the holiday name if matched or the empty string if not matched.
func SearchHolidayName(holidays []Holiday, date time.Time) string {
	y, m, d := date.In(JST).Date()
	for _, holiday := range holidays {
		hy, hm, hd := holiday.Date.Date()
		if y == hy && m == hm && d == hd {
			return holiday.Name
		}
	}
	return ""
}

// Returns holidays in the specified year and month.
func GetHolidays(year int, month time.Month) []Holiday {
	holidays := []Holiday{}
	switch month {
	case time.January:
		if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 1), "元日"})
		}

		if year >= 2000 {
			holidays = append(holidays,
				Holiday{nthWeekday(year, month, 2, time.Monday), "成人の日"})
		} else if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 15), "成人の日"})
		}
	case time.February:
		if year >= 1967 {
			holidays = append(holidays,
				Holiday{date(year, month, 11), "建国記念の日"})
		}
	case time.March:
		if year >= 1949 {
			holidays = append(holidays,
				Holiday{shunbunDate(year), "春分の日"})
		}
	case time.April:
		if year >= 2007 {
			holidays = append(holidays,
				Holiday{date(year, month, 29), "昭和の日"})
		} else if year >= 1989 {
			holidays = append(holidays,
				Holiday{date(year, month, 29), "みどりの日"})
		} else if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 29), "天皇誕生日"})
		}
	case time.May:
		if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 3), "憲法記念日"})
		}

		if year >= 2007 {
			holidays = append(holidays,
				Holiday{date(year, month, 4), "みどりの日"})
		} else if year >= 1988 {
			d := date(year, month, 4)
			wd := d.Weekday()
			if time.Tuesday <= wd && wd <= time.Friday {
				holidays = append(holidays,
					Holiday{d, "国民の休日"})
			}
		}

		if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 5), "こどもの日"})
		}
	case time.June:
		break
	case time.July:
		if year >= 2003 {
			holidays = append(holidays,
				Holiday{nthWeekday(year, month, 3, time.Monday), "海の日"})
		} else if year >= 1996 {
			holidays = append(holidays,
				Holiday{date(year, month, 20), "海の日"})
		}
	case time.August:
		break
	case time.September:
		var keirouD time.Time
		if year >= 2003 {
			keirouD = nthWeekday(year, month, 3, time.Monday)
			holidays = append(holidays,
				Holiday{keirouD, "敬老の日"})
		} else if year >= 1996 {
			holidays = append(holidays,
				Holiday{date(year, month, 15), "敬老の日"})
		}

		shuubunD := shuubunDate(year)
		if year >= 2003 {
			d := keirouD.AddDate(0, 0, 1)
			if d.AddDate(0, 0, 1) == shuubunD {
				wd := d.Weekday()
				if time.Tuesday <= wd && wd <= time.Friday {
					holidays = append(holidays,
						Holiday{d, "国民の休日"})
				}
			}
		}

		if year >= 1949 {
			holidays = append(holidays,
				Holiday{shuubunD, "秋分の日"})
		}
	case time.October:
		if year >= 2000 {
			holidays = append(holidays,
				Holiday{nthWeekday(year, month, 2, time.Monday), "体育の日"})
		} else if year >= 1949 {
			holidays = append(holidays,
				Holiday{date(year, month, 10), "体育の日"})
		}
	case time.November:
		if year >= 1948 {
			holidays = append(holidays,
				Holiday{date(year, month, 3), "文化の日"},
				Holiday{date(year, month, 23), "勤労感謝の日"},
			)
		}
	case time.December:
		if year >= 1989 {
			holidays = append(holidays,
				Holiday{date(year, month, 23), "天皇誕生日"})
		}
	}

	return addFurikaeKyuujitsu(holidays)
}

func addFurikaeKyuujitsu(holidays []Holiday) []Holiday {
	result := []Holiday{}
	for i := 0; i < len(holidays); i++ {
		holiday := holidays[i]
		result = append(result, holiday)
		if holiday.Date.Weekday() == time.Sunday {
			d := holiday.Date.AddDate(0, 0, 1)
			for j := i + 1; j < len(holidays); j++ {
				if d != holidays[j].Date {
					break
				}
				result = append(result, holidays[j])
				d = d.AddDate(0, 0, 1)
				i++
			}
			result = append(result, Holiday{d, "振替休日"})
		}
	}
	return result
}

func shunbunDate(year int) time.Time {
	var baseDate int
	if year <= 1999 {
		baseDate = 2213
	} else {
		baseDate = 2089
	}
	return date(year, time.March, shunbunShuubunDayHelper(year, baseDate))
}

func shuubunDate(year int) time.Time {
	var baseDate int
	if year <= 1999 {
		baseDate = 2525
	} else {
		baseDate = 2395
	}
	return date(year, time.September, shunbunShuubunDayHelper(year, baseDate))
}

// 指定した年の春分または秋分の日を天文学的に計算するためのヘルパ関数。
// @param year 年
// @param baseDate 計算に用いる基点の日
// @return 指定した年の春分または秋分の月内での日
//
// 参考文献： 将来の春分日・秋分日の計算
// http://blade.nagaokaut.ac.jp/cgi-bin/scat.rb/ruby/ruby-list/7112
//
// --(引用開始)-----------
//
// 年を y とした時、春分日（３月Ｘ日のＸ）、秋分日（９月Ｘ日のＸ）は以下
// の式で計算できます。ただし、除算の余りは切り捨てます。
//
// 春分日　(31y+2213)/128-y/4+y/100    (1851年-1999年通用)
// 　　　　(31y+2089)/128-y/4+y/100    (2000年-2150年通用)
//
// 秋分日　(31y+2525)/128-y/4+y/100    (1851年-1999年通用)
// 　　　　(31y+2395)/128-y/4+y/100    (2000年-2150年通用)
//
// --(引用終了)-----------
//
// 注：上の説明では1851年以降でないと計算できないようだが、実際は1948年以降で
//     無理矢理計算しても結果は以下のページ
//     http://www.asahi-net.or.jp/~CI5M-NMR/misc/equinox.html
//     で紹介されている理科年表の記録と一致した。春分の日が祝日になったのは
//     1949年、秋分の日は1948年なので、過去についてはこの式で問題ないことにな
//     る。2151年以降については、将来追加情報が得られたら計算式を更新すること。
func shunbunShuubunDayHelper(year, baseDate int) int {
	return (31*year+baseDate)/128 - year/4 + year/100
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, JST)
}

func nthWeekday(year int, month time.Month, nth int, weekday time.Weekday) time.Time {
	d := date(year, month, 1)
	for d.Weekday() != weekday {
		d = d.AddDate(0, 0, 1)
	}
	return d.AddDate(0, 0, (nth-1)*7)
}
