# jsonify dir
Print out a file system structure as json as follows:
```
{
  "folders": [
    {
      "name": ".git",
      "size": 408
    },
    {
      "name": ".idea",
      "size": 272
    },
    {
      "name": "jsonify",
      "size": 136
    }
  ]
}
```
# Why?
Why not?

# How do I get it running?
Pull the repo and run `go build`. Then run `./jsonifyDir .` where dir is the is name of the folder you want to jsonify.