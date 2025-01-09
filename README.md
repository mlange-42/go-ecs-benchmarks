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

Last run: Thu, 09 Jan 2025 17:42:03 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

Note that the Y axis has logarithmic scale in all plots.
So doubled bar or line height is not doubled time!

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/8a21ae09-eb82-4c74-bfa2-a343ef0650de)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.64ns | 45.58ns | 77.36ns | 76.23ns | 43.31ns | 16.75ns |
| 4 | 15.40ns | 13.53ns | 32.56ns | 64.16ns | 13.72ns | 6.52ns |
| 16 | 5.83ns | 5.37ns | 25.48ns | 66.72ns | 6.92ns | 3.96ns |
| 64 | 3.38ns | 3.27ns | 23.05ns | 65.09ns | 5.92ns | 3.31ns |
| 256 | 2.81ns | 2.76ns | 22.83ns | 73.28ns | 5.43ns | 3.18ns |
| 1k | 2.89ns | 2.65ns | 23.03ns | 83.61ns | 5.33ns | 3.47ns |
| 16k | 2.84ns | 2.82ns | 23.67ns | 114.60ns | 5.31ns | 3.65ns |
| 256k | 2.81ns | 2.81ns | 26.80ns | 181.42ns | 5.32ns | 3.15ns |
| 1M | 2.84ns | 2.85ns | 30.35ns | 327.12ns | 5.32ns | 3.19ns |


### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/5c3aa2c4-daaa-4224-b81e-0bed00d68a3f)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.12ns | 47.88ns | 69.45ns | 93.29ns | 46.11ns | 17.50ns |
| 4 | 16.12ns | 13.95ns | 32.65ns | 93.30ns | 14.68ns | 6.71ns |
| 16 | 6.06ns | 5.45ns | 24.97ns | 96.07ns | 7.31ns | 3.99ns |
| 64 | 3.94ns | 3.32ns | 23.07ns | 101.25ns | 5.88ns | 3.43ns |
| 256 | 2.86ns | 2.84ns | 22.74ns | 115.89ns | 5.42ns | 3.29ns |
| 1k | 2.70ns | 2.70ns | 22.63ns | 133.30ns | 5.33ns | 3.13ns |
| 16k | 2.78ns | 2.81ns | 25.44ns | 170.28ns | 5.33ns | 3.16ns |
| 256k | 2.78ns | 2.81ns | 34.30ns | 362.24ns | 5.32ns | 3.16ns |


### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/03871028-eb22-4d10-967a-f8b0a3ed9fe9)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.47ns | 47.26ns | 69.03ns | 66.48ns | 42.77ns | 16.77ns |
| 4 | 23.71ns | 17.61ns | 43.60ns | 63.33ns | 17.79ns | 15.11ns |
| 16 | 17.22ns | 9.58ns | 33.18ns | 65.04ns | 11.90ns | 20.87ns |
| 64 | 9.72ns | 6.15ns | 28.58ns | 65.63ns | 8.13ns | 11.79ns |
| 256 | 4.68ns | 3.72ns | 24.27ns | 71.77ns | 5.82ns | 5.44ns |
| 1k | 3.37ns | 3.07ns | 26.01ns | 81.26ns | 5.44ns | 4.01ns |
| 16k | 2.91ns | 2.81ns | 32.54ns | 116.03ns | 5.36ns | 3.31ns |
| 256k | 2.88ns | 2.77ns | 70.92ns | 187.17ns | 5.32ns | 3.18ns |
| 1M | 2.85ns | 2.81ns | 103.28ns | 289.54ns | 5.34ns | 3.19ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/d4b4cf38-bde0-4974-b46c-e4834d91dfdc)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 211.73ns | 238.80ns | 1.34us | 1.21us | 629.68ns | 1.01us |
| 4 | 82.95ns | 65.36ns | 607.31ns | 584.00ns | 285.79ns | 512.53ns |
| 16 | 49.23ns | 23.61ns | 404.45ns | 402.32ns | 202.18ns | 350.25ns |
| 64 | 39.91ns | 14.53ns | 311.70ns | 359.25ns | 164.75ns | 298.61ns |
| 256 | 34.35ns | 10.52ns | 251.00ns | 297.55ns | 139.65ns | 233.98ns |
| 1k | 29.21ns | 9.15ns | 240.60ns | 298.98ns | 124.99ns | 410.46ns |
| 16k | 26.94ns | 8.21ns | 276.83ns | 327.84ns | 137.38ns | 423.78ns |
| 256k | 26.95ns | 8.23ns | 254.45ns | 399.09ns | 139.77ns | 464.16ns |
| 1M | 26.21ns | 8.21ns | 256.66ns | 526.06ns | 217.99ns | 567.81ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/e0b587d9-b4b5-497f-8618-19c02f310ea9)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.56us | 9.72us | 4.50us | 1.72us | 39.45us | 3.35us |
| 4 | 2.47us | 2.38us | 1.49us | 715.56ns | 10.08us | 1.25us |
| 16 | 648.03ns | 609.61ns | 750.51ns | 489.21ns | 3.02us | 618.30ns |
| 64 | 196.12ns | 165.89ns | 476.33ns | 391.55ns | 938.73ns | 427.48ns |
| 256 | 75.48ns | 50.69ns | 373.75ns | 329.92ns | 485.82ns | 332.12ns |
| 1k | 49.52ns | 28.27ns | 333.80ns | 320.05ns | 290.75ns | 286.75ns |
| 16k | 65.34ns | 40.27ns | 434.04ns | 523.22ns | 280.27ns | 359.47ns |
| 256k | 132.01ns | 35.61ns | 565.27ns | 654.25ns | 710.85ns | 426.86ns |
| 1M | 119.75ns | 24.93ns | 565.45ns | 840.47ns | 1.66us | 422.86ns |


