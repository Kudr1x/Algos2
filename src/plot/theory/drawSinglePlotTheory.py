import matplotlib.pyplot as plt
import sys

if __name__ == "__main__":
    data = sys.argv[1]
    name = sys.argv[2]

    AllSortedXY = data.split(";")

    arrXInt = []
    arrYInt = []

    for i in range(0, 13):
        arr = AllSortedXY[i].split(",")
        arrXInt.append(int(arr[0]))
        arrYInt.append(int(arr[1]))

    plt.plot(arrXInt, arrYInt, label=name, color="RED")
    plt.xlabel('n')
    plt.ylabel('h(n)')
    plt.title(name)
    plt.legend()
    plt.grid(True)

    plt.savefig(f"image/theory/{name}.png")