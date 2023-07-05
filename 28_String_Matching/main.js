function stringMatching(words) {
  let result = []

  for (let w of words) {
    for (let i = 0; i < words.length; i++) {
      if (w === words[i]) {
        continue
      }

      if (w.includes(words[i]) && !result.includes(words[i])) {
        result.push(words[i])
      }
    }
  }

  return result
}

console.log(stringMatching(["leetcoder", "leetcode", "od", "hamlet", "am"]));