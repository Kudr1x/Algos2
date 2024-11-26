import numpy as np
import matplotlib.pyplot as plt
import sys

if __name__ == "__main__":
    data = sys.argv[1]
    name = sys.argv[2]

    AllSortedXY = data.split(";")

    arrXInt = []
    arrYInt = []

    for i in range(0, 5):
        arr = AllSortedXY[i].split(",")
        arrXInt.append(int(arr[0]))
        arrYInt.append(int(arr[1]))

    x_data = np.array(arrXInt)
    y_data = np.array(arrYInt)

    x_data = np.clip(x_data, a_min=1e-10, a_max=None)

    log_x_data = np.log(x_data)

    coefficients = np.polyfit(log_x_data, y_data, 1)

    poly = np.poly1d(coefficients)

    x_values = np.linspace(np.min(x_data), np.max(x_data), 100)
    log_x_values = np.log(x_values)
    y_values = poly(log_x_values)

    # Строим график
    plt.scatter(x_data, y_data, color="RED", label=name)
    plt.plot(x_values, y_values, label='Логарифмическая регрессия', color="RED")
    plt.xlabel('N')
    plt.ylabel('µs')
    plt.title(name)
    plt.legend()
    plt.grid(True)

    plt.savefig(f"image/{name}.png")
