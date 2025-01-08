# Go ECS Benchmarks

## Benchmark candidates

| ECS | Type | Version |
|-----|------|---------|
| [Arche](https://github.com/mlange-42/arche) | Archetype | v0.14.5 |
| [Donburi](https://github.com/yohamta/donburi) | Archetype | v1.15.6 |
| [ento](https://github.com/wwfranczyk/ento) | Sparse Set | v0.1.0 |
| [go-gameengine-ecs](https://github.com/marioolofo/go-gameengine-ecs) | Archetype | v0.9.0 |
| [unitoftime/ecs](https://github.com/unitoftime/ecs) | Archetype | v0.0.3 |

## Benchmarks

Last run: Tue, 07 Jan 2025 21:12:56 UTC  
CPU: AMD EPYC 7763 64-Core Processor

For each benchmark, the left plot panel and the table show time spent per entity,
while the right panel shows total time.

All components used in the benchmarks have two `float64` fields.

### Query

`N` entities with components `Position` and `Velocity`.
Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query2comp](https://github.com/user-attachments/assets/5f0ac10d-8f39-4d5c-9915-0dcae22a442c)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 51.00ns | 49.24ns | 68.65ns | 69.40ns | 44.17ns | 17.77ns |
| 4 | 15.09ns | 13.98ns | 32.57ns | 67.38ns | 13.81ns | 6.70ns |
| 16 | 5.79ns | 5.58ns | 25.07ns | 66.97ns | 6.86ns | 4.00ns |
| 64 | 4.45ns | 4.47ns | 23.07ns | 67.52ns | 5.53ns | 3.44ns |
| 256 | 3.03ns | 3.12ns | 22.93ns | 76.68ns | 5.19ns | 3.19ns |
| 1k | 2.68ns | 2.76ns | 22.59ns | 89.16ns | 5.02ns | 3.21ns |
| 16k | 2.72ns | 2.82ns | 23.52ns | 116.76ns | 4.99ns | 3.14ns |
| 256k | 2.72ns | 2.81ns | 26.33ns | 271.02ns | 5.02ns | 3.15ns |
| 1M | 2.76ns | 2.87ns | 29.45ns | 305.80ns | 5.03ns | 3.19ns |

### Query sparse

`N` entities with components `Position` and `Velocity`.
Additionally, there are `9*N` entities with only `Position`.
Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query1in10](https://github.com/user-attachments/assets/3a3510b3-e542-4fb0-b10a-5fdb4849eb10)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.53ns | 47.70ns | 69.29ns | 98.21ns | 46.11ns | 18.23ns |
| 4 | 16.24ns | 13.45ns | 32.68ns | 97.89ns | 14.74ns | 6.75ns |
| 16 | 6.25ns | 5.36ns | 25.07ns | 98.47ns | 7.11ns | 4.10ns |
| 64 | 3.52ns | 3.30ns | 23.69ns | 104.69ns | 5.57ns | 3.42ns |
| 256 | 2.85ns | 3.13ns | 22.96ns | 114.23ns | 5.11ns | 3.19ns |
| 1k | 2.70ns | 2.80ns | 22.72ns | 133.36ns | 5.02ns | 3.13ns |
| 16k | 2.76ns | 2.83ns | 24.74ns | 208.03ns | 5.00ns | 3.14ns |
| 256k | 2.78ns | 2.81ns | 37.11ns | 372.31ns | 5.04ns | 3.16ns |

### Query fragmented

`N` entities with components `Position` and `Velocity`.
Each of these `N` entities has some combination of components
`C1`, `C2`, ..., `C5`, so entities are fragmented over 32 archetypes.
Query all `[Position, Velocity]` entities, and add the velocity vector to the position vector.

![query32arch](https://github.com/user-attachments/assets/b1f915c9-8bd1-40be-8902-077323f18566)

| N | Arche | Arche (cached) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 53.24ns | 47.22ns | 68.69ns | 68.51ns | 43.24ns | 17.44ns |
| 4 | 22.93ns | 17.26ns | 38.48ns | 67.83ns | 17.92ns | 15.41ns |
| 16 | 16.27ns | 9.94ns | 33.11ns | 65.71ns | 11.96ns | 20.59ns |
| 64 | 9.64ns | 6.33ns | 28.89ns | 69.65ns | 8.20ns | 11.71ns |
| 256 | 4.68ns | 3.71ns | 24.60ns | 75.38ns | 5.71ns | 5.32ns |
| 1k | 3.31ns | 3.04ns | 24.78ns | 88.44ns | 5.17ns | 4.48ns |
| 16k | 2.97ns | 2.78ns | 31.30ns | 117.16ns | 5.06ns | 3.28ns |
| 256k | 2.82ns | 2.74ns | 86.71ns | 266.69ns | 5.03ns | 3.18ns |
| 1M | 2.87ns | 2.83ns | 104.67ns | 327.74ns | 5.02ns | 3.21ns |

### Create entities

Create `N` entities with components `Position` and `Velocity`.

The operation is performed once before benchmarking,
to exclude things like archetype creation and memory allocation.
See the benchmark below for entity creation with allocation.

### Create entities, allocating

Create `N` entities with components `Position` and `Velocity`.

Each round is performed on a fresh world.
Thus, low `N` values might be biased by things like archetype creation and memory allocation,
which may be handled differently by different implementations.
See the benchmark above for entity creation without allocation.

![create2comp](https://github.com/user-attachments/assets/f1b613a1-4609-4318-a7a0-98e511a3a881)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 9.57us | 9.65us | 4.70us | 1.62us | 35.25us | 3.27us |
| 4 | 2.38us | 2.34us | 1.51us | 720.00ns | 9.16us | 1.30us |
| 16 | 637.18ns | 598.26ns | 730.74ns | 473.11ns | 2.81us | 627.07ns |
| 64 | 190.94ns | 156.55ns | 469.59ns | 377.93ns | 1.06us | 431.40ns |
| 256 | 74.83ns | 48.26ns | 366.96ns | 330.32ns | 500.80ns | 346.81ns |
| 1k | 50.13ns | 26.19ns | 326.85ns | 304.42ns | 285.36ns | 311.92ns |
| 16k | 96.72ns | 34.01ns | 424.25ns | 526.58ns | 272.30ns | 340.31ns |
| 256k | 177.01ns | 33.33ns | 529.64ns | 670.21ns | 741.76ns | 437.81ns |
| 1M | 256.67ns | 23.51ns | 507.79ns | 751.97ns | 1.66us | 475.18ns |

### Add/remove component

`N` entities with components `Position`.
Query all `[Position]` entities and add `Velocity`.
Query all `[Position, Velocity]` entities and remove `Velocity`.

One iteration is performed before the benchmarking starts, to exclude memory allocation.

![add_remove](https://github.com/user-attachments/assets/4443b11b-ecb7-402c-825e-963601f19db9)

| N | Arche | Arche (batch) | Donburi | Ento | ggecs | uot |
| --- | --- | --- | --- | --- | --- | --- |
| 1 | 203.17ns | 1.16us | 565.42ns | 114.87ns | 556.74ns | 388.34ns |
| 4 | 141.46ns | 310.87ns | 496.56ns | 109.21ns | 519.29ns | 350.26ns |
| 16 | 122.23ns | 99.92ns | 480.65ns | 105.06ns | 553.42ns | 347.15ns |
| 64 | 115.92ns | 45.46ns | 472.42ns | 105.76ns | 541.97ns | 367.92ns |
| 256 | 115.05ns | 31.78ns | 465.99ns | 114.96ns | 551.39ns | 374.16ns |
| 1k | 112.36ns | 28.58ns | 477.82ns | 122.56ns | 552.86ns | 760.86ns |
| 16k | 114.70ns | 28.60ns | 522.67ns | 139.11ns | 580.68ns | 817.44ns |
| 256k | 115.21ns | 29.41ns | 505.87ns | 195.28ns | 717.39ns | 1.31us |
| 1M | 114.88ns | 30.65ns | 524.19ns | 246.51ns | 903.32ns | 1.48us |

## Running the benchmarks

Run the benchmarks using the following command:

```shell
go run . -test.benchtime=0.25s
```

> On PowerShell use this instead:  
> `go run . --% -test.benchtime=0.25s`

The `benchtime` limit is required for some of the benchmarks that have a high
overhead that is not measured. They would take forever otherwise.

To create the plots, run `plot/plot.py`. The following packages are required:
- numpy
- pandas
- matplotlib

```
pip install -r ./plot/requirements.txt
python plot/plot.py
```
