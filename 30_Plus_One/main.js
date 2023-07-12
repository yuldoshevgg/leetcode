function solution(digits) {
  let result = []
  let extra = 1
  let carry = 0

  for (let i = digits.length - 1; i >= 0; i--) {
    if (digits[i] + carry + extra == 10) {
      result.unshift(0)
      extra = 0
      carry = 1
    } else {
      digits[digits.length - 1] += extra
      digits[i] += carry
      result.unshift(digits[i])
      carry = 0
      extra = 0
    }
  }

  if (carry > 0) {
    result.unshift(carry)
  }

  return result
}

console.log(solution([9, 8, 7, 6, 5, 4, 3, 2, 1, 0]));