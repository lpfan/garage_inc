import Adafruit_DHT
from flask_restful import Resource


sensor = Adafruit_DHT.DHT11


class Index(Resource):
    def get(self):
        humidity, temperature = Adafruit_DHT.read_retry(sensor, 4)
        return {'temperature': temperature, 'humidity': humidity}