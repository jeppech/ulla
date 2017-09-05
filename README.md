# Ulla - UDP Load Latency Announcer
Ulla is a lightweight UDP server, that digests and announces
latencies/rpm and more, for services, that broadcasts data to Ulla.

---

Ulla accepts the simple CSV format, sent as an UDP packet.

```
[string] service;[int] ms;[string] method;[string] path(optional)
```

Example
```
Notification Service;120;post
```


---
The server will broadcast it's services by WebSocket, to any subscribers,
every 5th seconds.
```javascript
[
    {
        "service": "messaging",     // name of service
        "average": 120,             // Latency in ms, over a 5 minutes
        "weekly_average": 80        // Latency in ms over 1 week, if data is available
        "requests_per_minute": 20   //
    }
]
```