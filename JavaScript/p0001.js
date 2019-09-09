/**
 * @param {number[]} ns
 * @param {number} t
 * @return {number[]}
 */

// using hash
const twoSumHash = (ns, t) => {
  let hash = {};
  for (let i = 0; i < ns.length; i++) {
    if (ns[i] in hash) {
      return [hash[ns[i]], i];
    }
    hash[t - ns[i]] = i;
  }
};

// using for
const twoSumFor = (ns, t) => {
  for (i = 0; i < ns.length; i++) {
    for (j = 0; j < ns.length; j++) {
      if (i === j) {
        continue;
      }
      if (t - ns[i] === ns[j]) {
        return [i, j];
      }
    }
  }
};

// using some
const twoSumSome = (ns, t) => {
  let ans;
  ns.some((n1, i, ns2) => {
    ns2.some((n2, j) => {
      if (i === j) return false;
      if (n2 === t - n1) {
        return (ans = [i, j]);
      }
    });
    if (ans) {
      return true;
    }
  });
  return ans;
};
