## Installation and Run

```
$ cd aksservice
$ make build
$ make run
```

### Good URL examples
* AksValidator:
    * POST localhost:1323/aksvalidator Content-Type:text/plain body:Raw
    * Valid 1
    ```
    title: App w/ Invalid maintainer email
    version: 1.1.0
    maintainers:
    - name: firstmaintainer app1
    email: firstmaintainer@asd.com
    - name: secondmaintainer app1
    email: secondmaintainer@gmail.com
    company: Upbound Inc.
    website: https://upbound.io
    source: https://github.com/upbound/repo
    license: Apache-2.0
    description: |
    ### blob of markdown
    More markdown
    ```
    *Vaid 2
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
