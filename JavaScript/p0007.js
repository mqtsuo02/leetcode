// prettier-ignore

/**
 * @param {number} x
 * @return {number}
 */

const reverse = x => {
  const n = (x < 0 ? -1 : 1) * Number(String(Math.abs(x)).split("").reverse().join(""));
  return 2 ** 31 > n && n >= -1 * 2 ** 31 ? n : 0;
};
