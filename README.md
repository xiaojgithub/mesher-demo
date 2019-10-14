# mesher-demo
this is a mesher demo, how to transform a common go service to mesher service.

## steps
1. add build.sh in base dir, like this demo, to build
2. provide a Dockerfile

## api
1. get /demo/hello   
response: Hello, go mesher demo   
description: in java mesher demo: mesher-demo-server case:
mesher-demo-server's api /demo/greeting will invoke it
```java
RequestOptions requestOptions = new RequestOptions();
        requestOptions.setHost("127.0.0.1").setPort(30101).setURI("http://demo-mesher:8080/demo/hello");

        Promise<String> promise = Promise.promise();
        client.get(requestOptions, resp -> resp.bodyHandler(body -> {
            if (resp.statusCode() > 200 || resp.statusCode() > 299) {
                LOGGER.error("get data failed, {}", body.toString());
                promise.fail("get data failed");
            } else {
                promise.complete(body.toString());
                LOGGER.info("data is {}", body.toString());
            }
        })).end();
```
2. get /demo/greeting   
this api will invoke java demo mesher-demo-server's /demo/hello api:
```go
	resp, body, errs := proxy.Get("http://demo-mesher-server:8090/demo/hello").EndBytes()
	if errs != nil {
		return nil, fmt.Errorf(fmt.Sprintf("do request catch a err:%#v", errs))
	}
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("request status not ok, %d", resp.StatusCode)
	}
```
