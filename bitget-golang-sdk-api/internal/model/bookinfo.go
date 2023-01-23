package model

import (
	"hash/crc32"
	"log"
	"sort"
	"strings"
)

type BookInfo struct {
	Asks     []interface{}
	Bids     []interface{}
	Checksum string
	Ts       string
}

func (b *BookInfo) Merge(updateInfo BookInfo) *BookInfo {
	b.Asks = b.InnerMerge(b.Asks, updateInfo.Asks, false)
	b.Bids = b.InnerMerge(b.Bids, updateInfo.Bids, true)
	return b
}

func (b *BookInfo) InnerMerge(allList []interface{}, updateList []interface{}, isReverse bool) []interface{} {
	priceAndValue := make(map[string][]interface{})
	for _, val := range allList {
		valArray, ok := val.([]interface{})
		if !ok {
			// Handle the case where val is not of the expected type
			continue
		}
		priceAndValue[valArray[0].(string)] = valArray
	}
	for _, val := range updateList {
		valArray, ok := val.([]interface{})
		if !ok {
			continue
		}
		if valArray[1].(string) == "0" {
			delete(priceAndValue, valArray[0].(string))
			continue
		}
		priceAndValue[valArray[0].(string)] = valArray
	}
	result := make([]interface{}, 0, len(priceAndValue))
	for _, value := range priceAndValue {
		result = append(result, value)
	}
	if isReverse {
		sort.Slice(result, func(i, j int) bool {
			value1, ok1 := result[i].([]string)
			value2, ok2 := result[j].([]string)
			if ok1 && ok2 {
				return value1[0] > value2[0]
			}
			return false
		})
	} else {
		sort.Slice(result, func(i, j int) bool {
			value1, ok1 := result[i].([]string)
			value2, ok2 := result[j].([]string)
			if ok1 && ok2 {
				return value1[0] < value2[0]
			}
			return false
		})
	}
	return result
}

func (b *BookInfo) CheckSum(checkSum uint32, gear int) bool {
	var sb strings.Builder
	for i := 0; i < gear; i++ {
		bids := b.Bids[i].([]string)
		sb.WriteString(bids[0])
		sb.WriteString(":")
		sb.WriteString(bids[1])
		sb.WriteString(":")
		asks := b.Asks[i].([]string)
		sb.WriteString(asks[0])
		sb.WriteString(":")
		sb.WriteString(asks[1])
		sb.WriteString(":")
	}

	s := sb.String()
	str := s[:len(s)-1]

	crc32 := crc32.NewIEEE()
	crc32.Write([]byte(str))
	value := uint32(crc32.Sum32())

	log.Println("check val:", str)
	log.Println("start checknum mergeVal:", value, ",checkVal:", checkSum)

	return value == checkSum
}
