import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day10.input'),
  crlfDelay: Infinity,
});

let cycle = 1;
let x = 1;
let total = 0;
const display = new Set();

function exec() {
  const cyc = cycle % 40;
  if (Math.abs(x - cyc + 1) <= 1) {
    display.add(cycle);
  }
  cycle++;
}

rl.on('line', (line) => {
  const [op, value] = line.split(' ');
  // Part 1
  // if (op === 'addx') {
  //   if ([19, 59, 99, 139, 179, 219].indexOf(cycle) > -1) {
  //     total += x * (cycle+1);
  //   }

  //   x += Number(value);
  //   cycle += 2;
  // } else { // noop
  //   cycle += 1;
  // }

  // if ([20, 60, 100, 140, 180, 220].indexOf(cycle) > -1) {
  //   total += x * cycle;
  // }

  // Part 2
  if (op === 'addx') {
    exec();
    exec();
    x += Number(value);
  } else { // noop
    exec();
  }
});

rl.on('close', () => {
  let i = 0;
  while (true) {
    if (display.has(i+1)) {
      process.stdout.write('#');
    } else {
      process.stdout.write('.');
    }
    i++;
    if (i % 40 === 0) {
      process.stdout.write('\n');
    }
    if (i >= 240) break;
  }
});
