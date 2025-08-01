function countPrefixes(words, pref) {
  let result = 0

  for (let i of words) {
    if (i.startsWith(pref)) result++
  }

  return result
}

console.log(countPrefixes(["pay","attention","practice","attend"], "at"));