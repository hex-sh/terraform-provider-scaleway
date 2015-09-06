# Quirks

The Scaleway documentation is extremely lacking for now. I hope this improves later. This is a place where I collect quirks



# Server creation

## IPs
Upon server creation, if one doesn't explicitly provide a reserved IP, a dynamic
IP is created. The documentation says this isn't supported but it still seems
supported...

## Implict volumes
Upon server creation, if one doesn't provide any volumes,  a root volume (of the image provided) is added. Which has index 0.
Hence if you provide :


```
"volumes": {
  "0": {
    "id": <UUID>
  }
}
```

the request will fail as 0 already implictly exists..., one has to do instead. or it won't work

```
"volumes": {
  "1": {
    "id": <UUID>
  }
}
```



## Adding volumes to a server

One can create a new volume  by doing POST /volumes but one can directly inline define a volume as wel...
even in  PATCH I think!

```
"volumes": {
  "1": {
    "id": <UUID>
  },
  "2": {
    "name": "blah",
    "type": "l_ssd",
    "size": 5000000
  }
}
```


Servers can't be PUT'ed though that _is_ in the documentation. one uses PATCH instead :$
