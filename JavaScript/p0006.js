/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */

var convert = function(s, numRows) {
  if (!s) return "";
  let map = [];
  let cursol = 0;
  let isUp = true;
  for (let n = 0; n < s.length; n += 1) {
    if (!map[cursol]) map[cursol] = "";
    map[cursol] += s[n];
    cursol = isUp ? cursol + 1 : cursol - 1;
    if (cursol % (numRows - 1) === 0) {
      isUp = !isUp;
    }
  }
  return map.reduce((a, b) => {
    return a + b;
  });
};
