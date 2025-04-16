# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche) and [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Version |
|-----|---------|
| [Arche](https://github.com/mlange-42/arche) | v0.15.3 |
| [Ark](https://github.com/mlange-42/ark) | v0.4.0 |
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

Last run: Wed, 19 Mar 2025 10:54:57 UTC  
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

![query2comp](https://github.com/user-attachments/assets/7bda3560-726b-4879-b7f5-58b175dabc49)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 54.10ns | 46.15ns | 51.17ns | 48.92ns | 68.07ns | 45.02ns | 16.18ns | 200.09ns |
| 4 | 18.42ns | 16.48ns | 15.99ns | 14.55ns | 30.00ns | 16.35ns | 6.51ns | 51.99ns |
| 16 | 8.16ns | 8.13ns | 6.75ns | 6.10ns | 20.73ns | 11.12ns | 4.12ns | 16.62ns |
| 64 | 6.31ns | 6.12ns | 4.58ns | 4.30ns | 18.05ns | 9.55ns | 3.56ns | 7.50ns |
| 256 | 5.58ns | 5.48ns | 4.02ns | 3.75ns | 18.61ns | 8.98ns | 3.37ns | 5.25ns |
| 1k | 5.39ns | 5.37ns | 3.66ns | 3.59ns | 18.74ns | 8.93ns | 3.38ns | 4.78ns |
| 16k | 5.46ns | 5.35ns | 3.62ns | 3.63ns | 19.66ns | 8.97ns | 3.35ns | 4.66ns |
| 256k | 5.48ns | 5.49ns | 3.72ns | 3.78ns | 21.53ns | 9.20ns | 3.42ns | 4.67ns |
| 1M | 5.57ns | 5.51ns | 3.69ns | 3.70ns | 25.74ns | 9.32ns | 3.58ns | 4.83ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/db85bd0b-dc94-462e-973c-03e30df56b37)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.70ns | 47.45ns | 51.87ns | 49.80ns | 62.21ns | 44.19ns | 16.62ns | 210.85ns |
| 4 | 24.18ns | 18.21ns | 25.00ns | 21.40ns | 32.96ns | 19.89ns | 14.77ns | 98.60ns |
| 16 | 18.16ns | 11.54ns | 19.99ns | 15.30ns | 26.05ns | 14.25ns | 24.49ns | 70.41ns |
| 64 | 11.09ns | 8.28ns | 11.80ns | 8.84ns | 21.42ns | 10.64ns | 15.21ns | 34.88ns |
| 256 | 6.89ns | 6.16ns | 5.18ns | 4.62ns | 21.08ns | 9.34ns | 6.04ns | 12.13ns |
| 1k | 5.92ns | 5.62ns | 3.79ns | 3.94ns | 20.54ns | 9.04ns | 4.04ns | 6.42ns |
| 16k | 5.46ns | 5.58ns | 3.49ns | 3.74ns | 29.37ns | 8.77ns | 3.42ns | 4.79ns |
| 256k | 5.48ns | 5.46ns | 3.30ns | 3.59ns | 56.68ns | 8.69ns | 3.32ns | 4.53ns |
| 1M | 5.46ns | 5.48ns | 3.36ns | 3.69ns | 87.08ns | 8.84ns | 3.32ns | 4.65ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/b654d6dc-a2f5-445c-9536-eaf771a4ce91)

| N | Arche | Arche (cached) | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 62.41ns | 44.43ns | 57.43ns | 48.29ns | 62.08ns | 54.62ns | 16.62ns | 225.40ns |
| 4 | 28.74ns | 16.75ns | 26.98ns | 14.72ns | 29.39ns | 26.69ns | 9.27ns | 83.35ns |
| 16 | 20.21ns | 7.53ns | 18.47ns | 6.20ns | 20.69ns | 22.32ns | 4.76ns | 55.63ns |
| 64 | 18.81ns | 6.20ns | 16.66ns | 4.41ns | 19.10ns | 21.36ns | 3.77ns | 53.98ns |
| 256 | 8.82ns | 5.54ns | 6.66ns | 3.76ns | 18.80ns | 11.83ns | 3.39ns | 16.73ns |
| 1k | 6.28ns | 5.51ns | 4.10ns | 3.62ns | 19.04ns | 9.74ns | 3.30ns | 7.59ns |
| 16k | 5.51ns | 5.50ns | 3.42ns | 3.63ns | 19.70ns | 9.09ns | 3.32ns | 4.85ns |
| 256k | 5.49ns | 5.43ns | 3.32ns | 3.74ns | 22.18ns | 9.40ns | 3.48ns | 4.76ns |
| 1M | 5.64ns | 5.60ns | 3.38ns | 3.68ns | 25.46ns | 9.32ns | 3.35ns | 4.56ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/4aa9f2d5-7b59-4d62-9a4d-4e7a9554fe52)

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 4.71ns | 4.98ns | 9.18ns | 8.42ns | 37.10ns | 36.31ns |
| 4 | 4.13ns | 4.50ns | 8.75ns | 8.22ns | 39.59ns | 36.31ns |
| 16 | 3.98ns | 4.41ns | 8.59ns | 13.87ns | 36.97ns | 36.12ns |
| 64 | 4.00ns | 4.50ns | 8.79ns | 13.76ns | 38.84ns | 36.62ns |
| 256 | 3.93ns | 4.38ns | 9.11ns | 13.78ns | 39.95ns | 37.28ns |
| 1k | 4.37ns | 4.51ns | 11.76ns | 17.82ns | 41.00ns | 39.27ns |
| 16k | 10.01ns | 8.75ns | 35.40ns | 29.72ns | 58.55ns | 62.52ns |
| 256k | 12.39ns | 13.52ns | 67.48ns | 40.87ns | 69.21ns | 118.96ns |
| 1M | 56.04ns | 45.53ns | 235.55ns | 139.41ns | 203.43ns | 276.27ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/82df426e-b8a1-43d0-b13b-76e34cc40afa)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 216.20ns | 260.15ns | 264.94ns | 269.54ns | 1.07us | 342.14ns | 406.63ns | 791.57ns |
| 4 | 81.44ns | 71.18ns | 110.15ns | 72.27ns | 504.92ns | 161.64ns | 199.98ns | 389.88ns |
| 16 | 52.91ns | 24.83ns | 73.85ns | 24.67ns | 338.60ns | 132.14ns | 130.11ns | 262.84ns |
| 64 | 40.89ns | 13.81ns | 60.53ns | 13.45ns | 267.65ns | 100.48ns | 117.63ns | 210.81ns |
| 256 | 38.85ns | 10.42ns | 47.24ns | 10.02ns | 213.50ns | 97.07ns | 87.56ns | 171.30ns |
| 1k | 27.36ns | 9.69ns | 42.09ns | 8.94ns | 201.95ns | 102.68ns | 260.50ns | 184.05ns |
| 16k | 26.10ns | 8.26ns | 39.97ns | 7.62ns | 222.39ns | 107.20ns | 278.17ns | 205.26ns |
| 256k | 26.53ns | 8.43ns | 41.60ns | 7.60ns | 220.77ns | 168.68ns | 402.91ns | 312.14ns |
| 1M | 26.35ns | 8.26ns | 41.72ns | 7.63ns | 191.88ns | 220.78ns | 482.98ns | 363.70ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/41a21617-bea2-43bd-9c7a-30881e8da3c1)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 8.86us | 8.86us | 8.79us | 8.77us | 4.85us | 30.49us | 3.29us | 1.87us |
| 4 | 2.09us | 2.16us | 2.23us | 2.16us | 1.59us | 7.54us | 1.14us | 818.80ns |
| 16 | 562.75ns | 526.10ns | 608.05ns | 539.32ns | 680.93ns | 1.98us | 411.39ns | 424.57ns |
| 64 | 171.90ns | 143.38ns | 182.60ns | 143.45ns | 418.16ns | 691.07ns | 232.02ns | 282.89ns |
| 256 | 70.31ns | 44.35ns | 87.52ns | 45.05ns | 317.39ns | 275.29ns | 177.23ns | 256.83ns |
| 1k | 48.33ns | 26.26ns | 57.62ns | 20.39ns | 291.96ns | 206.03ns | 159.30ns | 222.11ns |
| 16k | 70.87ns | 32.62ns | 71.47ns | 43.10ns | 405.42ns | 253.92ns | 188.85ns | 372.58ns |
| 256k | 135.38ns | 29.43ns | 101.70ns | 41.50ns | 464.10ns | 651.66ns | 262.73ns | 496.73ns |
| 1M | 98.50ns | 19.47ns | 87.24ns | 31.76ns | 469.84ns | 1.44us | 366.44ns | 597.90ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/cafd0266-c8a9-440c-9842-f08103321190)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 346.94ns | 396.28ns | 355.85ns | 359.12ns | 2.17us | 501.05ns | 690.55ns | 2.85us |
| 4 | 156.16ns | 105.53ns | 171.65ns | 98.46ns | 1.31us | 263.05ns | 375.16ns | 2.10us |
| 16 | 107.85ns | 33.91ns | 119.41ns | 31.09ns | 1.11us | 205.00ns | 275.36ns | 1.57us |
| 64 | 92.47ns | 16.26ns | 104.44ns | 17.07ns | 810.58ns | 170.74ns | 232.07ns | 1.45us |
| 256 | 80.25ns | 11.40ns | 95.49ns | 11.33ns | 736.18ns | 195.62ns | 210.65ns | 1.35us |
| 1k | 73.14ns | 9.57ns | 85.67ns | 9.68ns | 742.76ns | 148.60ns | 457.33ns | 1.37us |
| 16k | 84.15ns | 8.35ns | 84.99ns | 8.08ns | 778.80ns | 160.47ns | 474.75ns | 1.41us |
| 256k | 73.60ns | 8.29ns | 84.57ns | 8.05ns | 704.23ns | 205.24ns | 673.50ns | 1.51us |
| 1M | 73.52ns | 8.26ns | 84.57ns | 8.07ns | 698.56ns | 255.74ns | 665.65ns | 1.54us |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/e91548da-ce89-43af-bac5-1d383217a321)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 215.07ns | 468.92ns | 222.77ns | 318.01ns | 505.47ns | 583.59ns | 350.18ns | 816.46ns |
| 4 | 144.95ns | 128.26ns | 150.61ns | 90.87ns | 446.75ns | 525.60ns | 333.74ns | 530.67ns |
| 16 | 120.31ns | 45.03ns | 122.24ns | 32.27ns | 432.53ns | 539.26ns | 355.14ns | 458.46ns |
| 64 | 113.89ns | 23.03ns | 123.99ns | 19.25ns | 424.44ns | 530.93ns | 368.22ns | 447.49ns |
| 256 | 113.48ns | 10.18ns | 113.77ns | 9.54ns | 425.19ns | 531.31ns | 351.86ns | 445.84ns |
| 1k | 113.48ns | 7.32ns | 113.78ns | 8.59ns | 427.75ns | 550.21ns | 720.43ns | 452.57ns |
| 16k | 114.80ns | 7.32ns | 114.63ns | 7.92ns | 476.15ns | 585.49ns | 767.90ns | 500.60ns |
| 256k | 114.53ns | 8.59ns | 116.64ns | 9.07ns | 471.32ns | 681.58ns | 1.09us | 694.14ns |
| 1M | 114.47ns | 9.73ns | 116.18ns | 10.11ns | 461.22ns | 904.77ns | 1.28us | 932.25ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/f9a0d8c1-a499-4a4a-988e-bdb8ce04078b)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 542.25ns | 825.16ns | 422.67ns | 622.57ns | 1.13us | 1.03us | 833.57ns | 2.51us |
| 4 | 498.18ns | 264.21ns | 396.22ns | 213.93ns | 1.10us | 989.31ns | 821.54ns | 2.06us |
| 16 | 480.70ns | 133.75ns | 377.17ns | 113.61ns | 1.06us | 972.03ns | 789.52ns | 1.94us |
| 64 | 456.55ns | 98.34ns | 377.75ns | 87.43ns | 1.04us | 951.50ns | 812.28ns | 1.88us |
| 256 | 445.29ns | 35.40ns | 391.05ns | 35.76ns | 1.06us | 952.18ns | 854.79ns | 1.87us |
| 1k | 461.44ns | 18.77ns | 395.80ns | 18.79ns | 1.15us | 990.54ns | 1.40us | 1.92us |
| 16k | 469.10ns | 24.96ns | 414.67ns | 25.56ns | 1.20us | 1.03us | 1.44us | 1.99us |
| 256k | 667.83ns | 40.69ns | 549.37ns | 41.79ns | 1.32us | 1.38us | 1.87us | 2.84us |
| 1M | 584.55ns | 44.26ns | 501.36ns | 46.61ns | 1.19us | 1.61us | 2.01us | 2.96us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/3d8f39dd-e9e3-481e-9c28-4e080044eec0)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 217.74ns | 563.05ns | 154.20ns | 720.62ns | 264.30ns | 337.77ns | 164.78ns | 465.15ns |
| 4 | 92.82ns | 164.82ns | 70.65ns | 198.66ns | 112.21ns | 160.07ns | 80.62ns | 240.13ns |
| 16 | 57.66ns | 55.42ns | 38.31ns | 60.98ns | 71.39ns | 137.14ns | 50.48ns | 187.40ns |
| 64 | 46.10ns | 25.84ns | 33.90ns | 25.07ns | 58.06ns | 105.13ns | 43.38ns | 138.07ns |
| 256 | 38.94ns | 12.02ns | 33.18ns | 10.25ns | 55.73ns | 107.71ns | 41.93ns | 127.58ns |
| 1k | 33.85ns | 7.35ns | 28.45ns | 5.99ns | 50.22ns | 102.77ns | 35.33ns | 130.93ns |
| 16k | 33.46ns | 6.37ns | 23.15ns | 5.21ns | 50.62ns | 117.12ns | 50.38ns | 161.51ns |
| 256k | 32.99ns | 6.86ns | 24.01ns | 5.67ns | 59.81ns | 235.91ns | 168.02ns | 331.00ns |
| 1M | 32.89ns | 7.87ns | 24.22ns | 6.54ns | 57.92ns | 311.43ns | 232.22ns | 397.64ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/10735abc-54a3-48e2-afd5-7ec9b8706310)

