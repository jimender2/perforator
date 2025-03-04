GO_LIBRARY()

LICENSE(
    Apache-2.0 AND
    BSD-3-Clause AND
    MIT
)

VERSION(v1.4.0)

SRCS(
    apic.go
    decode.go
    emitterc.go
    encode.go
    parserc.go
    readerc.go
    resolve.go
    scannerc.go
    sorter.go
    writerc.go
    yaml.go
    yamlh.go
    yamlprivateh.go
)

GO_XTEST_SRCS(
    decode_test.go
    encode_test.go
    example_embedded_test.go
    limit_test.go
    suite_test.go
)

END()

RECURSE(
    gotest
)
