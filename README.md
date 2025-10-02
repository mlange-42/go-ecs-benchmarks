# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of [Ark](https://github.com/mlange-42/ark).

## Benchmark candidates

| ECS | Version |
|-----|---------|
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

Last run: Wed, 10 Sep 2025 17:18:09 UTC  
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

![query2comp](https://github.com/user-attachments/assets/8ce53a8f-5b4d-46fb-94f9-4b7cf80199c2)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 57.70ns | 49.00ns | 62.53ns | 46.01ns | 16.22ns | 189.06ns |
| 4 | 17.45ns | 14.70ns | 28.96ns | 15.95ns | 6.73ns | 49.50ns |
| 16 | 6.74ns | 6.36ns | 20.44ns | 10.30ns | 4.23ns | 16.31ns |
| 64 | 4.87ns | 4.44ns | 18.52ns | 8.97ns | 3.71ns | 7.67ns |
| 256 | 3.93ns | 4.01ns | 19.91ns | 8.55ns | 3.50ns | 5.49ns |
| 1k | 3.80ns | 3.76ns | 19.43ns | 8.44ns | 3.46ns | 5.08ns |
| 16k | 3.81ns | 3.79ns | 20.28ns | 8.45ns | 3.49ns | 4.78ns |
| 256k | 3.76ns | 3.79ns | 21.95ns | 8.45ns | 3.49ns | 4.75ns |
| 1M | 3.78ns | 3.75ns | 25.18ns | 8.46ns | 3.59ns | 4.76ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/11a2ddd6-1d55-4b9b-84a7-fae6b5723cad)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 57.92ns | 49.14ns | 62.10ns | 44.91ns | 16.54ns | 199.77ns |
| 4 | 29.21ns | 21.46ns | 33.37ns | 20.28ns | 14.75ns | 100.08ns |
| 16 | 21.99ns | 14.40ns | 25.84ns | 14.70ns | 24.34ns | 73.64ns |
| 64 | 13.09ns | 9.27ns | 22.48ns | 10.82ns | 14.57ns | 36.69ns |
| 256 | 5.99ns | 4.54ns | 20.72ns | 9.08ns | 6.42ns | 13.06ns |
| 1k | 4.26ns | 3.76ns | 22.83ns | 8.48ns | 4.19ns | 7.30ns |
| 16k | 3.89ns | 3.55ns | 28.84ns | 8.21ns | 3.59ns | 5.02ns |
| 256k | 3.78ns | 3.50ns | 63.20ns | 8.14ns | 3.50ns | 4.78ns |
| 1M | 3.78ns | 3.49ns | 92.74ns | 8.15ns | 3.51ns | 4.78ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/802df1c4-f414-47be-b12c-7f196899da6a)

| N | Ark | Ark (cached) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 58.03ns | 50.19ns | 62.44ns | 56.44ns | 16.53ns | 223.55ns |
| 4 | 16.50ns | 14.79ns | 28.94ns | 27.25ns | 9.31ns | 85.97ns |
| 16 | 6.65ns | 6.38ns | 20.46ns | 21.61ns | 4.80ns | 56.79ns |
| 64 | 4.52ns | 4.13ns | 18.70ns | 19.70ns | 3.86ns | 55.80ns |
| 256 | 4.21ns | 3.60ns | 18.84ns | 11.13ns | 3.53ns | 17.66ns |
| 1k | 3.76ns | 3.53ns | 19.10ns | 8.88ns | 3.47ns | 8.07ns |
| 16k | 3.78ns | 3.48ns | 19.99ns | 8.20ns | 3.48ns | 5.05ns |
| 256k | 3.79ns | 3.46ns | 21.90ns | 8.14ns | 3.49ns | 4.80ns |
| 1M | 3.80ns | 3.47ns | 25.68ns | 8.15ns | 3.52ns | 4.77ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

