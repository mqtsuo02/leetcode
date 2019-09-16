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
  let stash = [];
  for (let i = 0; i < s.length; i += 1) {
    if (Object.keys(cs).includes(s[i])) {
      stash.push(cs[s[i]]);
      continue;
    }
    if (stash.length && s[i] === stash[stash.length - 1]) {
      stash.pop();
      continue;
    }
    return false;
  }
  return stash.length ? false : true;
};
