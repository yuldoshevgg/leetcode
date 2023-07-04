function kidsWithCandies(candies, extraCandies) {
  let result = []
  let max = Math.max(...candies)

  for (let c of candies) {
    if (c + extraCandies >= max) {
      result.push(true)
      continue
    }

    result.push(false)
  }

  return result
}

console.log(kidsWithCandies([2, 3, 5, 1, 3], 3));