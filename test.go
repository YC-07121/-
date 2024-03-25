package main

//import 需要的套件
import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 宣告Meat的類別
type Meat struct {
	name   string //肉品名稱
	amount int    //肉品數量
}

// 讓員工依序處理channel裡的資料
// 傳入員工編號、存放肉的Channel、WaitGroup
func worker(id int, meatChan <-chan Meat, wg *sync.WaitGroup) {
	defer wg.Done()
	for meat := range meatChan {
		fmt.Printf("%c 在 %s 取得%s肉\n", 'A'+id, time.Now().Format("2006-01-02 15:04:05"), meat.name)
		var processingTime int //宣告處理時間
		switch meat.name {
		case "牛":
			processingTime = 1
		case "豬":
			processingTime = 2
		case "雞":
			processingTime = 3
		}
		time.Sleep(time.Duration(processingTime) * time.Second) //模擬處理時間
		fmt.Printf("%c 在 %s 處理完%s肉\n", 'A'+id, time.Now().Format("2006-01-02 15:04:05"), meat.name)
	}
}

func main() {
	//宣告肉品種類和數量
	meats := []Meat{
		{"牛", 10},
		{"豬", 7},
		{"雞", 5},
	}

	//宣告存放肉品的Channel
	meatChan := make(chan Meat)
	var wg sync.WaitGroup

	go func() {
		total := meats[0].amount + meats[1].amount + meats[2].amount //設定肉品總量

		for total > 0 {
			total = meats[0].amount + meats[1].amount + meats[2].amount

			//將隨機肉品放入肉品Channel
			meatType := rand.Intn(3)
			if meats[meatType].amount > 0 {
				meats[meatType].amount--
				meatChan <- meats[meatType]
			} else {
				continue
			}

		}
		close(meatChan)
	}()

	//傳入5位員工
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(i, meatChan, &wg)
	}

	wg.Wait()
}
