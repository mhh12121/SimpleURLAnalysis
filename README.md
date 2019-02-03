## A simple URL Analysis tools

### 1. Description: 
>Filter out given quantity n of slowest response time URL
>( In such format:  ```<URL>, <Response Time>, <Response Status>``` )


### 2. Requirement: 
1. Only considering ```GET``` method and status with ```200``` URL
2. Ignore URL ending with ```".gif"```
3. Result should be decreasing
4. Assuming that Every Request strictly follows the rules:
    - **URL** : ```({GET/POST/PUT/DELETE} [url])//and url should begin with "/"```
    - **Response Time**: ```(minimum 0.001s, no space allowed between number and unit)```
    - **Response Status**: ```(only have one number)```


### 3. Environment
- Ubuntu 16.04
- Go 1.11


### 4. Procedure to run
+ Simply type following in your CLI:
    ```go run main.go```
+ Test with log file ```access.log```
+ Can modify path/filename and required quantity in ```run.conf```


