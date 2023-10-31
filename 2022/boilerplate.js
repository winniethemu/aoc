import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream(/* filename */),
  crlfDelay: Infinity,
});

rl.on('line', (line) => {
});

rl.on('close', () => {
});
