## Installation and Run

```
$ cd aksservice
$ make build
$ make run
```

### Good URL examples
* AksValidator:
    * POST localhost:1323/aksvalidator
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
