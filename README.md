# pgbouncer-authfile

Tiny project to create md5 authentication file for pgbouncer.

## Usage

Supply files with username and password using `-u` and `-p` command line options, specify output file with `-o` option.

```
pgbouncer-authfile -u=./username.txt -p=./password.txt -o=./userlist.txt
```

Plain text can be used instead of filenames:

```
pgbouncer-authfile -u=admin -p=./password.txt -o=./userlist.txt
```

Multiple usernames and passwords can be specified at the same time:

```
pgbouncer-authfile -u=admin -p=./password.txt -u=dummy -p=secret -o=./userlist.txt
```

By default passwords are written as plain text into output file:

```
$ pgbouncer-authfile -u admin -p password -u foo -p bar -u user -p secret
"admin" "password"
"foo" "bar"
"user" "secret"
```

But md5 hashing can be used if required (-t option):

```
./pgbouncer-authfile -u admin -p password -u foo -p bar -u user -p secret -t md5
"admin" "md580a19f669b02edfbc208a5386ab5036b"
"foo" "md596948aad3fcae80c08a35c9b5958cd89"
"user" "md520eb1b22a92b5c573dc1eb4331fc49ee"
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

* Add CircleCI automation to run tests and build & push images.
