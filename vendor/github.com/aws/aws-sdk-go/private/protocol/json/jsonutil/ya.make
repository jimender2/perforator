GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.46.7)

SRCS(
    build.go
    unmarshal.go
)

GO_XTEST_SRCS(
    build_test.go
    unmarshal_test.go
)

END()

RECURSE(
    gotest
)
