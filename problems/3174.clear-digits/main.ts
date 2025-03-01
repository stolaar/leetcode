function clearDigits(s: string): string {
  const alphaNums: string[] = []

  for (let i = 0; i < s.length; i++) {
    const charCode = s.charCodeAt(i)

    if (charCode >= 97 && charCode <= 122) {
      alphaNums.push(s[i])
      continue
    }
    alphaNums.pop()
  }

  return alphaNums.join("")
};

console.log("Result", clearDigits("cb34"))
console.log("Result", clearDigits("abc"))
console.log("Result", clearDigits("a1b2b3c"))
console.log("Result", clearDigits("ag3"))
