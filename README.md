# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.15.1 |
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

Last run: Fri, 10 Jan 2025 01:48:20 UTC  
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

![query2comp](https://github.com/user-attachments/assets/174b1c7b-3462-42ca-8935-a93937013ad0)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.65ns | 45.67ns | 68.73ns | 65.96ns | 43.56ns | 16.74ns |
| 4 | 15.43ns | 13.53ns | 32.95ns | 65.63ns | 13.69ns | 6.77ns |
| 16 | 5.99ns | 5.50ns | 24.99ns | 63.02ns | 7.93ns | 4.25ns |
| 64 | 3.43ns | 3.28ns | 22.69ns | 65.33ns | 5.53ns | 3.70ns |
| 256 | 2.80ns | 3.04ns | 22.67ns | 75.73ns | 5.10ns | 3.49ns |
| 1k | 2.79ns | 2.65ns | 23.03ns | 85.32ns | 5.00ns | 3.43ns |
| 16k | 2.74ns | 2.73ns | 23.92ns | 117.65ns | 5.03ns | 3.47ns |
| 256k | 2.75ns | 2.75ns | 27.43ns | 267.22ns | 5.02ns | 3.47ns |
| 1M | 2.80ns | 2.79ns | 30.34ns | 347.15ns | 5.03ns | 3.49ns |


### Query fragmented, inner

Query where the matching entities are fragmented over 32 archetypes.

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over up to 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/afc12f58-8ced-4297-9b60-3b9c8ddb45c7)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.30ns | 48.72ns | 68.59ns | 65.26ns | 43.14ns | 16.83ns |
| 4 | 23.01ns | 17.53ns | 39.68ns | 66.42ns | 18.18ns | 14.87ns |
| 16 | 16.28ns | 10.68ns | 34.22ns | 62.81ns | 12.74ns | 20.55ns |
| 64 | 9.54ns | 6.29ns | 27.95ns | 67.94ns | 8.30ns | 11.69ns |
| 256 | 4.66ns | 3.79ns | 28.81ns | 74.41ns | 5.75ns | 5.85ns |
| 1k | 3.29ns | 3.15ns | 26.03ns | 86.77ns | 5.19ns | 4.35ns |
| 16k | 2.87ns | 2.82ns | 31.88ns | 117.72ns | 5.05ns | 3.60ns |
| 256k | 2.82ns | 2.82ns | 84.36ns | 279.44ns | 5.02ns | 3.50ns |
| 1M | 2.87ns | 2.91ns | 109.06ns | 314.42ns | 5.03ns | 3.51ns |


### Query fragmented, outer

Query where there are 256 non-matching archetypes.

