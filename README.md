# Snakelet

A lightweight, easy-to-use environment variable loader.

## TODO
1. Parse name of fields to traditional `FOO_BAR` environment variables.
    - Will need to think of how to handle custom names
2. Loop through parsed names and get environment variables
    - Use traditional getenv way first
    - Ingest from .env file next
3. Load values gained from .3 into struct passed

