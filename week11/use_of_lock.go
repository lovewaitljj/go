package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type RankItem struct {
	Name    string
	FatRate float64
}

type FatRateRank struct {
	// 有序存储所有RankItem
	items []RankItem
	sync.Mutex
}

func (r *FatRateRank) updateRecord(name string, fatRate float64) int {
	r.Lock()
	defer r.Unlock()
	currentItem := RankItem{
		Name:    name,
		FatRate: fatRate,
	}

	// 一次遍历，查找先前idx以供删除，和将要插入的idx
	previousIdx, currentIdx := -1, -1
	for idx, item := range r.items {
		if previousIdx >= 0 && currentIdx >= 0 {
			break
		}

		if item.Name == name {
			previousIdx = idx
		}

		if currentIdx < 0 && item.FatRate > fatRate {
			// 只取第一次大于当前体脂率的idx
			currentIdx = idx
		}
	}
	if previousIdx >= 0 {
		// 如果找到了有先前的，就删掉它
		r.items = append(r.items[:previousIdx], r.items[previousIdx+1:]...)

		// 如果删除的idx在即将插入的之前，则将即将插入的idx-1
		if previousIdx < currentIdx {
			currentIdx -= 1
		}
	}

	if currentIdx == -1 {
		r.items = append(r.items, currentItem)
		currentIdx = len(r.items) - 1
	} else {
		r.items = insert(r.items, currentIdx, currentItem)
	}

	return currentIdx
}

func insert(a []RankItem, index int, value RankItem) []RankItem {
	last := len(a) - 1
	a = append(a, a[last])
	copy(a[index+1:], a[index:last])
	a[index] = value
	return a
}
func main() {
	wg := sync.WaitGroup{}
	r := &FatRateRank{}
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		i := i
		go func() {
			defer wg.Done()
			p := &Person{
				name:        fmt.Sprint(i),
				BaseFatRate: rand.Float64() * 0.4,
			}

			// 无限循环（不停地去更新自己的体脂信息）
			for {
				if err := p.changeFatRate(); err == nil {
					rank := r.updateRecord(p.name, p.CurrentFatRate)
					fmt.Println(p.name, p.CurrentFatRate, rank)
				} else {
					fmt.Println(p.name, err)
				}
				time.Sleep(time.Second)
			}
		}()
	}

	wg.Wait()
}

type Person struct {
	name string

	BaseFatRate    float64
	CurrentFatRate float64
}

func (p *Person) changeFatRate() error {
	min := -0.2
	max := 0.2
	delta := rand.Float64()*(max-min) + min
	result := p.BaseFatRate + delta

	if result < 0 || result > 0.4 {
		// 超出范围返回错误
		return fmt.Errorf("invalid fatrate: %f", result)
	} else {
		p.CurrentFatRate = result
		return nil
	}
}