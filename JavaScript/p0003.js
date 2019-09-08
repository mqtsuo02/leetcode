/**
 * @param {string} s
 * @return {number}
 */

const lengthOfLongestSubstring = s => {
  let answer = 0;
  let count = 0;
  let hash = {};

  if (s.length < 2) {
    return s.length;
  }

  for (let i = 0; i < s.length; i++) {
    if (hash[s[i]] == null) {
      count += 1;
    } else {
      count = Math.min(i - hash[s[i]], count + 1);
    }
    answer = Math.max(answer, count);
    hash[s[i]] = i;
  }
  return answer;
};
