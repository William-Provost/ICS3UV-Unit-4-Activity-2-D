/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-28
 * @fileoverview This program finds the threshold below which the
 * antibiotic is considered expired
 */

// variables
const EFFECTIVENESS_THRESHOLD: number = 50.0;
const DECREASE_RATE: number = 0.04; // percentage decrease per month
let effectiveness: number = 100.0;
let months: number = 0;

console.log("Calculating antibiotic effectiveness over time:\n");

while (effectiveness >= EFFECTIVENESS_THRESHOLD) {
  console.log(`Month ${months}: Effectiveness = ${effectiveness.toFixed(2)}%`);
  effectiveness = effectiveness - (effectiveness * DECREASE_RATE);
  months++;
}

// output result
console.log(`\nThe antibiotic is considered expired after ${months} months when effectiveness drops below ${EFFECTIVENESS_THRESHOLD}%.`);

console.log("\nDone.");
