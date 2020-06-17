def cons(a, b):
    def pair(f):
        return f(a, b)
    return pair

def car(pair):
    def car_pair(a, b):
        return a
    return pair(car_pair)

def cdr(pair):
    def cdr_pair(a, b):
        return b
    return pair(cdr_pair)

assert car(cons(3, 4)) == 3
assert cdr(cons(3, 4)) == 4