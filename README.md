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

<p align="center">
<a title="Star History Chart" href="https://star-history.com/#mlange-42/arche&yohamta/donburi&wwfranczyk/ento&marioolofo/go-gameengine-ecs&unitoftime/ecs&Date">
<img src="https://api.star-history.com/svg?repos=mlange-42/arche,yohamta/donburi,wwfranczyk/ento,marioolofo/go-gameengine-ecs,unitoftime/ecs&type=Date" alt="Star History Chart" width="600"/>
</a>
</p>

## Benchmarks

Last run: Wed, 08 Jan 2025 19:55:09 UTC  
CPU: AMD EPYC 7763 64-Core Processor


For each benchmark, the left plot panel and the table show the time spent per entity,
while the right panel shows the total time.

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/f0ba6cea-e904-4f88-a783-a0c8cf2d5a62)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.81ns | 45.98ns | 72.80ns | 70.46ns | 43.19ns | 16.85ns |
| 4 | 15.48ns | 13.30ns | 32.63ns | 66.70ns | 13.78ns | 6.51ns |
| 16 | 6.10ns | 5.48ns | 25.46ns | 64.85ns | 6.91ns | 3.96ns |
| 64 | 3.45ns | 3.32ns | 23.58ns | 67.94ns | 5.75ns | 3.42ns |
| 256 | 2.79ns | 3.19ns | 23.20ns | 76.44ns | 5.45ns | 3.18ns |
| 1k | 2.89ns | 3.22ns | 22.66ns | 83.55ns | 5.31ns | 3.15ns |
| 16k | 2.82ns | 3.11ns | 23.48ns | 115.84ns | 5.31ns | 3.15ns |
| 256k | 2.80ns | 2.79ns | 24.88ns | 135.51ns | 5.60ns | 3.13ns |
| 1M | 2.83ns | 2.85ns | 28.96ns | 306.03ns | 5.37ns | 3.15ns |


### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/cc0af017-13ac-4278-8cb4-6ca8cf240e46)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 52.05ns | 49.00ns | 69.31ns | 94.61ns | 46.14ns | 17.39ns |
| 4 | 18.20ns | 13.90ns | 32.80ns | 95.06ns | 14.67ns | 6.68ns |
| 16 | 6.68ns | 5.42ns | 24.97ns | 96.56ns | 7.23ns | 4.02ns |
| 64 | 3.63ns | 3.30ns | 23.19ns | 103.55ns | 5.86ns | 3.42ns |
| 256 | 2.84ns | 2.81ns | 22.90ns | 114.60ns | 5.45ns | 3.52ns |
| 1k | 2.70ns | 2.74ns | 22.95ns | 130.39ns | 5.34ns | 3.13ns |
| 16k | 3.00ns | 2.80ns | 25.01ns | 165.94ns | 5.32ns | 3.15ns |
| 256k | 2.78ns | 2.78ns | 24.88ns | 338.08ns | 5.30ns | 3.16ns |


### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.

- Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/863a3258-3a30-4b43-8866-e652f4f8310f)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.23ns | 48.58ns | 68.60ns | 68.10ns | 42.83ns | 16.74ns |
| 4 | 23.53ns | 17.23ns | 39.65ns | 66.26ns | 17.77ns | 15.13ns |
| 16 | 16.33ns | 9.93ns | 33.66ns | 66.09ns | 11.95ns | 21.40ns |
| 64 | 9.64ns | 6.12ns | 28.58ns | 67.80ns | 8.20ns | 11.92ns |
| 256 | 4.67ns | 3.83ns | 24.09ns | 75.55ns | 5.90ns | 5.38ns |
| 1k | 3.30ns | 3.07ns | 24.96ns | 83.82ns | 5.64ns | 3.84ns |
| 16k | 2.88ns | 2.81ns | 30.91ns | 121.81ns | 5.34ns | 3.33ns |
| 256k | 2.81ns | 2.74ns | 62.96ns | 165.39ns | 5.36ns | 3.19ns |
| 1M | 2.84ns | 2.76ns | 97.91ns | 279.34ns | 5.40ns | 3.18ns |


