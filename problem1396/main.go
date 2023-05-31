package problem1396

import (
	"fmt"
	"log"
	"strconv"
)

type tmpR struct {
	stationName string
	time        int
}

type UndergroundSystem struct {
	record    map[string]map[string][]float64
	tmpRecord map[int]tmpR
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		record:    map[string]map[string][]float64{},
		tmpRecord: map[int]tmpR{},
	}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	us.tmpRecord[id] = tmpR{
		stationName: stationName,
		time:        t,
	}
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	r, ok := us.tmpRecord[id]
	if !ok {
		log.Fatalln("error")
	}

	if _, ok := us.record[r.stationName]; !ok {
		us.record[r.stationName] = make(map[string][]float64)
	}

	tmp, ok := us.record[r.stationName][stationName]
	if !ok {
		us.record[r.stationName][stationName] = []float64{float64(t - r.time)}
	} else {
		us.record[r.stationName][stationName] = append(tmp, float64(t-r.time))
	}

	delete(us.tmpRecord, id)
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	record := us.record[startStation][endStation]

	var sum float64
	for _, v := range record {
		sum += v
	}

	ans, _ := strconv.ParseFloat(fmt.Sprintf("%.5f", sum/float64(len(record))), 64)
	return ans
}

/**
 * Your UndergroundSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.CheckIn(id,stationName,t);
 * obj.CheckOut(id,stationName,t);
 * param_3 := obj.GetAverageTime(startStation,endStation);
 */
