/**
 * @param {string} s
 * @return {boolean}
 */

const cs = {
  "(": ")",
  "{": "}",
  "[": "]",
};

const isValid = s => {
  let stacks = [];
  for (let i = 0; i < s.length; i += 1) {
    if (Object.keys(cs).includes(s[i])) {
      stacks.push(cs[s[i]]);
      continue;
    }
    if (stacks.length && s[i] === stacks[stacks.length - 1]) {
      stacks.pop();
      continue;
    }
    return false;
  }
  return stacks.length ? false : true;
};
