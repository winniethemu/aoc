import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day13.input'),
  crlfDelay: Infinity,
});

let total = 0;
let count = 1;
let first = null, second = null;
const lists = [];

// -1 = in order
// 0 = undecided
// 1 = not in order
function compare(left, right) {
  for (let i = 0; i < Math.max(left.length, right.length); i++) {
    let l = left[i], r = right[i];

    if (l === undefined) {
      return -1;
    } else if (r === undefined) {
      return 1;
    }

    if (typeof l === 'number' && typeof r === 'number') {
      if (l === r) {
        continue;
      } else {
        return (l < r) ? -1 : 1;
      }
    } else {
      if (typeof l === 'number') {
        l = [l];
      } else if (typeof r === 'number') {
        r = [r];
      }
      const result = compare(l, r);
      if (result === 0) continue;
      return result;
    }
  }

  return 0;
}

rl.on('line', (line) => {
  // Part 1
  // if (line.length < 1) {
  //   first = second = null;
  //   count++;
  //   return;
  // }

  // if (!first) {
  //   first = JSON.parse(line);
  // } else {
  //   second = JSON.parse(line);
  //   if (compare(first, second) < 2) {
  //     total += count;
  //   }
  // }

  // Part 2
  if (line.length > 0) {
    lists.push(JSON.parse(line));
  }
});

rl.on('close', () => {
  let prod = 1;
  lists.push([[2]]);
  lists.push([[6]]);
  lists.sort(compare);

  for (let i = 0; i < lists.length; i++) {
    if (lists[i].length === 1&& lists[i][0].length === 1) {
      if (lists[i][0][0] === 2 || lists[i][0][0] === 6) {
        prod *= i+1;
      }
    }
  }
  console.log(prod);
});