| N | Arche | Arche (batch) | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | 290.75ns | 590.76ns | 245.11ns | 1.07us | 462.89ns | 502.05ns | 180.68ns | 827.94ns |
| 4 | 179.85ns | 161.24ns | 135.92ns | 291.08ns | 202.22ns | 269.93ns | 78.16ns | 470.44ns |
| 16 | 138.77ns | 53.98ns | 108.53ns | 109.51ns | 152.45ns | 237.17ns | 53.20ns | 416.24ns |
| 64 | 127.40ns | 25.55ns | 85.82ns | 58.37ns | 138.09ns | 173.58ns | 42.46ns | 348.55ns |
| 256 | 112.63ns | 11.41ns | 82.06ns | 21.73ns | 124.28ns | 157.55ns | 39.25ns | 304.74ns |
| 1k | 109.07ns | 7.63ns | 72.72ns | 8.24ns | 138.02ns | 155.75ns | 34.46ns | 309.06ns |
| 16k | 108.83ns | 6.63ns | 71.35ns | 7.73ns | 138.29ns | 179.87ns | 50.29ns | 348.06ns |
| 256k | 116.76ns | 6.73ns | 86.11ns | 12.39ns | 158.91ns | 286.85ns | 106.66ns | 563.01ns |
| 1M | 115.25ns | 7.86ns | 83.65ns | 14.73ns | 127.40ns | 398.18ns | 220.72ns | 674.15ns |


### Create world

- Create a new world

| N | Arche | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.68us | 23.40us | 2.58us | 190.26us | 2.55us | 45.00us |


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

Developers of included frameworks are encouraged to review the benchmarks,
and to fix (or point to) misuse or potential optimizations.