### Create large entities

- Create `N` entities with 10 components `C1`, ..., `C10`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.

![create10comp](https://github.com/user-attachments/assets/d3da6683-5d9c-4095-adab-3f88106d8c31)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 450.77ns | 460.03ns | 1.92us | 2.20us | 507.49ns | 606.26ns |
| 4 | 176.00ns | 129.63ns | 1.30us | 1.63us | 267.83ns | 360.24ns |
| 16 | 112.49ns | 41.18ns | 1.04us | 1.22us | 249.12ns | 287.94ns |
| 64 | 94.46ns | 17.85ns | 827.96ns | 1.06us | 199.48ns | 261.00ns |
| 256 | 85.58ns | 11.78ns | 775.42ns | 1.11us | 149.71ns | 211.16ns |
| 1k | 77.54ns | 9.78ns | 773.61ns | 1.14us | 148.06ns | 469.85ns |
| 16k | 76.81ns | 8.27ns | 804.63ns | 1.22us | 161.39ns | 541.31ns |
| 256k | 76.72ns | 8.18ns | 797.44ns | 2.22us | 166.32ns | 596.78ns |
| 1M | 77.08ns | 8.24ns | 723.98ns | 2.39us | 247.24ns | 655.08ns |


### Add/remove component

`N` entities with component `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/a9b2962c-4f8e-4118-a954-afab5f6fd879)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 205.45ns | 1.14us | 553.14ns | 113.23ns | 559.03ns | 368.22ns |
| 4 | 141.94ns | 288.65ns | 489.26ns | 107.47ns | 522.67ns | 344.40ns |
| 16 | 122.42ns | 77.26ns | 480.61ns | 103.51ns | 528.92ns | 345.23ns |
| 64 | 114.68ns | 23.70ns | 471.38ns | 106.01ns | 535.43ns | 356.07ns |
| 256 | 113.46ns | 10.34ns | 464.74ns | 113.08ns | 553.53ns | 370.85ns |
| 1k | 113.58ns | 7.36ns | 468.73ns | 123.34ns | 552.94ns | 763.85ns |
| 16k | 114.79ns | 7.67ns | 507.93ns | 137.41ns | 585.65ns | 824.90ns |
| 256k | 114.05ns | 8.80ns | 503.28ns | 168.48ns | 689.46ns | 1.01us |
| 1M | 114.16ns | 10.21ns | 501.45ns | 236.76ns | 854.28ns | 1.47us |


### Add/remove component, large entity

`N` entities with component `Position` and 10 further components `C1`, ..., `C10`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove_large](https://github.com/user-attachments/assets/0edcf6a3-ae4d-43a5-a5c3-f338e468e13b)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 565.22ns | 5.78us | 1.44us | 115.42ns | 1.01us | 941.01ns |
| 4 | 512.66ns | 1.52us | 1.19us | 108.39ns | 1.02us | 889.99ns |
| 16 | 489.57ns | 369.07ns | 1.15us | 104.04ns | 1.00us | 858.15ns |
| 64 | 444.74ns | 103.80ns | 1.13us | 107.33ns | 992.16ns | 871.61ns |
| 256 | 450.70ns | 37.35ns | 1.16us | 112.31ns | 986.79ns | 947.89ns |
| 1k | 453.81ns | 19.26ns | 1.21us | 123.40ns | 1.01us | 1.48us |
| 16k | 474.83ns | 24.85ns | 1.30us | 135.73ns | 1.04us | 1.53us |
| 256k | 608.57ns | 44.38ns | 1.45us | 170.68ns | 1.43us | 2.02us |
| 1M | 581.30ns | 45.81ns | 1.28us | 251.01ns | 1.56us | 2.24us |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 10.19us | 6.81us | 71.22us | 508.46us | 6.04us |


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
The benchmarks should take around 20-30 minutes to complete.

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```
