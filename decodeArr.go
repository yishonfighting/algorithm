package main

//未知 整数数组 arr 由 n 个非负整数组成。
//
//经编码后变为长度为 n - 1 的另一个整数数组 encoded ，其中 encoded[i] = arr[i] XOR arr[i + 1] 。例如，arr = [1,0,2,1] 经编码后得到 encoded = [1,2,3] 。
//
//给你编码后的数组 encoded 和原数组 arr 的第一个元素 first（arr[0]）。
//
//请解码返回原数组 arr 。可以证明答案存在并且是唯一的。
//
// 
//
//示例 1：
//
//输入：encoded = [1,2,3], first = 1
//输出：[1,0,2,1]
//解释：若 arr = [1,0,2,1] ，那么 first = 1 且 encoded = [1 XOR 0, 0 XOR 2, 2 XOR 1] = [1,2,3]
//示例 2：
//
//输入：encoded = [6,2,7,3], first = 4
//输出：[4,2,0,7,4]




func decode(encoded []int, first int) []int {
	//原数组 \textit{arr}arr 的长度为 nn，对 \textit{arr}arr 编码后得到长度为 n-1n−1 的数组 \textit{encoded}encoded，编码规则为：\textit{encoded}[i]=\textit{arr}[i] \oplus \textit{arr}[i+1]encoded[i]=arr[i]⊕arr[i+1]，其中 \oplus⊕ 是异或运算符，0 \le i<n-10≤i<n−1。
	//
	//已知编码后的数组 \textit{encoded}encoded 和原数组 \textit{arr}arr 的第一个元素 \textit{arr}[0]=\textit{first}arr[0]=first，需要解码得到原数组 \textit{arr}arr。可以利用异或运算的性质实现。
	//
	//异或运算具有如下性质：
	//
	//异或运算满足交换律和结合律；
	//
	//任意整数和自身做异或运算的结果都等于 00，即 x \oplus x = 0x⊕x=0；
	//
	//任意整数和 00 做异或运算的结果都等于其自身，即 x \oplus 0 = 0 \oplus x = xx⊕0=0⊕x=x。
	rawArr := make([]int,len(encoded) + 1)
	rawArr[0] = first
	for i , val := range encoded {
		rawArr[i + 1] = rawArr[i] ^ val
	}

	return rawArr
}