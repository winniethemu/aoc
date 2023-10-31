import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day3.input'),
  crlfDelay: Infinity,
});

let total = 0;

function priority(letter) {
  if (letter === letter.toUpperCase()) {
    // A = 27
    return 27 + letter.charCodeAt(0) - 'A'.charCodeAt(0);
  } else { // lower case
    // a = 1
    return 1 + letter.charCodeAt(0) - 'a'.charCodeAt(0);
  }
}

function intersect(setA, setB) {
  const intersection = new Set();
  for (const item of setB) {
    if (setA.has(item)) {
      intersection.add(item);
    }
  }
  return intersection;
}

// Part 1
//
// rl.on('line', (line) => {
//   const s1 = new Set(), s2 = new Set();
//   for (let i = 0; i < line.length; i++) {
//     if (i < line.length / 2) {
//       s1.add(line[i]);
//     } else if (s1.has(line[i])) {
//       total += priority(line[i]);
//       break;
//     }
//   }
// });

// Part 2
let group = [];
rl.on('line', (line) => {
  group.push(line);
  if (group.length === 3) {
    const s1 = new Set(group[0].split(''));
    const s2 = new Set(group[1].split(''));
    const s3 = new Set(group[2].split(''));
    const badge = intersect(intersect(s1, s2), s3);
    total += priority(Array.from(badge)[0]);
    group = [];
  }
});

rl.on('close', () => {
  console.log(total);
});
