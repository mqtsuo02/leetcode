/**
 * @param {string} str
 * @return {number}
 */

const myAtoi = str => {
  let s = str
    .trim()
    .split(" ")[0]
    .match(/^[+-]?[0-9]+/);
  if (s === null) return 0;
  let n = Number(s);
  if (n > 2147483647) n = 2147483647;
  if (n < -2147483648) n = -2147483648;
  return n;
};
