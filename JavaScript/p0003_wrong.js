/**
 * @param {string} s
 * @return {number}
 */

const lengthOfLongestSubstring = s => {
  let stock = "";
  let longest = "";
  s.split("").forEach(c => {
    if (stock.indexOf(c) === -1) {
      stock += c;
      return;
    }
    if (stock.length > longest.length) {
      longest = stock;
    }
    stock = c;
  });
  return Math.max(longest.length, stock.length);
};
