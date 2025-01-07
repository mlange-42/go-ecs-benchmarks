import os

import numpy as np
import pandas as pd
from matplotlib import pyplot as plt

results_dir = "results"
plot_dir = "plots"

files = ["query.csv"]


def plot_all():
    for f in files:
        plot(f)

def plot(file: str):
    data = pd.read_csv(os.path.join(results_dir, file))
    fig, ax = plt.subplots(ncols=2, figsize=(10, 4))

    plot_bars(data, ax[0])
    plot_lines(data, ax[1])
    
    fig.tight_layout()

    fig.savefig(os.path.join(plot_dir, f"{file}.svg"))
    fig.savefig(os.path.join(plot_dir, f"{file}.png"))
    plt.close(fig)

def plot_bars(data: pd.DataFrame, ax):
    cols = data.columns[1:]
    width = 1.0/(1.5*len(cols))

    for i, col in enumerate(cols):
        col_data = data[col]
        x = np.arange(len(col_data)) + i * width - 0.375
        ax.bar(x, col_data, width=width, label=col)
    
    ax.set_ylabel("Time per entity")
    ax.set_yscale('log')
    ax.set_yticks([1, 10, 100, 1000])
    ax.set_yticklabels(["1ns", "10ns", "100ns", "1μs"])

    ax.set_xlabel("#Entities")
    ax.set_xticks(range(len(data.index)))
    ax.set_xticklabels(data.N)

    ax.legend(framealpha=0.5)


def plot_lines(data: pd.DataFrame, ax):
    cols = data.columns[1:]
    width = 1.0/(1.5*len(cols))

    for i, col in enumerate(cols):
        col_data = data[col]
        ax.plot(data.N, col_data * data.N, label=col)
    
    ax.set_ylabel("Total time")
    ax.set_xscale('log')
    ax.set_yscale('log')

    ax.set_yticks([1, 1000, 1000_000])
    ax.set_yticklabels(["1ns", "1μs", "1ms"])

    ax.set_xlabel("#Entities")

    ax.set_xticks(data.N)
    ax.set_xticklabels(data.N)

    ax.legend(framealpha=0.5)


if __name__ == "__main__":
    plot_all()
