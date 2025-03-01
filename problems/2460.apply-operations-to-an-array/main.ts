function applyOperations(nums: number[]): number[] {
  const result = new Array(nums.length).fill(0)

  let start = 0

  for (let i = 0; i < nums.length - 1; i++) {
    if (nums[i] === nums[i + 1]) {
      nums[i] *= 2
      nums[i + 1] = 0
    }

    if (nums[i] !== 0) {
      result[start] = nums[i]
      start++
    }
  }

  result[start] = nums[nums.length - 1]

  return result
};

console.log(applyOperations([]))
