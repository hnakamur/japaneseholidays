package japaneseholidays

import (
	"strings"
	"testing"
	"time"
)

func Test2002(t *testing.T) {
	checkHolidays(t, 2002, time.September,
		"2002-09-15:敬老の日, 2002-09-16:振替休日, 2002-09-23:秋分の日")
}

func Test2003(t *testing.T) {
	checkHolidays(t, 2003, time.April,
		"2003-04-29:みどりの日")
	checkHolidays(t, 2003, time.May,
		"2003-05-03:憲法記念日, 2003-05-05:こどもの日")
	checkHolidays(t, 2003, time.September,
		"2003-09-15:敬老の日, 2003-09-23:秋分の日")
}

func Test2004(t *testing.T) {
	checkHolidays(t, 2004, time.September,
		"2004-09-20:敬老の日, 2004-09-23:秋分の日")
}

func Test2005(t *testing.T) {
	checkHolidays(t, 2005, time.April,
		"2005-04-29:みどりの日")
	checkHolidays(t, 2005, time.May,
		"2005-05-03:憲法記念日, 2005-05-04:国民の休日, 2005-05-05:こどもの日")
}

func Test2006(t *testing.T) {
	checkHolidays(t, 2006, time.April,
		"2006-04-29:みどりの日")
	checkHolidays(t, 2006, time.May,
		"2006-05-03:憲法記念日, 2006-05-04:国民の休日, 2006-05-05:こどもの日")
}

func Test2007(t *testing.T) {
	checkHolidays(t, 2007, time.April,
		"2007-04-29:昭和の日, 2007-04-30:振替休日")
	checkHolidays(t, 2007, time.May,
		"2007-05-03:憲法記念日, 2007-05-04:みどりの日, 2007-05-05:こどもの日")
}

func Test2008(t *testing.T) {
	checkHolidays(t, 2008, time.April,
		"2008-04-29:昭和の日")
	checkHolidays(t, 2008, time.May,
		"2008-05-03:憲法記念日, 2008-05-04:みどりの日, 2008-05-05:こどもの日, 2008-05-06:振替休日")
}

func Test2009(t *testing.T) {
	checkHolidays(t, 2009, time.May,
		"2009-05-03:憲法記念日, 2009-05-04:みどりの日, 2009-05-05:こどもの日, 2009-05-06:振替休日")
	checkHolidays(t, 2009, time.September,
		"2009-09-21:敬老の日, 2009-09-22:国民の休日, 2009-09-23:秋分の日")
}

func Test2012(t *testing.T) {
	checkHolidays(t, 2012, time.December,
		"2012-12-23:天皇誕生日, 2012-12-24:振替休日")
}

func Test2013(t *testing.T) {
	checkHolidays(t, 2013, time.January,
		"2013-01-01:元日, 2013-01-14:成人の日")
	checkHolidays(t, 2013, time.February,
		"2013-02-11:建国記念の日")
	checkHolidays(t, 2013, time.March,
		"2013-03-20:春分の日")
	checkHolidays(t, 2013, time.April,
		"2013-04-29:昭和の日")
	checkHolidays(t, 2013, time.May,
		"2013-05-03:憲法記念日, 2013-05-04:みどりの日, 2013-05-05:こどもの日, 2013-05-06:振替休日")
	checkHolidays(t, 2013, time.June,
		"")
	checkHolidays(t, 2013, time.July,
		"2013-07-15:海の日")
	checkHolidays(t, 2013, time.August,
		"")
	checkHolidays(t, 2013, time.September,
		"2013-09-16:敬老の日, 2013-09-23:秋分の日")
	checkHolidays(t, 2013, time.October,
		"2013-10-14:体育の日")
	checkHolidays(t, 2013, time.November,
		"2013-11-03:文化の日, 2013-11-04:振替休日, 2013-11-23:勤労感謝の日")
	checkHolidays(t, 2013, time.December,
		"2013-12-23:天皇誕生日")
}

func checkHolidays(t *testing.T, year int, month time.Month, expected string) {
	holidays := GetHolidays(year, month)
	holidaysStr := make([]string, len(holidays))
	for i, holiday := range holidays {
		holidaysStr[i] = holiday.String()
	}
	joinedGot := strings.Join(holidaysStr, ", ")
	if joinedGot != expected {
		t.Errorf("%d-%02d\ngot     : %s\nexpected: %s", year, month, joinedGot, expected)
	}
}

