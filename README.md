## Momento Golang Sample

A simple repository that supports this [Article](https://www.binaryheap.com/caching-with-momento-and-golang). It highlights how to set up the Momento client and use it in a Golang Lambda Function.

### Up and Running

```bash
cdk deploy
# watch
cdk watch # this just is a nice little nugget
```

Then create the record route record

![DDB Record](https://www.binaryheap.com/wp-content/uploads/2023/05/ddb_record.png)

Lastly, let's run the Lambda with this event payload

```json
{
    "name": "sample",
    "correlationId": "abc"
}
```

Once you've done all of that, you'll get some output that looks like this.

First, run, you'll set the cache as a miss, the DDB is queried and then set.

![Hit Miss](https://www.binaryheap.com/wp-content/uploads/2023/05/mo_init_run.png)

Run it a second time, and you'll see the cache hit

![Output](https://www.binaryheap.com/wp-content/uploads/2023/05/mo_run.png)
