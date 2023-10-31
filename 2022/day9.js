import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day9.input'),
  crlfDelay: Infinity,
});

let currT = [0, 0];
let currH = [0, 0];
const visited = new Set(['0,0']);

// Part 2
const knots = [];
for (let i = 0; i < 10; i++) {
  knots.push([0, 0]);
}

const MOVE = {
  R: [0, 1],
  U: [-1, 0],
  L: [0, -1],
  D: [1, 0],
};

function add(v1, v2) {
  return [v1[0] + v2[0], v1[1] + v2[1]];
}

function touching(v1, v2) {
  const delta0 = v1[0] - v2[0];
  const delta1 = v1[1] - v2[1];
  return Math.abs(delta0) < 2 && Math.abs(delta1) < 2;
}

function catchup(v1, v2) {
  const delta0 = Math.abs(v1[0] - v2[0]);
  const delta1 = Math.abs(v1[1] - v2[1]);

  if (delta0 === 0) {
    return [v1[0], Math.max(v1[1], v2[1])-1];
  }

  if (delta1 === 0) {
    return [Math.max(v1[0], v2[0])-1, v1[1]];
  }

  if (delta0 === 2 && delta1 === 1) {
    if (v1[0] > v2[0]) {
      return [v1[0]-1, v2[1]];
    } else {
      return [v1[0]+1, v2[1]];
    }
  }

  if (delta0 === 1 && delta1 === 2) {
    if (v1[1] > v2[1]) {
      return [v2[0], v1[1]-1];
    } else {
      return [v2[0], v1[1]+1];
    }
  }

  if (delta0 === 2 && delta1 === 2) {
    return [Math.min(v1[0], v2[0])+1, Math.min(v1[1], v2[1])+1];
  }

  console.log('Should not get here', v1, v2);
}

rl.on('line', (line) => {
  const [direction, count] = line.split(' ');
  for (let i = 0; i < Number(count); i++) {
    const moveH = MOVE[direction];
    // Part 1
    // currH = add(moveH, currH);
    // if (!touching(currT, currH)) {
    //   currT = catchup(currT, currH);
    //   visited.add(currT.toString());
    // }

    // Part 2
    knots[0] = add(moveH, knots[0]); // update H
    for (let i = 1; i < 10; i++) {
      if (!touching(knots[i], knots[i-1])) {
        knots[i] = catchup(knots[i], knots[i-1]);
      }
    }
    visited.add(knots[9].toString());
  }
});

rl.on('close', () => {
  console.log(visited.size);
});