func TestSearchHolidayName(t *testing.T) {
	holidays := GetHolidays(2013, time.November)
	tests := []struct {
		date     time.Time
		expected string
	}{
		{time.Date(2013, time.November, 3, 0, 0, 0, 0, JST), "文化の日"},
		{time.Date(2013, time.November, 4, 1, 2, 3, 4, JST), "振替休日"},
		{time.Date(2013, time.November, 23, 23, 59, 0, 0, JST), "勤労感謝の日"},
		{time.Date(2013, time.November, 24, 23, 59, 0, 0, JST), ""},
	}
	for _, test := range tests {
		name := SearchHolidayName(holidays, test.date)
		if name != test.expected {
			t.Errorf("time=%v, got=%s, expected=%s", test.date, name, test.expected)
		}
	}

}

// 2000年から2030年までの範囲で春分の日と秋分の日の計算が合うかテスト。
// 参考文献： 何年後かの春分の日・秋分の日はわかるの？
// http://www.nao.ac.jp/QA/faq/a0301.html
// http://www.asahi-net.or.jp/~CI5M-NMR/misc/equinox.html
func TestShunbunShuubunDays(t *testing.T) {
	tests := []struct {
		year       int
		shunbunDay int
		shuubunDay int
	}{
		{1948, 21, 23},
		{1949, 21, 23},
		{1950, 21, 23},
		{1951, 21, 24},
		{1952, 21, 23},
		{1953, 21, 23},
		{1954, 21, 23},
		{1955, 21, 24},
		{1956, 21, 23},
		{1957, 21, 23},
		{1958, 21, 23},
		{1959, 21, 24},
		{1960, 20, 23},
		{1961, 21, 23},
		{1962, 21, 23},
		{1963, 21, 24},
		{1964, 20, 23},
		{1965, 21, 23},
		{1966, 21, 23},
		{1967, 21, 24},
		{1968, 20, 23},
		{1969, 21, 23},
		{1970, 21, 23},
		{1971, 21, 24},
		{1972, 20, 23},
		{1973, 21, 23},
		{1974, 21, 23},
		{1975, 21, 24},
		{1976, 20, 23},
		{1977, 21, 23},
		{1978, 21, 23},
		{1979, 21, 24},
		{1980, 20, 23},
		{1981, 21, 23},
		{1982, 21, 23},
		{1983, 21, 23},
		{1984, 20, 23},
		{1985, 21, 23},
		{1986, 21, 23},
		{1987, 21, 23},
		{1988, 20, 23},
		{1989, 21, 23},
		{1990, 21, 23},
		{1991, 21, 23},
		{1992, 20, 23},
		{1993, 20, 23},
		{1994, 21, 23},
		{1995, 21, 23},
		{1996, 20, 23},
		{1997, 20, 23},
		{1998, 21, 23},
		{1999, 21, 23},
		{2000, 20, 23},
		{2001, 20, 23},
		{2002, 21, 23},
		{2003, 21, 23},
		{2004, 20, 23},
		{2005, 20, 23},
		{2006, 21, 23},
		{2007, 21, 23},
		{2008, 20, 23},
		{2009, 20, 23},
		{2010, 21, 23},
		{2011, 21, 23},
		{2012, 20, 22},
		{2013, 20, 23},
		{2014, 21, 23},
		{2015, 21, 23},
		{2016, 20, 22},
		{2017, 20, 23},
		{2018, 21, 23},
		{2019, 21, 23},
		{2020, 20, 22},
		{2021, 20, 23},
		{2022, 21, 23},
		{2023, 21, 23},
		{2024, 20, 22},
		{2025, 20, 23},
		{2026, 20, 23},
		{2027, 21, 23},
		{2028, 20, 22},
		{2029, 20, 23},
		{2030, 20, 23},
	}
	for _, test := range tests {
		shunbunDay := shunbunDate(test.year).Day()
		if shunbunDay != test.shunbunDay {
			t.Errorf("shunbunDay: year=%d, got=%d, want=%d", test.year, shunbunDay, test.shunbunDay)
		}

		shuubunDay := shuubunDate(test.year).Day()
		if shuubunDay != test.shuubunDay {
			t.Errorf("shuubunDay: year=%d, got=%d, want=%d", test.year, shuubunDay, test.shuubunDay)
		}
	}
}
