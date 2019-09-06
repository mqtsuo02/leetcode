/**
 * @param {string} s
 * @return {number}
 */

const transMap = {
  IV: "v",
  IX: "x",
  XL: "l",
  XC: "c",
  CD: "d",
  CM: "m",
};

const romanMap = {
  I: 1,
  v: 4,
  V: 5,
  x: 9,
  X: 10,
  l: 40,
  L: 50,
  c: 90,
  C: 100,
  d: 400,
  D: 500,
  m: 900,
  M: 1000,
};

const trans = s => {
  let target = s;
  Object.keys(transMap).forEach(key => {
    target = target.replace(key, transMap[key]);
  });
  return target;
};

const sum = s => {
  return s.split("").reduce((sum, c) => {
    return sum + romanMap[c];
  }, 0);
};

const romanToInt = s => {
  return sum(trans(s));
};
