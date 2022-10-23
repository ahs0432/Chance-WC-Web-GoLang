function getList() {
    var list = new Array();

    $("input[name=mail]").each(function(index, item){
        list.push($(item).val());
    });
    $("#mailList").val(list);
}

function getWebList() {
    var list = new Array();

    $("input[name=web]").each(function(index, item){
        
        if ($(item).is(":checked") == true) {
            list.push($(item).val());
        }
    });
    $("#webList").val(list);
}

function isValidSubmit() {
    if ($("input[name=url]").val().indexOf("http://") == -1 && $("input[name=url]").val().indexOf("https://") == -1) {
        window.alert("URL에 http:// 또는 https://가 포함되어야 합니다.")
        return false;
    }

    return true;
}

var app;
var mailcnt=1;

function mailAppend() {
    if(mailcnt == 5) {
        window.alert("메일은 최대 5개까지만 추가할 수 있습니다.");
    } else {
        app = document.getElementById("mailContent");
        var oRow = app.insertRow();
        oRow.onmouseover = function() {
            app.clickedRowIndex=this.rowIndex;
        }
        var oCell = oRow.insertCell();
        oCell.innerHTML = '<input type="email" name="mail" placeholder="E-Mail (예: test1@test.com)" value="" required minlength="0" maxlength="500"> <input type="button" value="-" onclick="mailRemove()">';
        mailcnt+=1;
    }
}

function mailAppendMod(data) {
    if(mailcnt == 5) {
        window.alert("메일은 최대 5개까지만 추가할 수 있습니다.");
    } else {
        app = document.getElementById("mailContent");
        var oRow = app.insertRow();
        oRow.onmouseover = function() {
            app.clickedRowIndex=this.rowIndex;
        }
        var oCell = oRow.insertCell();
        oCell.innerHTML = '<input type="email" name="mail" placeholder="E-Mail (예: test1@test.com)" value="'+ data +'" required minlength="0" maxlength="500"> <input type="button" value="-" onclick="mailRemove()">';
        mailcnt+=1;
    }
}

function mailRemove() {
    app.deleteRow(app.clickedRowIndex);
    mailcnt-=1;
}