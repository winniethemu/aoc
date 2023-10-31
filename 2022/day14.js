import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day14.input'),
  crlfDelay: Infinity,
});

const cave = new Set(); // position of all rocks and sands at rest
let bottom = 0;
let count = 0;          // num of rocks come to rest

rl.on('line', (line) => {
  const vertices = line
    .split(' -> ')
    .map(xy => xy.split(',').map(value => Number(value)));

  for (let i = 1; i < vertices.length; i++) {
    const [x0, y0] = vertices[i-1], [x1, y1] = vertices[i];
    bottom = Math.max(bottom, y0, y1);
    if (x0 === x1) {
      for (let j = Math.min(y0, y1); j <= Math.max(y0, y1); j++) {
        cave.add(String([x0, j]));
      }
    } else if (y0 === y1) {
      for (let j = Math.min(x0, x1); j <= Math.max(x0, x1); j++) {
        cave.add(String([j, y0]));
      }
    }
  }
});

function drop(x, y) {
  // if (y >= bottom) return false;

  // Part 2
  if (y === bottom+1) {
    cave.add(String([x, y]));
    return true;
  }
  if (cave.has(String([x, y]))) return false;

  const down = cave.has(String([x, y+1]));
  const downLeft = cave.has(String([x-1, y+1]));
  const downRight = cave.has(String([x+1, y+1]));

  if (down && downLeft && downRight) { // come to rest
    cave.add(String([x, y]));
    return true;
  } else if (!down) {
    return drop(x, y+1);
  } else if (!downLeft) {
    return drop(x-1, y+1);
  } else if (!downRight) {
    return drop(x+1, y+1);
  }
}

rl.on('close', () => {
  while (true) {
    if (drop(500, 0)) {
      count++;
    } else {
      break;
    }
  }
  console.log(count);
});
