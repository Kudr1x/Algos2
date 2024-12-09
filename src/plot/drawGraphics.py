import numpy as np
import matplotlib.pyplot as plt
import sys


arr_x_int, arr_y_int = [], []


def read_data(file, num_entries):
    data = file.readline()
    all_sorted_xy = data.split(";")

    arr_x_int.clear()
    arr_y_int.clear()

    for i in range(num_entries):
        arr = all_sorted_xy[i].split(",")
        arr_x_int.append(int(arr[0]))
        arr_y_int.append(int(arr[1]))

    return arr_x_int, arr_y_int


if __name__ == "__main__":
    name = sys.argv[1]

    with open(f"/home/kudrix/GolandProjects/Algos2/raw/{name}Data.txt", 'r') as f:
        arr_x_int, arr_y_int = read_data(f, 13)

        x_data = np.array(arr_x_int)
        y_data = np.array(arr_y_int)

        x_data = np.clip(x_data, a_min=1e-10, a_max=None)

        log_x_data = np.log(x_data)

        coefficients = np.polyfit(log_x_data, y_data, 1)
        print(f"{coefficients[0]} {coefficients[1]}")
        poly = np.poly1d(coefficients)

        x_values = np.linspace(np.min(x_data), np.max(x_data), 100)
        log_x_values = np.log(x_values)
        y_values = poly(log_x_values)

        plt.scatter(x_data, y_data, color="RED", label="Экспериментальные значения")
        plt.plot(x_values, y_values, label='Логарифмическая регрессия', color="RED")

        for color, label in [("GREEN", "Теоритическая оценка сверху"), ("BLUE", "Теоритическая оценка снизу")]:
            arr_x_int, arr_y_int = read_data(f, 13)
            plt.plot(arr_x_int, arr_y_int, label=label, color=color)

        plt.xlabel('n')
        plt.ylabel('h(n)')
        plt.title(name)
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/{name}.png")
