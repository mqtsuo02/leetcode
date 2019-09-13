/**
 * @param {string[]} strs
 * @return {string}
 */

const longestCommonPrefix = function(strs) {
  if (!strs.length) return "";
  if (strs.length === 1) return strs[0];
  let prefix = "";
  let c = "";
  for (i = 0; i < strs[0].length; i += 1) {
    c = strs[0][i];
    for (j = 1; j < strs.length; j += 1) {
      if (c !== strs[j][i]) return prefix;
    }
    prefix += c;
  }
  return prefix;
};
