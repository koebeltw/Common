package orderArrayManager

// func Benchmark_OAM8(b *testing.B) {
// 	b.StopTimer() //調用該函數停止壓力測試的時間計數
// 	oam8 := NewOMM8(255)
// 	// //做一些初始化的工作,例如讀取文件數據,數據庫連接之類的,
// 	// //這樣這些時間不影響我們測試函數本身的性能

// 	b.StartTimer() //重新開始時間
// 	for i := 0; i < b.N; i++ {
// 		var randNum int
// 		rand.Seed(time.Now().Unix())
// 		count := rand.Intn(255)
// 		for j := 0; j < count; j++ {
// 			oam8.Add(o8{})

// 			randNum = rand.Intn(255)
// 			oam8.Get(uint8(randNum))
// 		}

// 		count = rand.Intn(255)
// 		for j := 0; j < count; j++ {
// 			randNum = rand.Intn(255)
// 			oam8.Del(uint8(randNum))

// 			randNum = rand.Intn(255)
// 			oam8.Get(uint8(randNum))
// 		}
// 	}
// }
