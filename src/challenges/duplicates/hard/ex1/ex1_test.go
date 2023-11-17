package ex1_test

import (
	"fmt"
	"sort"
	"sync"
)

type DB struct {
	Data map[int]int // 數字: count
	mu   sync.Mutex
}

func (d *DB) Add(key int) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if _, exists := d.Data[key]; exists {
		d.Data[key]++
	} else {
		d.Data[key] = 1
	}
}

// 能模擬一個大資料的處理過程(多個輸入、多個處理器，共用單一通道) -> 可以再想一個多通道的分配範例
// 輸入: 接收到文本內容時，切割文本內容後，依序向通道傳送資料
// 處理器: 隨時對管道進行監控，當有內容出現在管道之中，會將db的該資料其計數+1，表示出現了幾次
// 收尾:
// 輸入完全結束、且須等待所有處理器完成，顯示出db之中有重複的資料
func Example_ex1() {
	const chunk = 20            // 模擬溝通的通道大小
	ch := make(chan int, chunk) // 這邊模擬多輸入，多處理器大家都共用單管道，您也可以練習一個多管道的類型

	db := new(DB)
	db.Data = make(map[int]int)

	// handler
	const numWorkers = 4
	wgWorker := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wgWorker.Add(1)
		go func(workID int) {
			defer func() {
				wgWorker.Done()
			}()
			for {
				n, open := <-ch
				// log.Printf("workerID: %04d get: %d\n", workID, n)
				if !open {
					// log.Printf("chan closed. worker: %04d Stop\n", workID)
					return
				}
				db.Add(n)
			}
		}(i)
	}

	// randObj := rand.New(rand.NewSource(time.Now().Unix()))
	wgInput := sync.WaitGroup{}

	allData := [][]byte{ // 模擬大資料，這邊只是為了測試所以我們把它全部寫出來，實際上底下有模擬一次讀一點的情況
		{100, 233, 81, 120, 14, 45, 251, 208},
		{65, 246, 85, 225, 172, 26, 52, 14},
		{153, 106, 227, 226, 215, 148, 243, 27},
		{103, 249, 28, 111, 159, 73, 58, 15},
		{93, 129, 227, 176, 160, 74, 102, 139},
	}

	for i := 0; i < 5; i++ { // 假設這每一個i都是要讀一個檔案
		// 這種情況就可以模擬很巨量的資料，慢慢處理的過程

		wgInput.Add(1)
		// curData := make([]byte, 4*2)
		// _, _ = randObj.Read(curData) // 隨機產生資料，模擬每一個檔案的資料或者片段讀取檔案內容，存放於curData
		curData := allData[i] // 為了固定測試資料，我們用假資料替換

		/* 如果二進位資料需要特殊處理就可以考慮以下內容
		r := bytes.NewReader(curData)
		values := make([]uint32, len(curData)/4)
		if err := binary.Read(r, binary.BigEndian, values); err != nil {
			panic(err)
		}
		*/

		// log.Println(curData)

		// 將我們要處理的資料發送出去
		go func(inputID int, bs []byte) {
			// log.Printf("inputID: %04d assign Data: %v\n", inputID, bs)
			defer func() {
				wgInput.Done()
			}()
			for _, b := range bs {
				ch <- int(b)
			}
		}(i, curData)
	}

	// 等待所有輸入資料都送完
	wgInput.Wait()

	// 關閉與worker之間的溝通管道(確保每個worker都可以順利結束)
	close(ch)

	// 等待所有工作者結束工作
	wgWorker.Wait()

	// 查看統計結果
	// log.Println("有重複的數字:")

	// 排序: 為了讓輸出可以匹配我們預期的結果
	keys := make([]int, len(db.Data))
	for k, _ := range db.Data {
		keys = append(keys, k)
	}
	sort.Ints(keys) // 排序

	// for key, count := range db.Data { // 測試無序的，因此不能確保輸出
	for _, key := range keys {
		count := db.Data[key]
		if count > 1 {
			fmt.Println(key)
		}
	}

	// Output:
	// 14
	// 227
}
