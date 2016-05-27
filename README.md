# Count Words

Count words in a file, reporting the top 3 most frequent. This was done as a
tribute to the speed of Go.

## Usage

Run the executable passing in the desired files.

```
$ ./bin/countwords file1.txt file2.txt file3.txt
```

### Running on [Peter Norvig's big.txt](http://norvig.com/big.txt):

```
➜  countwords git:(master) ✗ time ./bin/countwords test/big.txt
2016/05/27 11:59:36 Opening file: test/big.txt
2016/05/27 11:59:36 Top 3 (or less) common words in test/big.txt are:
the:71744
of:39169
and:35968
./bin/countwords test/big.txt  0.32s user 0.01s system 108% cpu 0.309 total
```


### Specs on Big.txt

A few specs on Peter Norvig's big.txt file included in the test directory. 

#### 1095695 words in 128457 lines of the 6.2M file.

```
➜  test git:(master) ✗ du -sh big.txt
6.2M  big.txt
```


```
➜  countwords git:(master) ✗ wc -l test/big.txt
128457 test/big.txt
```

```
➜  countwords git:(master) ✗ wc -w test/big.txt
1095695 test/big.txt
```

## License

MIT © Skyler Layne, 2016
