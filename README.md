# pgbouncer-authfile

Tiny project to create md5 authentication file for pgbouncer.

## Usage

Supply files with username and password using `-u` and `-p` command line options, specify output file with `-o` option.

```
pgbouncer-authfile -u=./username.txt -p=./password.txt -o=./userlist.txt
```

## Sample deployment as init container

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: credentials
data:
  password: aWR6eEpuMG9Oem5rRFBnVg==
  username: NnJpdmVy
---
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: busybox
spec:
  replicas: 1
  selector:
    matchLabels:
      name: busybox
  template:
    metadata:
      labels:
        name: busybox
    spec:
      containers:
      - name: busybox
        image: busybox
        args:
        - sleep
        - "1000000"
        volumeMounts:
        - mountPath: /var/run/pgbouncer-auth
          name: pgbouncer-auth
      initContainers:
      - name: pgbouncer-auth
        image: gcr.io/plasma-column-128721/pgbouncer-authfile:0.0.1-scratch
        args:
        - -u=/var/run/in/username.txt
        - -p=/var/run/in/password.txt
        - -o=/var/run/out/userlist.txt
        imagePullPolicy: IfNotPresent
        volumeMounts:
        - mountPath: /var/run/out
          name: pgbouncer-auth
        - mountPath: /var/run/in
          name: credentials
      volumes:
      - name: pgbouncer-auth
        emptyDir: {}
      - name: credentials
        secret:
          secretName: credentials
          items:
          - key: username
            path: username.txt
          - key: password
            path: password.txt
```

## TODO

* Allow multiple username-password pairs to be specified.
* Allow username to be specified as plain text (not in file like now).
* Add command line option to specify password string format. -f txt|md5|scram. See [this link](https://www.pgbouncer.org/config.html#authentication-file-format) for details.