#!/bin/bash
## Created by Syldie Aldi Wijaya         ##
##                                       ##
########################################### 

set -e

Package=(
		"github.com/astaxie/beego"
		"github.com/beego/bee"
		"github.com/lib/pq"
		"github.com/mattn/go-sqlite3"		
		"golang.org/x/crypto/bcrypt"
		"gopkg.in/gomail.v2"		
		"github.com/dgrijalva/jwt-go"		
		"github.com/satori/go.uuid"
		"github.com/smartystreets/goconvey/convey"	
)

arrayPackage=${#Package[@]}

for ((k=0; k<${arrayPackage}; k++));
do
	echo "GO GET ${Package[$k]}"
	go get ${Package[$k]} 2>&1
done
