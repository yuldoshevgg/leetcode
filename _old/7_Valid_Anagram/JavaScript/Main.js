var validAnagram = function (s, t) {
  
  if (s.length !== t.length) {
    return false
  } else if (s.split("").sort().join() === t.split("").sort().join()) {
    return true
  }

  return false
};

console.log(validAnagram("race", "car"));