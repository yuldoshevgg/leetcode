function containsDuplicateJS(nums) {
  for (let i = 0; i < nums.length; i++) {
    for (let j = i + 1; j < nums.length; j++) {
      if (nums[i] === nums[j]) {
        return true
      }
    }
  }

  return false
}

console.log(containsDuplicateJS([1,1,1,3,3,4,3,2,4,2]));