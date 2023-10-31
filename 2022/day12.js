import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day12.input'),
  crlfDelay: Infinity,
});

const grid = [];
const directions = [
  [-1, 0], [0, 1], [1, 0], [0, -1]
];

rl.on('line', (line) => {
  grid.push(line);
});

function steep(curr, next) {
  const c = (curr === 'S') ? 'a' : curr;
  const n = (next === 'E') ? 'z' : next;
  return n.charCodeAt(0) - c.charCodeAt(0) > 1;
}

function bfs(x0, y0, dist) {
  const visited = new Set();
  const q = [[x0, y0, dist]];

  while (q.length > 0) {
    const [x, y, steps] = q.shift();
    if (grid[x][y] === 'E') return steps;
    for (const [dx, dy] of directions) {
      const nx = x + dx, ny = y + dy;
      if (nx < 0 || nx > grid.length-1) continue;
      if (ny < 0 || ny > grid[0].length-1) continue;
      if (visited.has([nx, ny].toString())) continue;
      if (steep(grid[x][y], grid[nx][ny])) continue;
      q.push([nx, ny, steps+1]);
      visited.add([nx, ny].toString());
    }
  }

  return Infinity;
}

rl.on('close', () => {
  // Part 1
  // for (let i = 0; i < grid.length; i++) {
  //   const pos = grid[i].indexOf('S');
  //   if (pos > -1) {
  //     const total = bfs(i, pos, 0);
  //     console.log(total);
  //     return;
  //   }
  // }

  // Part 2
  let total = Infinity;
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j] === 'a' || grid[i][j] === 'S') {
        const steps = bfs(i, j, 0);
        total = Math.min(total, steps);
      }
    }
  }
  console.log(total);
});
