import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day1.input'),
  crlfDelay: Infinity,
});

// Part 1
// let result = -Infinity;
// let current = 0;

// rl.on('line', (line) => {
//   if (line.length < 1) {
//     result = Math.max(current, result);
//     current = 0;
//   } else {
//     current += parseInt(line, 10);
//   }
// });

// Part 2
const elves = [];
let curr = 0;

rl.on('line', (line) => {
  if (line.length < 1) {
    elves.push(curr);
    curr = 0;
  } else {
    curr += parseInt(line, 10);
  }
});

rl.on('close', () => {
  elves.sort((a, b) => b - a);
  const top3 = elves.slice(0, 3);
  console.log(top3.reduce((cur, acc) => acc + cur, 0));
});
