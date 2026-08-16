# go-duration2

把"人话时长"解析成 Go 的 `time.Duration`。标准库的 `ParseDuration` 只认 `1h30m` 这种，这个在它基础上多了中文口语单位——`1天2小时30分钟` 也能吃，写配置、读用户输入更顺手。

纯数字按秒算，认不出的直接报错。解析出来的就是标准 `time.Duration`，后面怎么加减、格式化都随便。

## 装

```bash
go build -o duration2 ./cmd/duration2
```

## 用

```bash
echo "1h30m" | ./duration2          # 1h30m0s (540000000000 ns)
echo "1天2小时30分钟" | ./duration2  # 26h30m0s
echo "90" | ./duration2             # 1m30s（按秒）
```

## 当库用

```go
import "duration2"

d, _ := duration2.Parse("1天2小时30分钟")  // 26h30m0s
d, _ = duration2.Parse("1h30m")           // 90m0s
```

中文单位支持：天 / 小时 / 分钟 / 秒 / 毫秒。

## License

MIT
