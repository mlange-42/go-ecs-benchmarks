# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

## Benchmark candidates

| ECS | Tested | Latest | Activity |
|-----|--------|--------|----------|
| [Ark](https://github.com/mlange-42/ark) | v0.6.1 | ![GitHub Tag](https://img.shields.io/github/v/tag/mlange-42/ark?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/mlange-42/ark?label=date) | ![Last commit](https://img.shields.io/github/last-commit/mlange-42/ark) |
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

Last run: Wed, 15 Oct 2025 16:10:15 UTC  
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

![query2comp](https://github.com/user-attachments/assets/0a46269c-0661-48fa-ad42-3439c87da1ff)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.99ns | 58.45ns | 62.26ns | 47.90ns | 18.19ns | 113.37ns |
| 4 | 18.11ns | 16.34ns | 29.53ns | 15.34ns | 7.09ns | 29.75ns |
| 16 | 6.01ns | 5.46ns | 20.82ns | 7.78ns | 4.35ns | 8.10ns |
| 64 | 2.81ns | 2.65ns | 19.00ns | 6.02ns | 3.71ns | 2.97ns |
| 256 | 2.02ns | 1.99ns | 19.17ns | 5.49ns | 3.52ns | 1.71ns |
| 1k | 1.91ns | 1.93ns | 20.07ns | 5.36ns | 3.49ns | 1.40ns |
| 16k | 1.96ns | 1.96ns | 20.71ns | 5.36ns | 3.49ns | 1.28ns |
| 256k | 1.94ns | 1.94ns | 24.90ns | 5.35ns | 3.49ns | 1.27ns |
| 1M | 2.04ns | 2.04ns | 30.53ns | 5.36ns | 3.52ns | 1.35ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/f58295f7-6ab9-409e-952c-198530ecfd0e)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.02ns | 57.15ns | 59.87ns | 44.30ns | 17.78ns | 114.29ns |
| 4 | 31.65ns | 24.23ns | 32.74ns | 19.11ns | 15.08ns | 74.42ns |
| 16 | 23.94ns | 15.46ns | 25.65ns | 12.41ns | 30.10ns | 64.24ns |
| 64 | 13.13ns | 8.72ns | 23.22ns | 8.38ns | 15.04ns | 31.57ns |
| 256 | 4.85ns | 3.67ns | 20.96ns | 5.85ns | 6.33ns | 9.14ns |
| 1k | 2.69ns | 2.34ns | 22.17ns | 5.47ns | 4.35ns | 3.44ns |
| 16k | 2.10ns | 1.87ns | 29.94ns | 5.38ns | 3.59ns | 1.50ns |
| 256k | 1.93ns | 1.78ns | 78.60ns | 5.35ns | 3.49ns | 1.29ns |
| 1M | 2.06ns | 1.89ns | 105.88ns | 5.37ns | 3.52ns | 1.39ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/ec7cc347-9de0-455f-aa49-7ddc969f25bb)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 65.22ns | 58.60ns | 62.97ns | 55.14ns | 17.77ns | 137.48ns |
| 4 | 17.96ns | 16.34ns | 29.19ns | 25.21ns | 9.43ns | 64.34ns |
| 16 | 5.98ns | 5.36ns | 20.77ns | 18.73ns | 4.86ns | 48.73ns |
| 64 | 2.75ns | 2.57ns | 18.79ns | 17.24ns | 3.85ns | 51.06ns |
| 256 | 1.97ns | 1.91ns | 20.62ns | 8.42ns | 3.54ns | 13.77ns |
| 1k | 1.79ns | 1.78ns | 20.37ns | 6.16ns | 3.50ns | 4.42ns |
| 16k | 1.79ns | 1.77ns | 20.83ns | 5.41ns | 3.48ns | 1.51ns |
| 256k | 1.77ns | 1.77ns | 23.24ns | 5.36ns | 3.49ns | 1.28ns |
| 1M | 1.87ns | 1.89ns | 29.70ns | 5.37ns | 3.51ns | 1.39ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/cf708bbf-f83a-4a81-b1fc-5250528ff970)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 2.86ns | 7.95ns | 8.43ns | 36.65ns | 38.14ns |
| 4 | 3.00ns | 8.02ns | 8.25ns | 36.45ns | 38.36ns |
| 16 | 2.76ns | 8.05ns | 13.74ns | 36.77ns | 38.76ns |
| 64 | 2.76ns | 8.31ns | 13.99ns | 39.22ns | 38.32ns |
| 256 | 2.78ns | 8.74ns | 14.42ns | 41.17ns | 39.39ns |
| 1k | 2.85ns | 12.24ns | 18.56ns | 41.50ns | 41.63ns |
| 16k | 5.47ns | 38.99ns | 30.28ns | 58.77ns | 65.81ns |
| 256k | 8.35ns | 83.60ns | 50.89ns | 120.52ns | 173.76ns |
| 1M | 38.69ns | 249.60ns | 154.63ns | 234.55ns | 269.18ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/98baca11-43b1-4e46-9ce9-90e78d4a708e)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 237.75ns | 227.28ns | 1.02us | 375.31ns | 456.04ns | 805.67ns |
| 4 | 100.43ns | 65.50ns | 498.97ns | 178.06ns | 214.59ns | 382.89ns |
| 16 | 68.86ns | 23.47ns | 335.97ns | 144.83ns | 145.36ns | 259.12ns |
| 64 | 55.15ns | 12.89ns | 284.96ns | 118.50ns | 137.24ns | 253.03ns |
| 256 | 47.53ns | 10.26ns | 225.07ns | 112.15ns | 88.43ns | 190.09ns |
| 1k | 39.32ns | 9.79ns | 208.67ns | 110.85ns | 282.84ns | 204.34ns |
| 16k | 34.45ns | 7.78ns | 218.36ns | 118.95ns | 286.51ns | 221.05ns |
| 256k | 34.45ns | 7.59ns | 203.10ns | 118.20ns | 300.58ns | 225.54ns |
| 1M | 34.37ns | 7.73ns | 190.83ns | 167.80ns | 412.72ns | 333.44ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/ca2adb50-a0c1-40c5-b2a0-b99d88cd57ab)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.29us | 7.34us | 4.53us | 16.85us | 3.31us | 1.79us |
| 4 | 1.92us | 1.83us | 1.39us | 3.85us | 1.06us | 802.05ns |
| 16 | 522.03ns | 478.10ns | 641.57ns | 1.24us | 387.69ns | 418.88ns |
| 64 | 168.18ns | 130.44ns | 415.96ns | 498.59ns | 220.99ns | 285.80ns |
| 256 | 79.70ns | 45.03ns | 366.53ns | 283.41ns | 168.01ns | 251.56ns |
| 1k | 51.81ns | 20.14ns | 310.84ns | 199.75ns | 154.54ns | 236.66ns |
| 16k | 56.78ns | 40.32ns | 434.91ns | 243.79ns | 174.07ns | 402.50ns |
| 256k | 100.78ns | 45.32ns | 488.14ns | 584.53ns | 243.61ns | 475.76ns |
| 1M | 74.34ns | 34.81ns | 460.18ns | 1.40us | 285.96ns | 533.75ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/3b3d7a67-7d2f-431a-b765-264d864aa2b0)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 320.50ns | 305.86ns | 1.91us | 485.13ns | 720.97ns | 3.07us |
| 4 | 148.92ns | 88.73ns | 1.20us | 255.58ns | 419.22ns | 2.16us |
| 16 | 94.33ns | 29.00ns | 1.08us | 221.05ns | 276.73ns | 1.52us |
| 64 | 79.22ns | 13.91ns | 781.82ns | 213.26ns | 281.97ns | 1.35us |
| 256 | 78.79ns | 10.61ns | 739.30ns | 159.97ns | 223.79ns | 1.26us |
| 1k | 69.73ns | 8.76ns | 774.23ns | 162.48ns | 470.15ns | 1.29us |
| 16k | 67.49ns | 7.84ns | 764.18ns | 167.01ns | 484.08ns | 1.29us |
| 256k | 67.34ns | 9.22ns | 722.32ns | 177.66ns | 558.92ns | 1.37us |
| 1M | 67.29ns | 7.63ns | 715.23ns | 239.13ns | 652.26ns | 1.41us |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/93412c4f-a25d-41b9-ac27-306475dd14df)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 95.71ns | 199.08ns | 407.38ns | 506.36ns | 339.35ns | 439.44ns |
| 4 | 100.52ns | 63.83ns | 422.69ns | 515.95ns | 344.50ns | 446.39ns |
| 16 | 102.20ns | 27.68ns | 431.82ns | 552.36ns | 338.57ns | 440.64ns |
| 64 | 100.85ns | 20.82ns | 428.43ns | 546.92ns | 348.67ns | 442.46ns |
| 256 | 100.86ns | 11.18ns | 426.38ns | 552.57ns | 362.28ns | 448.81ns |
| 1k | 101.44ns | 8.69ns | 425.79ns | 569.88ns | 735.14ns | 457.73ns |
| 16k | 106.75ns | 9.49ns | 486.27ns | 599.52ns | 816.03ns | 515.63ns |
| 256k | 103.40ns | 10.34ns | 479.08ns | 882.60ns | 1.34us | 886.93ns |
| 1M | 101.90ns | 11.97ns | 476.83ns | 985.86ns | 1.41us | 995.48ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/6c5c17d0-9522-4bef-917c-c4a401513ce3)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 342.10ns | 470.62ns | 1.06us | 931.31ns | 805.83ns | 1.71us |
| 4 | 393.24ns | 187.73ns | 1.08us | 956.42ns | 827.33ns | 1.62us |
| 16 | 386.54ns | 118.22ns | 1.07us | 976.11ns | 816.44ns | 1.64us |
| 64 | 396.02ns | 96.22ns | 1.05us | 955.84ns | 854.08ns | 1.63us |
| 256 | 382.84ns | 36.84ns | 1.07us | 956.80ns | 885.61ns | 1.63us |
| 1k | 424.76ns | 20.66ns | 1.15us | 983.18ns | 1.44us | 1.70us |
| 16k | 414.22ns | 24.85ns | 1.30us | 1.13us | 1.51us | 2.03us |
| 256k | 556.85ns | 45.27ns | 1.46us | 1.57us | 2.00us | 2.69us |
| 1M | 491.22ns | 46.25ns | 1.39us | 1.74us | 2.19us | 2.80us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/7411de37-3ee8-45c4-86fb-1b32f1dc8477)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 175.98ns | 311.37ns | 264.61ns | 443.81ns | 224.26ns | 583.17ns |
| 4 | 85.29ns | 90.17ns | 111.85ns | 182.35ns | 86.07ns | 269.82ns |
| 16 | 49.15ns | 33.85ns | 72.18ns | 141.25ns | 53.59ns | 187.08ns |
| 64 | 40.22ns | 20.10ns | 60.16ns | 105.41ns | 42.34ns | 169.61ns |
| 256 | 36.16ns | 9.19ns | 57.75ns | 109.63ns | 44.25ns | 138.52ns |
| 1k | 27.72ns | 6.19ns | 57.93ns | 104.82ns | 36.65ns | 132.46ns |
| 16k | 27.29ns | 5.57ns | 51.85ns | 116.46ns | 50.73ns | 152.70ns |
| 256k | 26.98ns | 5.36ns | 54.30ns | 150.84ns | 81.16ns | 210.03ns |
| 1M | 27.12ns | 6.55ns | 51.45ns | 253.49ns | 157.71ns | 338.31ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/c2041e75-f778-4cd7-9cd4-ff8f5613a5e8)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 263.41ns | 403.77ns | 432.97ns | 557.02ns | 190.42ns | 929.54ns |
| 4 | 165.99ns | 140.61ns | 205.92ns | 266.08ns | 77.86ns | 521.50ns |
| 16 | 146.53ns | 74.24ns | 151.28ns | 234.29ns | 48.79ns | 492.06ns |
| 64 | 137.64ns | 53.99ns | 151.11ns | 209.35ns | 42.85ns | 355.11ns |
| 256 | 107.04ns | 18.42ns | 120.52ns | 146.75ns | 47.64ns | 314.01ns |
| 1k | 82.01ns | 8.19ns | 140.82ns | 161.82ns | 35.13ns | 316.34ns |
| 16k | 85.08ns | 7.92ns | 137.53ns | 173.99ns | 50.46ns | 337.56ns |
| 256k | 89.93ns | 12.43ns | 163.50ns | 262.07ns | 76.40ns | 484.97ns |
| 1M | 92.92ns | 14.74ns | 229.20ns | 375.41ns | 166.48ns | 674.11ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 25.73us | 2.61us | 220.57us | 2.43us | 43.46us |


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
