/**
 * @param {string} s
 * @return {string}
 */

const longestPalindrome = s => {
  if (s.length < 2) return s;
  let start = 0;
  let end = 0;
  for (let i = 0; i < s.length; i += 1) {
    const len1 = expandAroundCenter(s, i, i);
    const len2 = expandAroundCenter(s, i, i + 1);
    const len = Math.max(len1, len2);
    if (len > end + 1 - start) {
      start = i - parseInt((len - 1) / 2, 10);
      end = i + parseInt(len / 2, 10);
    }
  }
  return s.substring(start, end + 1);
};

const expandAroundCenter = (s, left, right) => {
  let l = left;
  let r = right;
  while (l >= 0 && r < s.length && s[l] === s[r]) {
    l--;
    r++;
  }
  return r - l - 1;
};
