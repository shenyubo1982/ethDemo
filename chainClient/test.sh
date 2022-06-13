#!/bin/bash

# 说明概述：单体测试工具。运行测试脚本，生成测试报告和覆盖率
# 作者：沈宇波
# 创建时间：2022.06.28
# 当前版本：v1.0
# 版本：v1.0
# 版本内容：


# go test -v -covermode=set -coverprofile=hint_test.out ./ >>hint_test.out
go test -v -covermode=set -coverprofile=../testing/hint_test_chainClient.out ./
go test -coverprofile=../testing/coverage_test_chainClient.out ./
go tool cover -html=../testing/coverage_test_chainClient.out
