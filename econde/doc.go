package econde

//[{"content": "123", "commitTime": "2021-11-23T06:08:34.736Z", "intervieweeId": "1041221229477908", "reserveEnterTime": "2021-11-22T22:08:21.105Z"}, {"content": "123", "commitTime": "2021-11-23T06:08:34.736Z", "intervieweeId": "1041221229477908", "reserveEnterTime": "2021-11-22T22:08:21.105Z"}, {"content": "123", "checkState": 3, "commitTime": "2021-11-23T06:08:34.736Z", "intervieweeId": "1041221229477908", "reserveEnterTime": "1970-01-01T00:00:00Z"}, {"content": "123", "checkState": 2, "commitTime": "2021-11-23T06:08:34.736Z", "intervieweeId": "1041221229477908", "reserveEnterTime": "2021-11-19T18:20:01.217Z"}]
//[{"content": "创建告警", "user_id": "MOCK", "update_at": {"nanos": 903267087, "seconds": 1631600221}}]

//对于时间戳：MustMarshalProtoArray 生成的是字符串形式的jsonb数据
//MustMarshal 生成毫秒数和纳秒数表达的值
//Marshal 使用的是 json解析，  MustMarshalProtoArray 使用的是反射  两者生成的jsonb数据不能互转。
