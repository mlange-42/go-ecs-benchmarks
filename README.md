# Go ECS Benchmarks

Comparative benchmarks for Go Entity Component System (ECS) implementations.

> Disclaimer: This repository is maintained by the author of
> [Arche](https://github.com/mlange-42/arche), an archetype-based ECS for Go.

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.15.0 |
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

Last run: Thu, 09 Jan 2025 13:19:17 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/a0b74fa0-7cb9-4365-b724-bd11ceb24f91)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.65ns | 46.30ns | 68.79ns | 66.41ns | 43.33ns | 16.79ns |
| 4 | 15.52ns | 13.30ns | 32.95ns | 63.86ns | 13.73ns | 6.50ns |
| 16 | 5.94ns | 5.36ns | 24.94ns | 63.66ns | 6.93ns | 3.96ns |
| 64 | 3.39ns | 3.26ns | 23.05ns | 65.05ns | 5.78ns | 3.32ns |
| 256 | 2.80ns | 2.77ns | 22.63ns | 74.19ns | 5.44ns | 3.18ns |
| 1k | 2.96ns | 2.83ns | 22.73ns | 82.54ns | 5.30ns | 3.14ns |
| 16k | 2.81ns | 2.93ns | 23.72ns | 115.81ns | 5.32ns | 3.15ns |
| 256k | 2.81ns | 2.84ns | 24.80ns | 153.18ns | 5.32ns | 3.15ns |
| 1M | 2.84ns | 2.84ns | 32.23ns | 284.30ns | 5.34ns | 3.18ns |


### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/05139999-9ebb-4091-b202-5fb5d76abe40)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.08ns | 48.53ns | 69.05ns | 93.28ns | 46.46ns | 17.40ns |
| 4 | 18.19ns | 13.79ns | 32.90ns | 93.39ns | 15.27ns | 6.66ns |
| 16 | 6.60ns | 5.43ns | 24.96ns | 96.92ns | 7.27ns | 3.99ns |
| 64 | 3.59ns | 3.31ns | 23.59ns | 102.09ns | 5.87ns | 3.43ns |
| 256 | 2.83ns | 3.12ns | 22.67ns | 113.05ns | 5.42ns | 3.19ns |
| 1k | 2.81ns | 2.85ns | 23.02ns | 132.72ns | 5.34ns | 3.13ns |
| 16k | 2.78ns | 2.79ns | 25.31ns | 165.42ns | 5.30ns | 3.18ns |
| 256k | 2.79ns | 2.79ns | 34.08ns | 351.19ns | 5.31ns | 3.17ns |


### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/7744a95d-d0c1-4447-8912-2010de31ca14)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.39ns | 49.27ns | 68.87ns | 66.19ns | 43.00ns | 16.74ns |
| 4 | 23.76ns | 17.26ns | 38.51ns | 63.63ns | 17.75ns | 15.19ns |
| 16 | 16.38ns | 9.55ns | 33.10ns | 63.25ns | 11.90ns | 20.16ns |
| 64 | 9.61ns | 6.23ns | 27.46ns | 68.23ns | 8.12ns | 11.78ns |
| 256 | 4.69ns | 3.72ns | 24.05ns | 72.91ns | 5.80ns | 5.30ns |
| 1k | 3.29ns | 3.06ns | 24.26ns | 81.82ns | 5.43ns | 3.72ns |
| 16k | 3.04ns | 2.83ns | 31.82ns | 115.88ns | 5.38ns | 3.30ns |
| 256k | 2.82ns | 2.77ns | 71.40ns | 145.66ns | 5.32ns | 3.17ns |
| 1M | 2.85ns | 2.80ns | 97.97ns | 279.11ns | 5.40ns | 3.20ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/0a5ec854-bff9-4fc2-813c-5db65e75cc79)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 224.39ns | 262.90ns | 1.41us | 1.21us | 682.03ns | 1.03us |
| 4 | 83.49ns | 75.06ns | 654.02ns | 580.10ns | 317.80ns | 540.05ns |
| 16 | 50.69ns | 24.36ns | 380.37ns | 413.68ns | 208.81ns | 375.11ns |
| 64 | 38.06ns | 13.59ns | 296.58ns | 352.40ns | 151.74ns | 270.98ns |
| 256 | 37.43ns | 10.70ns | 249.75ns | 301.95ns | 137.50ns | 235.51ns |
| 1k | 31.14ns | 9.41ns | 234.62ns | 298.24ns | 123.68ns | 419.77ns |
| 16k | 26.86ns | 8.44ns | 263.42ns | 326.58ns | 137.14ns | 438.67ns |
| 256k | 26.50ns | 8.19ns | 263.47ns | 387.44ns | 140.63ns | 520.76ns |
| 1M | 26.91ns | 8.33ns | 282.57ns | 510.31ns | 215.85ns | 552.44ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/a2edbe9f-e575-43b1-8033-3bc3ea0b99b8)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.72us | 9.77us | 4.64us | 1.62us | 38.12us | 3.45us |
| 4 | 2.53us | 2.39us | 1.53us | 688.89ns | 10.14us | 1.30us |
| 16 | 663.36ns | 638.86ns | 779.79ns | 484.62ns | 2.98us | 623.75ns |
| 64 | 201.64ns | 169.46ns | 478.93ns | 398.74ns | 933.20ns | 418.55ns |
| 256 | 77.05ns | 50.68ns | 361.54ns | 332.04ns | 466.86ns | 327.74ns |
| 1k | 49.00ns | 27.30ns | 331.07ns | 320.27ns | 313.25ns | 286.46ns |
| 16k | 73.72ns | 41.48ns | 420.95ns | 521.32ns | 262.10ns | 322.21ns |
| 256k | 161.70ns | 36.55ns | 501.38ns | 688.80ns | 682.35ns | 398.57ns |
| 1M | 111.66ns | 21.86ns | 485.69ns | 782.48ns | 1.62us | 421.64ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/eeb8413b-5a10-456e-a7c1-1e39de0d1bef)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 473.47ns | 504.68ns | 2.11us | 2.21us | 443.79ns | 625.98ns |
| 4 | 188.63ns | 137.13ns | 1.32us | 1.56us | 254.70ns | 369.86ns |
| 16 | 117.50ns | 42.79ns | 1.06us | 1.22us | 241.16ns | 306.15ns |
| 64 | 90.97ns | 17.97ns | 830.60ns | 1.02us | 196.21ns | 262.65ns |
| 256 | 85.97ns | 11.31ns | 769.41ns | 1.05us | 169.99ns | 222.19ns |
| 1k | 79.85ns | 9.16ns | 761.25ns | 1.14us | 148.35ns | 516.47ns |
| 16k | 76.76ns | 8.19ns | 798.17ns | 1.17us | 164.75ns | 474.13ns |
| 256k | 76.61ns | 8.22ns | 782.37ns | 2.13us | 164.44ns | 539.68ns |
| 1M | 76.56ns | 8.17ns | 743.87ns | 2.25us | 239.30ns | 688.64ns |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/1387d6e0-33bf-4e56-a8ec-072504fcb69d)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 231.96ns | 1.13us | 554.57ns | 110.97ns | 560.22ns | 374.58ns |
| 4 | 151.42ns | 303.20ns | 481.20ns | 105.18ns | 521.39ns | 348.45ns |
| 16 | 123.91ns | 100.15ns | 466.80ns | 102.73ns | 546.88ns | 346.32ns |
| 64 | 116.36ns | 45.31ns | 455.97ns | 104.94ns | 534.65ns | 358.77ns |
| 256 | 115.88ns | 32.45ns | 452.99ns | 111.14ns | 543.65ns | 370.71ns |
| 1k | 113.13ns | 28.83ns | 454.95ns | 120.37ns | 551.52ns | 741.95ns |
| 16k | 114.44ns | 28.70ns | 501.59ns | 136.40ns | 577.14ns | 807.84ns |
| 256k | 115.86ns | 28.89ns | 497.69ns | 157.09ns | 619.23ns | 934.41ns |
| 1M | 114.61ns | 30.50ns | 480.75ns | 207.90ns | 795.66ns | 1.38us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/352146b8-da64-4c4a-b41b-f901615adb82)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 557.94ns | 5.72us | 1.21us | 110.82ns | 997.42ns | 878.52ns |
| 4 | 523.22ns | 1.64us | 1.15us | 105.71ns | 999.67ns | 870.81ns |
| 16 | 490.77ns | 596.22ns | 1.12us | 102.78ns | 996.46ns | 858.55ns |
| 64 | 451.87ns | 299.71ns | 1.10us | 105.11ns | 980.59ns | 876.35ns |
| 256 | 437.20ns | 227.68ns | 1.11us | 110.40ns | 990.03ns | 937.60ns |
| 1k | 464.97ns | 212.52ns | 1.17us | 120.74ns | 995.80ns | 1.48us |
| 16k | 464.73ns | 208.90ns | 1.27us | 134.17ns | 1.03us | 1.55us |
| 256k | 621.90ns | 236.57ns | 1.43us | 170.64ns | 1.29us | 1.66us |
| 1M | 567.79ns | 221.09ns | 1.48us | 222.72ns | 1.47us | 2.18us |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 8.78us | 6.63us | 72.40us | 453.57us | 5.34us |


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

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```
