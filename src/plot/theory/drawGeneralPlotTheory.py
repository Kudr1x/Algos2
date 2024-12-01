import numpy as np
import matplotlib.pyplot as plt

colors = ["RED", "GREEN", "BLUE"]
names = ["AVL дерево", "RB дерево", "BS дерево"]

if __name__ == '__main__':
    f = open("/home/kudrix/GolandProjects/Algos2/raw/theoryData.txt", 'r')

    for i in range(0, 3):
        data = f.readline()

        AllSortedXY = data.split(";")

        arrXInt = []
        arrYInt = []

        for j in range(0, 13):
            arr = AllSortedXY[j].split(",")
            arrXInt.append(int(arr[0]))
            arrYInt.append(int(arr[1]))

        # Строим график
        plt.plot(arrXInt, arrYInt, color=colors[i], label=names[i])
        plt.xlabel('n')
        plt.ylabel('h(n)')
        plt.title("Сводный график")
        plt.legend()
        plt.grid(True)

        plt.savefig(f"image/theory/Сводный график.png")
