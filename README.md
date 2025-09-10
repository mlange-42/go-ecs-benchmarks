# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche) and [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Version |
|-----|---------|
| [Arche](https://github.com/mlange-42/arche) | v0.15.3 |
| [Ark](https://github.com/mlange-42/ark) | v0.5.1 |
| [Donburi](https://github.com/yottahmd/donburi) | v1.15.7 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | v0.0.3 |
| [Volt](https://github.com/akmonengine/volt) | v1.6.0 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.
See the section on [Contributing](#contributing) for details.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Benchmarks

Last run: Wed, 10 Sep 2025 15:45:41 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

All components used in the benchmarks have two `float64` fields.
The initial capacity of the world is set to 1024 where this is supported.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/21f27565-60a5-4712-a379-b46c13650a0f)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 54.56ns | 46.37ns | 57.75ns | 49.07ns | 62.31ns | 45.75ns | 16.29ns | 188.64ns |
| 4 | 18.75ns | 16.72ns | 16.78ns | 14.82ns | 28.99ns | 15.94ns | 6.64ns | 49.90ns |
| 16 | 9.00ns | 8.35ns | 6.70ns | 6.24ns | 21.46ns | 10.33ns | 4.24ns | 16.18ns |
| 64 | 7.04ns | 6.87ns | 4.53ns | 4.44ns | 19.09ns | 8.98ns | 3.70ns | 7.71ns |
| 256 | 6.44ns | 6.39ns | 3.99ns | 3.91ns | 19.02ns | 8.56ns | 3.50ns | 5.52ns |
| 1k | 6.35ns | 6.33ns | 3.84ns | 3.79ns | 19.21ns | 8.45ns | 3.49ns | 4.94ns |
| 16k | 6.28ns | 6.28ns | 3.81ns | 3.79ns | 22.38ns | 8.63ns | 3.48ns | 4.80ns |
| 256k | 6.31ns | 6.29ns | 3.80ns | 3.80ns | 22.87ns | 8.46ns | 3.50ns | 4.78ns |
| 1M | 6.31ns | 6.30ns | 3.81ns | 3.80ns | 27.15ns | 8.59ns | 3.54ns | 4.78ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/4fe8e837-8a7a-49d0-9bde-00e7b36e613f)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.66ns | 47.80ns | 57.72ns | 49.06ns | 62.10ns | 44.84ns | 16.53ns | 203.85ns |
| 4 | 24.58ns | 18.77ns | 28.90ns | 21.43ns | 31.56ns | 20.39ns | 14.73ns | 102.16ns |
| 16 | 19.03ns | 12.94ns | 22.08ns | 14.37ns | 25.74ns | 14.72ns | 24.63ns | 74.44ns |
| 64 | 12.06ns | 9.10ns | 13.52ns | 9.14ns | 22.33ns | 10.85ns | 15.26ns | 37.09ns |
| 256 | 7.69ns | 6.95ns | 5.77ns | 4.66ns | 20.71ns | 8.91ns | 6.39ns | 13.09ns |
| 1k | 6.63ns | 6.45ns | 4.26ns | 3.74ns | 22.20ns | 8.47ns | 4.15ns | 7.24ns |
| 16k | 6.33ns | 6.31ns | 3.88ns | 3.58ns | 30.53ns | 8.16ns | 3.61ns | 5.05ns |
| 256k | 6.31ns | 6.31ns | 3.82ns | 3.49ns | 78.08ns | 8.16ns | 3.51ns | 4.89ns |
| 1M | 6.34ns | 6.33ns | 3.84ns | 3.51ns | 103.31ns | 8.17ns | 3.53ns | 4.79ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/175c305f-617f-4050-8904-19574e6befbc)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 64.81ns | 46.57ns | 57.40ns | 50.03ns | 62.18ns | 56.51ns | 16.51ns | 226.73ns |
| 4 | 29.54ns | 16.99ns | 16.62ns | 14.78ns | 29.06ns | 27.14ns | 9.32ns | 87.12ns |
| 16 | 21.29ns | 8.39ns | 6.67ns | 6.14ns | 20.44ns | 21.59ns | 4.80ns | 57.30ns |
| 64 | 19.21ns | 6.91ns | 4.56ns | 4.70ns | 18.79ns | 19.79ns | 3.85ns | 56.35ns |
| 256 | 9.51ns | 6.38ns | 4.03ns | 3.74ns | 18.83ns | 11.14ns | 3.53ns | 17.94ns |
| 1k | 7.10ns | 6.28ns | 3.82ns | 3.52ns | 19.30ns | 8.86ns | 3.47ns | 8.03ns |
| 16k | 6.33ns | 6.29ns | 3.79ns | 3.49ns | 20.52ns | 8.22ns | 3.49ns | 5.03ns |
| 256k | 6.30ns | 6.29ns | 3.80ns | 3.50ns | 23.78ns | 8.15ns | 3.48ns | 4.77ns |
| 1M | 6.30ns | 6.29ns | 3.81ns | 3.51ns | 26.87ns | 8.16ns | 3.52ns | 4.79ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/621495ba-76a4-4392-a3e1-dcddace8b87d)

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 4.51ns | 4.89ns | 9.40ns | 8.47ns | 37.53ns | 39.32ns |
| 4 | 4.24ns | 4.56ns | 8.85ns | 8.26ns | 36.80ns | 39.04ns |
| 16 | 3.93ns | 4.42ns | 8.78ns | 13.65ns | 37.12ns | 39.21ns |
| 64 | 4.01ns | 4.49ns | 8.94ns | 14.24ns | 38.64ns | 38.98ns |
| 256 | 3.96ns | 4.42ns | 9.33ns | 14.08ns | 40.24ns | 40.34ns |
| 1k | 4.36ns | 4.56ns | 13.29ns | 17.61ns | 41.25ns | 41.84ns |
| 16k | 10.00ns | 8.48ns | 36.84ns | 28.68ns | 57.22ns | 63.00ns |
| 256k | 11.90ns | 13.48ns | 46.27ns | 39.30ns | 65.11ns | 86.62ns |
| 1M | 41.35ns | 24.90ns | 195.47ns | 110.79ns | 175.76ns | 229.24ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/269813e7-7b95-40a4-a279-c452574aa71c)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 199.20ns | 527.22ns | 276.89ns | 294.24ns | 1.11us | 359.19ns | 451.50ns | 808.17ns |
| 4 | 88.85ns | 141.43ns | 113.55ns | 88.72ns | 507.31ns | 167.87ns | 207.65ns | 357.05ns |
| 16 | 51.33ns | 43.92ns | 69.12ns | 27.31ns | 317.29ns | 139.69ns | 130.56ns | 243.00ns |
| 64 | 41.41ns | 16.86ns | 56.74ns | 13.72ns | 294.80ns | 140.44ns | 115.45ns | 215.21ns |
| 256 | 38.64ns | 11.41ns | 58.32ns | 9.98ns | 204.69ns | 114.25ns | 98.71ns | 201.76ns |
| 1k | 31.76ns | 9.96ns | 43.59ns | 8.31ns | 192.25ns | 115.62ns | 268.26ns | 209.26ns |
| 16k | 29.43ns | 8.54ns | 40.23ns | 7.63ns | 208.53ns | 119.03ns | 283.65ns | 213.19ns |
| 256k | 28.23ns | 8.31ns | 39.99ns | 7.62ns | 208.19ns | 128.80ns | 332.51ns | 253.51ns |
| 1M | 28.43ns | 8.25ns | 39.95ns | 7.57ns | 213.37ns | 176.85ns | 421.81ns | 332.52ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/6c7fc0d5-af93-4566-8011-1759c4b5c272)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.27us | 7.61us | 7.80us | 7.82us | 4.47us | 19.45us | 3.22us | 1.78us |
| 4 | 1.90us | 1.94us | 1.99us | 1.97us | 1.34us | 4.81us | 1.02us | 783.39ns |
| 16 | 497.73ns | 485.25ns | 554.32ns | 484.63ns | 623.78ns | 1.36us | 378.24ns | 406.56ns |
| 64 | 158.99ns | 131.77ns | 179.35ns | 135.97ns | 423.21ns | 513.64ns | 238.82ns | 278.55ns |
| 256 | 71.35ns | 42.59ns | 88.18ns | 43.46ns | 311.82ns | 256.35ns | 176.01ns | 245.00ns |
| 1k | 48.33ns | 26.02ns | 56.27ns | 20.21ns | 288.11ns | 242.47ns | 157.52ns | 215.80ns |
| 16k | 68.02ns | 22.69ns | 63.69ns | 38.15ns | 399.79ns | 265.33ns | 193.03ns | 389.24ns |
| 256k | 105.20ns | 26.52ns | 97.50ns | 45.99ns | 479.86ns | 614.25ns | 250.17ns | 489.41ns |
| 1M | 81.84ns | 15.24ns | 77.14ns | 28.52ns | 422.79ns | 1.38us | 299.28ns | 544.32ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/7a828d84-9225-4851-b92d-c5327f6f878c)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 396.36ns | 706.04ns | 385.76ns | 409.09ns | 2.03us | 485.61ns | 685.00ns | 3.06us |
| 4 | 164.97ns | 187.64ns | 178.66ns | 106.31ns | 1.20us | 230.98ns | 386.35ns | 2.14us |
| 16 | 116.40ns | 56.84ns | 120.40ns | 32.48ns | 1.03us | 207.37ns | 282.01ns | 1.76us |
| 64 | 101.27ns | 20.67ns | 106.45ns | 14.94ns | 832.09ns | 230.50ns | 258.63ns | 1.41us |
| 256 | 84.60ns | 12.19ns | 92.48ns | 11.07ns | 730.22ns | 155.15ns | 213.22ns | 1.35us |
| 1k | 80.35ns | 9.51ns | 86.35ns | 9.36ns | 712.05ns | 162.84ns | 493.96ns | 1.39us |
| 16k | 76.46ns | 8.41ns | 82.16ns | 7.78ns | 793.39ns | 168.98ns | 482.40ns | 1.39us |
| 256k | 76.09ns | 8.28ns | 81.63ns | 7.60ns | 709.37ns | 181.27ns | 560.30ns | 1.47us |
| 1M | 81.54ns | 8.23ns | 81.55ns | 7.56ns | 696.66ns | 243.57ns | 667.38ns | 1.55us |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/ef5d0979-3f37-480d-87d0-870bc7514e89)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 210.07ns | 512.97ns | 223.29ns | 470.47ns | 488.05ns | 591.20ns | 380.72ns | 922.58ns |
| 4 | 146.12ns | 140.67ns | 145.31ns | 125.95ns | 440.47ns | 535.47ns | 351.58ns | 571.57ns |
| 16 | 134.17ns | 46.49ns | 128.60ns | 44.14ns | 433.12ns | 559.04ns | 348.43ns | 488.32ns |
| 64 | 126.61ns | 23.42ns | 122.16ns | 22.68ns | 430.80ns | 558.64ns | 355.65ns | 461.37ns |
| 256 | 123.37ns | 10.23ns | 119.77ns | 11.81ns | 427.18ns | 567.63ns | 366.92ns | 462.55ns |
| 1k | 122.84ns | 7.27ns | 118.90ns | 8.43ns | 429.22ns | 571.26ns | 742.15ns | 465.35ns |
| 16k | 123.76ns | 7.34ns | 119.56ns | 8.08ns | 477.73ns | 606.53ns | 802.97ns | 515.25ns |
| 256k | 123.11ns | 8.28ns | 119.48ns | 8.88ns | 468.73ns | 654.35ns | 894.22ns | 608.38ns |
| 1M | 123.63ns | 10.01ns | 118.97ns | 9.89ns | 468.65ns | 845.19ns | 1.22us | 887.77ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/b40a706a-9556-4f78-99f1-4ae2ed378ee1)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 558.65ns | 848.68ns | 454.82ns | 774.32ns | 1.13us | 1.01us | 807.21ns | 2.47us |
| 4 | 501.22ns | 268.68ns | 436.94ns | 254.10ns | 1.08us | 981.49ns | 802.96ns | 1.96us |
| 16 | 475.55ns | 137.83ns | 408.38ns | 131.54ns | 1.05us | 972.59ns | 800.44ns | 1.80us |
| 64 | 465.88ns | 98.62ns | 410.78ns | 99.22ns | 1.03us | 958.27ns | 823.42ns | 1.77us |
| 256 | 457.30ns | 36.26ns | 398.71ns | 35.57ns | 1.05us | 974.56ns | 872.29ns | 1.77us |
| 1k | 468.16ns | 18.96ns | 416.86ns | 20.45ns | 1.11us | 985.98ns | 1.42us | 1.81us |
| 16k | 463.07ns | 21.19ns | 426.11ns | 22.02ns | 1.23us | 1.05us | 1.45us | 1.86us |
| 256k | 626.09ns | 38.35ns | 546.45ns | 38.88ns | 1.32us | 1.32us | 1.62us | 2.42us |
| 1M | 584.17ns | 41.60ns | 500.67ns | 42.52ns | 1.20us | 1.57us | 2.09us | 2.75us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/867d4d90-4709-4fb6-802d-3c99ba79b8db)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 211.30ns | 647.30ns | 171.79ns | 775.78ns | 273.46ns | 372.99ns | 207.87ns | 456.59ns |
| 4 | 89.13ns | 185.80ns | 75.22ns | 213.60ns | 116.24ns | 167.19ns | 83.90ns | 228.59ns |
| 16 | 52.65ns | 59.47ns | 41.17ns | 63.45ns | 74.18ns | 148.16ns | 54.53ns | 174.42ns |
| 64 | 44.59ns | 27.63ns | 36.85ns | 31.63ns | 61.74ns | 117.50ns | 42.90ns | 157.51ns |
| 256 | 42.85ns | 11.51ns | 28.79ns | 11.39ns | 57.05ns | 105.77ns | 44.34ns | 128.41ns |
| 1k | 35.58ns | 7.37ns | 25.46ns | 7.30ns | 51.55ns | 107.18ns | 36.60ns | 133.69ns |
| 16k | 33.40ns | 6.58ns | 24.70ns | 6.33ns | 51.90ns | 116.42ns | 50.80ns | 152.65ns |
| 256k | 39.08ns | 6.56ns | 24.57ns | 6.31ns | 58.86ns | 153.50ns | 88.24ns | 232.14ns |
| 1M | 33.51ns | 7.74ns | 24.76ns | 7.54ns | 68.30ns | 270.14ns | 188.60ns | 361.66ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/0609c251-8b95-478d-85ca-68c76da5f047)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 310.37ns | 677.98ns | 243.51ns | 988.47ns | 424.74ns | 507.67ns | 198.40ns | 923.17ns |
| 4 | 175.09ns | 177.38ns | 154.13ns | 279.24ns | 227.21ns | 260.67ns | 83.14ns | 569.90ns |
| 16 | 141.89ns | 56.82ns | 119.01ns | 116.39ns | 154.43ns | 220.43ns | 54.25ns | 450.02ns |
| 64 | 123.62ns | 27.35ns | 117.85ns | 72.44ns | 140.01ns | 220.48ns | 43.80ns | 362.75ns |
| 256 | 120.57ns | 11.68ns | 96.11ns | 23.01ns | 120.65ns | 165.81ns | 47.48ns | 329.96ns |
| 1k | 103.77ns | 8.08ns | 81.26ns | 8.53ns | 141.47ns | 158.65ns | 35.00ns | 310.73ns |
| 16k | 109.99ns | 6.67ns | 77.73ns | 9.00ns | 139.69ns | 178.95ns | 49.93ns | 334.63ns |
| 256k | 119.04ns | 6.65ns | 91.56ns | 13.28ns | 174.78ns | 280.55ns | 89.97ns | 511.04ns |
| 1M | 115.31ns | 7.83ns | 88.64ns | 15.95ns | 128.64ns | 392.16ns | 205.50ns | 632.98ns |


### Create world

- Create a new world

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.32us | 14.87us | 2.28us | 187.88us | 2.15us | 39.57us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&mlange-42/ark&yottahmd/donburi&marioolofo/go-gameengine-ecs&unitoftime/ecs&akmonengine/volt&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,mlange-42/ark,yottahmd/donburi,marioolofo/go-gameengine-ecs,unitoftime/ecs,akmonengine/volt&type=Date" alt="Star History Chart" width="600"/>
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
