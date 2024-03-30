FROM php:8.2-fpm
# node js
RUN curl -sL https://deb.nodesource.com/setup_20.x -o nodesource_setup.sh
RUN ["sh",  "./nodesource_setup.sh"]

# Install environment dependencies
RUN apt-get update \
	# gd
	&& apt-get install -y build-essential supervisor openssl nginx libfreetype6-dev libjpeg-dev libicu-dev  libpng-dev libwebp-dev zlib1g-dev libzip-dev gcc g++ make vim unzip curl git jpegoptim optipng pngquant gifsicle locales libonig-dev nodejs redis-server librdkafka-dev libsnappy-dev \
	&& docker-php-ext-configure gd  \
	&& docker-php-ext-install gd \
	# gmp
	&& apt-get install -y --no-install-recommends libgmp-dev \
	&& docker-php-ext-install gmp \
	# pdo_mysql
	&& docker-php-ext-install pdo_mysql mbstring \
	# pdo
	&& docker-php-ext-install pdo \
	# opcache
	&& docker-php-ext-enable opcache \
	# exif
    && docker-php-ext-install exif \
    && docker-php-ext-install sockets \
    && docker-php-ext-install pcntl \
    && docker-php-ext-install bcmath \
    && docker-php-ext-install intl \
	# zip
	&& docker-php-ext-install zip \
	&& apt-get autoclean -y \
	&& rm -rf /var/lib/apt/lists/* \
	&& rm -rf /tmp/pear/

# setup composer
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

## redis
RUN \
  curl 'http://pecl.php.net/get/redis-6.0.2.tgz' -o /tmp/redis-6.0.2.tgz  \
  && cd /tmp \
  && pecl install redis-6.0.2.tgz \
  && rm -rf redis-6.0.2.tgz \
  && docker-php-ext-enable redis

## Kafka
RUN pecl install rdkafka \
    && docker-php-ext-enable rdkafka

# copy files
COPY ./php/local.ini /usr/local/etc/php/local.ini
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf
COPY ./redis/redis.conf /etc/redis/redis.conf.default
COPY ./supervisord.conf /etc/supervisor/conf.d/supervisord.conf

# Start supervisord
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/conf.d/supervisord.conf"]

WORKDIR /app
