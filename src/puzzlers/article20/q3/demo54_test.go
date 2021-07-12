package q3

import "testing"

/*
这里我给出的典型回答是下面三个内容。
	对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明。
	对于性能测试函数来说，其名称必须以Benchmark为前缀，并且唯一参数的类型必须是*testing.B类型的。
	对于示例测试函数来说，其名称必须以Example为前缀，但对函数的参数列表没有强制规定。
*/
func BenchmarkGetPrimes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetPrimes(1000)
	}
}
