## Installation and Run

```
$ cd aksservice
$ make build
$ make run
$ make test // to test source code
```

### Api URL examples
* AksConfig:
    * POST localhost:1323/aksconfig 
    * Content-Type:text/plain body:Raw 
    * Valid 1
    ```
    title: Valid App 1
    version: 0.0.1
    maintainers:
    - name: firstmaintainer app1
      email: firstmaintainer@hotmail.com
    - name: secondmaintainer app1
      email: secondmaintainer@gmail.com
    company: Random Inc.
    website: https://website.com
    source: https://github.com/random/repo
    license: Apache-2.0
    description: |
     ### Interesting Title
     Some application content, and description
    ```
    * Valid 2
    ```
    title: Valid App 2
    version: 1.0.1
    maintainers:
    - name: AppTwo Maintainer
      email: apptwo@hotmail.com
    company: Upbound Inc.
    website: https://upbound.io
    source: https://github.com/upbound/repo
    license: Apache-2.0
    description: |
     ### Why app 2 is the best
     Because it simply is...
    ```  
    
 * AksConfig/:title:
    * GET by Title localhost:1323/aksconfig/:title 

* AksConfig:
    * GET All localhost:1323/aksconfig

* AksConfig:
    * DELETE All localhost:1323/aksconfig
