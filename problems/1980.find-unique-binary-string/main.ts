function findDifferentBinaryString(nums: string[]): string {
  const set = new Set<number>(nums.map((num) => Number.parseInt(num, 2)))

  const maxNum = Number.parseInt("".padStart(nums[0].length, "1"), 2)

  for (let i = maxNum; i >= 0; i--) {
    if (!set.has(i)) {
      return i.toString(2).padStart(nums[0].length, "0")
    }
  }
  return ""
};

console.log(findDifferentBinaryString(["01", "10"]))
console.log(findDifferentBinaryString(["111", "011", "001"]))
console.log(findDifferentBinaryString(["10", "11"]))
