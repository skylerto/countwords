# Count Words

Count words in a file, reporting the top 3 most frequent.

## Usage

Run the executable passing in the desired files.

```
$ countwords file1.txt file2.txt file3.txt
```

Running on [Peter Norvig's big.txt](http://norvig.com/big.txt):

```
➜  countwords git:(master) ✗ time ./bin/countwords test/big.txt
2016/05/27 11:59:36 Opening file: test/big.txt
2016/05/27 11:59:36 Top 3 (or less) common words in test/big.txt are:
the:71744
of:39169
and:35968
./bin/countwords test/big.txt  0.32s user 0.01s system 108% cpu 0.309 total
```

## License

MIT © Skyler Layne, 2016
