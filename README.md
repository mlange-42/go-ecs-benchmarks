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

Last run: Fri, 10 Jan 2025 00:43:41 UTC  
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

![query2comp](https://github.com/user-attachments/assets/752a4daf-a8d3-4dcb-bd58-4a6e0575661c)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 54.09ns | 45.40ns | 69.60ns | 66.32ns | 43.62ns | 16.76ns |
| 4 | 15.31ns | 13.66ns | 33.68ns | 66.40ns | 13.77ns | 6.76ns |
| 16 | 5.83ns | 5.37ns | 24.90ns | 64.67ns | 6.80ns | 4.26ns |
| 64 | 3.38ns | 3.27ns | 23.70ns | 66.69ns | 5.48ns | 3.71ns |
| 256 | 2.80ns | 2.77ns | 23.55ns | 73.36ns | 5.12ns | 3.49ns |
| 1k | 2.89ns | 2.74ns | 22.92ns | 89.20ns | 5.01ns | 3.48ns |
| 16k | 2.78ns | 2.79ns | 23.48ns | 116.65ns | 5.00ns | 3.46ns |
| 256k | 2.82ns | 2.82ns | 25.40ns | 179.57ns | 5.02ns | 3.46ns |
| 1M | 2.85ns | 2.84ns | 29.50ns | 310.48ns | 5.03ns | 3.48ns |


### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/cea7a0d5-eca3-46c0-babd-ce44bf752c21)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.33ns | 49.14ns | 69.58ns | 97.46ns | 45.96ns | 16.82ns |
| 4 | 16.10ns | 14.01ns | 32.93ns | 94.17ns | 14.75ns | 6.78ns |
| 16 | 6.19ns | 5.47ns | 25.42ns | 94.26ns | 7.14ns | 4.24ns |
| 64 | 3.43ns | 3.32ns | 23.41ns | 102.10ns | 5.55ns | 3.78ns |
| 256 | 2.82ns | 2.80ns | 22.71ns | 112.68ns | 5.24ns | 3.49ns |
| 1k | 2.81ns | 2.87ns | 22.87ns | 133.90ns | 5.01ns | 3.44ns |
| 16k | 2.80ns | 2.82ns | 24.87ns | 176.08ns | 5.01ns | 3.45ns |
| 256k | 2.79ns | 2.80ns | 35.21ns | 358.39ns | 5.01ns | 3.47ns |


### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/c2eb28c0-dfe6-468f-aeb1-8bfcd306ce1f)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.96ns | 47.10ns | 68.78ns | 65.52ns | 43.19ns | 16.53ns |
| 4 | 23.01ns | 17.37ns | 43.74ns | 63.96ns | 17.94ns | 14.71ns |
| 16 | 16.40ns | 9.99ns | 35.10ns | 65.53ns | 12.25ns | 20.98ns |
| 64 | 9.92ns | 6.13ns | 29.31ns | 65.97ns | 8.25ns | 11.93ns |
| 256 | 4.78ns | 3.74ns | 23.83ns | 73.99ns | 5.72ns | 5.46ns |
| 1k | 3.29ns | 3.12ns | 24.75ns | 83.09ns | 5.32ns | 4.36ns |
| 16k | 2.95ns | 2.80ns | 31.79ns | 116.36ns | 5.06ns | 3.61ns |
| 256k | 2.83ns | 2.76ns | 71.09ns | 217.84ns | 5.02ns | 3.51ns |
| 1M | 2.87ns | 2.80ns | 101.20ns | 301.71ns | 5.03ns | 3.51ns |


### Component random access

`N` entities with component `Position`.
All entities are collected into a slice, and the slice is shuffled.

* Iterate the shuffled entities.
* For each entity, get its `Position` and sum up their `X` fields.

Ento is left out here, as component access for a specific entity seems broken (see issue [ento/#2](https://github.com/wwfranczyk/ento/issues/2)).

![random](https://github.com/user-attachments/assets/f0aaa4cd-0839-42fd-a4a7-b7fa4e0192bc)

| N | Arche | Donburi | ggecs | uot |
| --- | --- | --- | --- | --- |
| 1 | 2.32ns | 8.25ns | 8.42ns | 32.51ns |
| 4 | 2.30ns | 8.04ns | 8.77ns | 33.13ns |
| 16 | 2.35ns | 8.01ns | 11.52ns | 32.59ns |
| 64 | 2.25ns | 8.11ns | 11.70ns | 37.25ns |
| 256 | 2.26ns | 8.55ns | 11.86ns | 37.34ns |
| 1k | 2.54ns | 12.04ns | 13.20ns | 38.29ns |
| 16k | 5.93ns | 38.67ns | 35.48ns | 57.41ns |
| 256k | 8.89ns | 109.70ns | 51.13ns | 89.08ns |
| 1M | 38.46ns | 261.73ns | 137.16ns | 252.49ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/e786401b-039c-45da-8b37-af341424552b)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 234.15ns | 271.10ns | 1.08us | 1.19us | 325.04ns | 398.29ns |
| 4 | 83.38ns | 72.24ns | 566.30ns | 571.16ns | 162.62ns | 177.50ns |
| 16 | 52.31ns | 26.51ns | 367.11ns | 404.57ns | 140.26ns | 139.16ns |
| 64 | 39.60ns | 14.56ns | 305.03ns | 332.20ns | 121.03ns | 109.97ns |
| 256 | 39.33ns | 10.64ns | 240.10ns | 285.45ns | 97.07ns | 84.52ns |
| 1k | 31.30ns | 9.67ns | 217.59ns | 288.40ns | 95.14ns | 266.12ns |
| 16k | 26.90ns | 8.51ns | 237.21ns | 325.29ns | 109.34ns | 275.84ns |
| 256k | 26.45ns | 8.24ns | 231.67ns | 406.16ns | 105.69ns | 321.65ns |
| 1M | 26.23ns | 8.21ns | 214.27ns | 487.61ns | 167.92ns | 412.09ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/8af6919d-b95f-42aa-a1bc-3c67a7fb8c12)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 7.46us | 7.72us | 4.29us | 1.51us | 32.26us | 3.00us |
| 4 | 1.97us | 1.92us | 1.35us | 653.45ns | 7.78us | 987.74ns |
| 16 | 524.00ns | 486.92ns | 678.13ns | 482.54ns | 2.43us | 366.77ns |
| 64 | 163.01ns | 131.79ns | 453.22ns | 403.68ns | 895.92ns | 228.04ns |
| 256 | 69.85ns | 40.93ns | 343.45ns | 343.96ns | 411.59ns | 164.77ns |
| 1k | 46.40ns | 25.56ns | 310.25ns | 315.53ns | 287.17ns | 153.33ns |
| 16k | 63.55ns | 18.23ns | 428.68ns | 546.66ns | 240.93ns | 192.63ns |
| 256k | 137.57ns | 23.62ns | 558.91ns | 668.45ns | 620.63ns | 246.39ns |
| 1M | 120.66ns | 25.05ns | 485.58ns | 744.06ns | 1.39us | 287.85ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/d1bad1eb-c356-4ca9-bbee-c366be9f5078)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 369.05ns | 468.03ns | 2.12us | 2.28us | 435.53ns | 605.16ns |
| 4 | 160.44ns | 127.57ns | 1.35us | 1.67us | 271.87ns | 366.24ns |
| 16 | 119.15ns | 41.30ns | 1.06us | 1.26us | 262.65ns | 314.31ns |
| 64 | 88.96ns | 18.87ns | 821.44ns | 1.02us | 184.96ns | 270.13ns |
| 256 | 90.18ns | 11.46ns | 777.21ns | 1.07us | 156.59ns | 219.96ns |
| 1k | 75.90ns | 10.12ns | 784.57ns | 1.16us | 148.31ns | 480.48ns |
| 16k | 76.26ns | 8.21ns | 840.40ns | 1.22us | 158.04ns | 484.59ns |
| 256k | 76.45ns | 8.57ns | 789.20ns | 2.17us | 163.63ns | 559.23ns |
| 1M | 76.61ns | 8.19ns | 819.62ns | 2.33us | 248.16ns | 631.11ns |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/886c0f81-9475-421f-91cb-a56c34a8edf2)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 206.63ns | 1.13us | 592.24ns | 112.37ns | 559.62ns | 373.92ns |
| 4 | 142.64ns | 285.12ns | 527.68ns | 107.69ns | 528.89ns | 350.62ns |
| 16 | 121.04ns | 80.08ns | 517.88ns | 105.42ns | 540.65ns | 348.99ns |
| 64 | 114.43ns | 23.73ns | 504.55ns | 107.10ns | 542.73ns | 356.81ns |
| 256 | 114.08ns | 10.33ns | 493.55ns | 113.62ns | 553.55ns | 367.28ns |
| 1k | 112.41ns | 7.18ns | 491.85ns | 121.07ns | 556.11ns | 742.85ns |
| 16k | 115.70ns | 7.51ns | 544.84ns | 137.68ns | 584.39ns | 803.28ns |
| 256k | 113.17ns | 8.28ns | 519.81ns | 166.59ns | 663.81ns | 1.05us |
| 1M | 113.64ns | 10.01ns | 543.16ns | 237.42ns | 856.56ns | 1.36us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/8350bf39-f0e4-4724-be9f-9f1fa484f834)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 535.47ns | 6.09us | 1.21us | 111.12ns | 981.41ns | 884.07ns |
| 4 | 484.96ns | 1.45us | 1.17us | 105.67ns | 993.35ns | 878.02ns |
| 16 | 471.07ns | 367.78ns | 1.18us | 102.72ns | 984.02ns | 858.31ns |
| 64 | 450.58ns | 111.28ns | 1.12us | 103.85ns | 992.04ns | 869.07ns |
| 256 | 455.15ns | 36.19ns | 1.13us | 112.48ns | 978.64ns | 962.08ns |
| 1k | 459.53ns | 18.75ns | 1.23us | 122.41ns | 1.00us | 1.53us |
| 16k | 474.52ns | 21.82ns | 1.29us | 132.41ns | 1.04us | 1.54us |
| 256k | 595.43ns | 40.85ns | 1.51us | 191.04ns | 1.42us | 1.98us |
| 1M | 571.88ns | 43.55ns | 1.29us | 242.12ns | 1.54us | 2.07us |


