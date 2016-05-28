As stated in the [README.md](./README.md) this is a tribute to speed. I wrote a
similar word counting program in an operating systems class, the point was to
study multi-threaded programs. This code can be found at
[skylerto/count-words](https://github.com/skylerto/count-words).


## test/file3.txt


### File Specs

About 1000 lines, and 9500 words, file size of 72K.

```
➜  countwords git:(master) ✗ wc -lw test/file3.txt
 1067  9520 test/file3.txt
➜  countwords git:(master) ✗ du -sh test/file3.txt
72K  test/file3.txt  
```

### Ranking

In C land, we're using 98% of our CPU to finish in 0.6s.
```
➜  count-words git:(master) ✗ time ./thread ~/work/src/countwords/test/file3.txt
/Users/sky/work/src/countwords/test/file3.txt: 817 bag art man
./thread ~/work/src/countwords/test/file3.txt  0.59s user 0.01s system 98% cpu
0.605 total
```

Over in the world of Go, we're only using 59% of our CPU to finish in 0.018s.
```
➜  countwords git:(master) ✗ time ./bin/countwords test/file3.txt
2016/05/28 12:11:07 Opening file: test/file3.txt
2016/05/28 12:11:07 Top 3 (or less) common words in test/file3.txt are:
bag:81
art:78
man:73
./bin/countwords test/file3.txt  0.01s user 0.00s system 59% cpu 0.018 total
```

## test/big.txt

### File Specs

128457 lines, 1095695 words, and 6.2M files size.
```
➜  countwords git:(master) ✗ wc -lw test/big.txt
128457 1095695 test/big.txt
➜  countwords git:(master) ✗ du -sh test/big.txt
6.2M test/big.txt
```

Some points to consider below are not just only the elapsed times, but the
amount of CPU usage. It's clear that Go is not afraid to use the CPU (108%)
where as the C equivalent is using only 30% of the CPU. The optimizations done
here by Go are amazing.

**The C threaded version**
```
➜  count-words git:(master) ✗ time ./thread ~/work/src/countwords/test/big.txt
/Users/sky/work/src/countwords/test/big.txt: 81397 the of and
./thread ~/work/src/countwords/test/big.txt  12676.58s user 104.41s system 30%
cpu 11:29:44.89 total
```

**The Go threaded version**
```
➜  countwords git:(master) ✗ time ./bin/countwords test/big.txt
2016/05/27 11:59:36 Opening file: test/big.txt
2016/05/27 11:59:36 Top 3 (or less) common words in test/big.txt are:
the:71744
of:39169
and:35968
./bin/countwords test/big.txt  0.32s user 0.01s system 108% cpu 0.309 total
```

## Conclusions

In a short amount of time we have seen that for the code presented Go's time to
CPU ratio far out ranks the equivalent C example.

For these reasons, go seems like it could be a viable option for systems
programming. Especially in nix land where everything is a file.
