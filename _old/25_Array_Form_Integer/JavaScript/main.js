function addToArrayForm(num, k) {

  return (BigInt(num.join("")) + BigInt(k)).toString().split("") 
}

console.log(addToArrayForm([1, 2, 0, 0], 34));