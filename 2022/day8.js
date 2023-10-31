import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day8.input'),
  crlfDelay: Infinity,
});

const G = [];

rl.on('line', (line) => {
  const values = line.split('').map(c => parseInt(c, 10));
  G.push(values);
});

rl.on('close', () => {
  const H = G.length, W = G[0].length;

  // Part 1
  // const V = [];
  // for (let i = 0; i < H; i++) {
  //   V.push([]);
  // }

  // for (let i = 0; i < H; i++) {
  //   for (let j = 0; j < W; j++) {
  //     if (!V[i][j]) V[i][j] = [];
  //     V[i][j][0] = i > 0 ? Math.max(V[i-1][j][0], G[i-1][j]) : -1;
  //     V[i][j][3] = j > 0 ? Math.max(V[i][j-1][3], G[i][j-1]) : -1;
  //   }
  // }

  // for (let i = H-1; i >=0; i--) {
  //   for (let j = W-1; j >= 0; j--) {
  //     V[i][j][1] = j < W-1 ? Math.max(V[i][j+1][1], G[i][j+1]) : -1;
  //     V[i][j][2] = i < H-1 ? Math.max(V[i+1][j][2], G[i+1][j]) : -1;
  //   }
  // }

  // let total = 0;
  // for (let i = 0; i < H; i++) {
  //   for (let j = 0; j < W; j++) {
  //     for (let k = 0; k < 4; k++) {
  //       if (V[i][j][k] < 0 || V[i][j][k] < G[i][j]) {
  //         total++;
  //         break;
  //       }
  //     }
  //   }
  // }

  // Part 2
  let maxScore = -Infinity;
  const directions = [[-1, 0], [0, 1], [1, 0], [0, -1]];

  for (let i = 0; i < H; i++) {
    for (let j = 0; j < W; j++) {
      let score = 1;
      for (const [dx, dy] of directions) {
        let x = i + dx, y = j + dy;
        let distance = 0;
        while (true) {
          if (x < 0 || x >= H || y < 0 || y >= W) break;
          if (G[x][y] >= G[i][j]) {
            distance++;
            break;
          }
          distance++;
          x += dx;
          y += dy;
        }
        score *= distance;
      }
      maxScore = Math.max(maxScore, score);
    }
  }

  console.log(maxScore);
});
