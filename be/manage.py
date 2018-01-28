#!/usr/bin/env python
import click

from application import create_application


@click.group()
def cli():
    pass

@cli.command()
def run_app():
    app = create_application()
    if app.config.debug == True:
        app.run(host='0.0.0.0')
    else:
        app.run()


if __name__ == '__main__':
    cli()
