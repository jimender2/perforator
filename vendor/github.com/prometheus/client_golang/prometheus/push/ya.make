GO_LIBRARY()

LICENSE(Apache-2.0)

VERSION(v1.20.5)

SRCS(
    push.go
)

GO_TEST_SRCS(push_test.go)

GO_XTEST_SRCS(
    example_add_from_gatherer_test.go
    examples_test.go
)

END()

RECURSE(
    gotest
)