`N` entities with components `Position` and `Velocity`.
Another `4 * N` entities with `Position` and some combination of 8 components
`C1`, ..., `C8`, so these entities are fragmented over up to 256 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query256arch](https://github.com/user-attachments/assets/347e61cd-9a67-4592-a543-24b1472fb412)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 60.90ns | 45.79ns | 71.71ns | 81.00ns | 52.34ns | 16.73ns |
| 4 | 27.71ns | 13.58ns | 33.18ns | 78.55ns | 24.92ns | 7.44ns |
| 16 | 18.52ns | 5.38ns | 24.80ns | 76.39ns | 18.77ns | 4.38ns |
| 64 | 16.26ns | 4.14ns | 23.74ns | 80.46ns | 17.33ns | 3.77ns |
| 256 | 6.55ns | 2.79ns | 22.71ns | 87.30ns | 8.16ns | 3.49ns |
| 1k | 3.75ns | 2.76ns | 23.98ns | 104.44ns | 5.81ns | 3.47ns |
| 16k | 2.77ns | 2.72ns | 24.09ns | 138.12ns | 5.10ns | 3.45ns |
| 256k | 2.72ns | 2.77ns | 27.78ns | 322.80ns | 5.01ns | 3.48ns |
| 1M | 2.77ns | 2.80ns | 31.31ns | 354.46ns | 5.03ns | 3.48ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

Ento is left out here, as component access for a specific entity seems broken (see issue [ento/#2](https://github.com/wwfranczyk/ento/issues/2)).

![random](https://github.com/user-attachments/assets/21f3e269-f398-4aa2-bfdc-90978b9d3d23)

| N | Arche | Donburi | ggecs | uot |
| --- | --- | --- | --- | --- |
| 1 | 2.49ns | 8.33ns | 8.99ns | 32.23ns |
| 4 | 2.30ns | 8.21ns | 9.00ns | 31.91ns |
| 16 | 2.25ns | 8.02ns | 11.97ns | 32.54ns |
| 64 | 2.23ns | 8.10ns | 12.00ns | 35.61ns |
| 256 | 2.25ns | 8.59ns | 12.02ns | 37.11ns |
| 1k | 2.56ns | 11.90ns | 13.24ns | 38.45ns |
| 16k | 5.96ns | 40.20ns | 36.16ns | 56.38ns |
| 256k | 11.03ns | 195.88ns | 93.16ns | 198.66ns |
| 1M | 41.37ns | 285.23ns | 150.36ns | 261.65ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/d8aaf55e-b599-442b-98c2-e9bcdaa7845c)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 206.81ns | 241.02ns | 1.17us | 1.17us | 366.29ns | 395.89ns |
| 4 | 87.16ns | 67.69ns | 550.62ns | 569.20ns | 181.13ns | 192.60ns |
| 16 | 51.95ns | 25.04ns | 364.04ns | 388.53ns | 139.91ns | 133.70ns |
| 64 | 43.58ns | 13.98ns | 304.61ns | 338.47ns | 128.42ns | 128.80ns |
| 256 | 49.61ns | 10.80ns | 249.21ns | 282.45ns | 99.31ns | 84.65ns |
| 1k | 28.16ns | 9.93ns | 223.27ns | 293.61ns | 94.91ns | 258.05ns |
| 16k | 26.05ns | 8.20ns | 241.56ns | 318.34ns | 108.67ns | 279.07ns |
| 256k | 26.14ns | 8.21ns | 230.67ns | 395.69ns | 106.64ns | 332.36ns |
| 1M | 26.21ns | 8.23ns | 216.43ns | 465.54ns | 166.21ns | 430.73ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/8ee15eab-08d0-4df0-8ae2-adeb45daf53a)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.98us | 8.35us | 4.21us | 1.51us | 38.26us | 3.09us |
| 4 | 2.10us | 2.10us | 1.37us | 697.33ns | 8.97us | 1.03us |
| 16 | 550.34ns | 534.10ns | 689.09ns | 491.39ns | 2.70us | 362.61ns |
| 64 | 167.81ns | 139.14ns | 444.75ns | 385.40ns | 914.92ns | 229.68ns |
| 256 | 76.77ns | 43.88ns | 346.36ns | 316.17ns | 422.44ns | 166.25ns |
| 1k | 47.18ns | 25.32ns | 330.35ns | 301.26ns | 245.70ns | 156.01ns |
| 16k | 59.32ns | 21.50ns | 434.46ns | 530.30ns | 233.95ns | 169.72ns |
| 256k | 151.72ns | 20.58ns | 481.34ns | 729.59ns | 625.84ns | 241.22ns |
| 1M | 185.05ns | 26.45ns | 478.05ns | 801.87ns | 1.40us | 291.65ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/9428185f-5931-4c06-a51e-f015f82fa1dd)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 367.82ns | 522.03ns | 2.06us | 2.27us | 454.29ns | 624.47ns |
| 4 | 161.53ns | 138.76ns | 1.36us | 1.72us | 253.00ns | 385.25ns |
| 16 | 109.83ns | 41.90ns | 937.97ns | 1.26us | 234.67ns | 281.71ns |
| 64 | 91.88ns | 17.69ns | 865.40ns | 1.02us | 184.92ns | 259.10ns |
| 256 | 81.50ns | 12.03ns | 783.58ns | 1.07us | 160.12ns | 219.56ns |
| 1k | 78.28ns | 9.85ns | 779.73ns | 1.14us | 145.24ns | 467.57ns |
| 16k | 76.62ns | 8.28ns | 850.48ns | 1.25us | 159.09ns | 496.64ns |
| 256k | 75.85ns | 8.33ns | 839.79ns | 2.11us | 170.85ns | 594.09ns |
| 1M | 75.78ns | 8.22ns | 726.22ns | 2.26us | 222.73ns | 652.70ns |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/e4c4cee8-379f-4b6d-add0-5b2cdd963bba)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 205.16ns | 1.11us | 591.87ns | 112.33ns | 558.58ns | 370.19ns |
| 4 | 140.08ns | 284.56ns | 521.02ns | 105.34ns | 523.58ns | 352.33ns |
| 16 | 120.17ns | 75.50ns | 514.24ns | 102.19ns | 539.58ns | 334.84ns |
| 64 | 113.99ns | 23.35ns | 502.70ns | 102.44ns | 539.02ns | 345.14ns |
| 256 | 112.13ns | 10.30ns | 493.37ns | 110.48ns | 555.37ns | 362.92ns |
| 1k | 112.67ns | 7.10ns | 491.88ns | 119.67ns | 552.36ns | 728.10ns |
| 16k | 113.27ns | 7.62ns | 538.59ns | 132.81ns | 584.11ns | 792.04ns |
| 256k | 115.66ns | 8.19ns | 530.48ns | 154.71ns | 704.27ns | 1.24us |
| 1M | 114.00ns | 9.75ns | 528.78ns | 218.82ns | 849.74ns | 1.34us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/c8f6974b-1a0e-49c7-bc99-df2640c5ef92)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 530.40ns | 5.66us | 1.24us | 110.62ns | 997.04ns | 860.09ns |
| 4 | 514.03ns | 1.45us | 1.20us | 105.96ns | 1.01us | 873.62ns |
| 16 | 472.79ns | 408.29ns | 1.16us | 102.92ns | 1.00us | 870.28ns |
| 64 | 461.12ns | 106.49ns | 1.14us | 105.31ns | 1.01us | 892.90ns |
| 256 | 451.19ns | 35.09ns | 1.19us | 111.51ns | 983.53ns | 946.98ns |
| 1k | 470.04ns | 18.53ns | 1.22us | 120.64ns | 991.63ns | 1.49us |
| 16k | 479.30ns | 25.67ns | 1.37us | 132.02ns | 1.08us | 1.55us |
| 256k | 619.06ns | 42.44ns | 1.42us | 209.29ns | 1.47us | 2.06us |
| 1M | 587.41ns | 44.62ns | 1.51us | 223.95ns | 1.55us | 2.07us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/27db6b7a-9f49-4c02-972b-987e46e0f205)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 169.35ns | 1.29us | 252.11ns | 371.64ns | 323.87ns | 169.66ns |
| 4 | 79.50ns | 318.44ns | 106.14ns | 194.86ns | 152.02ns | 78.21ns |
| 16 | 52.05ns | 84.90ns | 65.50ns | 131.57ns | 146.56ns | 53.73ns |
| 64 | 42.95ns | 26.76ns | 54.57ns | 115.98ns | 132.20ns | 41.70ns |
| 256 | 41.88ns | 12.08ns | 49.60ns | 123.09ns | 116.60ns | 40.51ns |
| 1k | 33.06ns | 8.01ns | 45.67ns | 122.06ns | 104.78ns | 35.46ns |
| 16k | 32.73ns | 6.85ns | 45.75ns | 147.40ns | 112.15ns | 48.97ns |
| 256k | 32.76ns | 7.06ns | 47.61ns | 202.09ns | 130.82ns | 118.54ns |
| 1M | 32.74ns | 7.90ns | 64.26ns | 339.52ns | 254.64ns | 236.91ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/5ce68e1f-e24b-47d9-a541-ba4d826016cd)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 263.11ns | 1.28us | 399.27ns | 985.62ns | 444.45ns | 160.34ns |
| 4 | 163.97ns | 310.01ns | 199.91ns | 663.17ns | 277.75ns | 75.96ns |
| 16 | 146.11ns | 82.05ns | 156.50ns | 619.42ns | 258.94ns | 52.65ns |
| 64 | 127.59ns | 29.30ns | 136.73ns | 533.46ns | 204.38ns | 40.82ns |
| 256 | 114.76ns | 11.66ns | 115.82ns | 611.04ns | 182.22ns | 40.92ns |
| 1k | 108.17ns | 7.30ns | 128.51ns | 645.50ns | 180.66ns | 35.78ns |
| 16k | 106.37ns | 6.59ns | 135.96ns | 771.60ns | 198.29ns | 49.09ns |
| 256k | 115.93ns | 7.27ns | 163.11ns | 1.81us | 337.78ns | 140.02ns |
| 1M | 111.74ns | 8.00ns | 121.42ns | 1.82us | 371.67ns | 244.88ns |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 9.46us | 6.78us | 67.24us | 101.07us | 2.28us |


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
