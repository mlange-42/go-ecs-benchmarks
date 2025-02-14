import glob
import os
from string import Template

import numpy as np
import pandas as pd
from matplotlib import pyplot as plt
from matplotlib.figure import Figure

results_dir = "results"
template = "docs/README-template.md"

default_colors = plt.rcParams["axes.prop_cycle"].by_key()["color"]

colors = {
    "Arche": default_colors[0],
    "Arche (batch)": default_colors[1],
    "Arche (cached)": default_colors[1],
    "Arche (unchecked)": default_colors[1],
    "Donburi": default_colors[2],
    "Ento": default_colors[3],  # TODO: remove
    "Volt": default_colors[3],
    "ggecs": default_colors[4],
    "uot": default_colors[5],
}


def plot_all():
    template_vars = {
        "info": "",
    }

    with open(os.path.join(results_dir, "info.md"), "r") as f:
        template_vars["info"] = f.read()

    files = glob.glob(os.path.join(results_dir, "*.csv"))
    files = [os.path.split(f)[-1][:-4] for f in files]
    print(f"{len(files)} files to plot: {' '.join(files)}")

    for f in files:
        data = pd.read_csv(os.path.join(results_dir, f"{f}.csv"))
        fig = plot(data)
        fig.savefig(os.path.join(results_dir, f"{f}.svg"))
        fig.savefig(os.path.join(results_dir, f"{f}.png"))
        plt.close(fig)

        md = to_markdown(data)
        template_vars[f] = md
        with open(os.path.join(results_dir, f"{f}.md"), "w") as f:
            f.write(md)

    readme = update_readme(template, template_vars)
    with open(os.path.join(results_dir, "README.md"), "w") as f:
        f.write(readme)


def plot(data: pd.DataFrame) -> Figure:
    multi = len(data.index) > 1
    if multi:
        fig, ax = plt.subplots(ncols=2, figsize=(10, 4))
        plot_bars(data, ax[0], legend=False)
        plot_lines(data, ax[1], legend=True)
    else:
        fig, ax = plt.subplots(ncols=1, figsize=(5, 4))
        plot_bars(data, ax, legend=True)

    fig.tight_layout()
    return fig


def plot_bars(data: pd.DataFrame, ax, legend: bool):
    cols = data.columns[1:]
    width = 1.0 / (1.5 * len(cols))
    max_value = data.max()[1:].max()

    for i, col in enumerate(cols):
        col_data = data[col]
        x = np.arange(len(col_data)) + i * width - 0.375
        ax.bar(x, col_data, width=width, color=colors[col], label=col)

    ax.set_ylabel("Time per entity")
    ax.set_yscale("log")

    if max_value > 400:
        ax.set_yticks([1, 10, 100, 1000])
        ax.set_yticklabels(["1ns", "10ns", "100ns", "1μs"])
    else:
        ax.set_yticks([1, 10, 100])
        ax.set_yticklabels(["1ns", "10ns", "100ns"])

    ax.set_xlabel("#Entities")

    ax.set_xticks(range(len(data.index)))

    labels = [
        str(n) if n < 1000 else f"{n//1000}k" if n < 1000000 else f"{n//1000000}M"
        for n in data.N
    ]
    ax.set_xticklabels(labels)

    if legend:
        ax.legend(framealpha=0.5)


def plot_lines(data: pd.DataFrame, ax, legend: bool):
    cols = data.columns[1:]
    width = 1.0 / (1.5 * len(cols))

    for i, col in enumerate(cols):
        col_data = data[col]
        ax.plot(data.N, col_data * data.N, color=colors[col], label=col)

    ax.set_ylabel("Total time")
    ax.set_xscale("log")
    ax.set_yscale("log")

    ax.set_yticks([1, 1000, 1000_000])
    ax.set_yticklabels(["1ns", "1μs", "1ms"])

    ax.set_xlabel("#Entities")

    ax.set_xticks(data.N)

    labels = [
        str(n) if n < 1000 else f"{n//1000}k" if n < 1000000 else f"{n//1000000}M"
        for n in data.N
    ]
    ax.set_xticklabels(labels)

    if legend:
        ax.legend(framealpha=0.5)


def to_markdown(data: pd.DataFrame) -> str:
    s = ""
    s += "| " + " | ".join(data.columns) + " |\n"
    s += "| " + " | ".join(["---"] * len(data.columns)) + " |\n"

    for i, row in data.iterrows():
        if row.N < 1000:
            n = "%d" % (row.N)
        elif row.N < 1_000_000:
            n = "%dk" % (row.N // 1000)
        else:
            n = "%dM" % (row.N // 1000000)

        vals = [to_time(v) for v in row.iloc[1:]]
        s += "| " + " | ".join([n] + vals) + " |\n"

    return s


def to_time(v: float) -> str:
    if v < 1_000:
        return f"{v:.2f}ns"
    if v < 1_000_000:
        return f"{(v/1_000):.2f}us"
    return f"{(v/1_000_000):.2f}ms"


def update_readme(template_file: str, values: dict) -> str:
    with open(template_file, "r") as file:
        file_content = file.read()

    s = Template(file_content)
    return s.substitute(values)


if __name__ == "__main__":
    plot_all()
