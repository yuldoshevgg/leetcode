function mostWordsFound(sentences) {
  let max = 0

  for (let i of sentences) {
    let s = i.split(" ")
    if (s.length > max) {
      max = s.length
    }
  }

  return max
}

console.log(mostWordsFound(["alice and bob love leetcode", "i think so too", "this is great thanks very much"]));