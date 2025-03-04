GO_LIBRARY()

LICENSE(MIT)

VERSION(v0.11.1)

SRCS(
    maxheap.go
    quadtree.go
)

GO_TEST_SRCS(
    benchmarks_test.go
    maxheap_test.go
    quadtree_test.go
)

GO_XTEST_SRCS(example_test.go)

END()

RECURSE(
    gotest
)
