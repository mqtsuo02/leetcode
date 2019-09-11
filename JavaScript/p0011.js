/**
 * @param {number[]} heights
 * @return {number}
 */

const maxArea = heights => {
  if (heights.length < 1) return 0;

  let max = 0;
  for (let i = 0; i < heights.length; i += 1) {
    for (let j = 0; j < i; j += 1) {
      max = Math.max(max, Math.min(heights[i], heights[j]) * (i - j));
    }
  }
  return max;
};
