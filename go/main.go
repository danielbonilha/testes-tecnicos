package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	timeConversion("12:05:45AM")
}

func timeConversion(s string) string {
	suffix := s[:]
	hourStr := s[:2]
	if suffix == "AM" {
		if hourStr == "12" {
			hourStr = "00"
		}
		return hourStr + s[2:len(s)-2]
	}
	hour, _ := strconv.Atoi(hourStr)
	if hour != 12 {
		hour += 12
	}
	hourStr = strconv.Itoa(hour)
	fmt.Println(hourStr + s[2:len(s)-2])
	return hourStr + s[2:len(s)-2]
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}

func birthdayCakeCandles(candles []int32) int32 {
	if len(candles) == 0 {
		return 0
	}
	if len(candles) == 1 {
		return 1
	}

	sort.Slice(candles, func(i, j int) bool { return candles[i] < candles[j] })
	lastIndex := len(candles) - 1
	lastCandle := candles[lastIndex]

	count := int32(0)
	for i := lastIndex; candles[i] == lastCandle; i-- {
		count += 1
	}
	return count
}

func solution(queries [][]string) []string {
	db := NewDB()
	result := make([]string, 0)

	for _, q := range queries {
		if q[0] == "SET_OR_INC" {
			r := db.Set(q[1], q[2], q[3])
			result = append(result, r)
			continue
		}

		if q[0] == "GET" {
			r := db.Get(q[1], q[2])
			result = append(result, r)
			continue
		}

		if q[0] == "DELETE" {
			r := db.Delete(q[1], q[2])
			result = append(result, strconv.FormatBool(r))
		}

		if q[0] == "TOP_N_KEYS" {
			r := db.TopKeys(q[1])
			result = append(result, r)
		}

	}
	return result
}

type (
	DB interface {
		Get(key, field string) string
		Set(key, field, value string) string
		Delete(key, field string) bool
		TopKeys(n string) string
	}

	InMemoryDB struct {
		m     map[string]map[string]string
		stats map[string]int
	}
)

func NewDB() DB {
	return &InMemoryDB{
		m:     map[string]map[string]string{},
		stats: map[string]int{}}
}

func (db *InMemoryDB) Get(key, field string) string {
	record, ok := db.m[key]
	if !ok {
		return ""
	}
	value, ok := record[field]
	if !ok {
		return ""
	}

	return value
}

func (db *InMemoryDB) Set(key, field, value string) string {
	db.addStats(key)
	record, ok := db.m[key]
	if !ok {
		db.m[key] = map[string]string{
			field: value,
		}
		return value
	}

	vStr, ok := record[field]
	if !ok {
		db.m[key][field] = value
		return value
	}

	v, _ := strconv.Atoi(vStr)
	valueInt, _ := strconv.Atoi(value)
	total := strconv.Itoa(v + valueInt)
	db.m[key][field] = total
	return total

}

func (db *InMemoryDB) Delete(key, field string) bool {
	record, ok := db.m[key]
	if !ok {
		return false
	}

	_, ok = record[field]
	if !ok {
		return false
	}

	db.addStats(key)
	delete(record, field)
	if len(record) == 0 {
		delete(db.m, key)
		delete(db.stats, key)
	}
	return true
}

func (db *InMemoryDB) TopKeys(n string) string {
	var result string

	m := db.sortMap()

	i := 0
	nInt, _ := strconv.Atoi(n)
	for k, v := range m {
		if i == nInt {
			break
		}
		result = result + k + "(" + strconv.Itoa(v) + "), "
		i++
	}
	result = result[:len(result)-2]
	return result
}

func (db *InMemoryDB) addStats(key string) {
	stats, ok := db.stats[key]
	if !ok {
		db.stats[key] = 1
	} else {
		db.stats[key] = stats + 1
	}
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }

func (db *InMemoryDB) sortMap() map[string]int {
	p := make(PairList, len(db.stats))
	result := make(map[string]int)

	i := 0
	for k, v := range db.stats {
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(sort.Reverse(p))
	for _, k := range p {
		result[k.Key] = k.Value
	}
	db.stats = result
	return result
}
