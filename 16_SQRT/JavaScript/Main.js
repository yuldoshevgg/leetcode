let mySqrt = function (x) {   
  let left = 0
  let right = x

  while (left <= right) {
    let mid = Math.floor((left + right) / 2)

    if (mid * mid <= x && (mid + 1) * (mid + 1) > x) {
      return mid
    } else if (mid * mid < x) {
      left = mid + 1
    } else {
      right = mid - 1
    }
  }

  return left
}

console.log(mySqrt(10));
