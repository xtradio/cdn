# cdn

[![Build Status](https://github.com/xtradio/cdn/workflows/build/badge.svg?branch=master)](https://github.com/xtradio/cdn/actions)

XTRadio CDN built with golang for ease of metrics instrumentation.

## ENV Vars needed

`IMG_FOLDER` - The folder from where we want to serve all the images.

## Endpoints

> `:10000/tracks/` - Folder to serve the files

> `:10001/metrics` - Metrics endpoint

> `:10001/v1/upload` - Endpoint to upload images with POST request