![random](https://github.com/user-attachments/assets/893dc8db-bf68-4880-bcae-a0129845365a)

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 4.83ns | 9.37ns | 8.75ns | 37.01ns | 39.62ns |
| 4 | 4.52ns | 8.86ns | 8.25ns | 36.75ns | 39.16ns |
| 16 | 4.42ns | 8.85ns | 13.71ns | 36.99ns | 39.15ns |
| 64 | 4.51ns | 8.89ns | 13.77ns | 38.40ns | 39.16ns |
| 256 | 4.42ns | 8.96ns | 13.94ns | 40.26ns | 40.24ns |
| 1k | 4.57ns | 13.32ns | 17.88ns | 41.01ns | 41.71ns |
| 16k | 8.60ns | 33.04ns | 28.43ns | 54.30ns | 63.19ns |
| 256k | 13.16ns | 65.98ns | 48.55ns | 77.64ns | 110.34ns |
| 1M | 59.56ns | 239.63ns | 122.29ns | 189.08ns | 253.41ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude memory allocation, archetype creation etc.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/97bff8fe-4c5c-44ba-9a80-3361f005468a)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 274.54ns | 262.26ns | 1.05us | 309.75ns | 380.99ns | 812.03ns |
| 4 | 115.31ns | 73.61ns | 506.44ns | 148.37ns | 187.37ns | 366.86ns |
| 16 | 70.04ns | 26.34ns | 330.43ns | 137.80ns | 127.74ns | 277.45ns |
| 64 | 60.07ns | 13.55ns | 274.17ns | 132.11ns | 110.35ns | 204.01ns |
| 256 | 62.01ns | 9.97ns | 209.60ns | 106.98ns | 92.10ns | 185.45ns |
| 1k | 45.56ns | 10.96ns | 196.40ns | 115.61ns | 276.15ns | 198.69ns |
| 16k | 40.32ns | 7.60ns | 212.12ns | 118.13ns | 284.51ns | 211.93ns |
| 256k | 41.26ns | 7.55ns | 198.46ns | 135.70ns | 367.02ns | 276.53ns |
| 1M | 45.34ns | 7.64ns | 194.32ns | 203.22ns | 461.00ns | 343.38ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
This reflects the creation of the first entities with a certain components set in your game or application.
As soon as things stabilize, the benchmarks for entity creation without allocation (above) apply.

Low `N` values might be biased by things like archetype creation and memory allocation,
which is handled differently by different implementations.

![create2comp_alloc](https://github.com/user-attachments/assets/c079ddd0-0c60-4123-b373-4219423efa0c)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 6.95us | 6.97us | 4.49us | 21.01us | 3.31us | 1.81us |
| 4 | 1.81us | 1.77us | 1.29us | 4.84us | 1.03us | 817.76ns |
| 16 | 495.20ns | 444.91ns | 599.90ns | 1.57us | 377.59ns | 420.15ns |
| 64 | 167.95ns | 125.74ns | 389.16ns | 548.28ns | 234.12ns | 303.12ns |
| 256 | 82.14ns | 42.73ns | 294.66ns | 299.64ns | 152.37ns | 249.12ns |
| 1k | 53.61ns | 19.69ns | 288.01ns | 224.32ns | 142.26ns | 214.65ns |
| 16k | 59.62ns | 28.67ns | 398.26ns | 240.83ns | 181.65ns | 377.09ns |
| 256k | 86.74ns | 40.57ns | 427.28ns | 613.18ns | 245.69ns | 486.08ns |
| 1M | 80.41ns | 29.28ns | 420.35ns | 1.32us | 309.06ns | 567.85ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/e2d1038c-3707-4b1d-b573-748183693ae2)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 392.61ns | 388.60ns | 2.01us | 480.66ns | 605.59ns | 2.86us |
| 4 | 165.65ns | 107.28ns | 1.17us | 247.34ns | 359.69ns | 2.16us |
| 16 | 115.58ns | 31.62ns | 961.39ns | 220.81ns | 268.24ns | 1.62us |
| 64 | 101.21ns | 15.57ns | 790.26ns | 208.85ns | 287.31ns | 1.36us |
| 256 | 90.44ns | 11.21ns | 712.19ns | 162.42ns | 224.88ns | 1.38us |
| 1k | 85.47ns | 10.13ns | 712.96ns | 168.16ns | 477.29ns | 1.40us |
| 16k | 81.82ns | 7.58ns | 740.40ns | 167.97ns | 482.97ns | 1.40us |
| 256k | 81.74ns | 7.57ns | 677.12ns | 196.91ns | 594.08ns | 1.51us |
| 1M | 83.41ns | 7.56ns | 679.29ns | 255.23ns | 665.02ns | 1.55us |


### Add/remove component

