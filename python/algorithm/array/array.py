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

    def insert(self, index: int, value: index)->bool:
        if len(self) >= self._capacity:
            return False
        else:
            return self._data.insert(index, value)

    def print(self):
        for item in self:
            print(item)
