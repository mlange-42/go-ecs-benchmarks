# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.15.2 |
| [Donburi](https://github.com/yohamta/donburi) | Archetype | v1.15.6 |
| [ento](https://github.com/wwfranczyk/ento) | Sparse Set | v0.1.0 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | Archetype | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | Archetype | v0.0.3 |

Candidates are always displayed in alphabetical order.

In case you develop or use a Go ECS that is not in the list and that want to see here,
please open an issue or make a pull request.

In case you are a developer or user of an implementation included here,
feel free to check the benchmarked code for any possible improvements.
Open an issue if you want a version update.

## Benchmarks

Last run: Mon, 13 Jan 2025 00:07:46 UTC  
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

![query2comp](https://github.com/user-attachments/assets/4b9b3127-f2b8-447c-8bc2-f4c9c61690bf)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.68ns | 45.54ns | 68.89ns | 65.93ns | 43.70ns | 17.04ns |
| 4 | 15.39ns | 13.54ns | 33.05ns | 64.76ns | 13.67ns | 6.74ns |
| 16 | 5.85ns | 5.58ns | 25.01ns | 63.95ns | 7.93ns | 4.25ns |
| 64 | 3.39ns | 4.38ns | 24.60ns | 67.30ns | 5.52ns | 3.62ns |
| 256 | 3.12ns | 2.76ns | 23.00ns | 75.03ns | 5.10ns | 3.49ns |
| 1k | 2.66ns | 2.68ns | 22.79ns | 83.70ns | 5.01ns | 3.48ns |
| 16k | 2.75ns | 2.73ns | 24.03ns | 115.22ns | 5.01ns | 3.46ns |
| 256k | 2.74ns | 2.74ns | 24.81ns | 159.08ns | 5.01ns | 3.45ns |
| 1M | 2.78ns | 2.80ns | 29.89ns | 315.63ns | 5.02ns | 3.50ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/26423e6e-db5e-4dec-8b9c-43ded028f7c0)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.52ns | 48.52ns | 68.61ns | 65.42ns | 43.33ns | 16.76ns |
| 4 | 22.79ns | 17.65ns | 38.86ns | 64.45ns | 17.95ns | 14.67ns |
| 16 | 16.51ns | 9.76ns | 33.37ns | 62.50ns | 11.84ns | 20.08ns |
| 64 | 9.59ns | 6.18ns | 27.43ns | 65.84ns | 8.29ns | 11.79ns |
| 256 | 4.66ns | 3.76ns | 24.58ns | 75.08ns | 5.65ns | 5.53ns |
| 1k | 3.42ns | 3.12ns | 25.80ns | 84.79ns | 5.18ns | 4.12ns |
| 16k | 2.90ns | 2.88ns | 32.09ns | 115.50ns | 5.05ns | 3.62ns |
| 256k | 2.82ns | 2.83ns | 66.19ns | 167.88ns | 5.02ns | 3.47ns |
| 1M | 2.84ns | 2.85ns | 102.38ns | 292.20ns | 5.02ns | 3.50ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/b40f4cee-366f-41b7-a6fe-dead3e5d14d2)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 61.02ns | 46.00ns | 71.90ns | 76.39ns | 53.01ns | 16.78ns |
| 4 | 27.38ns | 13.60ns | 33.10ns | 77.61ns | 24.85ns | 7.36ns |
| 16 | 18.61ns | 5.50ns | 24.89ns | 74.77ns | 18.67ns | 4.38ns |
| 64 | 16.07ns | 3.37ns | 22.76ns | 79.72ns | 17.48ns | 3.78ns |
| 256 | 6.12ns | 2.78ns | 22.74ns | 86.33ns | 8.16ns | 3.65ns |
| 1k | 3.68ns | 2.74ns | 22.80ns | 103.34ns | 5.81ns | 3.44ns |
| 16k | 2.76ns | 2.75ns | 23.35ns | 127.83ns | 5.07ns | 3.47ns |
| 256k | 2.73ns | 2.74ns | 25.67ns | 286.54ns | 5.01ns | 3.46ns |
| 1M | 2.74ns | 2.77ns | 28.93ns | 365.01ns | 5.02ns | 3.48ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

Ento is left out here, as component access for a specific entity seems broken (see issue [ento/#2](https://github.com/wwfranczyk/ento/issues/2)).

![random](https://github.com/user-attachments/assets/4d2c48af-86d3-4b8a-8add-1e427145d4a8)

| N | Arche | Donburi | ggecs | uot |
| --- | --- | --- | --- | --- |
| 1 | 2.49ns | 8.21ns | 9.01ns | 33.57ns |
| 4 | 2.30ns | 8.06ns | 8.88ns | 33.85ns |
| 16 | 2.25ns | 8.10ns | 11.97ns | 32.96ns |
| 64 | 2.23ns | 8.09ns | 12.08ns | 36.83ns |
| 256 | 2.25ns | 8.70ns | 12.10ns | 37.74ns |
| 1k | 2.53ns | 11.93ns | 13.23ns | 37.83ns |
| 16k | 5.90ns | 39.40ns | 36.44ns | 55.61ns |
| 256k | 6.96ns | 83.58ns | 40.77ns | 73.05ns |
| 1M | 35.72ns | 223.32ns | 121.70ns | 195.58ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/034463b7-7ab2-4549-9e16-81a968fa6428)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 232.83ns | 281.90ns | 1.07us | 1.16us | 316.65ns | 379.93ns |
| 4 | 84.23ns | 73.81ns | 514.73ns | 612.13ns | 157.42ns | 189.95ns |
| 16 | 49.87ns | 23.99ns | 341.85ns | 416.09ns | 148.16ns | 135.74ns |
| 64 | 37.45ns | 13.55ns | 291.24ns | 333.28ns | 112.97ns | 117.19ns |
| 256 | 33.82ns | 10.83ns | 239.23ns | 295.62ns | 108.81ns | 85.80ns |
| 1k | 30.63ns | 10.07ns | 221.68ns | 299.14ns | 94.90ns | 259.30ns |
| 16k | 26.40ns | 8.36ns | 234.81ns | 328.49ns | 108.49ns | 274.28ns |
| 256k | 26.90ns | 8.16ns | 227.35ns | 418.23ns | 107.97ns | 347.15ns |
| 1M | 27.05ns | 8.18ns | 238.31ns | 502.74ns | 174.53ns | 389.69ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/b0fc539a-f0fa-421e-806e-85619e998b0d)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.92us | 7.94us | 4.24us | 1.46us | 34.56us | 3.02us |
| 4 | 2.04us | 2.00us | 1.32us | 652.44ns | 8.23us | 1.02us |
| 16 | 533.65ns | 499.85ns | 667.05ns | 473.42ns | 2.44us | 373.26ns |
| 64 | 158.11ns | 136.81ns | 443.14ns | 403.16ns | 909.86ns | 224.53ns |
| 256 | 66.63ns | 43.20ns | 339.62ns | 331.19ns | 403.96ns | 165.64ns |
| 1k | 45.22ns | 25.11ns | 307.20ns | 304.97ns | 229.15ns | 148.85ns |
| 16k | 60.98ns | 17.65ns | 422.54ns | 547.09ns | 243.79ns | 175.11ns |
| 256k | 127.21ns | 21.86ns | 504.42ns | 705.97ns | 656.11ns | 244.40ns |
| 1M | 144.41ns | 21.08ns | 483.61ns | 732.32ns | 1.41us | 311.11ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/3ca88394-f91f-43bb-9b1f-9fcd82079cae)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 381.45ns | 405.97ns | 2.19us | 2.36us | 466.82ns | 641.74ns |
| 4 | 163.75ns | 116.45ns | 1.55us | 1.57us | 270.86ns | 379.09ns |
| 16 | 108.75ns | 34.12ns | 1.13us | 1.31us | 242.79ns | 306.33ns |
| 64 | 90.91ns | 18.02ns | 837.06ns | 1.03us | 199.16ns | 250.48ns |
| 256 | 83.45ns | 11.26ns | 787.92ns | 1.07us | 151.71ns | 221.65ns |
| 1k | 79.70ns | 9.46ns | 787.46ns | 1.17us | 147.08ns | 506.25ns |
| 16k | 75.93ns | 8.26ns | 838.31ns | 1.35us | 169.46ns | 491.60ns |
| 256k | 76.07ns | 8.21ns | 831.76ns | 2.31us | 222.17ns | 655.17ns |
| 1M | 77.12ns | 8.20ns | 741.33ns | 2.36us | 234.25ns | 677.30ns |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/1588a7f3-17a8-4301-8dc0-f8cee6311920)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 205.78ns | 544.95ns | 549.77ns | 108.83ns | 555.97ns | 390.28ns |
| 4 | 154.51ns | 149.45ns | 491.03ns | 105.77ns | 529.49ns | 366.09ns |
| 16 | 130.20ns | 57.24ns | 485.23ns | 102.25ns | 562.38ns | 350.18ns |
| 64 | 115.89ns | 26.75ns | 471.75ns | 105.10ns | 556.61ns | 368.09ns |
| 256 | 113.05ns | 10.37ns | 468.26ns | 111.19ns | 560.19ns | 380.73ns |
| 1k | 114.87ns | 7.21ns | 470.07ns | 119.95ns | 570.46ns | 751.55ns |
| 16k | 113.26ns | 7.27ns | 510.32ns | 135.51ns | 583.30ns | 807.56ns |
| 256k | 115.15ns | 8.23ns | 507.59ns | 158.14ns | 631.06ns | 1.04us |
| 1M | 117.95ns | 9.86ns | 506.28ns | 212.29ns | 808.20ns | 1.33us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/64f13698-3269-4461-b43c-0ad46bbce422)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 513.13ns | 863.67ns | 1.19us | 107.63ns | 982.02ns | 870.16ns |
| 4 | 482.15ns | 275.84ns | 1.13us | 109.06ns | 999.90ns | 857.10ns |
| 16 | 477.54ns | 145.68ns | 1.09us | 100.57ns | 1.00us | 839.46ns |
| 64 | 448.01ns | 97.92ns | 1.08us | 107.79ns | 1.02us | 863.01ns |
| 256 | 447.83ns | 34.67ns | 1.10us | 110.28ns | 977.82ns | 910.66ns |
| 1k | 452.44ns | 18.56ns | 1.16us | 120.14ns | 985.69ns | 1.46us |
| 16k | 469.08ns | 21.25ns | 1.26us | 130.90ns | 1.03us | 1.51us |
| 256k | 596.94ns | 39.15ns | 1.40us | 171.65ns | 1.32us | 1.79us |
| 1M | 569.66ns | 43.66ns | 1.44us | 228.90ns | 1.50us | 2.01us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/359c6918-4a4b-4097-b6ea-eae2d4860155)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 179.64ns | 655.92ns | 264.53ns | 409.83ns | 350.76ns | 186.59ns |
| 4 | 84.88ns | 194.43ns | 109.26ns | 209.99ns | 180.04ns | 84.40ns |
| 16 | 53.65ns | 61.05ns | 63.82ns | 132.67ns | 146.55ns | 53.82ns |
| 64 | 42.39ns | 29.76ns | 54.97ns | 107.89ns | 113.58ns | 42.13ns |
| 256 | 41.92ns | 11.80ns | 51.84ns | 120.93ns | 111.96ns | 44.26ns |
| 1k | 33.71ns | 7.13ns | 45.54ns | 123.60ns | 105.66ns | 34.90ns |
| 16k | 33.02ns | 6.94ns | 46.02ns | 144.35ns | 111.78ns | 48.51ns |
| 256k | 32.92ns | 7.41ns | 49.05ns | 292.88ns | 161.09ns | 108.54ns |
| 1M | 32.97ns | 7.93ns | 61.30ns | 353.08ns | 234.54ns | 211.21ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/78b38763-7ae2-493c-9e02-7f96f6ff1e5b)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 233.79ns | 643.56ns | 405.71ns | 995.89ns | 535.12ns | 171.71ns |
| 4 | 158.80ns | 184.86ns | 209.80ns | 634.32ns | 270.82ns | 75.87ns |
| 16 | 130.22ns | 59.33ns | 157.38ns | 545.30ns | 283.95ns | 52.35ns |
| 64 | 120.24ns | 27.29ns | 147.41ns | 524.64ns | 248.39ns | 44.49ns |
| 256 | 117.38ns | 11.31ns | 116.55ns | 601.05ns | 177.92ns | 43.84ns |
| 1k | 104.89ns | 7.86ns | 138.79ns | 649.77ns | 178.73ns | 34.06ns |
| 16k | 105.78ns | 6.63ns | 136.54ns | 782.35ns | 181.50ns | 48.46ns |
| 256k | 119.14ns | 6.88ns | 153.83ns | 1.72us | 301.13ns | 105.77ns |
| 1M | 113.10ns | 7.99ns | 119.87ns | 1.87us | 354.10ns | 244.57ns |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 9.76us | 6.51us | 65.15us | 102.31us | 2.34us |


### Popularity

Given that all tested projects are on Github, we can use the star history as a proxy here.

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&yohamta/donburi&wwfranczyk/ento&marioolofo/go-gameengine-ecs&unitoftime/ecs&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,yohamta/donburi,wwfranczyk/ento,marioolofo/go-gameengine-ecs,unitoftime/ecs&type=Date" alt="Star History Chart" width="600"/>
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
