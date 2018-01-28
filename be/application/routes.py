import RPi.GPIO as GPIO
from flask_restful import Resource

from .dht11 import DHT11


class Index(Resource):
    def get(self):
        GPIO.setwarnings(False)
        GPIO.setmode(GPIO.BCM)
        GPIO.cleanup()
        instance = DHT11(pin = 4)
        result = instance.read()
        return {'temperature': result.temperature, 'humidity': result.humidity}