`N` entities with component `Position`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/4ce16713-bf7e-4662-9dfb-81825adff9a5)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 223.76ns | 463.33ns | 482.75ns | 590.65ns | 379.38ns | 931.51ns |
| 4 | 147.55ns | 123.08ns | 433.99ns | 535.04ns | 346.35ns | 565.79ns |
| 16 | 129.43ns | 41.21ns | 432.27ns | 556.94ns | 344.73ns | 481.91ns |
| 64 | 121.35ns | 22.78ns | 419.83ns | 552.98ns | 352.82ns | 462.55ns |
| 256 | 124.13ns | 9.88ns | 421.73ns | 554.56ns | 363.58ns | 459.42ns |
| 1k | 120.31ns | 8.20ns | 422.79ns | 569.74ns | 726.62ns | 466.40ns |
| 16k | 123.37ns | 8.07ns | 475.08ns | 604.11ns | 795.69ns | 522.11ns |
| 256k | 118.94ns | 8.64ns | 468.32ns | 654.74ns | 888.16ns | 613.12ns |
| 1M | 119.38ns | 9.89ns | 472.47ns | 858.61ns | 1.24us | 904.23ns |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Add `Velocity` to all entities.
- Remove `Velocity` from all entities.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/494dc137-d6c8-4474-a0a0-85b126113431)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 459.41ns | 757.78ns | 1.10us | 1.02us | 805.43ns | 2.48us |
| 4 | 576.53ns | 271.04ns | 1.07us | 989.76ns | 811.92ns | 1.99us |
| 16 | 445.13ns | 134.74ns | 1.04us | 994.92ns | 803.02ns | 1.88us |
| 64 | 448.75ns | 104.32ns | 1.04us | 965.71ns | 841.11ns | 1.82us |
| 256 | 403.82ns | 34.90ns | 1.03us | 963.50ns | 877.61ns | 1.84us |
| 1k | 416.70ns | 19.36ns | 1.10us | 989.45ns | 1.42us | 1.87us |
| 16k | 420.32ns | 25.88ns | 1.20us | 1.06us | 1.45us | 1.94us |
| 256k | 606.00ns | 39.53ns | 1.29us | 1.32us | 1.61us | 2.58us |
| 1M | 510.96ns | 43.42ns | 1.40us | 1.57us | 1.94us | 2.80us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/0c1ed1f6-91a8-4910-89be-18ff6e64b353)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 144.68ns | 771.99ns | 312.32ns | 393.36ns | 165.76ns | 519.71ns |
| 4 | 66.95ns | 215.44ns | 118.43ns | 166.34ns | 81.85ns | 249.54ns |
| 16 | 40.87ns | 64.89ns | 73.75ns | 152.49ns | 55.32ns | 189.69ns |
| 64 | 37.46ns | 31.20ns | 60.97ns | 120.45ns | 40.54ns | 161.50ns |
| 256 | 31.94ns | 12.14ns | 55.96ns | 95.05ns | 41.66ns | 130.29ns |
| 1k | 25.95ns | 7.15ns | 51.54ns | 105.74ns | 36.00ns | 131.92ns |
| 16k | 24.37ns | 6.36ns | 50.69ns | 117.29ns | 52.48ns | 153.32ns |
| 256k | 24.23ns | 6.75ns | 50.67ns | 204.72ns | 132.39ns | 296.57ns |
| 1M | 24.45ns | 7.50ns | 51.27ns | 306.89ns | 231.55ns | 375.36ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/e4e8243d-8315-41f2-9052-eea1954dca3b)

| N | Ark | Ark (batch) | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 207.34ns | 941.36ns | 497.61ns | 589.93ns | 180.38ns | 928.53ns |
| 4 | 146.13ns | 276.64ns | 234.39ns | 277.21ns | 80.55ns | 536.08ns |
| 16 | 118.02ns | 110.41ns | 160.77ns | 243.08ns | 50.98ns | 414.75ns |
| 64 | 99.22ns | 75.73ns | 145.47ns | 204.81ns | 46.08ns | 359.17ns |
| 256 | 86.82ns | 22.91ns | 119.62ns | 165.62ns | 40.84ns | 303.33ns |
| 1k | 79.61ns | 9.13ns | 133.90ns | 161.01ns | 35.12ns | 307.98ns |
| 16k | 77.63ns | 8.70ns | 123.77ns | 178.44ns | 50.17ns | 347.00ns |
| 256k | 91.26ns | 14.73ns | 156.21ns | 320.71ns | 100.92ns | 552.17ns |
| 1M | 88.40ns | 15.57ns | 137.25ns | 415.33ns | 227.35ns | 661.04ns |


### Create world

- Create a new world

| N | Ark | Donburi | ggecs | uot | Volt |
| --- | --- | --- | --- | --- | --- |
| 1 | 15.39us | 2.34us | 208.91us | 2.16us | 38.91us |


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
