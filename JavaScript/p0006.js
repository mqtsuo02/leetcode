/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */

var convert = function(s, numRows) {
  if (!s) return "";
  let map = [];
  let cursor = 0;
  let isUp = true;
  for (let n = 0; n < s.length; n += 1) {
    if (!map[cursor]) map[cursor] = "";
    map[cursor] += s[n];
    cursor = isUp ? cursor + 1 : cursor - 1;
    if (cursor % (numRows - 1) === 0) {
      isUp = !isUp;
    }
  }
  return map.reduce((a, b) => {
    return a + b;
  });
};
