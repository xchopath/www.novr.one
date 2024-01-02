#### Download Source

```
wget https://ftp.postgresql.org/pub/source/v12.8/postgresql-12.8.tar.bz2
```

#### Extract

```
tar -xvjf postgresql-12.8.tar.bz2 
```

#### Go to

```
cd postgresql-12.8/
```

#### Install Dependencies

```
sudo apt-get install build-essential -y
sudo apt-get install libreadline-dev -y
sudo apt-get install zlib1g-dev -y
```

#### Configure

```
./configure --prefix=~/postgresql-12
```

#### Make & Make Install

```
make -j4 && make install
```
