function countPrefixes(words, s) {
  let result = 0

  for (let i of words) {
    if (s.startsWith(i)) result++
  }

  return result
}

console.log(countPrefixes(["a", "a"], "aa"));