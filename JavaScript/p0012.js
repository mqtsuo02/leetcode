/**
 * @param {number} num
 * @return {string}
 */

const romanMap = {
  1000: "M",
  900: "CM",
  500: "D",
  400: "CD",
  100: "C",
  90: "XC",
  50: "L",
  40: "XL",
  10: "X",
  9: "IX",
  5: "V",
  4: "IV",
  1: "I",
};

const valueArray = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

const intToRoman = num => {
  let n = num;
  let roman = "";
  valueArray.forEach(value => {
    const t = parseInt(n / value);
    for (let i = 0; i < t; i += 1) {
      roman += romanMap[value];
    }
    n -= value * t;
  });
  return roman;
};
