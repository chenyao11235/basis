#


class Array:
    def __init__(self, capacity: int):
        self._data = []
        self._capacity = capacity

    def __getitem__(self, position: int) -> int:
        return self._data[position]

    def __setitem__(self, index:int, value: int):
        self._data[index] = value

    def __len__(self)->int:
        return len(self._data)

    def __iter__(self):
        for item in self._data:
            yield item

    def find(self, index: int)->object:
        try:
            return self._data[index]
        except IndexError:
            return None

    def delete(self, index: int)->bool:
        try:
            self._data.pop(index)
            return True
        except IndexError:
            return False

    def insert(self, index: int, value: int)->bool:
        if len(self) >= self._capacity:
            return False
        else:
            return self._data.insert(index, value)

    def print(self):
        for item in self:
            print(item)


def test_array():
    arr = Array(5)
    arr.insert(0, 0)
    arr.insert(1, 1)
    arr.insert(2, 2)
    arr.insert(3, 3)
    arr.delete(1)
    print(arr.find(1))
    arr.insert(10, 10)
    print(arr.find(10))
    print(arr.find(7))
    arr.print()

if __name__ == "__main__":
    test_array()