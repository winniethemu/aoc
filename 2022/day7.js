import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day7.input'),
  crlfDelay: Infinity,
});

const f = {};
let curr = f;

rl.on('line', (line) => {
  if (line.startsWith('$')) {
    const [, cmd, arg] = line.split(' ');
    if (cmd === 'cd') {
      if (!(arg in curr)) curr[arg] = {};
      curr = curr[arg];
    }
  } else if (line.startsWith('dir')) {
    const [, dirname] = line.split(' ');
    if (!(dirname in curr)) curr[dirname] = {};
    curr[dirname]['..'] = curr;
  } else {
    const [size, filename] = line.split(' ');
    if (!(filename in curr)) curr[filename] = parseInt(size, 10);
  }
});

rl.on('close', () => {
  // Part 1
  // let total = 0;

  // function traverse(node) {
  //   if (typeof node === 'number') {
  //     return node;
  //   }

  //   let subtotal = 0;
  //   for (const child in node) {
  //     if (child !== '..') {
  //       subtotal += traverse(node[child]);
  //     }
  //   }

  //   if (subtotal <= 100000) {
  //     total += subtotal;
  //   }

  //   return subtotal;
  // }

  // traverse(f);

  // Part 2
  let ans = Infinity;
  function traverse(node, threshold) {
    if (typeof node === 'number') {
      return node;
    }

    let subtotal = 0;
    for (const child in node) {
      if (child !== '..') {
        subtotal += traverse(node[child], threshold);
      }
    }

    if (subtotal >= threshold) {
      ans = Math.min(ans, subtotal);
    }

    return subtotal;
  }

  const used = traverse(f, Infinity);
  const unused = 70000000 - used;
  const needed = 30000000;
  const delta = needed - unused;
  traverse(f, delta);
  console.log(ans);
});
