# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Tested | Latest | Activity |
|-----|--------|--------|----------|
| [Ark](https://github.com/mlange-42/ark) | v0.8.0 | ![GitHub Tag](https://img.shields.io/github/v/tag/mlange-42/ark?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/mlange-42/ark?label=date) | ![Last commit](https://img.shields.io/github/last-commit/mlange-42/ark) |
| [Donburi](https://github.com/yohamta0/donburi-ecs) | v1.15.7 | ![GitHub Tag](https://img.shields.io/github/v/tag/yohamta0/donburi-ecs?color=blue) ![GitHub Release Date](https://img.shields.io/github/release-date/yohamta0/donburi-ecs?label=date) | ![Last commit](https://img.shields.io/github/last-commit/yohamta0/donburi-ecs) |
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
| [Donburi](https://github.com/yohamta0/donburi-ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ❌ |
| [go‑gameengine‑ecs](https://github.com/marioolofo/go-gameengine-ecs) | ❌ | ✅ | ❌ | ❌ | ❌ | ❌ |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |
| [Volt](https://github.com/akmonengine/volt) | ✅ | ❌ | ❌ | ✅ | ❌ | ❌ |

[1] ECS lifecycle events, allowing to react to entity creation, component addition, ...  
[2] Faster batch operations for entity creation etc.

## Benchmarks

Last run: Tue, 07 Apr 2026 09:27:38 UTC  
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

![query2comp](https://github.com/user-attachments/assets/68e3463d-bc1f-4270-80ed-d27e98c54a75)

| N | Ark | Ark (tables) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 61.38ns | 68.42ns | 63.41ns | 48.11ns | 18.12ns | 77.89ns |
| 4 | 17.42ns | 17.81ns | 29.52ns | 15.44ns | 6.88ns | 20.30ns |
| 16 | 5.77ns | 4.99ns | 21.05ns | 7.60ns | 4.06ns | 6.04ns |
| 64 | 2.75ns | 1.94ns | 19.13ns | 5.69ns | 3.47ns | 2.49ns |
| 256 | 2.11ns | 1.07ns | 19.40ns | 5.17ns | 3.21ns | 1.60ns |
| 1k | 1.91ns | 0.90ns | 19.37ns | 5.08ns | 3.19ns | 1.36ns |
| 16k | 1.94ns | 0.81ns | 21.08ns | 5.05ns | 3.19ns | 1.28ns |
| 256k | 1.94ns | 0.80ns | 25.17ns | 5.05ns | 3.20ns | 1.27ns |
| 1M | 2.00ns | 0.93ns | 28.82ns | 5.06ns | 3.20ns | 1.34ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/0f30fa7e-6444-4954-9bef-166130fc3810)

| N | Ark | Ark (tables) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 63.91ns | 69.01ns | 63.18ns | 45.56ns | 17.17ns | 82.00ns |
| 4 | 31.67ns | 38.69ns | 32.31ns | 18.98ns | 14.95ns | 66.92ns |
| 16 | 23.53ns | 31.71ns | 25.25ns | 12.51ns | 29.26ns | 63.62ns |
| 64 | 12.74ns | 16.48ns | 22.59ns | 8.62ns | 17.63ns | 31.37ns |
| 256 | 4.81ns | 4.73ns | 21.29ns | 5.73ns | 6.14ns | 9.42ns |
| 1k | 2.62ns | 1.85ns | 22.89ns | 5.21ns | 4.00ns | 3.50ns |
| 16k | 1.95ns | 0.99ns | 30.24ns | 5.10ns | 3.29ns | 1.58ns |
| 256k | 1.79ns | 0.84ns | 69.28ns | 5.27ns | 3.19ns | 1.28ns |
| 1M | 1.84ns | 1.01ns | 91.05ns | 5.07ns | 3.19ns | 1.35ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/c2280dab-4ce0-4b85-af42-4a4848e11e8e)

| N | Ark | Ark (tables) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 61.81ns | 68.56ns | 63.66ns | 55.75ns | 17.17ns | 99.44ns |
| 4 | 17.45ns | 17.67ns | 29.97ns | 25.62ns | 9.67ns | 47.42ns |
| 16 | 5.80ns | 4.97ns | 21.13ns | 18.81ns | 4.72ns | 35.11ns |
| 64 | 2.72ns | 1.94ns | 19.42ns | 17.05ns | 3.67ns | 36.96ns |
| 256 | 1.96ns | 1.03ns | 20.20ns | 8.05ns | 3.26ns | 10.28ns |
| 1k | 1.82ns | 0.89ns | 19.90ns | 5.83ns | 3.18ns | 3.53ns |
| 16k | 1.78ns | 0.84ns | 21.35ns | 5.11ns | 3.18ns | 1.44ns |
| 256k | 1.78ns | 0.83ns | 22.40ns | 5.05ns | 3.18ns | 1.28ns |
| 1M | 1.85ns | 0.97ns | 25.57ns | 5.06ns | 3.19ns | 1.34ns |


> **Note:** Donburi, unitoftime/ecs, and Volt use a callback-based approach for their query loops.
As a result, iteration speed may degrade if the callback contains complex logic
and the Go compiler is unable to inline it.

### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/2ebd86de-618f-4729-bf45-5dcc843338a1)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 2.80ns | 8.33ns | 8.45ns | 37.23ns | 13.13ns |
| 4 | 2.77ns | 8.14ns | 8.50ns | 37.27ns | 12.66ns |
| 16 | 2.77ns | 8.10ns | 14.05ns | 37.40ns | 12.58ns |
| 64 | 2.77ns | 8.17ns | 14.46ns | 38.74ns | 12.63ns |
| 256 | 2.79ns | 8.69ns | 14.30ns | 40.46ns | 12.55ns |
| 1k | 2.87ns | 14.30ns | 18.61ns | 41.42ns | 12.53ns |
| 16k | 5.22ns | 35.95ns | 30.51ns | 59.65ns | 17.07ns |
| 256k | 8.69ns | 61.22ns | 41.85ns | 81.86ns | 21.50ns |
| 1M | 42.68ns | 245.76ns | 163.21ns | 230.99ns | 108.86ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/175bde67-2565-4391-9d03-547365f63e5e)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 768.52ns | 793.00ns | 1.84us | 1.07us | 1.12us | 1.34us |
| 4 | 168.99ns | 157.81ns | 742.34ns | 326.38ns | 383.96ns | 401.65ns |
| 16 | 94.77ns | 48.51ns | 411.59ns | 190.27ns | 167.87ns | 176.85ns |
| 64 | 56.36ns | 24.22ns | 318.33ns | 137.88ns | 112.46ns | 136.58ns |
| 256 | 46.57ns | 10.65ns | 214.00ns | 120.45ns | 99.94ns | 111.67ns |
| 1k | 42.07ns | 10.55ns | 194.67ns | 116.36ns | 287.95ns | 93.35ns |
| 16k | 35.21ns | 8.30ns | 218.68ns | 118.13ns | 287.44ns | 83.96ns |
| 256k | 34.97ns | 8.47ns | 204.38ns | 123.87ns | 331.32ns | 84.58ns |
| 1M | 35.19ns | 8.48ns | 186.27ns | 183.37ns | 425.46ns | 85.04ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/84cb0a0a-f2a6-4eaa-816d-2cae03a727a4)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.68us | 9.80us | 5.29us | 20.29us | 3.62us | 1.64us |
| 4 | 2.41us | 2.40us | 1.61us | 5.17us | 1.23us | 650.57ns |
| 16 | 650.77ns | 610.12ns | 645.02ns | 1.47us | 422.98ns | 282.04ns |
| 64 | 204.99ns | 161.41ns | 463.78ns | 574.39ns | 242.95ns | 174.39ns |
| 256 | 86.12ns | 52.36ns | 362.60ns | 287.86ns | 177.33ns | 154.21ns |
| 1k | 53.77ns | 22.74ns | 300.31ns | 254.01ns | 164.46ns | 126.38ns |
| 16k | 63.34ns | 44.71ns | 435.48ns | 273.24ns | 199.49ns | 155.85ns |
| 256k | 76.58ns | 33.05ns | 417.97ns | 570.50ns | 224.83ns | 161.77ns |
| 1M | 51.84ns | 21.10ns | 386.60ns | 1.13us | 237.03ns | 117.95ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/cf44b325-1f45-4186-b115-620bed59540c)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 680.42ns | 804.11ns | 2.54us | 1.18us | 1.23us | 3.42us |
| 4 | 260.08ns | 210.44ns | 1.51us | 366.45ns | 501.74ns | 1.88us |
| 16 | 129.23ns | 71.45ns | 1.14us | 282.53ns | 322.53ns | 1.15us |
| 64 | 89.67ns | 25.07ns | 775.68ns | 193.12ns | 260.48ns | 1.07us |
| 256 | 83.19ns | 11.36ns | 743.57ns | 156.51ns | 235.07ns | 895.16ns |
| 1k | 71.10ns | 11.38ns | 724.80ns | 164.67ns | 461.52ns | 889.63ns |
| 16k | 68.59ns | 8.02ns | 754.74ns | 168.14ns | 483.89ns | 890.81ns |
| 256k | 68.25ns | 7.92ns | 682.06ns | 182.08ns | 558.36ns | 868.40ns |
| 1M | 68.08ns | 7.90ns | 729.42ns | 250.23ns | 665.04ns | 870.63ns |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/fbf7e9d8-0aad-4386-839e-310896b83be5)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 92.53ns | 179.45ns | 420.07ns | 511.63ns | 344.92ns | 262.96ns |
| 4 | 102.92ns | 47.82ns | 430.85ns | 511.70ns | 348.93ns | 265.55ns |
| 16 | 108.45ns | 18.54ns | 438.58ns | 545.68ns | 355.32ns | 268.17ns |
| 64 | 101.83ns | 11.36ns | 432.29ns | 550.41ns | 354.06ns | 265.91ns |
| 256 | 101.82ns | 9.14ns | 433.20ns | 553.37ns | 376.49ns | 268.57ns |
| 1k | 110.24ns | 8.75ns | 434.75ns | 565.15ns | 746.33ns | 266.10ns |
| 16k | 106.99ns | 9.51ns | 486.67ns | 599.71ns | 804.85ns | 270.02ns |
| 256k | 112.42ns | 10.25ns | 482.58ns | 719.86ns | 970.52ns | 263.63ns |
| 1M | 104.26ns | 11.03ns | 473.68ns | 924.37ns | 1.32us | 266.03ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/3604eb16-31aa-4a00-9b1a-68912f8a074d)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 353.30ns | 475.55ns | 1.03us | 927.45ns | 811.55ns | 1.39us |
| 4 | 398.57ns | 125.90ns | 1.05us | 959.30ns | 814.05ns | 1.38us |
| 16 | 400.64ns | 49.19ns | 1.03us | 978.11ns | 815.63ns | 1.35us |
| 64 | 408.90ns | 25.95ns | 1.03us | 971.84ns | 847.99ns | 1.35us |
| 256 | 386.26ns | 22.13ns | 1.05us | 963.19ns | 883.21ns | 1.38us |
| 1k | 425.20ns | 21.04ns | 1.14us | 993.11ns | 1.43us | 1.38us |
| 16k | 416.54ns | 22.52ns | 1.20us | 1.05us | 1.47us | 1.43us |
| 256k | 530.23ns | 39.17ns | 1.32us | 1.45us | 1.79us | 1.54us |
| 1M | 465.03ns | 39.90ns | 1.38us | 1.65us | 2.04us | 1.56us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/04f7691a-6bc5-46e9-b3e2-f07b688267dc)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 234.40ns | 350.58ns | 351.68ns | 440.62ns | 254.82ns | 309.61ns |
| 4 | 95.97ns | 94.78ns | 131.63ns | 177.03ns | 103.96ns | 132.90ns |
| 16 | 48.98ns | 27.41ns | 70.89ns | 158.15ns | 57.25ns | 96.70ns |
| 64 | 39.98ns | 11.06ns | 61.95ns | 129.34ns | 47.12ns | 78.14ns |
| 256 | 30.88ns | 7.43ns | 55.89ns | 100.34ns | 43.49ns | 71.06ns |
| 1k | 27.53ns | 6.06ns | 52.43ns | 107.89ns | 37.17ns | 62.73ns |
| 16k | 24.28ns | 5.40ns | 52.04ns | 116.24ns | 48.86ns | 59.23ns |
| 256k | 23.66ns | 5.44ns | 60.52ns | 159.15ns | 83.94ns | 62.22ns |
| 1M | 23.68ns | 6.15ns | 54.51ns | 289.87ns | 189.84ns | 60.38ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/6f564e2a-f79b-4f4c-83bd-96bc1b891849)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 309.31ns | 400.13ns | 485.00ns | 545.66ns | 238.33ns | 681.73ns |
| 4 | 171.10ns | 110.74ns | 221.85ns | 259.77ns | 98.26ns | 436.91ns |
| 16 | 122.39ns | 34.07ns | 156.22ns | 236.81ns | 57.91ns | 356.29ns |
| 64 | 124.54ns | 14.28ns | 144.84ns | 179.58ns | 41.34ns | 338.06ns |
| 256 | 96.67ns | 10.25ns | 121.86ns | 177.67ns | 42.56ns | 263.42ns |
| 1k | 81.62ns | 8.56ns | 135.73ns | 165.21ns | 34.91ns | 241.83ns |
| 16k | 79.41ns | 7.98ns | 124.08ns | 175.73ns | 50.77ns | 238.25ns |
| 256k | 90.57ns | 11.89ns | 133.59ns | 287.50ns | 80.97ns | 246.44ns |
| 1M | 80.86ns | 12.52ns | 206.25ns | 372.85ns | 164.78ns | 245.25ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 27.91us | 2.74us | 258.71us | 2.65us | 23.30us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/ark&yohamta0/donburi-ecs&marioolofo/go-gameengine-ecs&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/ark,yohamta0/donburi-ecs,marioolofo/go-gameengine-ecs,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
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