### Delete entities

`N` entities with components `Position` and `Velocity`.

* Delete all entities

![delete2comp](https://github.com/user-attachments/assets/003951b7-2dca-4fd6-acd8-010e50342f61)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 175.27ns | 1.23us | 256.86ns | 359.03ns | 316.57ns | 157.78ns |
| 4 | 79.64ns | 320.83ns | 104.55ns | 193.85ns | 159.36ns | 75.48ns |
| 16 | 53.25ns | 81.69ns | 65.19ns | 133.09ns | 142.01ns | 50.43ns |
| 64 | 41.98ns | 25.67ns | 57.37ns | 118.70ns | 127.42ns | 43.92ns |
| 256 | 40.58ns | 10.59ns | 52.09ns | 120.87ns | 114.61ns | 42.26ns |
| 1k | 33.83ns | 7.43ns | 46.08ns | 122.24ns | 106.00ns | 36.44ns |
| 16k | 32.61ns | 6.93ns | 45.40ns | 146.35ns | 112.14ns | 49.70ns |
| 256k | 33.11ns | 6.66ns | 48.75ns | 241.18ns | 141.41ns | 82.54ns |
| 1M | 32.85ns | 7.88ns | 44.87ns | 333.80ns | 218.86ns | 183.86ns |


### Delete large entities

`N` entities with 10 components `C1`, ..., `C10`.

* Delete all entities

![delete10comp](https://github.com/user-attachments/assets/19d88591-5f15-4491-817b-1ae0ece93a61)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 245.95ns | 1.24us | 404.15ns | 938.47ns | 452.50ns | 152.63ns |
| 4 | 168.31ns | 304.04ns | 212.24ns | 690.78ns | 251.31ns | 69.07ns |
| 16 | 135.72ns | 84.96ns | 163.20ns | 554.06ns | 219.75ns | 51.24ns |
| 64 | 125.96ns | 28.18ns | 145.15ns | 491.04ns | 218.79ns | 47.40ns |
| 256 | 120.02ns | 11.34ns | 122.50ns | 600.89ns | 195.75ns | 40.95ns |
| 1k | 108.53ns | 7.59ns | 130.65ns | 639.09ns | 183.74ns | 34.74ns |
| 16k | 104.61ns | 6.46ns | 129.37ns | 731.98ns | 187.54ns | 49.05ns |
| 256k | 120.97ns | 6.60ns | 151.47ns | 1.74us | 249.48ns | 74.52ns |
| 1M | 113.29ns | 7.91ns | 220.49ns | 1.82us | 319.35ns | 153.36ns |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 10.77us | 6.48us | 67.12us | 110.31us | 2.24us |


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
