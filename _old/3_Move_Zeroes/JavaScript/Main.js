
var moveZeroes = function (nums) {
  let zeros = 0
  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === 0) {
      zeros++
      nums.splice(i, 1)
      i--
    }
  }
  
  for (let i = 0; i < zeros; i++) {
    nums.push(0)
  }

  console.log(nums);
};

moveZeroes([0, 0, 1])
// moveZeroes([0,1,0,3,12])