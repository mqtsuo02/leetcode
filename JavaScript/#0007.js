/**
 * @param {number} x
 * @return {number}
 */

const reverse = x => {
  const s =  x < 0 ? -1 : 1;
  const n = s * Number(String(Math.abs(x)).split("").reverse().join(""));
  return 2 ** 31 > n && n >= -1 * 2 ** 31 ? n : 0;
};
