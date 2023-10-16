# https://dev.to/rohit20001221/simple-pub-sub-system-using-redis-and-python-25jh
import redis

r = redis.Redis(host='localhost', port=6379, db=0)

while True:
    message = input("Enter the messages you want to send to soliders: ")
    r.publish("army-camp-1", message)

