// prettier-ignore

/**
 * @param {number} x
 * @return {boolean}
 */

const isPalindrome = x => {
  if (x < 0) {
    return false;
  }
  return x === Number(String(x).split("").reverse().join(""));
};
