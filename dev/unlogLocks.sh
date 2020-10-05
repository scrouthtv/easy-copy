#!/bin/bash

sed -E '/^[[:space:]\t]*fmt.Println\("\[[a-z]*.go:[0-9]*\] ((lock)|(unlock))"\);$/d' -i *.go
