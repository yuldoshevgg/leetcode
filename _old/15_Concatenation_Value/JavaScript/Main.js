let findConcatenationValue = function (nums) {           
  let result = 0

  for (let i = 0; i < nums.length; i++) {
    if (nums.length > 1) {
      let value = String(nums[0]) + String(nums[nums.length - 1])
      result += Number(value)
      nums.shift()
      nums.pop()
      i--
    } else {
      result += nums[i]
    }
  }

  return result
}

console.log(findConcatenationValue([5, 14, 13, 8, 12]));
