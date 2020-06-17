import random
import math

def main():
    inside = 0.0000001
    count = 0.0000001
    while True: 
        a = random.uniform(0, 1)
        b = random.uniform(0, 1)
        distance_center = (a**2 + b**2)**0.5
        if distance_center<1: 
            inside +=1
        count +=1
        current_pi = (inside/count)/0.5**2
        if math.isclose(current_pi, 3.14159):
            print(current_pi)
            break
        


if __name__ == "__main__":
    main()