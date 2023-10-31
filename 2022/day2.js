import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day2.input'),
  crlfDelay: Infinity,
});

/**
 * Part 1
 * ======
 * A = Rock     = X
 * B = Paper    = Y
 * C = Scissors = Z
 *
 * totalScore = sum((round) => {
 *   const [opponent, self] = round;
 *   return value(self) + outcome(round);
 * });
 */
const hand = {
  A: 'rock',
  X: 'rock',
  B: 'paper',
  Y: 'paper',
  C: 'scissors',
  Z: 'scissors',
};

const value = { rock: 1, paper: 2, scissors: 3 };

// const outcome = (opponent, self) => {
//   if (opponent === self) return 3;

//   if (opponent === 'rock') {
//     if (self === 'paper') {
//       return 6;
//     } else {
//       return 0;
//     }
//   } else if (opponent === 'paper') {
//     if (self === 'scissors') {
//       return 6;
//     } else {
//       return 0;
//     }
//   } else { // 'scissors'
//     if (self === 'rock') {
//       return 6;
//     } else {
//       return 0;
//     }
//   }
// };

let total = 0;

// rl.on('line', (line) => {
//   const round = line.split(' ');
//   const opponent = hand[round[0]];
//   const self = hand[round[1]];
//   total += value[self] + outcome(opponent, self);
// });

/**
 * Part 2
 * ======
 * X = Lose = 0
 * Y = Draw = 3
 * Z = Win  = 6
 */
const score = {
  X: 0,
  Y: 3,
  Z: 6,
};

const response = {
  rock: {
    X: 'scissors',
    Y: 'rock',
    Z: 'paper',
  },
  paper: {
    X: 'rock',
    Y: 'paper',
    Z: 'scissors',
  },
  scissors: {
    X: 'paper',
    Y: 'scissors',
    Z: 'rock',
  },
};

rl.on('line', (line) => {
  const round = line.split(' ');
  const opponent = hand[round[0]];
  const outcome = round[1];
  total += value[response[opponent][outcome]] + score[outcome];
});

rl.on('close', () => {
  console.log(total);
});
