/**
 * 解析返回结果，每个项目规范不一致
 * {
"errorCode": 0,
"errorMsg": "成功",
"result": {}
}
 * @param data
 *
 * 解析为
 * {
 * success: true,
 * err: 123,
 * errmsg: "msg...",
 * rows: []
 * }
 */
function parseResult(data) {
    console.log("parseResult(data):")
    console.log(data)
    var result = data
    if (data.errorCode == 0) {
        result.success = true
    } else if (typeof data.errorCode != "undefined") {
        result.success = false
        result.msg = data.errorMsg
    } else {
        result.msg = "服务端返回结果不正确"
    }
    if (typeof data.result != "undefined") {
        result.rows = data.result.data
        result.total = data.result.total
        result.adminuid = data.result.pn
    }
    return result
}