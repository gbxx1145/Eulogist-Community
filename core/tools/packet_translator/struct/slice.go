package packet_translate_struct

// 将切片 from([]From) 转换为 []To。
// converter 是用于转换切片内的单个元素的函数
func ConvertSlice[From any, To any](
	from []From,
	converter func(from From) To,
) []To {
	to := make([]To, len(from))
	for index, value := range from {
		to[index] = converter(value)
	}
	return to
}
