# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Tested | Latest | Activity |
|-----|--------|--------|----------|
| [Ark](https://github.com/mlange-42/ark) | v0.6.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/mlange-42/ark?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/mlange-42/ark?label=date) | ![Last commit](https://img.shields.io/github/last-commit/mlange-42/ark) |
| [Donburi](https://github.com/yottahmd/donburi) | v1.15.7 | ![GitHub Tag](https://img.shields.io/github/v/tag/yottahmd/donburi?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/yottahmd/donburi?label=date) | ![Last commit](https://img.shields.io/github/last-commit/yottahmd/donburi) |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/marioolofo/go-gameengine-ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/marioolofo/go-gameengine-ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/marioolofo/go-gameengine-ecs) |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 | ![GitHub Tag](https://img.shields.io/github/v/tag/unitoftime/ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/unitoftime/ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/unitoftime/ecs) |
| [Volt](https://github.com/akmonengine/volt) | v1.7.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/akmonengine/volt?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/akmonengine/volt?label=date) | ![Last commit](https://img.shields.io/github/last-commit/akmonengine/volt) |

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
| [Volt](https://github.com/akmonengine/volt) | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ |

[1] ECS lifecycle events, allowing to react to entity creation, component addition, ...  
[2] Faster batch operations for entity creation etc.

## Benchmarks

Last run: Thu, 23 Oct 2025 15:26:42 UTC  
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

![query2comp](https://github.com/user-attachments/assets/c1f7c399-6885-4ded-9f07-5b461f459d05)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 63.43ns | 57.72ns | 61.91ns | 46.51ns | 18.11ns | 88.77ns |
| 4 | 18.27ns | 15.92ns | 30.50ns | 14.86ns | 6.86ns | 23.08ns |
| 16 | 5.78ns | 5.26ns | 22.34ns | 7.38ns | 4.05ns | 6.60ns |
| 64 | 2.70ns | 2.55ns | 19.62ns | 5.90ns | 3.46ns | 2.61ns |
| 256 | 1.93ns | 1.90ns | 20.12ns | 5.45ns | 3.31ns | 1.63ns |
| 1k | 1.81ns | 1.82ns | 19.71ns | 5.37ns | 3.18ns | 1.37ns |
| 16k | 1.76ns | 1.76ns | 21.23ns | 5.39ns | 3.17ns | 1.28ns |
| 256k | 1.76ns | 1.76ns | 24.88ns | 5.34ns | 3.17ns | 1.26ns |
| 1M | 1.85ns | 1.84ns | 31.05ns | 5.34ns | 3.20ns | 1.41ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/a91b4b67-4d38-4dc4-980d-10e856ecc2b2)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 63.21ns | 56.19ns | 61.99ns | 44.31ns | 17.45ns | 85.32ns |
| 4 | 31.51ns | 24.43ns | 33.68ns | 18.89ns | 14.94ns | 65.68ns |
| 16 | 22.91ns | 16.48ns | 26.62ns | 12.39ns | 24.70ns | 56.88ns |
| 64 | 12.97ns | 9.10ns | 24.29ns | 8.38ns | 15.08ns | 29.24ns |
| 256 | 4.76ns | 3.69ns | 22.32ns | 5.84ns | 6.37ns | 8.72ns |
| 1k | 2.67ns | 2.34ns | 23.58ns | 5.46ns | 4.04ns | 3.32ns |
| 16k | 2.08ns | 1.87ns | 29.89ns | 5.37ns | 3.28ns | 1.47ns |
| 256k | 1.94ns | 1.76ns | 71.52ns | 5.38ns | 3.19ns | 1.28ns |
| 1M | 2.04ns | 1.87ns | 103.98ns | 5.36ns | 3.20ns | 1.36ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/9560e1e3-178c-48a2-95d8-ad4ebe9534a3)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 63.01ns | 57.02ns | 62.28ns | 54.74ns | 17.44ns | 105.17ns |
| 4 | 17.46ns | 15.75ns | 30.10ns | 24.94ns | 9.67ns | 49.02ns |
| 16 | 5.86ns | 5.27ns | 21.90ns | 19.37ns | 4.70ns | 36.74ns |
| 64 | 2.79ns | 2.61ns | 19.81ns | 17.41ns | 3.72ns | 37.30ns |
| 256 | 2.12ns | 1.98ns | 20.74ns | 8.39ns | 3.24ns | 10.33ns |
| 1k | 1.92ns | 1.90ns | 22.44ns | 6.13ns | 3.18ns | 3.55ns |
| 16k | 1.93ns | 1.93ns | 20.34ns | 5.40ns | 3.25ns | 1.44ns |
| 256k | 1.92ns | 1.92ns | 23.40ns | 5.35ns | 3.17ns | 1.27ns |
| 1M | 2.01ns | 2.04ns | 26.46ns | 5.38ns | 3.20ns | 1.37ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/a5b2a5d7-7bf6-48bc-a260-5901aa9ad320)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 2.70ns | 8.32ns | 8.43ns | 37.38ns | 13.12ns |
| 4 | 2.58ns | 8.06ns | 8.44ns | 36.99ns | 12.64ns |
| 16 | 2.56ns | 8.03ns | 13.96ns | 37.22ns | 12.52ns |
| 64 | 2.62ns | 8.24ns | 14.10ns | 38.38ns | 12.69ns |
| 256 | 2.57ns | 8.64ns | 14.25ns | 39.83ns | 12.58ns |
| 1k | 2.69ns | 12.88ns | 18.49ns | 41.28ns | 12.50ns |
| 16k | 6.17ns | 40.56ns | 31.54ns | 59.23ns | 17.29ns |
| 256k | 15.14ns | 231.83ns | 142.03ns | 224.63ns | 96.49ns |
| 1M | 49.06ns | 294.73ns | 195.64ns | 285.56ns | 136.66ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/996c6fe2-ca21-41b2-8f32-c7387a7b6e88)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 231.79ns | 230.38ns | 884.45ns | 290.30ns | 336.13ns | 529.49ns |
| 4 | 93.70ns | 64.26ns | 428.11ns | 144.50ns | 174.22ns | 218.51ns |
| 16 | 55.46ns | 22.89ns | 319.59ns | 128.12ns | 129.37ns | 139.97ns |
| 64 | 44.52ns | 12.89ns | 266.02ns | 115.72ns | 107.54ns | 126.04ns |
| 256 | 41.59ns | 10.79ns | 214.21ns | 96.64ns | 92.78ns | 89.64ns |
| 1k | 36.21ns | 10.49ns | 201.64ns | 93.73ns | 265.12ns | 78.30ns |
| 16k | 30.78ns | 8.48ns | 228.06ns | 96.51ns | 273.51ns | 74.57ns |
| 256k | 30.62ns | 8.44ns | 205.18ns | 154.82ns | 422.46ns | 75.33ns |
| 1M | 31.67ns | 8.63ns | 204.04ns | 199.93ns | 455.14ns | 76.95ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/780ab5be-f2b1-404a-890b-985b2c0aa3f8)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.22us | 7.12us | 4.15us | 15.95us | 2.81us | 1.10us |
| 4 | 1.98us | 1.94us | 1.32us | 4.04us | 1.08us | 501.94ns |
| 16 | 533.03ns | 495.48ns | 616.15ns | 1.24us | 374.90ns | 248.01ns |
| 64 | 169.86ns | 138.63ns | 402.21ns | 476.71ns | 214.35ns | 153.18ns |
| 256 | 77.78ns | 45.74ns | 307.47ns | 247.02ns | 166.83ns | 118.81ns |
| 1k | 46.69ns | 21.41ns | 294.11ns | 162.30ns | 154.27ns | 111.73ns |
| 16k | 54.10ns | 40.99ns | 404.75ns | 225.72ns | 177.56ns | 139.39ns |
| 256k | 81.92ns | 41.59ns | 513.18ns | 702.65ns | 291.02ns | 162.40ns |
| 1M | 77.84ns | 35.17ns | 476.15ns | 1.93us | 359.62ns | 170.36ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/11cb313b-309a-453f-a1e7-fd74ec237036)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 277.25ns | 277.23ns | 1.52us | 403.94ns | 524.35ns | 1.76us |
| 4 | 125.10ns | 76.17ns | 951.37ns | 230.25ns | 323.47ns | 1.32us |
| 16 | 86.57ns | 26.05ns | 880.32ns | 177.58ns | 274.03ns | 1.08us |
| 64 | 82.29ns | 13.84ns | 646.67ns | 176.34ns | 243.51ns | 809.58ns |
| 256 | 67.79ns | 10.33ns | 580.20ns | 162.90ns | 187.71ns | 803.08ns |
| 1k | 57.32ns | 10.67ns | 590.91ns | 143.40ns | 422.58ns | 797.11ns |
| 16k | 56.00ns | 8.92ns | 652.37ns | 145.22ns | 440.64ns | 772.04ns |
| 256k | 55.33ns | 8.47ns | 559.30ns | 240.33ns | 591.82ns | 750.70ns |
| 1M | 55.98ns | 8.86ns | 535.68ns | 259.17ns | 617.34ns | 764.55ns |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/a1562226-99a3-4352-a135-1ab11ce3725c)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 91.49ns | 190.59ns | 419.14ns | 501.11ns | 335.42ns | 263.07ns |
| 4 | 99.16ns | 58.01ns | 436.54ns | 512.46ns | 344.76ns | 272.20ns |
| 16 | 101.74ns | 27.18ns | 447.08ns | 552.16ns | 351.86ns | 269.78ns |
| 64 | 100.65ns | 21.47ns | 435.97ns | 552.79ns | 362.92ns | 269.98ns |
| 256 | 99.20ns | 11.20ns | 439.33ns | 551.96ns | 374.86ns | 270.09ns |
| 1k | 100.19ns | 9.19ns | 442.50ns | 567.51ns | 740.31ns | 270.96ns |
| 16k | 101.75ns | 9.90ns | 501.43ns | 604.68ns | 916.01ns | 273.86ns |
| 256k | 102.79ns | 11.69ns | 481.89ns | 946.68ns | 1.38us | 262.96ns |
| 1M | 102.82ns | 12.56ns | 469.35ns | 1.04us | 1.48us | 263.69ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/2b10748e-34fe-4092-8b2c-f0d0838969f9)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 349.79ns | 462.92ns | 1.07us | 933.60ns | 780.55ns | 1.53us |
| 4 | 397.75ns | 182.46ns | 1.10us | 964.83ns | 781.18ns | 1.44us |
| 16 | 386.16ns | 115.97ns | 1.08us | 987.26ns | 776.58ns | 1.41us |
| 64 | 388.29ns | 95.23ns | 1.07us | 956.30ns | 807.85ns | 1.38us |
| 256 | 401.63ns | 37.53ns | 1.10us | 965.59ns | 855.14ns | 1.47us |
| 1k | 416.21ns | 21.54ns | 1.17us | 980.68ns | 1.41us | 1.41us |
| 16k | 434.71ns | 28.14ns | 1.32us | 1.11us | 1.50us | 1.50us |
| 256k | 599.14ns | 47.53ns | 1.45us | 1.65us | 2.08us | 1.73us |
| 1M | 508.32ns | 48.36ns | 1.41us | 1.69us | 2.13us | 1.70us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/f996bfba-58ca-4156-8035-c031fa0327b5)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 165.48ns | 265.99ns | 251.52ns | 378.74ns | 178.09ns | 298.65ns |
| 4 | 79.45ns | 80.68ns | 112.50ns | 163.79ns | 77.43ns | 134.48ns |
| 16 | 48.46ns | 32.00ns | 71.21ns | 141.72ns | 53.30ns | 94.31ns |
| 64 | 38.23ns | 19.41ns | 62.14ns | 123.21ns | 42.94ns | 68.13ns |
| 256 | 40.07ns | 9.17ns | 56.04ns | 109.83ns | 46.33ns | 64.95ns |
| 1k | 30.15ns | 5.86ns | 51.56ns | 105.47ns | 37.05ns | 62.22ns |
| 16k | 25.96ns | 5.68ns | 52.64ns | 120.36ns | 50.04ns | 59.11ns |
| 256k | 25.77ns | 5.65ns | 56.64ns | 165.65ns | 90.84ns | 59.54ns |
| 1M | 25.34ns | 6.63ns | 51.65ns | 300.75ns | 214.00ns | 60.29ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/8f2e9b15-7e43-4f14-8ed0-869bf78d5ba9)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 254.36ns | 372.00ns | 439.82ns | 505.39ns | 176.09ns | 653.92ns |
| 4 | 166.29ns | 131.09ns | 210.60ns | 246.62ns | 78.81ns | 410.40ns |
| 16 | 140.30ns | 71.77ns | 161.67ns | 227.15ns | 48.31ns | 356.93ns |
| 64 | 119.92ns | 60.71ns | 144.96ns | 218.06ns | 43.74ns | 303.94ns |
| 256 | 116.33ns | 18.79ns | 129.13ns | 168.20ns | 40.41ns | 277.85ns |
| 1k | 82.75ns | 8.70ns | 137.53ns | 162.40ns | 35.45ns | 260.91ns |
| 16k | 81.94ns | 7.90ns | 122.33ns | 177.61ns | 50.38ns | 237.99ns |
| 256k | 91.33ns | 13.32ns | 138.68ns | 286.03ns | 98.00ns | 252.75ns |
| 1M | 89.81ns | 14.49ns | 132.03ns | 390.62ns | 199.32ns | 266.58ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 21.32us | 1.97us | 231.24us | 1.83us | 16.27us |


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
