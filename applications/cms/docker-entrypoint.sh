#!/bin/bash

set -e

if [ ! -d "vendor" ]; then
    composer install --no-interaction --optimize-autoloader
fi

if [ ! -f ".env" ]; then
    cp .env.example .env
    php artisan key:generate
fi

php artisan optimize
php artisan migrate --force

exec "$@"
