1. Set up a loop device using `test.img`:
```
 # losetup --show --find test.img
/dev/loop0
 # sudo blockdev --getsize /dev/loop0
262144
```
2. Set up a delayed device using 
```
 # sudo sh -c 'echo "0 262144 delay /dev/loop0 0 200" | dmsetup create dm-slow'
```
Parameters are:
 - start sector
 - end sector (gotten from blockdev)
 - type of mapper (`delay`: delayed device)
 - real device for `delay` to use 
 - offset in the real device for `delay` to use
 - ms to `delay` operations on the real device

3. Mount it:
```
 # mount /dev/mapper/dm-slow mnt/
```

ToDo: io cache has to be disabled or else it'll fast again
