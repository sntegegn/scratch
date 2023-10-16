import redis

r = redis.Redis(host='localhost', port=6379, db=0)

wakitaki = r.pubsub()

wakitaki.subscribe("army-camp-1")

for message in wakitaki.listen():
    print(message)