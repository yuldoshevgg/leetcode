function truncateSentence(s, k) {
  
  return s.split(" ").slice(0, k).join(" ")
}

console.log(truncateSentence("What is the solution to this problem", 4));