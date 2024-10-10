load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

gazelle_version = "v0.37.0"

http_archive(
    name = "bazel_gazelle",
    sha256 = "d76bf7a60fd8b050444090dfa2837a4eaf9829e1165618ee35dceca5cbdf58d5",
    urls = [
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/{version}/bazel-gazelle-{version}.tar.gz".format(version = gazelle_version),
    ],
)

http_archive(
    name = "io_bazel_rules_go",
    integrity = "sha256-M6zErg9wUC20uJPJ/B3Xqb+ZjCPn/yxFF3QdQEmpdvg=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.48.0/rules_go-v0.48.0.zip",
    ],
)

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")



go_rules_dependencies()
go_register_toolchains(version = "1.23.1")
gazelle_dependencies()

go_repository(
    name = "com_github_cockroachdb_apd_v3",
    importpath = "github.com/cockroachdb/apd/v3",
    sum = "h1:U+8j7t0axsIgvQUqthuNm82HIrYXodOV2iWLWtEaIwg=",
    version = "v3.2.1",
)

go_repository(
    name = "com_github_rs_zerolog",
    importpath = "github.com/rs/zerolog",
    sum = "h1:1cU2KZkvPxNyfgEmhHAz/1A9Bz+llsdYzklWFzgp0r8=",
    version = "v1.33.0",
)

go_repository(
    name = "com_github_coreos_go_systemd_v22",
    importpath = "github.com/coreos/go-systemd/v22",
    sum = "h1:RrqgGjYQKalulkV8NGVIfkXQf6YYmOyiJKk8iXXhfZs=",
    version = "v22.5.0",
)

go_repository(
    name = "com_github_godbus_dbus_v5",
    importpath = "github.com/godbus/dbus/v5",
    sum = "h1:9349emZab16e7zQvpmsbtjc18ykshndd8y2PG3sgJbA=",
    version = "v5.0.4",
)

go_repository(
    name = "com_github_lib_pq",
    importpath = "github.com/lib/pq",
    sum = "h1:p7ZhMD+KsSRozJr34udlUrhboJwWAgCg34+/ZZNvZZw=",
    version = "v1.10.7",
)

go_repository(
    name = "com_github_mattn_go_colorable",
    importpath = "github.com/mattn/go-colorable",
    sum = "h1:fFA4WZxdEF4tXPZVKMLwD8oUnCTTo08duU7wxecdEvA=",
    version = "v0.1.13",
)

go_repository(
    name = "com_github_mattn_go_isatty",
    importpath = "github.com/mattn/go-isatty",
    sum = "h1:xfD0iDuEKnDkl03q4limB+vH+GxLEtL/jb4xVJSWWEY=",
    version = "v0.0.20",
)

go_repository(
    name = "com_github_pkg_errors",
    importpath = "github.com/pkg/errors",
    sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
    version = "v0.9.1",
)

go_repository(
    name = "com_github_rs_xid",
    importpath = "github.com/rs/xid",
    sum = "h1:mKX4bl4iPYJtEIxp6CYiUuLQ/8DYMoz0PUdtGgMFRVc=",
    version = "v1.5.0",
)

go_repository(
    name = "org_golang_x_sys",
    importpath = "golang.org/x/sys",
    sum = "h1:KHjCJyddX0LoSTb3J+vWpupP9p0oznkqVk/IfjymZbo=",
    version = "v0.26.0",
)

# ### Go External Dependencies - BEGINS HERE
# go_repository(
#     name = "lib_golang_zerolog",
#     importpath = "github.com/rs/zerolog",
#     sum = "h1:1cU2KZkvPxNyfgEmhHAz/1A9Bz+llsdYzklWFzgp0r8=",
#     version = "v1.33.0",
# )

# go_repository(
#     name = "lib_golang_cockroachdb_apb",
#     importpath = "github.com/cockroachdb/apd/v3",
#     sum = "h1:U+8j7t0axsIgvQUqthuNm82HIrYXodOV2iWLWtEaIwg=",
#     version = "v3.2.1",
# )
# ### Go External Dependencies - Ends Here
