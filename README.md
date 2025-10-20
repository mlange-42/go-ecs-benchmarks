# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

## Benchmark candidates

| ECS | Tested | Latest | Activity |
|-----|--------|--------|----------|
| [Ark](https://github.com/mlange-42/ark) | v0.6.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/mlange-42/ark?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/mlange-42/ark?label=date) | ![Last commit](https://img.shields.io/github/last-commit/mlange-42/ark) |
| [Donburi](https://github.com/yottahmd/donburi) | v1.15.7 | ![GitHub Tag](https://img.shields.io/github/v/tag/yottahmd/donburi?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/yottahmd/donburi?label=date) | ![Last commit](https://img.shields.io/github/last-commit/yottahmd/donburi) |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/marioolofo/go-gameengine-ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/marioolofo/go-gameengine-ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/marioolofo/go-gameengine-ecs) |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/unitoftime/ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/unitoftime/ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/unitoftime/ecs) |
| [Volt](https://github.com/akmonengine/volt) | v1.6.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/akmonengine/volt?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/akmonengine/volt?label=date) | ![Last commit](https://img.shields.io/github/last-commit/akmonengine/volt) |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.
See the section on [Contributing](#contributing) for details.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Features

| ECS | Type-safe API | ID-based API | Relations | Events<sup>[1]</sup> | Batches<sup>[2]</sup> | Command buffer |
|-----|:-------------:|:------------:|:---------:|:-------:|:-------:|:--------------:|
| [Ark](https://github.com/mlange-42/ark) | ✅ | ✅ | ✅ | ✅ | ✅ | ❌ |
| [Donburi](https://github.com/yottahmd/donburi) | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| [Volt](https://github.com/akmonengine/volt) | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |

[1] ECS lifecycle events, allowing to react to entity creation, component addition, ...  
[2] Faster batch operations for entity creation etc.

## Benchmarks

Last run: Mon, 20 Oct 2025 20:10:09 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

All components used in the benchmarks have two `float64` fields.
The initial capacity of the world is set to 1024 where this is supported.

### Query

`N` entities with components `Position` and `Velocity`.
10x `N` entities with components `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/4a14df74-1a09-451f-933c-9d96cf3d7905)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.73ns | 55.98ns | 62.28ns | 67.87ns | 18.11ns | 112.62ns |
| 4 | 18.14ns | 15.73ns | 29.20ns | 14.88ns | 6.86ns | 28.98ns |
| 16 | 5.80ns | 5.22ns | 21.02ns | 7.60ns | 4.05ns | 8.10ns |
| 64 | 2.72ns | 2.54ns | 19.10ns | 6.04ns | 3.46ns | 2.98ns |
| 256 | 2.06ns | 1.90ns | 19.22ns | 5.47ns | 3.24ns | 1.71ns |
| 1k | 1.89ns | 1.88ns | 19.45ns | 5.36ns | 3.24ns | 1.39ns |
| 16k | 1.78ns | 1.79ns | 21.67ns | 5.34ns | 3.17ns | 1.27ns |
| 256k | 1.76ns | 1.77ns | 27.95ns | 5.34ns | 3.18ns | 1.27ns |
| 1M | 1.83ns | 1.84ns | 25.65ns | 5.37ns | 3.20ns | 1.38ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/60cdd6df-c133-4a73-ac99-444f1a252623)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 62.22ns | 56.06ns | 62.25ns | 44.83ns | 17.54ns | 116.11ns |
| 4 | 31.55ns | 26.42ns | 32.67ns | 18.95ns | 14.94ns | 72.10ns |
| 16 | 23.33ns | 17.17ns | 26.05ns | 12.24ns | 24.47ns | 59.56ns |
| 64 | 13.14ns | 9.85ns | 22.32ns | 8.24ns | 14.49ns | 29.30ns |
| 256 | 4.70ns | 3.72ns | 20.78ns | 5.84ns | 6.07ns | 8.75ns |
| 1k | 2.61ns | 2.41ns | 22.08ns | 5.48ns | 4.04ns | 3.31ns |
| 16k | 1.91ns | 2.05ns | 29.65ns | 5.37ns | 3.27ns | 1.48ns |
| 256k | 1.78ns | 1.93ns | 74.30ns | 5.35ns | 3.18ns | 1.28ns |
| 1M | 1.85ns | 2.05ns | 98.89ns | 5.35ns | 3.21ns | 1.33ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/5f6921f0-6473-4151-8ef7-7c3f4e6facc7)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 62.85ns | 55.90ns | 62.33ns | 54.39ns | 17.45ns | 131.66ns |
| 4 | 17.68ns | 15.89ns | 30.28ns | 25.21ns | 9.72ns | 54.83ns |
| 16 | 5.89ns | 5.30ns | 22.44ns | 19.19ns | 4.71ns | 36.72ns |
| 64 | 2.78ns | 2.61ns | 19.11ns | 17.39ns | 3.60ns | 37.58ns |
| 256 | 2.13ns | 1.99ns | 19.82ns | 8.37ns | 3.49ns | 10.38ns |
| 1k | 1.97ns | 1.95ns | 19.67ns | 6.11ns | 3.18ns | 3.59ns |
| 16k | 1.94ns | 1.93ns | 20.62ns | 5.41ns | 3.17ns | 1.45ns |
| 256k | 1.92ns | 1.92ns | 24.76ns | 5.36ns | 3.17ns | 1.28ns |
| 1M | 2.06ns | 2.03ns | 27.05ns | 5.36ns | 3.21ns | 1.38ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/8d3ba26f-370e-4edb-a076-add1889dde83)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 2.67ns | 7.94ns | 8.42ns | 36.95ns | 38.12ns |
| 4 | 2.55ns | 8.01ns | 8.23ns | 36.65ns | 38.22ns |
| 16 | 2.55ns | 8.02ns | 13.77ns | 36.80ns | 38.52ns |
| 64 | 2.61ns | 8.52ns | 13.94ns | 38.42ns | 38.25ns |
| 256 | 2.56ns | 8.71ns | 14.26ns | 39.60ns | 40.19ns |
| 1k | 2.68ns | 12.30ns | 18.65ns | 40.97ns | 41.47ns |
| 16k | 5.70ns | 38.59ns | 30.49ns | 57.39ns | 65.79ns |
| 256k | 9.45ns | 80.19ns | 45.17ns | 81.45ns | 124.60ns |
| 1M | 27.03ns | 215.10ns | 130.69ns | 201.67ns | 254.51ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/6025577d-2d17-4ef7-acb5-539fef0798a9)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 232.67ns | 225.59ns | 925.72ns | 380.24ns | 418.86ns | 833.12ns |
| 4 | 94.40ns | 60.32ns | 459.23ns | 165.55ns | 197.39ns | 368.00ns |
| 16 | 59.74ns | 24.20ns | 337.16ns | 147.39ns | 139.82ns | 262.80ns |
| 64 | 46.78ns | 11.83ns | 300.85ns | 127.97ns | 130.84ns | 203.60ns |
| 256 | 41.89ns | 9.62ns | 224.32ns | 112.28ns | 98.32ns | 186.96ns |
| 1k | 38.26ns | 8.38ns | 202.22ns | 113.56ns | 277.49ns | 202.54ns |
| 16k | 34.84ns | 7.83ns | 213.24ns | 119.64ns | 285.14ns | 212.53ns |
| 256k | 34.71ns | 7.67ns | 211.77ns | 118.96ns | 319.86ns | 244.01ns |
| 1M | 34.69ns | 7.68ns | 215.65ns | 196.67ns | 420.55ns | 338.19ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/40851a26-510e-4071-9162-b048c7550d99)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.83us | 7.78us | 4.42us | 16.18us | 3.18us | 1.83us |
| 4 | 2.02us | 1.96us | 1.29us | 4.62us | 1.03us | 793.26ns |
| 16 | 595.16ns | 545.42ns | 612.34ns | 1.25us | 368.21ns | 417.58ns |
| 64 | 172.40ns | 139.00ns | 386.68ns | 470.44ns | 236.41ns | 303.87ns |
| 256 | 81.70ns | 46.12ns | 301.06ns | 276.45ns | 171.07ns | 245.95ns |
| 1k | 52.60ns | 21.43ns | 284.55ns | 228.02ns | 150.48ns | 224.99ns |
| 16k | 58.03ns | 38.83ns | 404.87ns | 252.80ns | 174.57ns | 373.73ns |
| 256k | 88.02ns | 49.76ns | 480.53ns | 596.13ns | 257.44ns | 491.75ns |
| 1M | 73.01ns | 31.99ns | 448.72ns | 1.38us | 295.58ns | 519.10ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/81cbb3d0-19b3-43ea-ac0d-19eb3e1fb082)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 327.89ns | 298.59ns | 1.85us | 485.18ns | 629.19ns | 2.81us |
| 4 | 143.51ns | 84.57ns | 1.21us | 243.48ns | 366.82ns | 2.16us |
| 16 | 97.50ns | 27.93ns | 1.03us | 222.97ns | 307.28ns | 1.40us |
| 64 | 82.85ns | 13.87ns | 766.79ns | 198.12ns | 269.51ns | 1.39us |
| 256 | 74.19ns | 10.43ns | 714.07ns | 170.91ns | 214.81ns | 1.32us |
| 1k | 69.46ns | 8.17ns | 709.53ns | 166.83ns | 478.94ns | 1.32us |
| 16k | 67.59ns | 7.73ns | 757.00ns | 167.32ns | 483.01ns | 1.32us |
| 256k | 67.73ns | 7.61ns | 764.65ns | 176.40ns | 600.28ns | 1.43us |
| 1M | 68.01ns | 7.57ns | 685.91ns | 245.85ns | 611.79ns | 1.49us |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/0241c9c1-2907-47c0-8ea8-0f71a9cfee07)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 90.68ns | 198.81ns | 393.67ns | 502.92ns | 332.62ns | 436.17ns |
| 4 | 99.55ns | 58.94ns | 404.54ns | 509.79ns | 333.38ns | 447.86ns |
| 16 | 102.01ns | 27.39ns | 415.05ns | 551.30ns | 348.08ns | 453.02ns |
| 64 | 101.15ns | 20.71ns | 407.28ns | 547.60ns | 349.60ns | 453.57ns |
| 256 | 99.54ns | 11.62ns | 406.18ns | 550.64ns | 360.89ns | 453.30ns |
| 1k | 100.83ns | 8.91ns | 413.49ns | 567.72ns | 728.99ns | 462.64ns |
| 16k | 103.76ns | 9.50ns | 469.09ns | 600.34ns | 801.75ns | 510.82ns |
| 256k | 103.33ns | 11.11ns | 462.63ns | 875.70ns | 1.28us | 844.68ns |
| 1M | 102.31ns | 11.72ns | 459.23ns | 970.22ns | 1.42us | 957.19ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/6d29fe09-42ea-4c8b-90a5-5bd696c3f9ab)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 346.59ns | 473.86ns | 1.05us | 933.90ns | 762.46ns | 1.65us |
| 4 | 397.57ns | 184.14ns | 1.07us | 962.38ns | 774.21ns | 1.63us |
| 16 | 388.55ns | 114.21ns | 1.06us | 980.38ns | 775.74ns | 1.63us |
| 64 | 387.72ns | 95.03ns | 1.04us | 958.85ns | 812.65ns | 1.63us |
| 256 | 384.08ns | 35.22ns | 1.06us | 963.91ns | 849.19ns | 1.62us |
| 1k | 416.38ns | 20.83ns | 1.15us | 992.88ns | 1.39us | 1.70us |
| 16k | 434.28ns | 27.25ns | 1.29us | 1.11us | 1.47us | 1.90us |
| 256k | 570.73ns | 43.25ns | 1.35us | 1.55us | 1.97us | 2.58us |
| 1M | 499.36ns | 45.90ns | 1.34us | 1.68us | 2.10us | 2.72us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/d65d301c-5a9b-4cff-ac93-727718de2ac4)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 219.87ns | 286.35ns | 265.58ns | 389.85ns | 196.23ns | 485.90ns |
| 4 | 90.68ns | 85.21ns | 111.24ns | 167.87ns | 85.73ns | 240.54ns |
| 16 | 51.20ns | 32.46ns | 73.30ns | 151.44ns | 52.77ns | 187.96ns |
| 64 | 44.92ns | 18.90ns | 60.32ns | 134.53ns | 43.40ns | 147.77ns |
| 256 | 42.76ns | 8.72ns | 57.19ns | 99.07ns | 41.73ns | 141.99ns |
| 1k | 27.82ns | 6.41ns | 52.10ns | 106.72ns | 36.98ns | 134.63ns |
| 16k | 26.25ns | 5.38ns | 50.22ns | 115.75ns | 48.54ns | 152.60ns |
| 256k | 25.30ns | 5.47ns | 53.21ns | 155.74ns | 63.74ns | 196.29ns |
| 1M | 25.13ns | 6.42ns | 50.64ns | 277.60ns | 171.90ns | 326.98ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/7281b1e0-362e-4c54-8e84-36b6eaecf2b8)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 278.54ns | 377.32ns | 405.22ns | 548.74ns | 207.67ns | 892.79ns |
| 4 | 177.08ns | 134.16ns | 208.47ns | 265.43ns | 84.37ns | 591.87ns |
| 16 | 138.96ns | 73.61ns | 160.03ns | 240.99ns | 54.16ns | 494.22ns |
| 64 | 127.50ns | 60.33ns | 143.97ns | 183.77ns | 46.68ns | 385.73ns |
| 256 | 110.05ns | 17.95ns | 123.71ns | 169.23ns | 49.03ns | 334.03ns |
| 1k | 80.16ns | 8.88ns | 137.73ns | 161.97ns | 37.75ns | 315.86ns |
| 16k | 82.82ns | 8.31ns | 143.74ns | 174.89ns | 50.63ns | 339.98ns |
| 256k | 85.08ns | 12.53ns | 131.80ns | 274.18ns | 70.52ns | 476.15ns |
| 1M | 85.99ns | 13.73ns | 128.39ns | 382.87ns | 167.48ns | 643.12ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 28.70us | 2.63us | 224.73us | 2.46us | 43.17us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/ark&yottahmd/donburi&marioolofo/go-gameengine-ecs&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/ark,yottahmd/donburi,marioolofo/go-gameengine-ecs,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
</a>
</p>

## Running the benchmarks

Run the benchmarks using the following command:

```shell
go run . -test.benchtime=0.25s
```

> On PowerShell use this instead:  
> `go run . --% -test.benchtime=0.25s`

The `benchtime` limit is required for some of the benchmarks that have a high
setup cost which is not timed. They would take forever otherwise.
The benchmarks can take up to one hour to complete.

To run a selection of benchmarks, add their names as arguments:

```shell
go run . query2comp query1in10 query32arch
```

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```

## Contributing

Developers of ECS frameworks are welcome to add their implementation to the benchmarks.
However, there are a few (quality) criteria that need to be fulfilled for inclusion:

- All benchmarks must be implemented, which means that the ECS must have the required features
- The ECS must be working properly and not exhibit serious flaws; it will undergo a basic review by maintainers
- The ECS must be sufficiently documented so that it can be used without reading the code
- There must be at least basic unit tests
- Unit tests must be run in the CI of the repository
- The ECS *must not* be tightly coupled to a particular game engine, particularly graphics stuff
- There must be tagged release versions; only tagged versions will be included here

Developers of included frameworks are encouraged to review the benchmarks,
and to fix (or point to) misuse or potential optimizations.
