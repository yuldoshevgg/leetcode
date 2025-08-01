let searchInsert = function (nums, target) {             // O(Log n) function 
  let start = 0
  let end = nums.length - 1

  while (start <= end) {
    let mid = Math.floor((start + end) / 2)
    if (nums[mid] === target) {
      return mid
    } 
    
    if (nums[mid] < target) {
      start = mid + 1
    } 
    
    if (nums[mid] > target) {
      end = mid - 1
    }
  }

  return start
}

console.log(searchInsert([1], 0));
