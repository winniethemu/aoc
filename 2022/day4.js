import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day4.input'),
  crlfDelay: Infinity,
});

let count = 0;

rl.on('line', (line) => {
  const [a1, a2] = line.split(',');
  const [s1, e1] = a1.split('-').map(a => parseInt(a, 10));
  const [s2, e2] = a2.split('-').map(a => parseInt(a, 10));
  // Part 1
  // if ((s1 <= s2 && e1 >= e2) || (s2 <= s1 && e2 >= e1)) {
  //   count++;
  // }

  // Part 2
  if (s1 <= s2) {
    if (s2 <= e1) count++;
  } else {
    if (s1 <= e2) count++;
  }
});

rl.on('close', () => {
  console.log(count);
});
