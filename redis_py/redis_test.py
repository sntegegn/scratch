# run redis-server
import redis

pool = redis.ConnectionPool(host='localhost', port=6379, db=0)
redis = redis.Redis(connection_pool=pool)

redis.set('mykey', 'Hello from Python!')
value = redis.get('mykey')
print(value)

redis.zadd('vehicles', {'car' : 0})
redis.zadd('vehicles', {'bike' : 0})
vehicles = redis.zrange('vehicles', 0, -1)
print(vehicles)

redis.incr("delivered")
redis.incr("delivered", 5)
val = redis.get("delivered")
print(val)

print("_____")
pipe = redis.pipeline()
pipe.set('foo', 5)
pipe.set('bar', 18.5)
pipe.set('blee', "hello-world")
pipe.execute()

foo_val = redis.get('foo')
bar_val = redis.get('bar')
blee_val = redis.get('blee')
print(foo_val, bar_val, blee_val)