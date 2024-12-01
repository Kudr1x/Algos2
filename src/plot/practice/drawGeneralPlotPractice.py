import numpy as np
import matplotlib.pyplot as plt

colors = ["RED", "GREEN", "BLUE"]
names = ["AVL дерево", "RB дерево", "BS дерево"]

if __name__ == '__main__':
    f = open("/home/kudrix/GolandProjects/Algos2/raw/practiceData.txt", 'r')

    for i in range(0, 3):
        data = f.readline()

        AllSortedXY = data.split(";")

        arrXInt = []
        arrYInt = []

        for j in range(0, 13):
            arr = AllSortedXY[j].split(",")
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
        plt.scatter(x_data, y_data, color=colors[i], label=names[i])
        plt.plot(x_values, y_values, color=colors[i])
        plt.xlabel('n')
        plt.ylabel('h(n)')
        plt.title("Сводный график")
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/practice/Сводный график.png")
