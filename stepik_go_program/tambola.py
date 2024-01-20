from typing import Any


class TestSequence(list):
    __match_args__ = ("a", "b", "c")

    def __getattr__(self, __name: str) -> Any:
        if __name in self.__match_args__:
            indx = self.__match_args__.index(__name)
            # return self.__getitem__(indx)
            return self[indx]
        else:
            raise AttributeError(f"No such attribute: {__name}")

    def __setattr__(self, __name: str, __value: Any) -> None:
        if __name in self.__match_args__:
            raise AttributeError(f"You can't change: {__name}")
        else:
            return super().__setattr__(__name, __value)


tlist = TestSequence([1, 2, 3, 4, 5])
tlist.d = 2
print(tlist.a)
