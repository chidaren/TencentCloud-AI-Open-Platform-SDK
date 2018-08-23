####  EXAMPLE - Image Filter Effects
    client := tcloud.NewClient(appid, appkey, connTimeout, transportTimeout, maxCoonPerHost, maxIdleCoons)
    data, err := ioutil.ReadFile("example.jpg")
    if err != nil {
      log.Fatal(err)
    }
    
    req := &FilterRequest{4, []byte(base64.StdEncoding.EncodeToString(data))}
    res, err := client.Filter(req)
    if err != nil {
      log.Fatal(err)
    }
    // done!
    log.Info(res)

