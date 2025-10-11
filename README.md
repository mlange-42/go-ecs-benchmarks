# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

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

Last run: Thu, 02 Oct 2025 20:54:43 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

Also note that Go's benchmarking prevents some compiler optimizations.
Hence, benchmark results are conservative and performance may be better in real-world use cases.

All components used in the benchmarks have two `float64` fields.
The initial capacity of the world is set to 1024 where this is supported.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/bad1e06b-4849-4993-aebb-5c0ca4de79ba)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 57.85ns | 50.47ns | 62.40ns | 45.62ns | 16.23ns | 178.72ns |
| 4 | 17.36ns | 14.65ns | 29.52ns | 15.82ns | 6.40ns | 48.42ns |
| 16 | 6.53ns | 6.10ns | 20.49ns | 10.01ns | 3.99ns | 15.69ns |
| 64 | 4.27ns | 4.16ns | 18.51ns | 8.79ns | 3.44ns | 7.60ns |
| 256 | 3.64ns | 3.62ns | 19.05ns | 8.25ns | 3.20ns | 5.48ns |
| 1k | 3.49ns | 3.54ns | 19.20ns | 8.19ns | 3.15ns | 4.93ns |
| 16k | 3.48ns | 3.49ns | 20.13ns | 8.13ns | 3.18ns | 4.77ns |
| 256k | 3.48ns | 3.49ns | 21.50ns | 8.14ns | 3.18ns | 4.78ns |
| 1M | 3.49ns | 3.50ns | 25.51ns | 8.14ns | 3.20ns | 4.76ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/540a4560-3393-4edc-a1b1-6ef1979a9e08)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 57.84ns | 49.85ns | 62.18ns | 45.14ns | 16.22ns | 194.53ns |
| 4 | 29.43ns | 21.58ns | 32.14ns | 20.72ns | 14.63ns | 99.10ns |
| 16 | 21.99ns | 14.91ns | 26.24ns | 14.99ns | 28.37ns | 72.16ns |
| 64 | 13.09ns | 8.84ns | 22.11ns | 11.14ns | 14.53ns | 35.68ns |
| 256 | 5.68ns | 4.58ns | 21.29ns | 9.19ns | 6.04ns | 12.92ns |
| 1k | 4.02ns | 3.71ns | 22.16ns | 8.80ns | 3.94ns | 7.06ns |
| 16k | 3.57ns | 3.55ns | 29.30ns | 8.51ns | 3.28ns | 5.04ns |
| 256k | 3.50ns | 3.49ns | 65.23ns | 8.45ns | 3.20ns | 4.78ns |
| 1M | 3.53ns | 3.51ns | 100.88ns | 8.47ns | 3.22ns | 4.78ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/fd6bf4fe-f34a-4a61-bc18-c4486d106c00)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 58.11ns | 50.12ns | 61.68ns | 56.93ns | 16.21ns | 207.65ns |
| 4 | 16.57ns | 14.90ns | 28.86ns | 27.28ns | 9.58ns | 78.85ns |
| 16 | 6.66ns | 6.14ns | 22.00ns | 21.52ns | 4.66ns | 45.42ns |
| 64 | 4.52ns | 4.15ns | 18.63ns | 20.08ns | 3.59ns | 43.01ns |
| 256 | 3.95ns | 3.62ns | 19.08ns | 11.26ns | 3.24ns | 14.26ns |
| 1k | 3.85ns | 3.50ns | 19.43ns | 8.91ns | 3.16ns | 7.07ns |
| 16k | 3.80ns | 3.49ns | 21.89ns | 8.21ns | 3.18ns | 4.97ns |
| 256k | 3.79ns | 3.49ns | 21.97ns | 8.14ns | 3.22ns | 4.76ns |
| 1M | 3.81ns | 3.49ns | 26.53ns | 8.15ns | 3.20ns | 4.78ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/6e7e3c1c-092a-480a-855f-4c8d0a24e267)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 4.95ns | 9.14ns | 8.45ns | 37.23ns | 38.52ns |
| 4 | 4.57ns | 8.75ns | 8.24ns | 36.96ns | 38.60ns |
| 16 | 4.44ns | 8.67ns | 14.18ns | 37.13ns | 39.31ns |
| 64 | 4.49ns | 8.87ns | 13.81ns | 38.59ns | 39.04ns |
| 256 | 4.44ns | 8.99ns | 14.01ns | 40.44ns | 39.81ns |
| 1k | 4.57ns | 13.52ns | 17.97ns | 41.51ns | 41.82ns |
| 16k | 7.55ns | 36.18ns | 29.02ns | 58.59ns | 64.22ns |
| 256k | 13.79ns | 140.54ns | 52.61ns | 136.21ns | 201.90ns |
| 1M | 64.96ns | 257.07ns | 167.69ns | 247.43ns | 309.79ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/01d52fb5-13a3-4a3b-b9ff-7c02bd1af6f1)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 252.34ns | 278.46ns | 1.01us | 393.66ns | 424.61ns | 823.02ns |
| 4 | 114.99ns | 75.58ns | 498.21ns | 165.38ns | 203.47ns | 395.17ns |
| 16 | 67.32ns | 25.11ns | 287.53ns | 129.55ns | 134.41ns | 254.89ns |
| 64 | 54.79ns | 12.97ns | 256.78ns | 116.72ns | 120.90ns | 219.98ns |
| 256 | 58.85ns | 10.16ns | 212.08ns | 107.41ns | 95.64ns | 198.84ns |
| 1k | 46.33ns | 8.69ns | 196.11ns | 115.59ns | 272.78ns | 209.08ns |
| 16k | 40.13ns | 7.56ns | 209.98ns | 117.59ns | 285.51ns | 216.53ns |
| 256k | 40.37ns | 7.82ns | 207.01ns | 127.10ns | 319.38ns | 257.19ns |
| 1M | 40.47ns | 7.63ns | 188.12ns | 190.90ns | 428.36ns | 352.27ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/7e9c7ea1-b800-454d-96fb-6ca879296257)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.00us | 7.97us | 4.45us | 19.38us | 3.17us | 1.84us |
| 4 | 2.10us | 1.95us | 1.36us | 5.00us | 1.05us | 805.39ns |
| 16 | 554.15ns | 514.05ns | 628.96ns | 1.39us | 383.95ns | 423.12ns |
| 64 | 184.50ns | 135.87ns | 404.76ns | 494.52ns | 224.77ns | 288.24ns |
| 256 | 86.30ns | 44.48ns | 300.09ns | 255.69ns | 171.10ns | 255.71ns |
| 1k | 56.59ns | 20.15ns | 304.95ns | 244.13ns | 154.51ns | 219.76ns |
| 16k | 65.42ns | 38.06ns | 395.85ns | 258.18ns | 186.18ns | 392.18ns |
| 256k | 100.06ns | 45.86ns | 432.06ns | 605.70ns | 260.30ns | 453.41ns |
| 1M | 79.92ns | 28.74ns | 451.06ns | 1.35us | 305.77ns | 553.73ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/2b8bbac2-a2be-43f2-9eea-cac2142af236)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 367.73ns | 412.67ns | 1.86us | 479.85ns | 657.75ns | 3.05us |
| 4 | 169.92ns | 117.00ns | 1.21us | 253.33ns | 368.79ns | 1.87us |
| 16 | 115.47ns | 34.40ns | 985.80ns | 224.06ns | 301.52ns | 1.55us |
| 64 | 109.79ns | 15.74ns | 797.45ns | 190.69ns | 276.89ns | 1.39us |
| 256 | 93.15ns | 10.81ns | 713.92ns | 181.72ns | 229.22ns | 1.39us |
| 1k | 86.11ns | 8.67ns | 734.90ns | 162.53ns | 478.22ns | 1.40us |
| 16k | 81.24ns | 7.70ns | 754.09ns | 170.99ns | 485.36ns | 1.37us |
| 256k | 81.36ns | 7.78ns | 772.26ns | 171.56ns | 558.76ns | 1.46us |
| 1M | 81.11ns | 7.62ns | 759.29ns | 229.28ns | 643.40ns | 1.53us |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/84ad054e-92fe-4ae5-9d7e-1c38edc8bf9c)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 100.99ns | 496.22ns | 421.58ns | 514.61ns | 340.52ns | 441.44ns |
| 4 | 112.76ns | 133.56ns | 435.68ns | 523.06ns | 351.66ns | 451.88ns |
| 16 | 115.11ns | 44.37ns | 440.44ns | 561.19ns | 354.19ns | 447.27ns |
| 64 | 111.58ns | 22.85ns | 433.79ns | 545.95ns | 369.50ns | 447.49ns |
| 256 | 111.56ns | 10.24ns | 432.18ns | 552.25ns | 378.80ns | 453.02ns |
| 1k | 117.24ns | 7.00ns | 434.01ns | 571.75ns | 740.21ns | 460.31ns |
| 16k | 112.67ns | 7.93ns | 492.54ns | 603.05ns | 798.73ns | 510.04ns |
| 256k | 112.83ns | 9.55ns | 478.57ns | 824.94ns | 1.32us | 832.72ns |
| 1M | 112.42ns | 10.30ns | 468.15ns | 987.98ns | 1.44us | 986.77ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/a3826438-1601-48fd-9c6b-802e252699bd)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 346.99ns | 760.90ns | 1.03us | 993.10ns | 806.44ns | 1.83us |
| 4 | 402.21ns | 255.29ns | 1.06us | 1.02us | 807.24ns | 1.77us |
| 16 | 392.15ns | 133.93ns | 1.04us | 1.04us | 835.30ns | 1.73us |
| 64 | 394.75ns | 98.50ns | 1.04us | 1.03us | 841.08ns | 1.75us |
| 256 | 386.36ns | 35.98ns | 1.05us | 1.04us | 884.15ns | 1.74us |
| 1k | 399.30ns | 19.22ns | 1.11us | 1.06us | 1.43us | 1.86us |
| 16k | 415.55ns | 26.82ns | 1.24us | 1.15us | 1.47us | 2.15us |
| 256k | 600.10ns | 47.77ns | 1.30us | 1.65us | 2.06us | 2.92us |
| 1M | 508.95ns | 46.71ns | 1.46us | 1.78us | 2.21us | 2.94us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/b15a5082-e55a-4671-9a4d-0d3e3d121351)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 155.88ns | 724.96ns | 261.06ns | 355.10ns | 204.99ns | 462.39ns |
| 4 | 75.48ns | 203.77ns | 111.90ns | 162.99ns | 85.26ns | 231.12ns |
| 16 | 42.08ns | 60.55ns | 71.38ns | 137.40ns | 51.45ns | 175.39ns |
| 64 | 35.73ns | 28.83ns | 59.51ns | 108.96ns | 44.17ns | 146.29ns |
| 256 | 31.05ns | 11.87ns | 56.66ns | 106.59ns | 44.31ns | 131.51ns |
| 1k | 29.43ns | 7.23ns | 51.37ns | 104.31ns | 35.84ns | 130.03ns |
| 16k | 25.13ns | 6.40ns | 52.29ns | 116.22ns | 50.38ns | 154.30ns |
| 256k | 24.26ns | 6.34ns | 53.72ns | 161.54ns | 101.20ns | 247.52ns |
| 1M | 24.56ns | 7.46ns | 55.44ns | 289.13ns | 216.68ns | 356.69ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/34268a33-a978-4a32-8e71-7373502ba119)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 229.15ns | 891.11ns | 420.12ns | 524.77ns | 182.16ns | 846.91ns |
| 4 | 153.43ns | 278.74ns | 213.60ns | 262.49ns | 80.64ns | 507.87ns |
| 16 | 113.58ns | 108.11ns | 162.10ns | 238.49ns | 50.06ns | 468.63ns |
| 64 | 127.99ns | 71.69ns | 158.36ns | 191.47ns | 42.19ns | 334.24ns |
| 256 | 87.16ns | 21.27ns | 124.53ns | 160.23ns | 38.82ns | 312.59ns |
| 1k | 79.17ns | 9.57ns | 145.00ns | 166.62ns | 35.00ns | 314.43ns |
| 16k | 77.39ns | 8.61ns | 142.24ns | 179.67ns | 50.26ns | 334.87ns |
| 256k | 90.10ns | 14.91ns | 159.28ns | 299.00ns | 94.69ns | 505.80ns |
| 1M | 88.21ns | 15.57ns | 134.15ns | 401.77ns | 204.70ns | 644.24ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 15.30us | 2.33us | 213.60us | 2.14us | 32.74us |


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
