import fs from 'node:fs';
import readline from 'node:readline';

const rl = readline.createInterface({
  input: fs.createReadStream('day15.input'),
  crlfDelay: Infinity,
});

const TARGET = 2000000;
const locations = [];
let intervals = [];

rl.on('line', (line) => {
  // Source = {x, y}0, Beacon = {x, y}1
  const [x0, y0, x1, y1] = line
    .match(/-?\d+/g)
    .map(value => Number(value));

  locations.push([x0, y0, x1, y1]);
});

rl.on('close', () => {
  for (let target = 0; target <= 4000000; target++) {
    intervals = [];

    for (const [x0, y0, x1, y1] of locations) {
      const distanceToBeacon = Math.abs(x0 - x1) + Math.abs(y0 - y1);
      const distanceToTarget = Math.abs(y0 - target);

      if (distanceToTarget <= distanceToBeacon) {
        const delta = Math.abs(distanceToTarget - distanceToBeacon);
        const [left, right] = [x0 - delta, x0 + delta];
        intervals.push([left, right]);
      }
    }

    // Merge intervals
    intervals.sort((a, b) => a[0] - b[0]);

    // Safe to assume there are at least 2 intervals
    const merged = [intervals[0]];
    for (let i = 1; i < intervals.length; i++) {
      const last = merged[merged.length-1];
      const curr = intervals[i];
      if (curr[0] <= last[1]) {
        last[1] = Math.max(last[1], curr[1]);
      } else {
        merged.push(curr);
      }
    }

    // let count = 0;
    // for (const [start, end] of merged) {
    //   count += (end - start);
    // }
    // console.log(count);

    // Part 2
    if (merged.length > 1) {
      console.log(target, merged);
    }
  }
});
