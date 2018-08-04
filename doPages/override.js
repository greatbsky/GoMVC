/**
 * 解析返回结果，每个项目规范不一致
 * {
"errorCode": 0,
"errorMsg": "成功",
"result": {}
}
 * @param data
 */
function parseResult(data) {
    var result = data
    if (data.errorCode == 0) {
        result.success = true
    }
    return result
}