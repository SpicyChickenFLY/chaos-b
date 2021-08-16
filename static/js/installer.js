function addServer() {
    $(this).before($("#template .server-instance").clone());
}

function addInstance() {
    $(this).before($(this).siblings(".instance:last").clone());
}

function compileServInstInfoStr (mode) {
    let dataStr = "";
    $("div#" + mode + " div.server-instance").each(function() {
        let server = $(this).children("div.server");
        let name = server.children("name").val();
        let port = server.children("port").val();
        let user = server.children("user").val();
        let pwd = server.children("pwd").val();
        dataStr += name + ":" + port + "@" + user + ":" + pwd + "#";

        $(this).children("div.instance input.port").each(function() {
            dataStr += $(this).val() + "|"
        });
    })
    return dataStr.slice(0,-1)
}

function installStandardInstances() {
    let servInstInfoStr = compileServInstInfoStr("standard")
    let urlStr = "/api/v1/auto-mysql/standard";
    $.ajax({
        type: "POST",
        url: urlStr,
        data: {
            "InfoStr": servInstInfoStr,
            "SrcSQLFile": $("div#general input.src-mysql-file").val(),
            "SrcCnfFile": $("div#general input.src-cnf-file").val(),
            "MysqlPwd": $("div#general input.mysql-pwd").val()
        },
        dataType: "json",
        success: function (data, textStatus) {
            alert("Post\nurl:" + urlStr + "\nresult: success");
        },
        error: function (XMLHttpRequest, textStatus, errThrown) {
            alert("error");
        }
    });
}

$(document).ready(function () {
    $("#template").hide();
    $("#standard").children("button.add").click(addServer);
    $("#standard div.server-instance").children("button.add").click(addInstance);
});
