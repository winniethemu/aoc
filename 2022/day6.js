import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day6.input'),
  crlfDelay: Infinity,
});

const d = {};

function check() {
  const prod = Object.values(d).reduce((cur, acc) => cur * acc, 1);
  return prod === 1;
}

rl.on('line', (line) => {
  for (let i = 0; i < 14; i++) {
    const ch = line[i];
    if (ch in d) {
      d[ch] += 1;
    } else {
      d[ch] = 1;
    }
  }
  if (check()) return 14;

  for (let i = 14; i < line.length; i++) {
    const ch = line[i];
    if (ch in d) {
      d[ch] += 1;
    } else {
      d[ch] = 1;
    }

    d[line[i-14]] -= 1;
    if (d[line[i-14]] < 1) delete d[line[i-14]];

    if (check()) {
      console.log(i+1);
      return;
    }
  }
});