### Create entities

- Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

![create2comp](https://github.com/user-attachments/assets/00ae0147-821b-4158-ad61-b9d2382df638)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 211.75ns | 254.98ns | 1.22us | 1.24us | 711.25ns | 1.10us |
| 4 | 90.94ns | 78.49ns | 571.99ns | 543.86ns | 301.77ns | 532.53ns |
| 16 | 55.08ns | 24.16ns | 392.88ns | 406.71ns | 204.06ns | 332.38ns |
| 64 | 44.21ns | 13.82ns | 318.97ns | 348.65ns | 155.07ns | 267.01ns |
| 256 | 36.48ns | 10.33ns | 265.25ns | 289.28ns | 137.97ns | 237.36ns |
| 1k | 31.68ns | 9.40ns | 249.27ns | 292.91ns | 124.30ns | 426.89ns |
| 16k | 26.39ns | 8.18ns | 280.25ns | 325.10ns | 137.66ns | 429.09ns |
| 256k | 26.42ns | 8.16ns | 266.25ns | 389.67ns | 136.70ns | 493.92ns |
| 1M | 26.35ns | 8.18ns | 269.46ns | 495.70ns | 207.40ns | 542.83ns |


### Create entities, allocating

- Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp_alloc](https://github.com/user-attachments/assets/d8c7ab00-a915-4ef4-935a-79c11cd409e0)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.15us | 9.31us | 4.66us | 1.70us | 41.80us | 3.50us |
| 4 | 2.40us | 2.39us | 1.50us | 699.78ns | 11.14us | 1.28us |
| 16 | 646.54ns | 617.39ns | 782.83ns | 480.36ns | 3.25us | 595.57ns |
| 64 | 194.89ns | 165.71ns | 487.92ns | 412.32ns | 996.69ns | 454.86ns |
| 256 | 81.97ns | 52.32ns | 389.46ns | 341.10ns | 528.68ns | 329.95ns |
| 1k | 50.36ns | 27.75ns | 345.76ns | 327.45ns | 296.70ns | 301.59ns |
| 16k | 71.89ns | 40.99ns | 447.17ns | 518.62ns | 280.57ns | 330.04ns |
| 256k | 127.19ns | 35.17ns | 614.87ns | 667.53ns | 732.83ns | 403.89ns |
| 1M | 115.92ns | 25.30ns | 516.19ns | 842.45ns | 1.65us | 433.72ns |


### Add/remove component

`N` entities with components `Position`.

- Query all `[Position]` entities and add `Velocity`.
- Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/c4f99d59-ecbd-475a-9f56-8b6e1e977820)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 225.60ns | 1.13us | 547.49ns | 113.46ns | 554.93ns | 364.09ns |
| 4 | 152.49ns | 302.10ns | 491.27ns | 109.98ns | 512.60ns | 344.89ns |
| 16 | 124.50ns | 101.46ns | 503.52ns | 110.51ns | 521.14ns | 356.94ns |
| 64 | 122.68ns | 47.72ns | 480.40ns | 109.14ns | 540.19ns | 362.88ns |
| 256 | 113.99ns | 32.09ns | 485.77ns | 116.05ns | 558.33ns | 378.97ns |
| 1k | 115.47ns | 29.31ns | 473.52ns | 125.73ns | 571.51ns | 795.20ns |
| 16k | 115.08ns | 29.04ns | 521.38ns | 137.66ns | 577.74ns | 811.14ns |
| 256k | 114.54ns | 29.69ns | 506.37ns | 184.80ns | 650.92ns | 1.11us |
| 1M | 114.03ns | 30.26ns | 498.45ns | 228.50ns | 885.73ns | 1.42us |


### Create world

- Create a new world
- Register two components (if applicable)

| N | Arche | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- |
| 1 | 9.53us | 6.34us | 65.35us | 479.00us | 5.13us |


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
