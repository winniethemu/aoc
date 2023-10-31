import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day5.input'),
  crlfDelay: Infinity,
});

const crates = [];

rl.on('line', (line) => {
  if (line.startsWith(' 1')) return;
  if (line.length < 1) return;

  // parse moving instructions
  if (line.startsWith('move')) {
    const words = line.split(' ');
    const count = words[1];
    const src = words[3];
    const dest = words[5];

    // Part 1
    // for (let i = 0; i < count; i++) {
    //   const item = crates[src].pop();
    //   crates[dest].push(item);
    // }

    const st = [];
    for (let i = 0; i < count; i++) {
      st.push(crates[src].pop());
    }
    for (let i = 0; i < count; i++) {
      crates[dest].push(st.pop());
    }
    return;
  }

  // parse crate diagram
  const size = Math.floor(line.length / 3);
  for (let i = 1; i <= size; i++) {
    if (crates[i] === undefined) crates[i] = [];
    const item = line[(i-1)*4 + 1];
    if (item !== ' ') {
      crates[i].unshift(item);
    }
  }
});

rl.on('close', () => {
  let message = '';
  for (let i = 1; i < crates.length; i++) {
    const top = crates[i][crates[i].length-1] || ' ';
    message += top;
  }
  console.log(message);
});
