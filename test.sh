#!/bin/bash  
  
sum=0  
  
i=1  
  
  
for(( i = 1; i <= 100; i = i + 2 ))  
do  
  sleep 1
#  curl -X POST http://127.0.0.1:1234/v1/socket/notice -d '{"module":"uflow","type":"pipeline","content":{"aa":"aa"}}'
  curl -X POST http://127.0.0.1:1234/v1/socket/notice -d "$i"
done
