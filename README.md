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

Last run: Wed, 15 Oct 2025 12:42:37 UTC  
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

![query2comp](https://github.com/user-attachments/assets/975c3ecc-b62f-4060-a760-41c02f24a334)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 65.73ns | 59.37ns | 62.09ns | 47.33ns | 18.20ns | 112.30ns |
| 4 | 19.13ns | 16.34ns | 29.09ns | 15.25ns | 7.12ns | 28.74ns |
| 16 | 5.95ns | 5.45ns | 20.82ns | 7.76ns | 4.36ns | 8.03ns |
| 64 | 2.87ns | 2.67ns | 19.04ns | 6.00ns | 3.75ns | 2.95ns |
| 256 | 2.03ns | 1.99ns | 19.28ns | 5.48ns | 3.51ns | 1.70ns |
| 1k | 1.97ns | 1.94ns | 19.24ns | 5.36ns | 3.52ns | 1.39ns |
| 16k | 1.94ns | 1.94ns | 21.12ns | 5.34ns | 3.48ns | 1.27ns |
| 256k | 1.93ns | 1.93ns | 24.10ns | 5.34ns | 3.51ns | 1.28ns |
| 1M | 2.03ns | 2.03ns | 36.71ns | 5.43ns | 3.50ns | 1.32ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/2f8a810d-46e3-43d0-a4b4-943aaae34a59)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.85ns | 58.04ns | 62.20ns | 44.34ns | 17.85ns | 114.50ns |
| 4 | 33.39ns | 27.02ns | 32.55ns | 18.90ns | 15.06ns | 74.15ns |
| 16 | 23.84ns | 15.87ns | 25.58ns | 12.42ns | 25.06ns | 61.69ns |
| 64 | 13.17ns | 9.61ns | 22.00ns | 8.36ns | 14.31ns | 30.55ns |
| 256 | 4.92ns | 3.77ns | 20.69ns | 5.85ns | 6.23ns | 9.10ns |
| 1k | 2.70ns | 2.34ns | 22.49ns | 5.47ns | 4.47ns | 3.42ns |
| 16k | 2.07ns | 1.89ns | 29.56ns | 5.39ns | 3.57ns | 1.48ns |
| 256k | 1.93ns | 1.77ns | 66.05ns | 5.37ns | 3.49ns | 1.28ns |
| 1M | 2.00ns | 1.83ns | 92.11ns | 5.36ns | 3.51ns | 1.33ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/737ed68e-16a9-476e-87cd-360eed48ae13)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 65.74ns | 59.35ns | 62.79ns | 54.89ns | 17.80ns | 136.21ns |
| 4 | 17.89ns | 16.30ns | 29.13ns | 26.63ns | 9.57ns | 63.16ns |
| 16 | 5.87ns | 5.36ns | 20.75ns | 18.74ns | 4.87ns | 48.41ns |
| 64 | 2.71ns | 2.57ns | 19.18ns | 17.39ns | 3.94ns | 50.94ns |
| 256 | 1.96ns | 1.91ns | 19.81ns | 8.46ns | 3.59ns | 13.81ns |
| 1k | 1.84ns | 1.80ns | 19.57ns | 6.11ns | 3.49ns | 4.43ns |
| 16k | 1.77ns | 1.77ns | 20.87ns | 5.41ns | 3.46ns | 1.50ns |
| 256k | 1.76ns | 1.76ns | 22.24ns | 5.35ns | 3.48ns | 1.28ns |
| 1M | 1.82ns | 1.82ns | 26.03ns | 5.37ns | 3.50ns | 1.36ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/de9f5e9a-f09d-4af7-a73c-ae196859121f)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 5.19ns | 10.34ns | 10.61ns | 39.13ns | 40.34ns |
| 4 | 3.31ns | 8.42ns | 8.81ns | 37.25ns | 39.11ns |
| 16 | 2.95ns | 8.34ns | 13.93ns | 36.93ns | 39.35ns |
| 64 | 2.91ns | 8.04ns | 13.95ns | 38.70ns | 39.03ns |
| 256 | 2.85ns | 8.60ns | 13.96ns | 39.92ns | 40.13ns |
| 1k | 2.92ns | 12.12ns | 17.49ns | 41.43ns | 41.53ns |
| 16k | 6.02ns | 36.56ns | 29.23ns | 57.87ns | 62.96ns |
| 256k | 8.56ns | 54.52ns | 42.00ns | 73.88ns | 92.11ns |
| 1M | 22.91ns | 211.81ns | 123.42ns | 204.29ns | 254.19ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/220233b1-41f6-4dfa-ac77-9a2ff35b519b)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 199.66ns | 199.18ns | 792.60ns | 245.28ns | 292.66ns | 631.44ns |
| 4 | 93.62ns | 57.92ns | 399.52ns | 131.60ns | 166.78ns | 328.13ns |
| 16 | 60.84ns | 22.85ns | 269.46ns | 129.66ns | 130.72ns | 217.24ns |
| 64 | 50.66ns | 14.21ns | 253.67ns | 126.48ns | 126.06ns | 211.33ns |
| 256 | 47.16ns | 12.14ns | 188.96ns | 118.75ns | 98.15ns | 170.09ns |
| 1k | 45.47ns | 11.46ns | 171.00ns | 119.99ns | 290.64ns | 193.25ns |
| 16k | 44.11ns | 11.41ns | 187.00ns | 124.39ns | 302.95ns | 205.07ns |
| 256k | 44.18ns | 11.40ns | 181.61ns | 148.76ns | 425.20ns | 328.61ns |
| 1M | 43.87ns | 11.41ns | 188.84ns | 246.85ns | 553.08ns | 389.67ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/e978873c-9e2e-42be-be36-c05e740383cc)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.41us | 7.62us | 4.26us | 17.77us | 2.73us | 1.35us |
| 4 | 1.86us | 1.78us | 1.31us | 4.33us | 901.32ns | 635.93ns |
| 16 | 503.24ns | 479.25ns | 573.67ns | 1.33us | 388.48ns | 345.35ns |
| 64 | 163.24ns | 126.78ns | 389.81ns | 493.77ns | 218.55ns | 256.45ns |
| 256 | 73.96ns | 43.50ns | 282.08ns | 248.91ns | 171.08ns | 233.70ns |
| 1k | 49.44ns | 19.93ns | 265.98ns | 239.35ns | 165.77ns | 213.56ns |
| 16k | 57.78ns | 43.25ns | 394.08ns | 259.56ns | 191.79ns | 382.39ns |
| 256k | 91.66ns | 52.20ns | 448.97ns | 647.56ns | 290.04ns | 536.22ns |
| 1M | 77.08ns | 29.81ns | 418.24ns | 1.26us | 361.99ns | 619.11ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/ccfc57a8-58d6-4c38-92eb-cfded5a9b609)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 271.80ns | 280.38ns | 1.60us | 331.23ns | 446.14ns | 2.33us |
| 4 | 148.46ns | 77.72ns | 1.08us | 200.50ns | 323.31ns | 2.06us |
| 16 | 112.64ns | 26.34ns | 1.07us | 198.79ns | 267.68ns | 1.50us |
| 64 | 100.98ns | 13.84ns | 761.57ns | 204.30ns | 231.79ns | 1.24us |
| 256 | 97.76ns | 10.86ns | 733.23ns | 175.06ns | 209.94ns | 1.27us |
| 1k | 98.36ns | 9.83ns | 721.46ns | 174.69ns | 477.72ns | 1.28us |
| 16k | 98.32ns | 8.93ns | 766.43ns | 176.04ns | 488.85ns | 1.31us |
| 256k | 98.76ns | 8.91ns | 732.28ns | 241.71ns | 704.33ns | 1.45us |
| 1M | 98.33ns | 11.09ns | 710.36ns | 277.48ns | 731.50ns | 1.50us |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/0c949bf4-2331-44b2-acdc-1b3c466c0ffb)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 96.35ns | 199.73ns | 414.06ns | 501.06ns | 335.56ns | 440.70ns |
| 4 | 99.74ns | 64.88ns | 429.37ns | 511.53ns | 338.43ns | 450.74ns |
| 16 | 102.22ns | 28.34ns | 441.79ns | 547.22ns | 346.37ns | 448.15ns |
| 64 | 100.86ns | 20.94ns | 439.67ns | 546.29ns | 355.08ns | 448.79ns |
| 256 | 99.64ns | 10.82ns | 436.35ns | 551.37ns | 367.74ns | 448.43ns |
| 1k | 100.45ns | 8.94ns | 440.28ns | 566.76ns | 732.76ns | 458.98ns |
| 16k | 102.02ns | 9.31ns | 499.05ns | 602.25ns | 796.23ns | 504.69ns |
| 256k | 102.44ns | 10.43ns | 476.47ns | 792.13ns | 1.26us | 755.78ns |
| 1M | 101.66ns | 12.29ns | 474.01ns | 1.04us | 1.51us | 1.09us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/96c83e8c-d828-4acb-930e-2935ea98e4f0)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 334.23ns | 466.63ns | 1.06us | 926.39ns | 813.93ns | 1.67us |
| 4 | 391.82ns | 188.77ns | 1.10us | 973.98ns | 827.04ns | 1.60us |
| 16 | 385.44ns | 113.58ns | 1.07us | 979.26ns | 832.52ns | 1.59us |
| 64 | 385.62ns | 94.94ns | 1.05us | 961.19ns | 851.55ns | 1.62us |
| 256 | 381.14ns | 36.49ns | 1.07us | 960.10ns | 897.93ns | 1.60us |
| 1k | 407.10ns | 20.72ns | 1.15us | 985.17ns | 1.44us | 1.71us |
| 16k | 416.80ns | 23.81ns | 1.25us | 1.06us | 1.48us | 1.77us |
| 256k | 582.71ns | 48.97ns | 1.40us | 1.69us | 2.11us | 2.81us |
| 1M | 504.37ns | 51.33ns | 1.48us | 1.78us | 2.18us | 2.87us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/a8e74ba5-7612-4b17-b524-feb79e42bb30)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 197.62ns | 306.44ns | 256.56ns | 354.99ns | 189.46ns | 415.48ns |
| 4 | 80.58ns | 92.20ns | 108.35ns | 163.74ns | 77.64ns | 218.88ns |
| 16 | 46.76ns | 32.02ns | 70.56ns | 136.17ns | 50.13ns | 181.09ns |
| 64 | 43.15ns | 19.30ns | 61.99ns | 118.99ns | 41.36ns | 142.44ns |
| 256 | 40.27ns | 9.08ns | 55.82ns | 113.01ns | 42.13ns | 128.85ns |
| 1k | 28.86ns | 6.62ns | 51.13ns | 104.32ns | 36.07ns | 131.05ns |
| 16k | 26.96ns | 5.38ns | 51.91ns | 115.86ns | 49.34ns | 152.23ns |
| 256k | 27.17ns | 5.46ns | 50.73ns | 168.52ns | 182.63ns | 397.42ns |
| 1M | 27.01ns | 6.70ns | 50.77ns | 312.75ns | 226.28ns | 367.45ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/18b08934-9e3e-4411-8944-cc9a3ff41285)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 265.57ns | 436.75ns | 425.02ns | 440.81ns | 171.79ns | 781.14ns |
| 4 | 158.23ns | 147.76ns | 208.04ns | 250.73ns | 73.05ns | 540.10ns |
| 16 | 147.56ns | 75.95ns | 164.92ns | 216.17ns | 49.69ns | 471.84ns |
| 64 | 108.71ns | 59.23ns | 142.85ns | 197.84ns | 45.61ns | 381.24ns |
| 256 | 103.75ns | 18.67ns | 124.10ns | 143.76ns | 40.87ns | 312.09ns |
| 1k | 85.00ns | 9.10ns | 137.64ns | 156.30ns | 34.90ns | 318.73ns |
| 16k | 84.52ns | 7.84ns | 142.16ns | 176.25ns | 53.27ns | 447.46ns |
| 256k | 91.58ns | 13.07ns | 143.79ns | 295.37ns | 94.59ns | 548.91ns |
| 1M | 91.57ns | 14.82ns | 237.33ns | 423.03ns | 220.87ns | 668.20ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 17.89us | 1.89us | 224.24us | 1.77us | 28.79us |


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
