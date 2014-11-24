#!/usr/bin/env bash
thrift -gen go:thrift_import=git-wip-us.apache.org/repos/asf/thrift.git/lib/go/thrift contact.thrift
thrift -gen py contact.thrift
