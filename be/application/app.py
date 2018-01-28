import os

from flask import Flask
from flask_restful import Api

from .routes import Index


def create_app():
    app = Flask(__name__)
    config_app(app)
    return app


def config_app(app):
    config_path = os.environ.get('CONFIG_PATH', 'application.config.develop')
    app.config.from_object(config_path)
    api = Api(app)
    api.add_resource(Index, '/')