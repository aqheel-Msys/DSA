class Heap:
    def __init__(self):
        self._max_size = 10
        self._data = [-1] * self._max_size
        self._cur_size = 0

    def __len__(self):
        return len(self._data)

    def is_empty(self):
        return len(self._data) == 0

    def insert(self, e):
        if self._cur_size == self._max_size:
            print('No Space')
            return

        self._cur_size += 1
        hi = self._cur_size

        while hi > 1 and e > self._data[hi // 2]:
            self._data[hi] = self._data[hi // 2]
            hi //= 2

        self._data[hi] = e

    def max(self):
        if self._cur_size == 0:
            print('Heap is Empty')
            return
        return self._data[1]

if __name__ == '__main__':
    z = Heap()

    z.insert(23)
    z.insert(1)
    z.insert(21)
    z.insert(35)
    z.insert(20)
    z.insert(31)
    print(z.max())

    print(z._data, len(z))