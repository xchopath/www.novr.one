# Install PHP 8.2, composer, and Voyager with Laravel 9

Install PHP 8.2
```
sudo apt install php8.2 php8.2-dev php8.2-mbstring php8.2-gd -y
```

Install Composer
```
php -r "copy('https://getcomposer.org/installer', 'composer-setup.php');"
sudo php composer-setup.php --install-dir=/usr/local/bin --filename=composer
```

Create Voyager Project with Laravel 9.5.2
```
composer create-project laravel/laravel=9.5.2 voyagerproject
cd voyagerproject/
```

Create a Database
```
mysql -u root -p --protocol=tcp
mysql> CREATE DATABASE myapp;
```

Adjust the `.env`
```
APP_URL=http://localhost:8000
DB_CONNECTION=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE=myapp
DB_USERNAME=username
DB_PASSWORD=password
```

Add Voyager dependencies to the Project
```
composer require tcg/voyager
```

Install Voyager
```
php artisan voyager:install
```

Setup Voyager Admin
```
php artisan voyager:admin admin@email.com --create
```

Start the App
```
php artisan serve --host=0.0.0.0 --port=8000 
```
