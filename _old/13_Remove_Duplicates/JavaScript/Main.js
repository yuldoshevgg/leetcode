let removeDuplicates = function (nums) {
  let element = nums[0]
  let l = 1

  for (let i = 1; i < nums.length; i++) {
    if (nums[i] !== element) {
      nums[l] = nums[i]
      l++
    }

    element = nums[i]
  }

  return l
}

console.log(removeDuplicates([1, 1, 2]));
