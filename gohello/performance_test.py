import time
import random
sum =0
r =random.randint(10,13)
print(r)
t = time.time()
for i in range(40000+r):
    for j in range(40000):
        sum += i*j

print(sum)
print(time.time()-t)
