#!/usr/bin/env python
import click

from application import create_application


@click.group()
def cli():
    pass

@cli.command()
def run_app():
    app = create_application()
    app.run()


if __name__ == '__main__':
    cli()
