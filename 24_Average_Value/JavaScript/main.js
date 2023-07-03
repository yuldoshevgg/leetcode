function averageValue(nums) {
  nums = nums.filter(i => i % 6 === 0)
  
  return nums.length ? parseInt(nums.reduce((a, b) => a + b) / nums.length) : 0
}

console.log(averageValue([1, 3, 6, 10, 12, 15]));