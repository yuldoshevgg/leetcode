function missingNumber(nums) {
  let total = nums.length * (nums.length + 1) / 2
  let sum = 0

  for (let i = 0; i < nums.length; i++) {
    sum += nums[i]
  }

  return total - sum
}

console.log(missingNumber([3,0,1]));