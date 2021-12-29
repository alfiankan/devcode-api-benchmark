E[Sentry] 2021/12/28 08:13:49 Integration installed: ContextifyFrames
 @[Sentry] 2021/12/28 08:13:49 Integration installed: Environment
 <[Sentry] 2021/12/28 08:13:49 Integration installed: Modules
 A[Sentry] 2021/12/28 08:13:49 Integration installed: IgnoreErrors
 =2021/12/28 08:13:49 172.17.0.1 challenge_2_be devcode 123456
 ?[Sentry] 2021/12/28 08:13:49 Sending info event [0d26c1043ba94c86a36941edfcfeef01] to o913414.ingest.sentry.io project: 6125852
 ?[Sentry] 2021/12/28 08:13:49 Sending error event [267a1594d55f45ada47b5720d0ee1655] to o913414.ingest.sentry.io project: 6125852
 ?[Sentry] 2021/12/28 08:13:49 Sending info event [f9d9c9fb622e437fbd2f57497764f3de] to o913414.ingest.sentry.io project: 6125852
 ?[Sentry] 2021/12/28 08:13:49 Sending error event [126c8410b8bd4ea0aa0756fdb62a9599] to o913414.ingest.sentry.io project: 6125852
 ?[Sentry] 2021/12/28 08:13:49 Sending info event [3897ee5c0116487191d931094d35d434] to o913414.ingest.sentry.io project: 6125852
  
 ? ????????????????????????????????????????????????????? 
 < ?                   Devcode Todo                    ? 
 < ?                   Fiber v2.23.0                   ? 
 < ?               http://127.0.0.1:3030               ? 
 < ?       (bound on host 0.0.0.0 and port 3030)       ? 
 < ?                                                   ? 
 < ? Handlers ............ 14  Processes ........... 1 ? 
 < ? Prefork ....... Disabled  PID ................. 1 ? 
 ? ????????????????????????????????????????????????????? 
  
 ?[Sentry] 2021/12/28 08:13:59 Sending info event [f7aaac5951af498b90b8457f93b0896b] to o913414.ingest.sentry.io project: 6125852
 ?[Sentry] 2021/12/28 08:13:59 Sending info event [3eea8ec2b90844afa131491cb5b4f65a] to o913414.ingest.sentry.io project: 6125852
 Hpanic: runtime error: invalid memory address or nil pointer dereference
 H[signal SIGSEGV: segmentation violation code=0x1 addr=0x18 pc=0x733fae]
  
  goroutine 26 [running]:
 wdevcode/repository.(*TodoRepository).Add(0x0?, {0x0, 0x1, {0xc00047e230, 0xb}, 0x1, {0x989863, 0x9}, {0x0, 0x0}, ...})
 . /app/repository/todo_repository.go:53 +0x1ce
 Hdevcode/service.(*TodoService).Add(0xc8?, {0xc00047e230?, 0xb?}, 0x20?)
 ( /app/service/todo_service.go:42 +0x135
 Edevcode/controller.(*TodoController).Add(0xc00020ccf0, 0xc00010e840)