// Package addremove benchmarks adding and removing components
//
// Setup:
//   - 1000 entities with Position{f64, f64}
//
// Benchmark:
//   - Iterate all entities with Position, and add Velocity
//   - Iterate all entities with Position and Velocity, and remove Velocity
package addremove
